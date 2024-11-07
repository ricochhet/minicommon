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

package aflag

import (
	"flag"
	"strconv"
)

func StrVar(ptr *string, name, value, usage string, keyvalues map[string]string) {
	flag.StringVar(ptr, name, value, usage)
	setStr(ptr, name, keyvalues)
}

func BoolVar(ptr *bool, name string, value bool, usage string, keyvalues map[string]string) {
	flag.BoolVar(ptr, name, value, usage)
	setBool(ptr, name, keyvalues)
}

func setStr(flag *string, key string, kvp map[string]string) {
	val := kvp[key]

	if val == "" {
		return
	}

	*flag = val
}

func setBool(flag *bool, key string, kvp map[string]string) {
	val := kvp[key]

	if val == "" {
		return
	}

	b, parseErr := strconv.ParseBool(val)

	*flag = b

	if parseErr != nil {
		panic(parseErr)
	}
}
