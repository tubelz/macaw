package entity

import (
	"testing"
)

func TestManager_CreateIncreaseID(t *testing.T) {
	m := &Manager{}
	entity1 := m.Create("")
	entity2 := m.Create("")

	if entity2.GetID() < entity1.GetID() {
		t.Errorf("Auto increment of entity manager not working. ID_1: %d, ID_2: %d", entity1.GetID(), entity2.GetID())
	}
}

func TestManager_GetType(t *testing.T) {
	m := &Manager{}
	entity1 := m.Create("type1")
	entity2 := m.Create("type2")

	if entity1.GetType() != "type1" {
		t.Errorf("GetType() returning %s. Should return type1", entity1.GetType())
	}
	if entity2.GetType() != "type2" {
		t.Errorf("GetType() returning %s. Should return type2", entity2.GetType())
	}
}

func TestManager_Get(t *testing.T) {
	m := &Manager{}
	_ = m.Create("")
	_ = m.Create("")

	if m.Get(1).GetID() < m.Get(0).GetID() {
		t.Error("Get() getting wrong entity")
	}
	if m.Get(2) != nil {
		t.Error("Get(n) should return nil if N is greater than len(entities)")
	}
}

func TestManager_IterAvailable(t *testing.T) {
	m := &Manager{}
	_ = m.Create("")
	_ = m.Create("")
	_ = m.Create("")
	// Delete object with ID 1 so we expect to return objects with ID 0 and 2
	m.Delete(1)
	it := m.IterAvailable(-1)
	if obj, i := it(); i == -1 {
		t.Error("Expecting true, returned false")
	} else if obj.GetID() != 0 {
		t.Error("First object in test should have id 0")
	}

	if obj, i := it(); i == -1 {
		t.Error("Expecting true, returned false")
	} else if obj.GetID() != 2 {
		t.Error("Second object in test should have id 2")
	}

	if _, i := it(); i != -1 {
		t.Error("Expecting iterator to end")
	}
}

func TestManager_IterAvailable_noItem(t *testing.T) {
	m := &Manager{}

	it := m.IterAvailable(-1)
	if _, i := it(); i != -1 {
		t.Errorf("IterAvailable should not return true for empty items")
	}
}

func TestManager_IterAvailable_twoIterators(t *testing.T) {
	m := &Manager{}
	_ = m.Create("")
	_ = m.Create("")
	_ = m.Create("")
	// Delete object with ID 1 so we expect to return objects with ID 0 and 2
	m.Delete(1)
	it1 := m.IterAvailable(-1)
	it2 := m.IterAvailable(-1)

	obj1, _ := it1()
	obj2, _ := it2()
	if obj1.GetID() != obj2.GetID() {
		t.Errorf(`IterAvailable getting different objects when it should get the
			same object Expecting 0 == 0 got %d == %d`, obj1.GetID(), obj2.GetID())
	}
	obj1, _ = it1()
	obj2, _ = it2()
	if obj1.GetID() != obj2.GetID() {
		t.Errorf(`IterAvailable getting different objects when it should get the
			same object. Expecting 0 == 0 got %d == %d`, obj1.GetID(), obj2.GetID())
	}
}

type TestComponent struct {
	fakeData string
}

func TestManager_IterFilter_simple(t *testing.T) {
	m := &Manager{}
	_ = m.Create("e1")
	e2 := m.Create("e2")
	_ = m.Create("e3")
	e4 := m.Create("e4")
	// Create components
	c1 := &TestComponent{fakeData: "c1"}
	c2 := &TestComponent{fakeData: "c2"}
	// Add components
	e2.AddComponent(c1)
	e4.AddComponent(c2)
	// Create list of components we are interested in
	listComponents := []Component{&TestComponent{}}
	// Verify that we filtered them correctly
	it := m.IterFilter(listComponents, -1)
	count := 0
	for entityObj, i := it(); i != -1; entityObj, i = it() {
		count++
		if entityObj.GetID() == 0 || entityObj.GetID() == 2 {
			t.Error("IterFilter([]Component) not filtering correctly.")
		}
	}
	if count < 2 {
		t.Error(`IterFilter([]Component) not filtering correctly.
			Filtering more than it should`)
	}
}

type TestComponent2 struct{}

