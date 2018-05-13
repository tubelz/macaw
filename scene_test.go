package macaw

import (
	// "github.com/tubelz/macaw/entity"
	// "github.com/tubelz/macaw/input"
	// "github.com/tubelz/macaw/internal/utils"
	"github.com/tubelz/macaw/system"
	"github.com/veandco/go-sdl2/sdl"
	"testing"
	// "time"
)

func TestScene_InitNoSceneOptions(t *testing.T) {
	scene := &Scene{}
	scene.AddRenderSystem(&system.RenderSystem{})
	scene.Init()
}

func TestScene_InitAllSceneOptions(t *testing.T) {
	scene := &Scene{SceneOptions: SceneOptions{
		HideCursor: true,
		Music:      "",
		BgColor:    sdl.Color{0, 0, 0, 0},
	}}
	scene.AddRenderSystem(&system.RenderSystem{})
	scene.Init()
}

func TestScene_InitInitFunc(t *testing.T) {
	isExecuted := false
	initFunc := func() { isExecuted = true }
	scene := &Scene{InitFunc: initFunc}
	scene.AddRenderSystem(&system.RenderSystem{})
	scene.Init()

	if isExecuted != true {
		t.Error("InitFunc not being executed when scene is initiated.")
	}
}

func TestSceneManager(t *testing.T) {
	// initialize scenes
	scene1 := &Scene{Name: "scene1", RenderSystem: &system.RenderSystem{}}
	scene2 := &Scene{Name: "scene2", RenderSystem: &system.RenderSystem{}}
	// AddScene func
	t.Run("AddScene", func(t *testing.T) {
		sm := SceneManager{}
		sm.AddScene(scene1)
		sm.AddScene(scene2)
		// check if the first scene is indeed scene1
		if sm.Current().Name != "scene1" {
			t.Errorf("Got %s; want scene1", sm.Current().Name)
		}
	})
	t.Run("AddScene with InitFunc", func(t *testing.T) {
		val := false
		scene1.InitFunc = func() { val = true }
		sm := SceneManager{}
		sm.AddScene(scene1)
		if val != true {
			t.Error("AddScene not calling InitFunc for the first scene")
		}
	})
	// NextScene func
	t.Run("NextScene", func(t *testing.T) {
		sm := SceneManager{}
		sm.AddScene(scene1)
		sm.AddScene(scene2)
		// check if the next scene is working
		sm.NextScene()
		if sm.Current().Name != "scene2" {
			t.Errorf("Got %s; want scene2", sm.Current().Name)
		}
		// needs to go to first scene now
		sm.NextScene()
		if sm.Current().Name != "scene1" {
			t.Errorf("Got %s; want scene1", sm.Current().Name)
		}
	})
	t.Run("NextScene with InitFunc", func(t *testing.T) {
		sm := SceneManager{}
		val1 := false
		val2 := false
		// add InitFunc to our scenes
		scene1.InitFunc = func() { val1 = true }
		scene2.InitFunc = func() { val2 = true }
		// now we add the scenes to do the checking
		sm.AddScene(scene1)
		sm.AddScene(scene2)
		sm.NextScene()
		if val2 != true {
			t.Error("InitFunc in NextScene not being called")
		}
		// we have to change val value again, because AddScene should change val1 value
		val1 = false
		sm.NextScene()
		if val1 != true {
			t.Error("InitFunc in NextScene not being called")
		}
	})
	t.Run("NextScene with ExitFunc", func(t *testing.T) {
		sm := SceneManager{}
		val1 := false
		val2 := false
		scene1.ExitFunc = func() { val1 = true }
		scene2.ExitFunc = func() { val2 = true }
		sm.AddScene(scene1)
		sm.AddScene(scene2)
		sm.NextScene()
		if val1 != true {
			t.Error("ExitFunc in NextScene not being called")
		}
		sm.NextScene()
		if val2 != true {
			t.Error("ExitFunc in NextScene not being called")
		}
	})
	t.Run("ChangeScene", func(t *testing.T) {
		sm := SceneManager{}
		sm.AddScene(scene1)
		sm.AddScene(scene2)
		sm.ChangeScene("scene2")
		if sm.Current().Name != "scene2" {
			t.Errorf("Got %s; want scene2", sm.Current().Name)
		}
		// needs to go to first scene now
		sm.ChangeScene("scene1")
		if sm.Current().Name != "scene1" {
			t.Errorf("Got %s; want scene1", sm.Current().Name)
		}
	})
	t.Run("ChangeScene with InitFunc", func(t *testing.T) {
		sm := SceneManager{}
		val1 := false
		val2 := false
		// add InitFunc to our scenes
		scene1.InitFunc = func() { val1 = true }
		scene2.InitFunc = func() { val2 = true }
		// now we add the scenes to do the checking
		sm.AddScene(scene1)
		sm.AddScene(scene2)
		sm.ChangeScene("scene2")
		if val2 != true {
			t.Error("InitFunc in NextScene not being called")
		}
		// we have to change val value again, because AddScene should change val1 value
		val1 = false
		sm.ChangeScene("scene1")
		if val1 != true {
			t.Error("InitFunc in NextScene not being called")
		}
	})
	t.Run("ChangeScene with ExitFunc", func(t *testing.T) {
		sm := SceneManager{}
		val1 := false
		val2 := false
		scene1.ExitFunc = func() { val1 = true }
		scene2.ExitFunc = func() { val2 = true }
		sm.AddScene(scene1)
		sm.AddScene(scene2)
		sm.ChangeScene("scene2")
		if val1 != true {
			t.Error("ExitFunc in NextScene not being called")
		}
		sm.ChangeScene("scene1")
		if val2 != true {
			t.Error("ExitFunc in NextScene not being called")
		}
	})
}
