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
	"fmt"
)

func HexStringToBytes(hexStr string) ([]byte, error) {
	var bytes []byte

	for i := 0; i < len(hexStr); i += 2 {
		var newByte byte

		_, err := fmt.Sscanf(hexStr[i:i+2], "%02X", &newByte)
		if err != nil {
			return nil, err
		}

		bytes = append(bytes, newByte)
	}

	return bytes, nil
}

func FindAllByteOccurrences(data []byte, pattern []byte) []int {
	var indices []int

	for i := range data {
		if bytes.HasPrefix(data[i:], pattern) {
			indices = append(indices, i)
		}
	}

	return indices
}

func ReplaceByteOccurrences(original []byte, expectedBytes []byte, replacement []byte, occurrenceToReplace int) []byte {
	var result []byte

	remaining := original
	occurrenceCount := 0

	for {
		index := bytes.Index(remaining, expectedBytes)
		if index == -1 {
			result = append(result, remaining...)
			break
		}

		result = append(result, remaining[:index]...)

		occurrenceCount++

		if occurrenceToReplace == 0 || occurrenceCount == occurrenceToReplace {
			replacementLen := len(replacement)

			if replacementLen > len(expectedBytes) {
				replacementLen = len(expectedBytes)
			}

			result = append(result, replacement[:replacementLen]...)
		} else {
			result = append(result, remaining[index:index+len(expectedBytes)]...)
		}

		remaining = remaining[index+len(expectedBytes):]
	}

	return result
}
