package utils

// MockSystem is a system used for test
type MockSystem struct {
	MockFunc func()
}

// Init mock
func (m *MockSystem) Init() {
}

// Update executes a function if it was declared
func (m *MockSystem) Update() {
	if m.MockFunc != nil {
		m.MockFunc()
	}
}
