package macaw

import (
	"log"

	"github.com/tubelz/macaw/internal/utils"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	// WinWidth has the default screen width
	WinWidth = 800
	// WinHeight has the default screen height
	WinHeight = 600
	// WinTitle has the default screen title
	WinTitle = "macaw"
)

// We only have one input manager for now and one window, so they are going to be globals
var (
	// Window is the window of our game
	Window *sdl.Window
)

// Initialize SDL
func Initialize() error {
	var window *sdl.Window
	var err error
	// flags available:
	// INIT_TIMER          = 0x00000001 // timer subsystem
	// INIT_AUDIO          = 0x00000010 // audio subsystem
	// INIT_VIDEO          = 0x00000020 // video subsystem; automatically initializes the events subsystem
	// INIT_JOYSTICK       = 0x00000200 // joystick subsystem; automatically initializes the events subsystem
	// INIT_HAPTIC         = 0x00001000 // haptic (force feedback) subsystem
	// INIT_GAMECONTROLLER = 0x00002000 // controller subsystem; automatically initializes the joystick subsystem
	// INIT_EVENTS         = 0x00004000 // events subsystem
	// INIT_NOPARACHUTE    = 0x00100000 // compatibility; this flag is ignored
	// INIT_EVERYTHING = INIT_TIMER | INIT_AUDIO | INIT_VIDEO | INIT_EVENTS | INIT_JOYSTICK | INIT_HAPTIC | INIT_GAMECONTROLLER // all of the above subsystems
	sdl.Init(sdl.INIT_EVERYTHING)

	// TODO: enable / disable log per module
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// load support for the JPG and PNG image formats
	imgFlags := img.INIT_JPG | img.INIT_PNG
	if err := img.Init(imgFlags); err != imgFlags {
		utils.LogFatalf("Failed to initialize IMG: %d (%s)\n", err, img.GetError())
	}

	// load ttf support
	if err := ttf.Init(); err != nil {
		utils.LogFatalf("Failed to initialize TTF: %s\n", err)
	}

	// load sound support
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		utils.LogFatalf("Failed to initialize MIX: %s\n", err)
	}
	soundFlags := mix.INIT_FLAC | mix.INIT_OGG
	if err := mix.Init(soundFlags); err != nil {
		log.Println(err)
	}
	if err := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Println(err)
	}

	// we are only creating one window for now, so Window will be a global
	if window, err = sdl.CreateWindow(WinTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WinWidth, WinHeight, sdl.WINDOW_SHOWN); err != nil {
		utils.LogFatalf("Failed to create window: %s\n", err)
	}
	Window = window

	// Whenever there's a text input from user the text input should be activate to start accepting Unicode characters.
	sdl.StopTextInput()
	return err
}

// QuitImg unloads libraries loaded with img.Init()
func QuitImg() {
	img.Quit()
}

// QuitSound close open files and unloads libraries loaded with mix.Init()
func QuitSound() {
	for _, _, _, open, err := mix.QuerySpec(); err == nil && open > 0; _, _, _, open, err = mix.QuerySpec() {
		log.Println("closing")
		mix.CloseAudio()
	}
	mix.Quit()
}

// QuitFont unloads libraries loaded with ttf.Init()
func QuitFont() {
	ttf.Quit()
}

// Quit cleans up all initialized subsystems
func Quit() {
	// Close SDL
	sdl.Quit()
}
