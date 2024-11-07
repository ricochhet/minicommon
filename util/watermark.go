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
	"fmt"
	"regexp"
	"strings"
)

func DrawWatermark(text []string, draw func(string)) {
	result := []string{}

	longestLength := 0

	for _, textItem := range text {
		length := textLength(textItem)
		if length > longestLength {
			longestLength = length
		}
	}

	line := strings.Repeat("-", longestLength)
	result = append(result, fmt.Sprintf("┌─%s─┐", line))

	for _, textItem := range text {
		spacingSize := longestLength - textLength(textItem)
		spacingText := textItem + strings.Repeat(" ", spacingSize)
		result = append(result, fmt.Sprintf("│ %s │", spacingText))
	}

	result = append(result, fmt.Sprintf("└─%s─┘", line))

	for _, textItem := range result {
		draw(textItem)
	}
}

func textLength(s string) int {
	re := regexp.MustCompile(`[\p{Han}\p{Katakana}\p{Hiragana}\p{Hangul}]`)
	processedString := re.ReplaceAllString(s, "ab")

	return len(processedString)
}
