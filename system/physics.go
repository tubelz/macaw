package system

import (
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
	var component interface{}
	phyComp := &entity.PhysicsComponent{}

	requiredComponents := []entity.Component{phyComp}
	it := p.EntityManager.IterFilter(requiredComponents, -1)
	for obj, i := it(); i != -1; obj, i = it() {
		component = obj.GetComponent(phyComp)
		physics := component.(*entity.PhysicsComponent)

		// To use Semi-implicit Euler, we first update the velocity, then we update the position.
		// FuturePos is used so we can interpolate with current position
		physics.Vel = math.SumFPoint(physics.Vel, physics.Acc)
		physics.FuturePos = math.SumFPoint(physics.FuturePos, physics.Vel)
	}
}
