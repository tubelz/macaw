package entity

import (
	"testing"
)

func TestManager_CreateIncreaseID(t *testing.T) {
	m := &Manager{}
	entity1 := m.Create()
	entity2 := m.Create()

	if entity2.GetID() < entity1.GetID() {
		t.Errorf("Auto increment of entity manager not working. ID_1: %d, ID_2: %d", entity1.GetID(), entity2.GetID())
	}
}

func TestManager_Get(t *testing.T) {
	m := &Manager{}
	_ = m.Create()
	_ = m.Create()

	if m.Get(1).GetID() < m.Get(0).GetID() {
		t.Error("Get() getting wrong entity")
	}
	if m.Get(2) != nil {
		t.Error("Get(n) should return nil if N is greater than len(entities)")
	}
}

func TestManager_IterAvailable(t *testing.T) {
	m := &Manager{}
	_ = m.Create()
	_ = m.Create()
	_ = m.Create()
	// Delete object with ID 1 so we expect to return objects with ID 0 and 2
	m.Delete(1)
	it := m.IterAvailable()
	if obj, ok := it(); !ok {
		t.Error("Expecting true, returned false")
	} else if obj.GetID() != 0 {
		t.Error("First object in test should have id 0")
	}

	if obj, ok := it(); !ok {
		t.Error("Expecting true, returned false")
	} else if obj.GetID() != 2 {
		t.Error("Second object in test should have id 2")
	}

	if _, ok := it(); ok {
		t.Error("Expecting iterator to end")
	}
}

func TestManager_Delete(t *testing.T) {
	m := &Manager{}
	_ = m.Create()
	_ = m.Create()

	if ok := m.Delete(1); !ok {
		t.Error("Delete() not removing element")
	}
	if len(m.availableSlots) != 1 {
		t.Fatal("Delete() not creating queue")
		if m.availableSlots[0] != 1 {
			t.Error("Delete() not adding correct index to the queue")
		}
	}
	if m.Get(1) != nil {
		t.Errorf("Delete() not working. Should return nil. Returning %v", m.Get(1))
	}
	m.Delete(1)
	if ok := m.Delete(1); ok {
		t.Error("Delete() removing element that doesn't exist")
	}
}

func TestManager_ReuseSlot(t *testing.T) {
	m := &Manager{}
	_ = m.Create()
	entity2 := m.Create()
	m.Delete(0)
	entity1 := m.Create()

	if entity2.GetID() < entity1.GetID() {
		t.Errorf("Reuse slot not working. ID_1: %d, ID_2: %d", entity1.GetID(), entity2.GetID())
	}
}
