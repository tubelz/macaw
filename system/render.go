package system

import (
	"github.com/tubelz/macaw/entity"
	"github.com/tubelz/macaw/math"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

// RenderSystem is probably one of the most important system. It is responsible to render (draw) the entities
type RenderSystem struct {
	Entities    []entity.Entitier
	Renderer    *sdl.Renderer
	BgColor     sdl.Color
	accumulator uint32 // used for interpolation
	time        uint32 // used for animation
	Name        string
}

// Init initializes the render system using the current window
func (r *RenderSystem) Init(window *sdl.Window) {
	var err error
	if r.Renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		log.Fatalf("Renderer could not be created! SDL Error: %s\n", sdl.GetError())
	} else {
		//Initialize renderer color
		r.BgColor = sdl.Color{0xFF, 0xFF, 0xFF, 0xFF}
		r.Renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
	}
}

// Assign assign entities with this system
func (r *RenderSystem) Assign(entities []entity.Entitier) {
	r.Entities = entities
}

// UpdateAccumulator update the accumulator time.
func (r *RenderSystem) UpdateAccumulator(accumulator uint32) {
	r.accumulator = accumulator
}

// UpdateTime update the time
func (r *RenderSystem) UpdateTime(time uint32) {
	r.time = time
}

// Update will draw the entities accordingly to their position.
// it can render animated sprites, fonts or geometry
func (r *RenderSystem) Update() {
	var ok bool
	var component interface{}

	r.Renderer.SetDrawColor(r.BgColor.R, r.BgColor.G, r.BgColor.B, r.BgColor.A)
	r.Renderer.Clear()

	// interpolation variable
	alpha := float32(r.accumulator) / UpdateTickLength

	for _, obj := range r.Entities {
		// Position component
		components := obj.GetComponents()
		component, ok = components["position"]
		if !ok {
			continue
		}
		position := component.(*entity.PositionComponent)

		// Do interpolation if necessary - requires physics component (physics)
		component, ok = components["physics"]
		if ok {
			physics := component.(*entity.PhysicsComponent)
			if physics.FuturePos == nil {
				position.Pos.X = 0
				position.Pos.Y = 0
			} else {
				pos1 := &sdl.Point{int32(physics.FuturePos.X), int32(physics.FuturePos.Y)}
				position.Pos = lerp(position.Pos, pos1, alpha)
			}
		}

		// Geometry component
		component, ok = components["geometry"]
		if ok {
			drawGeometry(r.Renderer, position, component)
			continue
		}

		// Render component
		component, ok = components["render"]
		if !ok {
			continue
		}
		render := component.(*entity.RenderComponent)

		// Font Component
		component, ok = components["font"]
		if ok {
			font := component.(*entity.FontComponent)
			if font.Modified {
				generateTextureFromFont(render, font)
				font.Modified = false
			}
		}

		// Check for animation component
		component, ok = components["animation"]
		if ok {
			animation := component.(*entity.AnimationComponent)
			render.Crop = nextAnimation(r.time, animation, render.Crop)
		}

		// Draw
		r.Renderer.Copy(render.Texture, render.Crop, &sdl.Rect{position.Pos.X, position.Pos.Y, render.Crop.W, render.Crop.H})
	}
	r.Renderer.Present()
}

// generateTextureFromFont generate Texture from Font component
func generateTextureFromFont(render *entity.RenderComponent, font *entity.FontComponent) {
	var newTexture *sdl.Texture
	var solid *sdl.Surface
	var color sdl.Color
	var err error
	// Error checking
	if render == nil {
		log.Fatal("Error: Render cannot be null")
	}
	if font == nil {
		log.Fatal("Error: Font cannot be null")
	}
	// Get color. If color is not set, make it black
	if font.Color == nil {
		color = sdl.Color{0, 0, 0, 255}
	} else {
		color = *font.Color
	}
	//Load image at specified path
	if solid, err = font.Font.RenderUTF8Solid(font.Text, color); err != nil {
		log.Fatalf("Failed to render text: %s\n", err)
	}
	defer solid.Free()
	//Create texture from surface pixels
	newTexture, err = render.Renderer.CreateTextureFromSurface(solid)
	if err != nil {
		log.Fatalf("Unable to create texture from %s! SDL Error: %s\n", font.Text, sdl.GetError())
	}
	render.Texture = newTexture
	render.Crop = &sdl.Rect{0, 0, solid.W, solid.H}
}

// drawGeometry draws on the renderer the geometry. We don't use texture, because it's faster to draw directly using the renderer
func drawGeometry(render *sdl.Renderer, pos *entity.PositionComponent, geometryComponent interface{}) {
	switch g := geometryComponent.(type) {
	case *entity.RectangleComponent:
		render.SetDrawColor(g.Color.R, g.Color.G, g.Color.B, g.Color.A)
		x := pos.Pos.X
		y := pos.Pos.Y
		w := g.Size.X
		h := g.Size.Y
		// log.Printf("W: %v \\ H: %v", x, y)
		rect := &sdl.Rect{x, y, w, h}
		if g.Filled {
			render.FillRect(rect)
		} else {
			render.DrawRect(rect)
		}
	default:
		log.Fatal("Geometry component not implemented in render function")
	}
}

// lerp is the linear interpolation. pos0 is the old position, pos1 is the new position,
// alpha is the coeficient of the linear interpolation
func lerp(pos0, pos1 *sdl.Point, alpha float32) *sdl.Point {
	x := math.Round(float32(pos1.X)*alpha + float32(pos0.X)*(1.0-alpha))
	y := math.Round(float32(pos1.Y)*alpha + float32(pos0.Y)*(1.0-alpha))
	return &sdl.Point{x, y}
}

// nextAnimation returns the crop for the next animation
func nextAnimation(now uint32, anim *entity.AnimationComponent, currentRect *sdl.Rect) *sdl.Rect {
	dt := now - anim.PreviousTime
	// log.Printf("diff time: %v\n", anim.PreviousTime)
	animations := dt * uint32(anim.AnimationSpeed) / 1000
	if animations < 1 {
		// don't do anything. No animations to be done
		return currentRect
	}
	anim.Current += int(animations)
	anim.PreviousTime = now
	// log.Printf("Animation frame: %d\n", anim.Current)
	if lastElement := anim.Frames; anim.Current >= lastElement {
		anim.Current %= anim.Frames
	}
	xMultiplier := anim.Current % anim.RowLength
	yMultiplier := anim.Current / anim.RowLength
	x := int32(xMultiplier)*currentRect.W + anim.InitialPos.X
	y := int32(yMultiplier)*currentRect.H + anim.InitialPos.Y
	return &sdl.Rect{x, y, currentRect.W, currentRect.H}
}
