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

func MoveEntry[T comparable](slice []T, entry T, newIndex int) []T {
	currentIndex := -1

	for i, s := range slice {
		if s == entry {
			currentIndex = i
			break
		}
	}

	if currentIndex == -1 {
		return slice
	}

	slice = append(slice[:currentIndex], slice[currentIndex+1:]...)

	if newIndex >= len(slice) {
		slice = append(slice, entry)
	} else {
		slice = append(slice[:newIndex], append([]T{entry}, slice[newIndex:]...)...)
	}

	return slice
}
