package entity

import (
	"log"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// MFont is the font struct that has the file, size and font pointer
type MFont struct {
	File string
	Size uint8
	font *ttf.Font
}

// Open opens the font
func (f *MFont) Open() *ttf.Font {
	var font *ttf.Font
	var err error

	if font, err = ttf.OpenFont(f.File, int(f.Size)); err != nil {
		log.Fatal("Failed to open font: %s\n", err)
	}
	f.font = font
	return font
}

// Close closes the font
func (f *MFont) Close() {
	f.font.Close()
}

// MText is the struct that has the rendering information
type MText struct {
	renderer *sdl.Renderer
	font *ttf.Font
	Text string
	Color *sdl.Color
}

// Init initialize the MText struct
func (t *MText) Init(renderer *sdl.Renderer, font *ttf.Font) {
	t.renderer = renderer
	t.font = font
}

// GenerateRenderComponent generate the render component used to draw the text
func (t *MText) GenerateRenderComponent() *RenderComponent {
	var newTexture *sdl.Texture
	var solid *sdl.Surface
	var color sdl.Color
	var err error

	if t.renderer == nil {
		log.Fatal("Error: Render cannot be null")
	}
	if t.font == nil {
		log.Fatal("Error: Font cannot be null")
	}

	// Get color. If color is not set, make it black
	if t.Color == nil {
		color = sdl.Color{0, 0, 0, 255}
	} else {
		color = *t.Color
	}

	//Load image at specified path
	if solid, err = t.font.RenderUTF8Solid(t.Text, color); err != nil {
		log.Fatal("Failed to render text: %s\n", err)
	}
	defer solid.Free()

	//Create texture from surface pixels
	newTexture, err = t.renderer.CreateTextureFromSurface(solid)
	if( err != nil ) {
		log.Fatal("Unable to create texture from %s! SDL Error: %s\n", t.Text, sdl.GetError())
	}

	component := &RenderComponent{Renderer: t.renderer, Texture: newTexture, Crop: &sdl.Rect{0, 0, solid.W, solid.H}}
	return component
}
