package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("chip8", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 400, 400, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			switch t := event.(type) {
			case sdl.QuitEvent:
				running = false
				break
			case sdl.KeyboardEvent:
				keyCode := t.Keysym.Sym
				log.Printf("%d\n", keyCode)
				break
			}

		}

	}
}
