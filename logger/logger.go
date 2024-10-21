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

package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	GoRoutineErrorLevel
)

type Logger struct {
	CallerDepth int
	MinLevel    LogLevel
	Writer      io.Writer
	Flag        int
}

func NewLogger(callerDepth int, minLevel LogLevel, writer io.Writer, flag int) *Logger {
	log.SetOutput(writer)
	log.SetFlags(flag)

	return &Logger{CallerDepth: callerDepth, MinLevel: minLevel, Writer: writer, Flag: flag}
}

func (l *Logger) log(level LogLevel, message string) {
	if l.Writer == nil {
		fmt.Fprintln(os.Stdout, "Logger STDOUT is nil")
		return
	}

	levelName := "DEBUG"

	switch level {
	case DebugLevel:
		levelName = "DEBUG"
	case InfoLevel:
		levelName = "INFO"
	case WarnLevel:
		levelName = "WARN"
	case ErrorLevel:
		levelName = "ERROR"
	case FatalLevel:
		levelName = "FATAL"
	case GoRoutineErrorLevel:
		levelName = "GOROUTINE_ERROR"
	}

	if level >= l.MinLevel {
		oRaw := fmt.Sprintf("[%s] %s\n", levelName, message)

		if err := log.Output(l.CallerDepth, oRaw); err != nil {
			panic(err)
		}
	}
}

func (l *Logger) Debug(format string) {
	l.log(DebugLevel, format)
}

func (l *Logger) Debugf(format string, a ...any) {
	l.log(DebugLevel, fmt.Sprintf(format, a...))
}

func (l *Logger) Info(format string) {
	l.log(InfoLevel, format)
}

func (l *Logger) Infof(format string, a ...any) {
	l.log(InfoLevel, fmt.Sprintf(format, a...))
}

func (l *Logger) Warn(format string) {
	l.log(WarnLevel, format)
}

func (l *Logger) Warnf(format string, a ...any) {
	l.log(WarnLevel, fmt.Sprintf(format, a...))
}

func (l *Logger) Error(format string) {
	l.log(ErrorLevel, format)
}

func (l *Logger) Errorf(format string, a ...any) {
	l.log(ErrorLevel, fmt.Sprintf(format, a...))
}

func (l *Logger) Fatal(format string) {
	l.log(FatalLevel, format)
	// os.Exit(1)
	l.wait()
}

func (l *Logger) Fatalf(format string, a ...any) {
	l.log(FatalLevel, fmt.Sprintf(format, a...))
	// os.Exit(1)
	l.wait()
}

func (l *Logger) wait() {
	l.log(FatalLevel, "Press CTRL+C to Exit.")

	_, _ = fmt.Scanln()
}

var SharedLogger *Logger //nolint:gochecknoglobals // wontfix
