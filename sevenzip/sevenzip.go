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

package sevenzip

import (
	"github.com/ricochhet/minicommon/filesystem"
	"github.com/ricochhet/minicommon/process"
)

func SzExtract(src, dest string, silent bool) (ErrorCode, error) {
	if filesystem.Exists("redist/win64/7z.exe") {
		if err := process.RunFile("redist/win64/7z.exe", true, true, false, "x", src, "-o"+dest+"/*"); err != nil {
			return CouldNotCompress, err
		}

		return NoError, nil
	}

	if !process.DoesFileExist("7z") {
		return ProcessNotFound, ErrSevenZipNotFound
	}

	if err := process.RunFile("7z", true, false, silent, "x", src, "-o"+dest+"/*"); err != nil {
		return CouldNotExtract, err
	}

	return NoError, nil
}

func SzBinExtract(src, dest, bin string, silent bool) (ErrorCode, error) {
	if !filesystem.Exists(bin) {
		return ProcessNotFound, ErrSevenZipNotFound
	}

	if err := process.RunFile(bin, true, true, silent, "x", src, "-o"+dest+"/*"); err != nil {
		return CouldNotCompress, err
	}

	return NoError, nil
}

func SzCompress(src, dest string, silent bool, opts ...Options) (ErrorCode, error) {
	opt := assureOptions(opts...)

	if !process.DoesFileExist("7z") {
		return ProcessNotFound, ErrSevenZipNotFound
	}

	//nolint:lll // wontfix
	if err := process.RunFile("7z", true, false, silent, "a", "-t"+opt.SzCompressionFormat, dest, src+"/*", opt.SzCompressionLevel, opt.SzCompressionMethod, opt.SzCompressionDictionarySize, opt.SzCompressionFastBytes, opt.SzCompressionSolidBlockSize, opt.SzCompressionMultithreading, opt.SzCompressionMemory); err != nil {
		return CouldNotCompress, err
	}

	return NoError, nil
}

func SzBinCompress(src, dest, bin string, silent bool, opts ...Options) (ErrorCode, error) {
	opt := assureOptions(opts...)

	if !filesystem.Exists(bin) {
		return ProcessNotFound, ErrSevenZipNotFound
	}

	//nolint:lll // wontfix
	if err := process.RunFile(bin, true, true, silent, "a", "-t"+opt.SzCompressionFormat, dest, src+"/*", opt.SzCompressionLevel, opt.SzCompressionMethod, opt.SzCompressionDictionarySize, opt.SzCompressionFastBytes, opt.SzCompressionSolidBlockSize, opt.SzCompressionMultithreading, opt.SzCompressionMemory); err != nil {
		return CouldNotCompress, err
	}

	return NoError, nil
}
