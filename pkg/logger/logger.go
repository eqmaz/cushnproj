package logger

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Logger is a wrapper around logrus.Logger
type Logger struct {
	suspended bool
	*logrus.Logger
}

// NewLogger creates a new logger with the default settings
func NewLogger() *Logger {
	log := logrus.New()
	log.Out = os.Stdout

	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
	})

	return &Logger{false, log}
}

// WithContextFields returns a new logger with the given fields added to the context
func (l *Logger) WithContextFields(fields map[string]interface{}) *Logger {
	return &Logger{false, l.WithFields(fields).Logger}
}

// SetOutput sets the output destination for the logger
// destination can be "stdout", "stderr" or a writable file path
func (l *Logger) SetOutput(destination string) error {
	var writer io.Writer

	switch destination {
	case "stdout":
		writer = os.Stdout
	case "stderr":
		writer = os.Stderr
	default:
		file, err := os.OpenFile(destination, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		writer = file
	}

	//l.Logger.Out = writer
	l.Logger.SetOutput(writer)
	return nil
}

// CloseOutput safely closes the output destination for the logger
// You MUST call this function manually if the logger's lifetime is not the same as the application's
func (l *Logger) CloseOutput() error {
	if closer, ok := l.Logger.Out.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

// Suspend stops the logger from logging until resumed
func (l *Logger) Suspend() {
	l.suspended = true
}

// Resume resumes the logger from logging
func (l *Logger) Resume() {
	l.suspended = false
}

// Info logs a message with the Info level
func (l *Logger) Info(msg string, fields map[string]interface{}) {
	if !l.suspended {
		l.WithFields(fields).Log(logrus.InfoLevel, msg)
	}
}

// Error logs a message with the Error level
func (l *Logger) Error(msg string, fields map[string]interface{}) {
	if !l.suspended {
		l.WithFields(fields).Log(logrus.ErrorLevel, msg)
	}
}

// Warn logs a message with the Warn level
func (l *Logger) Warn(msg string, fields map[string]interface{}) {
	if !l.suspended {
		l.WithFields(fields).Log(logrus.WarnLevel, msg)
	}
}

// Debug logs a message with the Debug level
func (l *Logger) Debug(msg string, fields map[string]interface{}) {
	if !l.suspended {
		l.WithFields(fields).Log(logrus.DebugLevel, msg)
	}
}
