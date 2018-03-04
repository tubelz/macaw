package macaw

import (
	"github.com/tubelz/macaw/input"
	"github.com/tubelz/macaw/system"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

// GameLoop is the data structure which will execute our systems in order.
// inspired by http://www.koonsolo.com/news/dewitters-gameloop/
type GameLoop struct {
	InputManager  *input.Manager
	updateSystems []system.Systemer    // responsible to update the game
	renderSystem  *system.RenderSystem // responsible to render the game
	now           uint32
	nextTick      uint32
	fps           uint32
}

// AddGameUpdateSystem adds the systems which will run in the game loop
func (g *GameLoop) AddGameUpdateSystem(system system.Systemer) {
	g.updateSystems = append(g.updateSystems, system)
}

// AddRenderSystem adds the render system to our game loop
func (g *GameLoop) AddRenderSystem(system *system.RenderSystem) {
	g.renderSystem = system
}

// update game systems that can be updated every couple frames
func (g *GameLoop) gameUpdate() {
	for _, system := range g.updateSystems {
		system.Update()
	}
	g.InputManager.PopEvent()
}

// make it be generic like game update
func (g *GameLoop) render() {
	g.renderSystem.UpdateTime(g.now)
	// the accumulator is important for linear interpolation.
	// If the value of next tick is greater it means we can interpolate,
	// if not it means we hit the max frame so our render is as up-to-date with the physics as possible
	if g.nextTick > g.now {
		g.renderSystem.UpdateAccumulator((g.now + system.UpdateTickLength) - g.nextTick)
	} else {
		g.renderSystem.UpdateAccumulator(0)
	}
	g.renderSystem.Update()
	g.fps++
}

// Run executes the game loop
func (g *GameLoop) Run() {
	fpsTick := sdl.GetTicks()
	g.nextTick = fpsTick
	for running := true; running; {
		running = g.InputManager.HandleEvents(running)
		g.now = sdl.GetTicks()
		for loops := 0; g.now >= g.nextTick && loops < system.MaxFrameskip; loops++ {
			g.gameUpdate()
			g.nextTick += system.UpdateTickLength
		}
		g.render()

		if g.now >= fpsTick+1000 {
			log.Printf("FPS: %d\n", g.fps)
			g.fps = 0
			fpsTick += 1000
		}
	}
}
