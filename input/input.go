package input

import (
	"log"
	"github.com/veandco/go-sdl2/sdl"
)

// Manager hold data of events
type Manager struct {
	// TODO: 1) Implement Command Pattern?
	//       2) Implement Mouse Handler
	button []*sdl.KeyboardEvent
	// mouseHandler
}

// HandleEvents handle the events such as key pressed and mouse movements.
func (i *Manager) HandleEvents(running bool) bool {
	var event sdl.Event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

		switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					running = false
				}
				i.button = append(i.button, t)
				// keyboard map: https://github.com/veandco/go-sdl2/blob/master/sdl/keycode.go#L11
				// log.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c (%d)\tmodifiers:%d\tstate:%d\trepeat:%d\n",
				// 						t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)

			case *sdl.MouseMotionEvent:
				//TODO
			case *sdl.MouseButtonEvent:
				//TODO
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
		s[len(s)-1] = nil
		i.button = s[:len(s)-1]
	}
}

// Button returns the first button pressed. Usefull to use in multiple systems
func (i *Manager) Button() *sdl.KeyboardEvent {
	if len(i.button) > 0 {
		return i.button[0]
	}
	return nil
}
