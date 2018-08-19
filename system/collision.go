package system

import (
	"github.com/tubelz/macaw/entity"
	"github.com/tubelz/macaw/math"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

// CollisionSystem is the system responsible to handle collisions
type CollisionSystem struct {
	EntityManager *entity.Manager
	Name          string
	Subject
}

// Init initializes this system. So far it does nothing.
func (c *CollisionSystem) Init() {}

// Update check for collision and notify observers
func (c *CollisionSystem) Update() {
	var component interface{}

	requiredComponents := []entity.Component{&entity.PositionComponent{},
		&entity.CollisionComponent{}}
	it := c.EntityManager.IterFilter(requiredComponents)
	for obj, itok := it(); itok; obj, itok = it() {
		component = obj.GetComponent(&entity.PositionComponent{})
		position := component.(*entity.PositionComponent)
		component := obj.GetComponent(&entity.CollisionComponent{})
		collision := component.(*entity.CollisionComponent)
		// check collision with border
		c.checkBorderCollision(obj, position, collision)

		// check collision with other entities
		it2 := c.EntityManager.IterFilter(requiredComponents)
		for obj2, itok2 := it2(); itok2; obj2, itok2 = it2() {
			if obj == obj2 {
				continue
			}
			component = obj2.GetComponent(&entity.PositionComponent{})
			position2 := component.(*entity.PositionComponent)
			component = obj2.GetComponent(&entity.CollisionComponent{})
			collision2 := component.(*entity.CollisionComponent)

			if c.checkCollisionBetweenAreas(position, collision, position2, collision2) {
				c.NotifyEvent(&CollisionEvent{Ent: obj, With: obj2})
			}
		}
	}
}

func (c *CollisionSystem) checkCollisionBetweenAreas(pos1 *entity.PositionComponent,
	col1 *entity.CollisionComponent,
	pos2 *entity.PositionComponent,
	col2 *entity.CollisionComponent) bool {
	var rect1, rect2 *sdl.Rect
	for _, area1 := range col1.CollisionAreas {
		rect1 = &sdl.Rect{pos1.Pos.X + area1.X, pos1.Pos.Y + area1.Y, area1.W, area1.H}
		for _, area2 := range col2.CollisionAreas {
			rect2 = &sdl.Rect{pos2.Pos.X + area2.X, pos2.Pos.Y + area2.Y, area2.W, area2.H}
			if rect1.HasIntersection(rect2) {
				return true
			}
		}
	}

	return false
}

// BorderEvent has the entity (Ent) that transpassed the border and which border
type BorderEvent struct {
	Ent  *entity.Entity
	Side string
}

// Name returns the border event name
func (b *BorderEvent) Name() string {
	return "border event"
}

func (c *CollisionSystem) checkBorderCollision(obj *entity.Entity,
	position *entity.PositionComponent,
	collision *entity.CollisionComponent) {
	for _, area := range collision.CollisionAreas {
		// check each side. top and left don't require collision size
		if position.Pos.X+area.W > 799 {
			c.NotifyEvent(&BorderEvent{Ent: obj, Side: "right"})
		} else if position.Pos.X < 1 {
			c.NotifyEvent(&BorderEvent{Ent: obj, Side: "left"})
		}

		if position.Pos.Y < 1 {
			c.NotifyEvent(&BorderEvent{Ent: obj, Side: "top"})
		} else if position.Pos.Y+area.H > 599 {
			c.NotifyEvent(&BorderEvent{Ent: obj, Side: "bottom"})
		}
	}
}

// CollisionEvent has the entity (Ent) that produced the collision and the entity that got collided (With)
type CollisionEvent struct {
	Ent  *entity.Entity
	With *entity.Entity
}

// Name returns the collision event name
func (c *CollisionEvent) Name() string {
	return "collision event"
}

/*
	----
	Util functions for handling collision events
	----
*/

// InvertVel invert the vel of the collided object.
func InvertVel(event Event) {
	collision := event.(*CollisionEvent)
	log.Printf("Inverting pos and mov of obj %d", collision.Ent.GetID())

	component := collision.Ent.GetComponent(&entity.PositionComponent{})
	position := component.(*entity.PositionComponent)

	component = collision.Ent.GetComponent(&entity.PhysicsComponent{})
	physics := component.(*entity.PhysicsComponent)

	intersectRect := intersection(collision.Ent, collision.With)
	displacementPos := &sdl.Point{intersectRect.W, intersectRect.H}

	// TODO: Clean this a little bit...
	if displacementPos.X < displacementPos.Y {
		physics.Vel.X *= -1
		physics.Acc.X *= -1
		if physics.Vel.X > 0 {
			position.Pos.X = position.Pos.X + displacementPos.X
		} else if physics.Vel.X < 0 {
			position.Pos.X = position.Pos.X - displacementPos.X
		}
		physics.FuturePos.X = float32(position.Pos.X) + physics.Vel.X
	} else if displacementPos.Y < displacementPos.X {
		physics.Vel.Y *= -1
		physics.Acc.Y *= -1
		if physics.Vel.Y > 0 {
			position.Pos.Y = position.Pos.Y + displacementPos.Y
		} else if physics.Vel.Y < 0 {
			position.Pos.Y = position.Pos.Y - displacementPos.Y
		}
		physics.FuturePos.Y = float32(position.Pos.Y) + physics.Vel.Y
	} else {
		physics.Vel = math.MulFPointWithFloat(physics.Vel, -1)
		physics.Acc = math.MulFPointWithFloat(physics.Acc, -1)
		if physics.Vel.X > 0 {
			position.Pos.X = position.Pos.X + displacementPos.X
		} else if physics.Vel.X < 0 {
			position.Pos.X = position.Pos.X - displacementPos.X
		}
		if physics.Vel.Y > 0 {
			position.Pos.Y = position.Pos.Y + displacementPos.Y
		} else if physics.Vel.Y < 0 {
			position.Pos.Y = position.Pos.Y - displacementPos.Y
		}
		physics.FuturePos = math.ConvertPointToFPoint(math.SumPointWithFPoint(position.Pos, physics.Vel))
	}
}

// intersection get the intersection rectangle between two objects
func intersection(obj1, obj2 *entity.Entity) sdl.Rect {
	posComp := &entity.PositionComponent{}
	position1 := obj1.GetComponent(posComp).(*entity.PositionComponent)
	position2 := obj2.GetComponent(posComp).(*entity.PositionComponent)

	colComp := &entity.CollisionComponent{}
	collision1 := obj1.GetComponent(colComp).(*entity.CollisionComponent)
	collision2 := obj2.GetComponent(colComp).(*entity.CollisionComponent)

	for _, area1 := range collision1.CollisionAreas {
		rect1 := &sdl.Rect{position1.Pos.X + area1.X, position1.Pos.Y + area1.Y, area1.W, area1.H}
		for _, area2 := range collision2.CollisionAreas {
			rect2 := &sdl.Rect{position2.Pos.X + area2.X, position2.Pos.Y + area2.Y, area2.W, area2.H}
			if displacement, ok := rect1.Intersect(rect2); ok {
				return displacement
			}
		}
	}

	return sdl.Rect{0, 0, 0, 0}
}
