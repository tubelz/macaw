package macaw

import (
	"github.com/tubelz/macaw/cmd"
	"github.com/tubelz/macaw/input"
	"github.com/tubelz/macaw/internal/utils"
	"github.com/tubelz/macaw/system"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

// GameLoop is the data structure which will execute our systems in order.
// inspired by http://www.koonsolo.com/news/dewitters-gameloop/
type GameLoop struct {
	InputManager *input.Manager
	SceneManager
	now      uint32
	nextTick uint32
	fps      uint32
}

// update game systems that can be updated every couple frames
func (g *GameLoop) gameUpdate() {
	for _, system := range g.Current().UpdateSystems {
		system.Update()
	}
	g.InputManager.PopEvent()
	g.InputManager.Mouse.ClearMouseEvent()
}

// make it be generic like game update
func (g *GameLoop) render() {
	current := g.Current()
	current.RenderSystem.UpdateTime(g.now)
	// the accumulator is important for linear interpolation.
	// If the value of next tick is greater it means we can interpolate,
	// if not it means we hit the max frame so our render is as up-to-date with the physics as possible
	if g.nextTick > g.now {
		current.RenderSystem.UpdateAccumulator((g.now + system.UpdateTickLength) - g.nextTick)
	} else {
		current.RenderSystem.UpdateAccumulator(0)
	}
	current.RenderSystem.Update()
	g.fps++
}

// Run executes the game loop
func (g *GameLoop) Run() {
	if g.Current() == nil {
		utils.LogFatal("You need to add at least one scene")
	}
	fpsTick := sdl.GetTicks()
	g.nextTick = fpsTick
	for running := true; running; {
		running = g.InputManager.HandleEvents()
		g.now = sdl.GetTicks()
		for g.now >= g.nextTick {
			g.gameUpdate()
			g.nextTick += system.UpdateTickLength
		}
		g.render()

		if cmd.Parser.Debug() && g.now >= fpsTick+1000 {
			log.Printf("FPS: %d\n", g.fps)
			g.fps = 0
			fpsTick += 1000
		}
	}
}
