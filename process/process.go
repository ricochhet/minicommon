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

package process

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"
)

func DoesFileExist(name string) bool {
	if _, err := exec.LookPath(name); err != nil {
		return false
	}

	return true
}

func RunFile(name string, hideWindow, relativeExecutable, redirectStd bool, arg ...string) error {
	path := name

	if relativeExecutable {
		cwd, err := os.Executable()
		if err != nil {
			return err
		}

		path = filepath.Join(filepath.Dir(cwd), name)
	}

	cmd := exec.Command(path, arg...)

	if redirectStd {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if runtime.GOOS == "windows" {
		// cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: hideWindow} //nolint:exhaustruct // wontfix
	}

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
