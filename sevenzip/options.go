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

type Options struct {
	SzCompressionFormat         string
	SzCompressionLevel          string
	SzCompressionMethod         string
	SzCompressionDictionarySize string
	SzCompressionFastBytes      string
	SzCompressionSolidBlockSize string
	SzCompressionMultithreading string
	SzCompressionMemory         string
}

func getDefaultOptions() Options {
	return Options{
		SzCompressionFormat:         "7z",
		SzCompressionLevel:          "-mx9",
		SzCompressionMethod:         "-m0=lzma2",
		SzCompressionDictionarySize: "-md=64m",
		SzCompressionFastBytes:      "-mfb=64",
		SzCompressionSolidBlockSize: "-ms=4g",
		SzCompressionMultithreading: "-mmt=2",
		SzCompressionMemory:         "-mmemuse=26g",
	}
}

func assureOptions(opts ...Options) Options {
	defopt := getDefaultOptions()

	if len(opts) == 0 {
		return defopt
	}

	return opts[0]
}
