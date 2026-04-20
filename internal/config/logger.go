package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	err     *log.Logger
	info    *log.Logger
	warning *log.Logger
	writer  io.Writer
}

func NewLogger() *Logger {
	writer := io.Writer(os.Stdout)
	flags := log.Ldate | log.Ltime

	return &Logger{
		debug:   log.New(writer, "[DEBUG] ", flags),
		err:     log.New(writer, "[ERROR] ", flags),
		info:    log.New(writer, "[IFO] ", flags),
		warning: log.New(writer, "[WARNING] ", flags),
		writer:  writer,
	}
}

// Helpers para logs não formatados
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}
func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

// Helpers para logs formatadoa
func (l *Logger) DebugF(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Logger) ErrorF(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
func (l *Logger) InfoF(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}
func (l *Logger) WarningF(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}
