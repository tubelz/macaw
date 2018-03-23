package entity

import (
	"github.com/tubelz/macaw/math"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// entityCounter is responsible for count the amount of entities
var entityCounter uint16

// Component is the abstract type for each component
type Component interface{}

// Entitier has all the behaviours entities should have
type Entitier interface {
	GetID() uint16
	Init()
	GetComponents() map[string]Component
	AddComponent(string, Component)
	DelComponent(key string)
	GetComponent(componentName string) (Component, bool)
}

// Entity is the struct that contains the components. Right now the id's are not being used
type Entity struct {
	// components
	id         uint16
	components map[string]Component
}

// GetID the id of the entity
func (e *Entity) GetID() uint16 {
	return e.id
}

// Init adds the id and increments the counter.
func (e *Entity) Init() {
	e.id = entityCounter
	entityCounter++
	if e.components == nil {
		e.components = make(map[string]Component)
	}
}

// GetComponents returns a list of all the components of the entity
func (e *Entity) GetComponents() map[string]Component {
	return e.components
}

// GetComponent returns the given component
func (e *Entity) GetComponent(componentName string) (Component, bool) {
	val, ok := e.components[componentName]
	return val, ok
}

// AddComponent adds a component to the component map
func (e *Entity) AddComponent(name string, c Component) {
	e.components[name] = c
}

// DelComponent removes the given component
func (e *Entity) DelComponent(key string) {
	delete(e.components, key)
}

/////////////////////////////////////////////////
/// ... Basic Components ...
/////////////////////////////////////////////////

//PositionComponent is responsible for the position of the entity
type PositionComponent struct {
	Pos *sdl.Point
}

// PhysicsComponent is responsible for some of the physics
type PhysicsComponent struct {
	FuturePos *math.FPoint // TODO: move this to PositionComponent
	// velocity
	Vel *math.FPoint
	// acceleration
	Acc *math.FPoint
}

// RenderComponent is responsible for the rendering of the entity
type RenderComponent struct {
	Renderer *sdl.Renderer
	Texture  *sdl.Texture
	Crop     *sdl.Rect // part of the texture which will be displayed
	Scroll   bool
	Angle    float64
	Center   *sdl.Point
	Flip     sdl.RendererFlip
}

// CameraComponent is responsible to render only the content of the viewport
type CameraComponent struct {
	ViewportSize sdl.Point
	WorldSize    sdl.Point
	IsActive     bool
}

// AnimationComponent is responsible for animate the entity
type AnimationComponent struct {
	InitialPos     sdl.Point // frame reference
	AnimationSpeed uint8     //animations per second
	PreviousTime   uint32    //last animation time
	Current        int
	Frames         int // total sprites
	RowLength      int // number of sprites per row
	SpriteMap      map[string]int
}

// FontComponent holds the font and text information
type FontComponent struct {
	Font     *ttf.Font
	Text     string
	Modified bool
	Color    *sdl.Color
}

// RectangleComponent has the information to draw a rectangle
type RectangleComponent struct {
	Size   *sdl.Point
	Color  *sdl.Color
	Filled bool
}

// CollisionComponent makes the entity notify if it hits something else
// TODO: Add other type of information such as Shape, Density, Friction etc...
type CollisionComponent struct {
	// Size is duplicating data a little bit... we have this information
	// in the render and geometry component, but we will use this attribute for now
	Size *sdl.Point
}
