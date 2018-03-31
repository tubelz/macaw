package entity

import (
	"github.com/tubelz/macaw/internal/utils"
	"testing"
)

func TestManager_CreateIncreaseID(t *testing.T) {
	utils.SetupLog(t)

	m := &Manager{}
	entity1 := m.Create()
	entity2 := m.Create()

	if entity2.GetID() < entity1.GetID() {
		t.Errorf("Auto increment of entity manager not working. ID_1: %d, ID_2: %d", entity1.GetID(), entity2.GetID())
	}

	utils.TeardownLog()
}

func TestManager_Get(t *testing.T) {
	utils.SetupLog(t)

	m := &Manager{}
	_ = m.Create()
	_ = m.Create()

	if m.Get(1).GetID() < m.Get(0).GetID() {
		t.Error("Get() getting wrong entity")
	}
	if m.Get(2) != nil {
		t.Error("Get(n) should return nil if N is greater than len(entities)")
	}

	utils.TeardownLog()
}

func TestManager_Delete(t *testing.T) {
	utils.SetupLog(t)

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

	utils.TeardownLog()
}

func TestManager_ReuseSlot(t *testing.T) {
	utils.SetupLog(t)

	m := &Manager{}
	_ = m.Create()
	entity2 := m.Create()
	m.Delete(0)
	entity1 := m.Create()

	if entity2.GetID() < entity1.GetID() {
		t.Errorf("Reuse slot not working. ID_1: %d, ID_2: %d", entity1.GetID(), entity2.GetID())
	}

	utils.TeardownLog()
}
