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

package lang

import (
	"errors"
	"sync"

	"golang.org/x/text/language"
)

var (
	langMap  = map[string]map[string]string{} //nolint:gochecknoglobals // wontfix
	fallBack = "en"                           //nolint:gochecknoglobals // wontfix
	lang     map[string]string                //nolint:gochecknoglobals // wontfix
	lock     = sync.RWMutex{}                 //nolint:gochecknoglobals // wontfix
)

func SetLanguage(languge, fallback string, langmap map[string]map[string]string) error {
	fallBack = fallback
	langMap = langmap

	lock.Lock()

	defer lock.Unlock()

	tag := language.Make(languge)

	selected, _ := tag.Base()

	if langMap[selected.String()] == nil {
		return errLanguageNotFound
	}

	lang = langMap[selected.String()]

	return nil
}

func Lang(key string) string {
	lock.RLock()

	defer lock.RUnlock()

	word, ok := lang[key]
	if !ok {
		return langMap[fallBack][key]
	}

	return word
}

var errLanguageNotFound = errors.New("language not found")
