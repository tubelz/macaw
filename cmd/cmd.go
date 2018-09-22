// Package cmd is used to parse the command line arguments.
package cmd

import "flag"

// ConfigParser is the struct that contains the information passed via cmd
type ConfigParser struct {
	// Debug flag
	debug *bool
}

// Debug is responsible to return whether we are in debug mode or not
func (c *ConfigParser) Debug() bool {
	return *c.debug
}

var (
	// Parser is the singleton in this package that allocates a ConfigParser struct,
	// thus it holds the information passed in the command line to Macaw
	Parser = new(ConfigParser)
)

// init is the responsible to initialize the values of Parser (ConfigParser) using flag.Parse
func init() {
	Parser.debug = flag.Bool("debug", false, "Flag to enable debug messages")
	flag.Parse()
}
