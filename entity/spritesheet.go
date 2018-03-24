package entity

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

// Spritesheet has the information about the image
type Spritesheet struct {
	// renderer
	Renderer *sdl.Renderer
	Texture  *sdl.Texture
	filename string
}

// Init initialize the spritesheet. It generates the texture of that image
func (s *Spritesheet) Init(renderer *sdl.Renderer, fname string) {
	var newTexture *sdl.Texture
	var newSurface *sdl.Surface
	var err error
	if renderer == nil {
		logFatal("Render cannot be null")
	}
	if _, err = os.Stat(fname); os.IsNotExist(err) {
		logFatalf("File **%s** does not exist", fname)
	}
	//Load image at specified path
	newSurface, err = img.Load(fname)
	defer newSurface.Free()
	if err != nil {
		logFatalf("Unable to load image %s! SDL_image Error: %s\n", fname, img.GetError())
	} else {
		//Create texture from surface pixels
		newTexture, err = renderer.CreateTextureFromSurface(newSurface)
		if err != nil {
			logFatalf("Unable to create texture from %s! SDL Error: %s\n", fname, sdl.GetError())
		}
	}
	// set values on the struct
	s.Texture = newTexture
	s.filename = fname
	s.Renderer = renderer
}

// LoadSprite add the information of the sprite to the render component
func (s *Spritesheet) LoadSprite(crop *sdl.Rect) RenderComponent {
	sprite := RenderComponent{Renderer: s.Renderer, Texture: s.Texture, Crop: crop}
	return sprite
}
