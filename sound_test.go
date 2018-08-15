package macaw

import (
	"testing"
)

func TestPlaySound_noFile(t *testing.T) {
	err := PlaySound("nofilehere.mp3")
	if err == nil {
		t.Errorf("Error expected. None found")
	}
}
