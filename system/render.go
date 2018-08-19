package system

import (
	"github.com/tubelz/macaw/entity"
	"github.com/tubelz/macaw/math"
	"github.com/veandco/go-sdl2/sdl"
)

// RenderSystem is probably one of the most important system. It is responsible to render (draw) the entities
type RenderSystem struct {
	EntityManager *entity.Manager
	Window        *sdl.Window
	Renderer      *sdl.Renderer
	BgColor       sdl.Color
	Camera        entity.Entitier // entity that have the camera component
	accumulator   uint32          // used for interpolation
	time          uint32          // used for animation
	Name          string
}

// Init initializes the render system using the current window
func (r *RenderSystem) Init() {
	var err error
	if r.Renderer, err = sdl.CreateRenderer(r.Window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		logFatalf("Renderer could not be created! SDL Error: %s\n", sdl.GetError())
	} else {
		//Initialize renderer color
		r.BgColor = sdl.Color{0xFF, 0xFF, 0xFF, 0xFF}
		r.Renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
	}
}

// UpdateAccumulator update the accumulator time.
func (r *RenderSystem) UpdateAccumulator(accumulator uint32) {
	r.accumulator = accumulator
}

// UpdateTime update the time
func (r *RenderSystem) UpdateTime(time uint32) {
	r.time = time
}

// SetCamera sets the camera which controls what will be rendered
func (r *RenderSystem) SetCamera(camera entity.Entitier) {
	r.Camera = camera
}

// GetCameraPosition gets the camera position
func (r *RenderSystem) GetCameraPosition() (int32, int32) {
	if component := r.Camera.GetComponent(&entity.PositionComponent{}); component != nil {
		position := component.(*entity.PositionComponent)
		return position.Pos.X, position.Pos.Y
	}
	return 0, 0
}

// OffsetPosition changes the cartesian position according to the camera
func (r *RenderSystem) OffsetPosition(x, y int32) (int32, int32) {
	camX, camY := r.GetCameraPosition()
	x -= camX
	y -= camY
	return x, y
}

func (r *RenderSystem) isRenderable(pos *sdl.Point, size *sdl.Rect) bool {
	if r.Camera == nil {
		return false
	}
	if component := r.Camera.GetComponent(&entity.PositionComponent{}); component != nil {
		position := component.(*entity.PositionComponent)
		c := r.Camera.GetComponent(&entity.CameraComponent{})
		camera := c.(*entity.CameraComponent)

		// check
		objRect := &sdl.Rect{pos.X, pos.Y, size.W, size.H}
		cameraRect := sdl.Rect{position.Pos.X, position.Pos.Y, camera.ViewportSize.X, camera.ViewportSize.Y}
		return cameraRect.HasIntersection(objRect)
	}
	return false
}

