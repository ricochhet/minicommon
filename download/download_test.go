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

package download_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ricochhet/minicommon/download"
)

const testDownloadURL = "https://raw.githubusercontent.com/ricochhet/minicommon/main/LICENSE"

func TestGenericDownload(t *testing.T) {
	t.Parallel()

	testMessenger := download.Messenger{
		StartDownload: func(fname string) {
			fmt.Printf("Test download: %s\n", fname)
		},
	}

	if bytes, err := download.Download(testDownloadURL); err != nil || len(bytes) == 0 {
		t.Fatal("download fail")
	}

	if bytes, err := download.WithContext(context.TODO(), testMessenger, testDownloadURL); err != nil || len(bytes) == 0 {
		t.Fatal("download fail")
	}
}

func TestFileDownload(t *testing.T) {
	t.Parallel()

	if err := download.File(testDownloadURL, "LICENSE", "./.test/"); err != nil {
		t.Fatal(err)
	}

	if bytes, err := download.FileWithBytes(testDownloadURL, "LICENSE", "./.test/"); err != nil || len(bytes) == 0 {
		t.Fatal("download fail")
	}
}

//nolint:lll // test only
func TestFileValidated(t *testing.T) {
	t.Parallel()

	if err := download.FileValidated(testDownloadURL, "aaabbbccc", "LICENSE", "./.test/"); err == nil {
		t.Fatal("download fail")
	}

	if bytes, err := download.FileWithBytesValidated(testDownloadURL, "aaabbbccc", "LICENSE", "./.test/"); err == nil || len(bytes) != 0 {
		t.Fatal("download fail")
	}

	if err := download.FileValidated(testDownloadURL, "8486a10c4393cee1c25392769ddd3b2d6c242d6ec7928e1414efff7dfb2f07ef", "LICENSE", "./.test/"); err != nil {
		t.Fatal(err)
	}

	if bytes, err := download.FileWithBytesValidated(testDownloadURL, "8486a10c4393cee1c25392769ddd3b2d6c242d6ec7928e1414efff7dfb2f07ef", "LICENSE", "./.test/"); err != nil || len(bytes) == 0 {
		t.Fatal("download fail")
	}
}

//nolint:lll // test only
func TestFileDownloadWithHash(t *testing.T) {
	t.Parallel()

	testMessenger := download.Messenger{
		StartDownload: func(fname string) {
			fmt.Printf("Test download: %s\n", fname)
		},
	}

	if err := download.FileWithContext(context.TODO(), testMessenger, testDownloadURL, "8486a10c4393cee1c25392769ddd3b2d6c242d6ec7928e1414efff7dfb2f07ef", "LICENSE", "./.test/", download.DefaultHashValidator); err != nil {
		t.Fatal(err)
	}

	if err := download.FileWithContext(context.TODO(), testMessenger, testDownloadURL, "", "LICENSE", "./.test/", download.DefaultHashValidator); err == nil {
		t.Fatal("empty hash has validated successfully")
	}

	if bytes, err := download.FileWithContextAndBytes(context.TODO(), testMessenger, testDownloadURL, "", "LICENSE", "./.test/", nil); err != nil || len(bytes) == 0 {
		t.Fatal("download fail")
	}
}
