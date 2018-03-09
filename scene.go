package macaw

import (
	"github.com/tubelz/macaw/system"
	"github.com/veandco/go-sdl2/sdl"
)

// SceneManager manges the scenes in the game
type SceneManager struct {
	Scenes     []*Scene
	currentPos int
	// SceneMap has the position of the scene in the array
	SceneMap map[string]int
}

// AddScene adds a new scene
func (s *SceneManager) AddScene(scene *Scene) {
	s.Scenes = append(s.Scenes, scene)
	if len(s.Scenes) == 1 {
		scene.Init()
	}
	if s.SceneMap == nil {
		s.SceneMap = make(map[string]int)
	}
	if scene.Name != "" {
		s.SceneMap[scene.Name] = len(s.Scenes) - 1
	}
}

// Current returns the current Scene
func (s *SceneManager) Current() *Scene {
	if len(s.Scenes) > 0 {
		return s.Scenes[s.currentPos]
	}
	return nil
}

// RemoveScene removes a scene
func (s *SceneManager) RemoveScene() {

}

// NextScene goes to the next scene if it exists
func (s *SceneManager) NextScene() *Scene {
	if s.currentPos < (len(s.Scenes) - 1) {
		s.currentPos++
	}
	s.Current().Init()
	return s.Current()
}

// ChangeScene changes to a specific scene by its name
func (s *SceneManager) ChangeScene(sceneName string) *Scene {
	if pos, ok := s.SceneMap[sceneName]; ok {
		s.currentPos = pos
	}
	s.Current().Init()
	return s.Current()
}

// Scene is reponsible to hold the systems in a scene
type Scene struct {
	Name          string
	UpdateSystems []system.Systemer    // responsible to update the game
	RenderSystem  *system.RenderSystem // responsible to render the game
	SceneOptions
}

// Init initializes the scene
func (s *Scene) Init() {
	show := 1
	if s.HideCursor {
		show = 0
	}
	sdl.ShowCursor(show)
}

// AddGameUpdateSystem adds the systems which will run in the game loop
func (s *Scene) AddGameUpdateSystem(system system.Systemer) {
	s.UpdateSystems = append(s.UpdateSystems, system)
}

// AddRenderSystem adds the render system to our game loop
func (s *Scene) AddRenderSystem(system *system.RenderSystem) {
	s.RenderSystem = system
}


// SceneOptions contains the options for the scene
type SceneOptions struct {
	HideCursor bool // true - hides, false - shows
	Music string
}
