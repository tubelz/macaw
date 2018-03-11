package macaw

import (
	"github.com/veandco/go-sdl2/mix"
	// "github.com/veandco/go-sdl2/sdl"
	"io/ioutil"
	"log"
)

// PlaySound plays the file once
func PlaySound(file string) (*mix.Chunk, error) {
	// Load entire WAV data from file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Load WAV from data (memory)
	chunk, err := mix.QuickLoadWAV(data)
	if err != nil {
		log.Println(err)
	}
	// Play the sound one time
	chunk.Play(1, 1)

	return chunk, nil
}

// PlayMusic plays the file and leave it as background music
func PlayMusic(file string) {
	if music, err := mix.LoadMUS(file); err != nil {
		log.Println(err)
	} else if err = music.Play(-1); err != nil {
		log.Println(err)
	}
}

// StopMusic stops the music
func StopMusic() {
	mix.HaltMusic()
}
