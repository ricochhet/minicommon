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

import "strings"

func MapKeyValuePairs(lines []string) (map[string]string, error) {
	keyvalues := map[string]string{}

	for _, line := range lines {
		keyvalue := strings.SplitN(line, "=", 2) //nolint:mnd // wontfix

		if len(keyvalue) != 2 { //nolint:mnd // wontfix
			continue
		}

		keyvalues[keyvalue[0]] = keyvalue[1]
	}

	return keyvalues, nil
}
