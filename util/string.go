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

package util

import (
	"bytes"
	"unicode/utf16"
	"unsafe"
)

func StringToBytes(input string) []byte {
	tmp := []byte(input)
	tmp = append(tmp, bytes.Repeat([]byte{0}, 2-(len(tmp)%2))...) //nolint:mnd // wontfix

	return tmp
}

func GetStringFromBytes(data []byte, start, end int) string {
	var contentID string

	if end > len(data) {
		end = len(data)
	}

	rawSlice := data[start:end]
	u16Slice := ((*[1 << 30]uint16)(unsafe.Pointer(&rawSlice[0])))[:len(rawSlice)/2]

	nullIndex := -1

	for i, c := range u16Slice {
		if c == 0 {
			nullIndex = i
			break
		}
	}

	if nullIndex != -1 {
		contentID = string(utf16.Decode(u16Slice[:nullIndex]))
	} else {
		contentID = string(utf16.Decode(u16Slice))
	}

	return contentID
}