func TestManager_IterFilter_multiple(t *testing.T) {
	m := &Manager{}
	_ = m.Create("e1")
	e2 := m.Create("e2")
	_ = m.Create("e3")
	e4 := m.Create("e4")
	e5 := m.Create("e5")

	// Create components
	c1 := &TestComponent{fakeData: "c1"}
	c2 := &TestComponent2{}
	// Add components
	e2.AddComponent(c1)
	e2.AddComponent(c2)
	e4.AddComponent(c1)
	e5.AddComponent(c2)
	// Create list of components we are interested in
	listComponents := []Component{&TestComponent{}}
	// Verify that we filtered them correctly
	it := m.IterFilter(listComponents, -1)
	count := 0
	for entityObj, i := it(); i != -1; entityObj, i = it() {
		count++
		if entityObj.GetID() != 1 && entityObj.GetID() != 3 {
			t.Error("IterFilter([]Component) not filtering correctly.")
		}
	}
	if count < 2 {
		t.Error(`IterFilter([]Component) not filtering correctly.
			Filtering more than it should`)
	}

	count = 0
	listComponents2 := []Component{&TestComponent{}, &TestComponent2{}}
	it2 := m.IterFilter(listComponents2, -1)
	for entityObj, i := it2(); i != -1; entityObj, i = it2() {
		count++
		if entityObj.GetID() != 1 {
			t.Error("IterFilter([]Component) not filtering correctly.")
		}
	}
	if count != 1 {
		t.Error(`IterFilter([]Component) not filtering correctly.
			Filtering more than it should`)
	}
}

func TestManager_Delete(t *testing.T) {
	m := &Manager{}
	_ = m.Create("")
	_ = m.Create("")

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
	_ = m.Create("")        // 0
	entity1 := m.Create("") // 1
	m.Delete(0)
	entity0 := m.Create("") // 0

	// simple check
	if entity1.GetID() < entity0.GetID() {
		t.Errorf("Reuse slot not working. ID_1: %d, ID_0: %d", entity1.GetID(), entity0.GetID())
	}
	// check if it the sequencing keeps working
	entity2 := m.Create("")
	if entity2.GetID() != 2 {
		t.Errorf("Not using right slot. ID_2: %d; want 2", entity2.GetID())
	}

	// check if the reuse works on the middle as well
	for i := uint16(0); i < 3; i++ {
		m.Delete(i)
	}
	_ = m.Create("")
	m.Delete(0)
	entity0 = m.Create("")
	if entity0.GetID() != 0 {
		t.Errorf("Reuse slot not working. ID_0: %d; want 0", entity0.GetID())
	}
}

func TestBinarySearchInsert(t *testing.T) {
	size := func(arr []uint16) int {
		return len(arr) - 1
	}
	cases := []struct {
		inArray []uint16
		inVal   uint16
		want    int
	}{
		{[]uint16{0}, 1, 1},
		{[]uint16{1}, 0, 0},
		{[]uint16{1}, 1, 0},
		{[]uint16{1, 5, 6, 7, 8, 9}, 10, 6},
		{[]uint16{1, 5, 6, 7, 8, 9}, 2, 1},
		{[]uint16{1, 5, 6, 7, 8, 9}, 4, 1},
		{[]uint16{1, 5, 6, 7, 8, 9}, 0, 0},
	}
	for _, c := range cases {
		got := binarySearch(c.inArray, 0, size(c.inArray), c.inVal)
		if got != c.want {
			t.Errorf("binarySearchInsert(%v, 0, %d, %d) == %d; want %d", c.inArray, size(c.inArray), c.inVal, got, c.want)
		}
	}
}

func BenchmarkBinSearchInsert(b *testing.B) {
	// run the Fib function b.N times
	arr := []uint16{1, 5, 6, 7, 8, 9}
	size := len(arr) - 1
	val := uint16(10)
	for n := 0; n < b.N; n++ {
		binarySearch(arr, 0, size, val)
	}
}

func BenchmarkEntityAddComponentSimple(b *testing.B) {
	entityManager := &Manager{}
	entityTest := entityManager.Create("entity")
	for n := 0; n < b.N; n++ {
		entityTest.AddComponent(&PositionComponent{})
	}
}

func BenchmarkEntityAddComponentMultiple(b *testing.B) {
	entityManager := &Manager{}
	entityTest := entityManager.Create("entity")
	for n := 0; n < b.N; n++ {
		entityTest.AddComponent(&PositionComponent{})
		entityTest.AddComponent(&RenderComponent{})
		entityTest.AddComponent(&AnimationComponent{})
		entityTest.AddComponent(&FontComponent{})
	}
}

func BenchmarkEntityReadComponentSimple(b *testing.B) {
	entityManager := &Manager{}

	posComponent := &PositionComponent{}
	e1 := entityManager.Create("entity")

	e1.AddComponent(&posComponent)

	for n := 0; n < b.N; n++ {
		_ = e1.GetComponent(posComponent)
	}
}

func BenchmarkEntityReadComponentMultiple(b *testing.B) {
	entityManager := &Manager{}

	posComponent := &PositionComponent{}
	e1 := entityManager.Create("entity")

	e1.AddComponent(&AnimationComponent{})
	e1.AddComponent(&PhysicsComponent{})
	e1.AddComponent(&RenderComponent{})
	e1.AddComponent(&CollisionComponent{})
	e1.AddComponent(&posComponent)

	for n := 0; n < b.N; n++ {
		_ = e1.GetComponent(posComponent)
	}
}
