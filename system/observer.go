package system

/*
 This is the implementation of the Observer Pattern so we can
 track events and execute functions we are interested on for
 specific events
*/

// EventHandler is the interface function which will handle the event
// Think as the Notify() func from the Observer
type EventHandler func(event Event)

// Event adds data to the struct that implements event so you can communicate with the system
type Event interface {
	Name() string
}

// Subject holds the listerners and their handlers. An event can have multiple handlers.
type Subject struct {
	listeners map[string][]EventHandler
}

// AddHandler adds an event handler
func (s *Subject) AddHandler(eventName string, handler EventHandler) {
	if s.listeners == nil {
		s.listeners = make(map[string][]EventHandler)
	}
	s.listeners[eventName] = append(s.listeners[eventName], handler)
}

// RemHandler removes an event handler
func (s *Subject) RemHandler() {
	//TODO
}

// ClearEvents clear all events
func (s *Subject) ClearEvents() {
	s.listeners = nil
}

// NotifyEvent executes all event handlers for a specific event
func (s *Subject) NotifyEvent(event Event) {
	handlers := s.listeners[event.Name()]
	for _, handler := range handlers {
		handler(event)
	}
}
