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

package charmbracelet

import (
	"io"
	"time"

	"github.com/charmbracelet/log"
)

type MultiLogger struct {
	writers []io.Writer
	loggers []*log.Logger
}

var SharedLogger *MultiLogger //nolint:gochecknoglobals // wontfix

func NewMultiLogger(wrs ...io.Writer) *MultiLogger {
	loggers := new(MultiLogger)

	loggers.writers = make([]io.Writer, len(wrs))
	loggers.loggers = make([]*log.Logger, len(wrs))

	for i, w := range wrs {
		loggers.writers[i] = w
		loggers.loggers[i] = log.NewWithOptions(w, log.Options{ //nolint:exhaustruct // wontfix
			ReportCaller:    true,
			ReportTimestamp: true,
			TimeFormat:      time.Kitchen,
		})
	}

	return loggers
}

func RegisterLogger(wrs ...io.Writer) {
	SharedLogger = NewMultiLogger(wrs...)
}

func (ml *MultiLogger) Debug(msg any, kvs ...any) {
	for _, l := range ml.loggers {
		l.Debug(msg, kvs...)
	}
}

func (ml *MultiLogger) Info(msg any, kvs ...any) {
	for _, l := range ml.loggers {
		l.Info(msg, kvs...)
	}
}

func (ml *MultiLogger) Warn(msg any, kvs ...any) {
	for _, l := range ml.loggers {
		l.Warn(msg, kvs...)
	}
}

func (ml *MultiLogger) Error(msg any, kvs ...any) {
	for _, l := range ml.loggers {
		l.Error(msg, kvs...)
	}
}

func (ml *MultiLogger) Fatal(msg any, kvs ...any) {
	for _, l := range ml.loggers {
		l.Fatal(msg, kvs...)
	}
}

func (ml *MultiLogger) Print(msg any, kvs ...any) {
	for _, l := range ml.loggers {
		l.Print(msg, kvs...)
	}
}