// Update will draw the entities accordingly to their position.
// it can render animated sprites, fonts or geometry
func (r *RenderSystem) Update() {
	var component entity.Component

	if r.Camera == nil {
		logFatal("Please, assign at least one camera to the render system")
	}

	r.Renderer.SetDrawColor(r.BgColor.R, r.BgColor.G, r.BgColor.B, r.BgColor.A)
	r.Renderer.Clear()

	// interpolation variable
	alpha := float32(r.accumulator) / UpdateTickLength

	requiredComponents := []entity.Component{&entity.PositionComponent{}}
	it := r.EntityManager.IterFilter(requiredComponents)
	for obj, itok := it(); itok; obj, itok = it() {
		// Grid component
		if component = obj.GetComponent(&entity.GridComponent{}); component != nil {
			grid := component.(*entity.GridComponent)
			r.drawGrid(grid)
			continue
		}
		// Position component
		component = obj.GetComponent(&entity.PositionComponent{})
		if component == nil {
			continue
		}
		position := component.(*entity.PositionComponent)

		// Do interpolation if necessary - requires physics component (physics)
		if component = obj.GetComponent(&entity.PhysicsComponent{}); component != nil {
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
		if component = obj.GetComponent(&entity.RectangleComponent{}); component != nil {
			r.drawGeometry(position, component)
			continue
		}

		// Render component
		component = obj.GetComponent(&entity.RenderComponent{})
		if component == nil {
			continue
		}
		render := component.(*entity.RenderComponent)

		// Font Component
		component = obj.GetComponent(&entity.FontComponent{})
		if component != nil {
			font := component.(*entity.FontComponent)
			if font.Modified {
				generateTextureFromFont(render, font)
				font.Modified = false
			}
		}

		// Check for animation component
		component = obj.GetComponent(&entity.AnimationComponent{})
		if component != nil {
			animation := component.(*entity.AnimationComponent)
			render.Crop = nextAnimation(r.time, animation, render.Crop)
		}

		// Draw
		// Offset according to the camera
		crop := *render.Crop
		var x, y int32
		if render.Scroll {
			crop.X, crop.Y = r.GetCameraPosition()
			x = position.Pos.X
			y = position.Pos.Y
		} else if !r.isRenderable(position.Pos, render.Crop) {
			// check if it is necessary to render
			continue
		} else {
			x, y = r.OffsetPosition(position.Pos.X, position.Pos.Y)
		}
		dst := &sdl.Rect{x, y, render.Crop.W, render.Crop.H}
		r.Renderer.CopyEx(render.Texture, &crop, dst, render.Angle, render.Center, render.Flip)
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
		logFatal("Error: Render cannot be null")
	}
	if font == nil {
		logFatal("Error: Font cannot be null")
	}
	// Get color. If color is not set, make it black
	if font.Color == nil {
		color = sdl.Color{0, 0, 0, 255}
	} else {
		color = *font.Color
	}
	//Load image at specified path
	if solid, err = font.Font.RenderUTF8Solid(font.Text, color); err != nil {
		logFatalf("Failed to render text: %s\n", err)
	}
	defer solid.Free()
	//Create texture from surface pixels
	newTexture, err = render.Renderer.CreateTextureFromSurface(solid)
	if err != nil {
		logFatalf("Unable to create texture from %s! SDL Error: %s\n", font.Text, sdl.GetError())
	}
	render.Texture = newTexture
	render.Crop = &sdl.Rect{0, 0, solid.W, solid.H}
}

// drawGeometry draws on the renderer the geometry. We don't use texture, because it's faster to draw directly using the renderer
func (r *RenderSystem) drawGeometry(pos *entity.PositionComponent, geometryComponent interface{}) {
	render := r.Renderer
	switch g := geometryComponent.(type) {
	case *entity.RectangleComponent:
		render.SetDrawColor(g.Color.R, g.Color.G, g.Color.B, g.Color.A)
		x := pos.Pos.X
		y := pos.Pos.Y
		w := g.Size.X
		h := g.Size.Y
		// Offset position according to camera
		x, y = r.OffsetPosition(x, y)
		// Result of rectangle to draw
		rect := &sdl.Rect{x, y, w, h}
		// check if it is necessary to render
		if !r.isRenderable(pos.Pos, rect) {
			return
		}
		if g.Filled {
			render.FillRect(rect)
		} else {
			render.DrawRect(rect)
		}
	default:
		logFatal("Geometry component not implemented in render function")
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

// drawGrid is used to draw a grid to help debugging
func (r *RenderSystem) drawGrid(grid *entity.GridComponent) {
	render := r.Renderer
	var area sdl.Point
	area.X, area.Y = r.Window.GetSize()

	if grid.Color != nil {
		render.SetDrawColor(grid.Color.R, grid.Color.G, grid.Color.B, grid.Color.A)
	} else {
		render.SetDrawColor(0x0, 0x0, 0x0, 0xFF)
	}

	xIterations := area.X / grid.Size.X
	yIterations := area.Y / grid.Size.Y
	for i := int32(0); i < xIterations; i++ {
		for j := int32(0); j < yIterations; j++ {
			rect := &sdl.Rect{i * grid.Size.X, j * grid.Size.Y, grid.Size.X, grid.Size.Y}
			render.DrawRect(rect)
		}
	}
}
