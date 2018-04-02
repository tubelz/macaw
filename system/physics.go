package system

import (
	// "log"
	"github.com/tubelz/macaw/entity"
	"github.com/tubelz/macaw/math"
)

// PhysicsSystem is responsible to update the physics in the game.
type PhysicsSystem struct {
	EntityManager *entity.Manager
	Name          string
	Subject
}

// Init initializes this system. So far it does nothing.
func (p *PhysicsSystem) Init() {}

// Update change the position and velocity accordingly. We are using Semi-implicit Euler
func (p *PhysicsSystem) Update() {
	var ok bool
	var component interface{}

	it := p.EntityManager.IterAvailable()
	for obj, itok := it(); itok; obj, itok = it() {
		components := obj.GetComponents()
		component, ok = components["physics"]
		if !ok {
			continue
		}
		physics := component.(*entity.PhysicsComponent)

		// To use Semi-implicit Euler, we first update the velocity, then we update the position.
		// FuturePos is used so we can interpolate with current position
		physics.Vel = math.SumFPoint(physics.Vel, physics.Acc)
		physics.FuturePos = math.SumFPoint(physics.FuturePos, physics.Vel)
	}
}
