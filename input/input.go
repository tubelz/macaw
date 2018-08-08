// Package input provides a wrapper to handle user input events (keyboard and mouse for now).
package input

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

// Manager hold data of events
type Manager struct {
	// TODO: 1) Implement Command Pattern?
	button []sdl.KeyboardEvent
	Mouse  MouseEvent
}

// MouseEvent has the position and button pressed by the mouse
type MouseEvent struct {
	Pos    *sdl.Point
	Button uint8 // BUTTON_LEFT, BUTTON_MIDDLE, BUTTON_RIGHT, BUTTON_X1, BUTTON_X2
	State  uint8 // PRESSED, RELEASE
}

// HandleEvents handle the events such as key pressed and mouse movements.
func (i *Manager) HandleEvents() bool {
	running := true
	var event sdl.Event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

		switch t := event.(type) {
		case *sdl.QuitEvent:
			running = false
		case *sdl.KeyboardEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				running = false
			}
			i.button = append(i.button, *t)
			// keyboard map: https://github.com/veandco/go-sdl2/blob/master/sdl/keycode.go#L11
			// log.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c (%d)\tmodifiers:%d\tstate:%d\trepeat:%d\n",
			// 						t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)

		case *sdl.MouseMotionEvent:
			pos := &sdl.Point{t.X, t.Y}
			i.Mouse.Pos = pos
		case *sdl.MouseButtonEvent:
			i.Mouse.Button = t.Button
			i.Mouse.State = t.State
		default:
			log.Printf("Event not mapped")
		}
	}
	return running
}

// PopEvent removes the first element added.
// Usually pop returns the element popped, but this one doesn't. I couldn't come up with a good name.
func (i *Manager) PopEvent() {
	s := i.button
	if len(s) > 0 {
		copy(s[0:], s[1:])
		i.button = s[:len(s)-1]
	}
}

// ClearMouseEvent sets the id of the button to 0 which no mouse button is assigned to
func (m *MouseEvent) ClearMouseEvent() {
	m.Button = 0
}

// Button returns the first button pressed. Usefull to use in multiple systems
func (i *Manager) Button() sdl.KeyboardEvent {
	if len(i.button) > 0 {
		return i.button[0]
	}
	return sdl.KeyboardEvent{}
}
