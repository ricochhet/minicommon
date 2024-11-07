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

package filesystem

import (
	"embed"
	"encoding/json"
	"fmt"
)

type EmbeddedFileSystem struct {
	Initial string
	FS      embed.FS
}

func (e EmbeddedFileSystem) BytesToMap(data []byte) (map[string]interface{}, error) {
	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (e EmbeddedFileSystem) FilenameToMap(filename string) (map[string]interface{}, error) {
	data, err := e.FS.ReadFile(e.Initial + filename)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (e EmbeddedFileSystem) FilenameToBytes(filename string) ([]byte, error) {
	data, err := e.FS.ReadFile(e.Initial + filename)
	if err != nil {
		return nil, err
	}

	var rawMessage json.RawMessage
	if err := json.Unmarshal(data, &rawMessage); err != nil {
		return nil, fmt.Errorf("invalid JSON format: %w", err)
	}

	return data, nil
}
