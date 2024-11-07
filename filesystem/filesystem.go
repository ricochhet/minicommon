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

package filesystem

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/otiai10/copy"
)

var (
	errNoNameFound = errors.New("name could not be found in file path")
	errNoPathFound = errors.New("file path could not be found in file name")
	errFileExists  = errors.New("file exists in destination path")
)

var ReservedHostnames = []string{ //nolint:gochecknoglobals // wontfix
	"COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9",
	"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9",
	"PRN", "AUX", "NUL",
}

func Combine(path1 string, path2 ...string) string {
	path := append([]string{path1}, path2...)
	return filepath.Join(path...)
}

func FromCwd(path1 ...string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path := append([]string{wd}, path1...)

	return filepath.Join(path...), nil
}

func GetDirectoryName(fileName string) string {
	return filepath.Dir(fileName)
}

func GetFileName(fileName string) string {
	return strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
}

func GetFileExtension(fileName string) string {
	return filepath.Ext(fileName)
}

func GetRelativePath(directories ...string) string {
	result := "./" + directories[0]

	for _, dir := range directories[1:] {
		result = path.Join(result, dir)
	}

	return result
}

func TrimPath(input string) string {
	if strings.HasPrefix(input, "./") || strings.HasPrefix(input, ".\\") {
		return input[2:]
	} else if strings.HasPrefix(input, "/") || strings.HasPrefix(input, "\\") {
		return input[1:]
	}

	return input
}

func Copy(a, b string, opts ...copy.Options) error {
	if err := copy.Copy(a, b, opts...); err != nil {
		return err
	}

	return nil
}

func CopyAndRename(files []string, oldPath, newPath, oldName, newName string) error {
	found := false

	for _, file := range files {
		if strings.Contains(file, oldName) {
			found = true
			break
		}
	}

	if !found {
		return errNoNameFound
	}

	for _, file := range files {
		newFileName := strings.ReplaceAll(file, oldName, newName)

		if !strings.Contains(newFileName, TrimPath(oldPath)) {
			return errNoPathFound
		}

		newFilePath := strings.ReplaceAll(newFileName, TrimPath(oldPath), newPath)

		if Exists(newFilePath) {
			return errFileExists
		}

		if err := Copy(file, newFilePath); err != nil {
			return err
		}
	}

	return nil
}

func Exists(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func ReadFile(fileName string) ([]byte, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func ReadAllLines(file *os.File) ([]string, error) {
	return Scan(bufio.NewScanner(file))
}

func ReadAllStringLines(input string) ([]string, error) {
	return Scan(bufio.NewScanner(strings.NewReader(input)))
}

func WriteFile(fileName string, data []byte, perm fs.FileMode) error {
	err := os.WriteFile(fileName, data, perm)
	if err != nil {
		return err
	}

	return nil
}

func WriteToFile(file *os.File, entries []string) error {
	for _, entry := range entries {
		if _, err := file.WriteString(entry); err != nil {
			return err
		}
	}

	return nil
}

func OverwriteFile(file *os.File) error {
	if err := file.Truncate(0); err != nil {
		return err
	}

	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	return nil
}

func Scan(scanner *bufio.Scanner) ([]string, error) {
	entries := []string{}

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}

		entries = append(entries, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

func DeleteDirectory(fileName string) error {
	err := os.RemoveAll(fileName)
	if err != nil {
		return err
	}

	return nil
}

func DeleteEmptyDirectories(root string) error {
	dirs := []string{}

	err := filepath.WalkDir(root, func(path string, dir os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if dir.IsDir() && path != root {
			dirs = append(dirs, path)
		}

		return nil
	})
	if err != nil {
		return err
	}

	for i := len(dirs) - 1; i >= 0; i-- {
		dir := dirs[i]

		empty, err := IsEmpty(dir)
		if err != nil {
			return err
		}

		if empty {
			err = os.Remove(dir)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func SortFileNames(files []string) []string {
	sort.Slice(files, func(i, j int) bool {
		parentA := filepath.Dir(files[i])
		parentB := filepath.Dir(files[j])

		if parentA == parentB {
			return filepath.Base(files[i]) < filepath.Base(files[j])
		}

		return parentA < parentB
	})

	return files
}

func GetFiles(filePath string) []string {
	var files []string

	err := filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		return []string{}
	}

	return SortFileNames(files)
}

func GetDirectories(filePath string) []string {
	var directories []string

	err := filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			directories = append(directories, path)
		}

		return nil
	})
	if err != nil {
		return []string{}
	}

	return SortFileNames(directories)
}

func IsEmpty(dir string) (bool, error) {
	file, err := os.Open(dir)
	if err != nil {
		return false, err
	}
	defer file.Close()

	_, err = file.Readdir(1)

	if err == nil {
		return false, nil
	}

	if errors.Is(err, os.ErrNotExist) || err.Error() == "EOF" {
		return true, nil
	}

	return false, err
}

func BytesToMap(data []byte) (map[string]interface{}, error) {
	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	return jsonData, nil
}

func FilenameToMap(initial, filename string) (map[string]interface{}, error) {
	data, err := os.ReadFile(initial + filename)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	return jsonData, nil
}

func FilenameToBytes(initial, filename string) ([]byte, error) {
	data, err := os.ReadFile(initial + filename)
	if err != nil {
		return nil, err
	}

	var rawMessage json.RawMessage
	if err := json.Unmarshal(data, &rawMessage); err != nil {
		return nil, fmt.Errorf("invalid JSON format: %w", err)
	}

	return data, nil
}

func IsValidHostname(hostname string) bool {
	if len(hostname) < 1 || len(hostname) > 15 {
		return false
	}

	validNameRegex := regexp.MustCompile(`^[a-zA-Z0-9-]+$`)

	if !validNameRegex.MatchString(hostname) {
		return false
	}

	for _, reserved := range ReservedHostnames {
		if strings.EqualFold(hostname, reserved) {
			return false
		}
	}

	return true
}
