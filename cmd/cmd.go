// Package cmd is used to parse the command line arguments.
package cmd

import "flag"

type ConfigParser struct {
	// Debug flag
	debug *bool
}

// Debug is responsible to return whether we are in debug mode or not
func (c *ConfigParser) Debug() bool {
	return *c.debug
}

var (
	Parser = new(ConfigParser)
)

func init() {
	Parser.debug = flag.Bool("debug", false, "Flag to enable debug messages")
	flag.Parse()
}
