package macaw

import (
	"github.com/tubelz/macaw/entity"
	"github.com/tubelz/macaw/input"
	"github.com/tubelz/macaw/internal/utils"
	"github.com/tubelz/macaw/system"
	"github.com/veandco/go-sdl2/sdl"
	"testing"
	"time"
)

// TestGameLoop_Run checks if our gameloop runs with a basic setup
func TestGameLoop_Run(t *testing.T) {
	utils.SetupLog(t)
	// This is also consider setup. We are adding what it needs to be added so our game loop can be ran
	input := &input.Manager{}
	em := &entity.Manager{}
	sceneGame := &Scene{Name: "game"}
	sceneGame.AddRenderSystem(&system.RenderSystem{Camera: &entity.Entity{}, EntityManager: em})
	// Initialize the mocked systems to test whether systems are executed in sequence work or not
	var val int
	mockFunc1 := func() {
		if val != 0 {
			t.Errorf("Got %d; want 0", val)
		}
		val++
	}
	mockFunc2 := func() {
		if val != 1 {
			t.Errorf("Got %d; want 1", val)
		}
		// add big delay to avoid the systems to be ran again
		duration := time.Millisecond * 10
		time.Sleep(duration)
	}
	mockSystem1 := &utils.MockSystem{MockFunc: mockFunc1}
	mockSystem2 := &utils.MockSystem{MockFunc: mockFunc2}
	sceneGame.AddGameUpdateSystem(mockSystem1)
	sceneGame.AddGameUpdateSystem(mockSystem2)
	gameLoop := &GameLoop{InputManager: input}
	gameLoop.AddScene(sceneGame)

	// close game after 100 milliseconds
	go func() {
		duration := time.Millisecond
		time.Sleep(duration)
		quitEvent := &sdl.QuitEvent{Type: sdl.QUIT, Timestamp: sdl.GetTicks()}
		sdl.PushEvent(quitEvent)
	}()
	gameLoop.Run()

	utils.TeardownLog()
}
