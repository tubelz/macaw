// Package utils contains some internal tools we use in macaw.
package utils

import (
	"log"
	"testing"
)

var (
	origLogFatal  = func(text ...interface{}) {}
	origLogFatalf = func(format string, args ...interface{}) {}
	// LogFatal to log and exit. Replace for variable so we can change in the test
	LogFatal = log.Fatal
	// LogFatalf to log with format and exit. Replace for variable so we can change in the test
	LogFatalf = log.Fatalf
)

// SetupLog will replace the log function for testing
func SetupLog(t *testing.T) {
	// replace log.Fatal functions
	origLogFatal = LogFatal
	LogFatal = func(args ...interface{}) {
		t.Error(args)
		t.FailNow()
	}

	origLogFatalf = LogFatalf
	LogFatalf = func(format string, args ...interface{}) {
		t.Errorf(format, args)
		t.FailNow()
	}
}

// TeardownLog will return the original log function
func TeardownLog() {
	// bring back original log.Fatal functions
	LogFatal = origLogFatal
	LogFatalf = origLogFatalf
}
