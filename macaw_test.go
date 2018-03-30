package macaw

import (
	"os/exec"
	"testing"
)

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
