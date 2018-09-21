package entity

import (
	"github.com/tubelz/macaw/internal/utils"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

// Spritesheet has the information about the image
type Spritesheet struct {
	// renderer
	Renderer *sdl.Renderer
	Texture  *sdl.Texture
	Filepath string
}

// Init initialize the spritesheet. It generates the texture of that image
func (s *Spritesheet) Init() {
	var newTexture *sdl.Texture
	var newSurface *sdl.Surface
	var err error
	if s.Renderer == nil {
		utils.LogFatal("Render cannot be null")
	}
	if _, err = os.Stat(s.Filepath); os.IsNotExist(err) {
		utils.LogFatalf("File **%s** does not exist", s.Filepath)
	}
	//Load image at specified path
	newSurface, err = img.Load(s.Filepath)
	defer newSurface.Free()
	if err != nil {
		utils.LogFatalf("Unable to load image %s! SDL_image Error: %s\n", s.Filepath, img.GetError())
	} else {
		//Create texture from surface pixels
		newTexture, err = s.Renderer.CreateTextureFromSurface(newSurface)
		if err != nil {
			utils.LogFatalf("Unable to create texture from %s! SDL Error: %s\n", s.Filepath, sdl.GetError())
		}
	}
	// set values on the struct
	s.Texture = newTexture
}

// LoadSprite add the information of the sprite to the render component
func (s *Spritesheet) LoadSprite(crop *sdl.Rect) RenderComponent {
	sprite := RenderComponent{Texture: s.Texture, Crop: crop}
	return sprite
}
