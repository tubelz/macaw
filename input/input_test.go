package input

import (
	"github.com/veandco/go-sdl2/sdl"
	"testing"
)

// TestManager_HandleEvents should check if we push a keyboard event it will register in our queue
func TestManager_HandleEventsBasic(t *testing.T) {
	m := &Manager{}
	buttonDownPressed := &sdl.KeyboardEvent{Type: sdl.KEYDOWN, State: sdl.PRESSED, Keysym: sdl.Keysym{Sym: sdl.K_DOWN}}

	sdl.PushEvent(buttonDownPressed)
	m.HandleEvents()
	if len(m.buttons) < 0 {
		t.Error("Expecting 1 button in event queue, got 0")
	} else if m.buttons[0] != *buttonDownPressed {
		t.Errorf("Expecting queue event to be %v, got %v", buttonDownPressed, m.buttons[0])
	}
}

// TestManager_HandleEventsMultiple should check if we push multiple events of the same type (keyboard)
// the queue of objects will be good
func TestManager_HandleEventsMultiple(t *testing.T) {
	m := &Manager{}
	buttonDownPressed := &sdl.KeyboardEvent{Type: sdl.KEYDOWN, State: sdl.PRESSED, Keysym: sdl.Keysym{Sym: sdl.K_DOWN}}
	buttonLeftPressed := &sdl.KeyboardEvent{Type: sdl.KEYDOWN, State: sdl.PRESSED, Keysym: sdl.Keysym{Sym: sdl.K_LEFT}}
	buttonDownReleased := &sdl.KeyboardEvent{Type: sdl.KEYDOWN, State: sdl.RELEASED, Keysym: sdl.Keysym{Sym: sdl.K_DOWN}}

	sdl.PushEvent(buttonDownPressed)
	sdl.PushEvent(buttonLeftPressed)
	sdl.PushEvent(buttonDownReleased)
	m.HandleEvents()
	if len(m.buttons) != 3 {
		t.Errorf("Expecting 3 buttons in event queue, got %d", len(m.buttons))
	} else if m.buttons[0] != *buttonDownPressed {
		t.Errorf("Expecting queue event to be %v, got %v", buttonDownPressed, m.buttons[0])
	} else if m.buttons[2] != *buttonDownReleased {
		t.Errorf("Expecting queue event to be %v, got %v", buttonDownReleased, m.buttons[2])
	}
}

// TestManager_PopEvent checks if we our pop works appropriately
func TestManager_PopEvent(t *testing.T) {
	buttonDownPressed := &sdl.KeyboardEvent{Type: sdl.KEYDOWN, Timestamp: 0, State: sdl.PRESSED, Keysym: sdl.Keysym{Sym: sdl.K_DOWN}}
	buttonLeftPressed := &sdl.KeyboardEvent{Type: sdl.KEYDOWN, Timestamp: 0, State: sdl.PRESSED, Keysym: sdl.Keysym{Sym: sdl.K_LEFT}}
	empty := []sdl.KeyboardEvent{}

	cases := []struct {
		in   []*sdl.KeyboardEvent
		want []sdl.KeyboardEvent
	}{
		{[]*sdl.KeyboardEvent{}, empty},
		{[]*sdl.KeyboardEvent{buttonDownPressed}, empty},
		{[]*sdl.KeyboardEvent{buttonDownPressed, buttonLeftPressed}, []sdl.KeyboardEvent{*buttonLeftPressed}},
	}

	m := &Manager{}
	for i, c := range cases {
		m.buttons = []sdl.KeyboardEvent{}
		for _, event := range c.in {
			sdl.PushEvent(event)
		}
		m.HandleEvents()
		m.PopEvent()
		got := m.buttons
		// compare slices
		if len(got) != len(c.want) {
			t.Errorf("Case %d failing. Got %v want %v", i, got, c.want)
		} else {
			for j, button := range got {
				// We have to set the timestamp to 0, otherwise if our code takes longer
				// it will throw an error
				button.Timestamp = 0
				if button != c.want[j] {
					t.Errorf("Case %d failing. Got %v want %v", i, got, c.want)
				}
			}
		}
	}
}

func TestManager_Button(t *testing.T) {
	m := &Manager{}
	// check one button first
	buttonDownPressed := &sdl.KeyboardEvent{Type: sdl.KEYDOWN, Timestamp: 0, State: sdl.PRESSED, Keysym: sdl.Keysym{Sym: sdl.K_DOWN}}
	sdl.PushEvent(buttonDownPressed)
	m.HandleEvents()
	if m.Button() != *buttonDownPressed {
		t.Errorf("Button() == %v; want %v", m.Button(), *buttonDownPressed)
	}
	// check if we have two buttons in the queue.
	// It should show the first button since we haven't cleared the queue
	buttonLeftPressed := &sdl.KeyboardEvent{Type: sdl.KEYDOWN, Timestamp: 0, State: sdl.PRESSED, Keysym: sdl.Keysym{Sym: sdl.K_LEFT}}
	sdl.PushEvent(buttonLeftPressed)
	m.HandleEvents()
	if m.Button() != *buttonDownPressed {
		t.Errorf("Button() == %v; want %v", m.Button(), *buttonDownPressed)
	}
}
