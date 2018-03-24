package macaw

import "testing"
import "os/exec"

var (
	origLogFatal  = func(text ...interface{}) {}
	origLogFatalf = func(format string, args ...interface{}) {}
)

// setup will initialize some basic configuration
func setup(t *testing.T) {
	// replace log.Fatal functions
	origLogFatal = logFatal
	logFatal = func(args ...interface{}) {
		t.Error(args)
		t.FailNow()
	}

	origLogFatalf = logFatalf
	logFatalf = func(format string, args ...interface{}) {
		t.Errorf(format, args)
		t.FailNow()
	}
}

// teardown frees what setup has initialized
func teardown() {
	// bring back original log.Fatal functions
	logFatal = origLogFatal
	logFatalf = origLogFatalf
}

func TestGoFmt(t *testing.T) {
	cmd := exec.Command("gofmt", "-l", ".")

	if out, err := cmd.Output(); err != nil {
		if len(out) > 0 {
			t.Errorf("Exit error: %v", err)
		}
	} else {
		if len(out) > 0 {
			t.Error("You need to run go fmt")
		}

	}
}
