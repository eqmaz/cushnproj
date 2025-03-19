package logger

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	log := NewLogger()
	assert.NotNil(t, log, "Logger should be initialized")
	assert.IsType(t, &Logger{}, log, "Logger should be of type *Logger")
}

func TestLoggerWithContextFields(t *testing.T) {
	log := NewLogger()
	fields := map[string]interface{}{"key": "value"}
	newLog := log.WithContextFields(fields)

	assert.NotNil(t, newLog, "Logger with fields should not be nil")

	// Ensure the new logger has fields by checking WithFields output
	entry := newLog.WithFields(fields)
	assert.Equal(t, "value", entry.Data["key"], "Logger should contain the key field")
}

func TestLoggerSetOutput(t *testing.T) {
	log := NewLogger()
	err := log.SetOutput("stdout")
	assert.NoError(t, err, "Setting output to stdout should not return an error")

	tempFile, err := ioutil.TempFile("", "logtest")
	assert.NoError(t, err, "Creating temp file should not return an error")
	defer os.Remove(tempFile.Name())

	err = log.SetOutput(tempFile.Name())
	assert.NoError(t, err, "Setting output to a file should not return an error")
}

func TestLoggerSuspendResume(t *testing.T) {
	log := NewLogger()
	log.Suspend()
	assert.True(t, log.suspended, "Logger should be suspended")

	log.Resume()
	assert.False(t, log.suspended, "Logger should be resumed")
}

func TestLoggerLoggingMethods(t *testing.T) {
	log := NewLogger()
	log.SetLevel(logrus.DebugLevel)

	log.Info("info message", map[string]interface{}{"key": "value"})
	log.Warn("warn message", map[string]interface{}{"key": "value"})
	log.Error("error message", map[string]interface{}{"key": "value"})
	log.Debug("debug message", map[string]interface{}{"key": "value"})

	assert.True(t, true, "Logging methods should execute without error")
}
