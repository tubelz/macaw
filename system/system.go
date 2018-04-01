package system

import (
	"github.com/tubelz/macaw/internal/utils"
)

const (
	// If you want your game to be faster increase this number.
	ticksPerSecond = 50
	// UpdateTickLength is the length per update. The game will be updated at a steady 50 times per second,
	UpdateTickLength = 1000 / ticksPerSecond
)

var (
	logFatal  = utils.LogFatal // replace for variable so we can change in the test
	logFatalf = utils.LogFatal // replace for variable so we can change in the test
)

// Systemer is the interface containing behaviours that every system should have
type Systemer interface {
	// Update will be run in the game loop whenever possible
	Update()
	// Init is used to do whatever the system should do before being updated
	Init()
}
