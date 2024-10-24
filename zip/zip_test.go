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

package zip_test

import (
	"testing"

	"github.com/ricochhet/minicommon/zip"
)

func TestZip(t *testing.T) { //nolint:paralleltest // dependant
	if err := zip.Zip("./", "./.test/simplezip-src.zip"); err != nil {
		t.Fatal(err)
	}
}

func TestUnzip(t *testing.T) { //nolint:paralleltest // dependant
	if err := zip.Unzip("./.test/simplezip-src.zip", "./.test/src"); err != nil {
		t.Fatal(err)
	}
}
