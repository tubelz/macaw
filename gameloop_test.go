package macaw

import (
	"github.com/tubelz/macaw/entity"
	"github.com/tubelz/macaw/input"
	"github.com/tubelz/macaw/system"
	"github.com/veandco/go-sdl2/sdl"
	"testing"
	"time"
)

// TestGameLoop_Run checks if our gameloop runs with a basic setup
func TestGameLoop_Run(t *testing.T) {
	setup(t)
	// This is also consider setup. We are adding what it needs to be added so our game loop can be ran
	input := &input.Manager{}
	sceneGame := &Scene{Name: "game"}
	sceneGame.AddRenderSystem(&system.RenderSystem{Camera: &entity.Entity{}})
	gameLoop := &GameLoop{InputManager: input}
	gameLoop.AddScene(sceneGame)
	// close game after 100 milliseconds
	go func() {
		duration := time.Millisecond
		time.Sleep(duration * 100)
		quitEvent := &sdl.QuitEvent{Type: sdl.QUIT, Timestamp: sdl.GetTicks()}
		sdl.PushEvent(quitEvent)
	}()
	gameLoop.Run()

	teardown()
}
