/*
 * minicommon
 * Copyright (C) 2024 minicommon contributors
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.

 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package data

import (
	"encoding/json"

	"github.com/ricochhet/minicommon/filesystem"
	"github.com/tidwall/gjson"
	"github.com/vmihailenco/msgpack"
)

func Encode(value interface{}) ([]byte, error) {
	pk, err := msgpack.Marshal(value)
	if err != nil {
		return nil, err
	}

	return pk, nil
}

func Decode(data []byte) (interface{}, error) {
	var content interface{}

	err := msgpack.Unmarshal(data, &content)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func DecodeFile(filename string) (string, error) {
	content, err := filesystem.ReadFile(filename)
	if err != nil {
		return "", err
	}

	decdata, err := Decode(content)
	if err != nil {
		return "", err
	}

	jsoncontent, err := json.MarshalIndent(decdata, "", " ")
	if err != nil {
		return "", err
	}

	return string(jsoncontent), nil
}

func EncodeFile(filename string) ([]byte, error) {
	content, err := filesystem.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	encdata, err := Encode(gjson.ParseBytes(content).Value())
	if err != nil {
		return nil, err
	}

	return encdata, nil
}
