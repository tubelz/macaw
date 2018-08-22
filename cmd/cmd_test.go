package cmd

import (
	"flag"
	"testing"
)

func TestConfigParser_Debug(t *testing.T) {
	if Parser.Debug() != false {
		t.Errorf("Parser.Debug() returning false as default value")
	}

	flag.Set("debug", "true")
	if Parser.Debug() != true {
		t.Errorf("Parser.Debug() returning false. Should return true when debug flag is true")
	}
}
