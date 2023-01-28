package main

import (
	"chip8/internal/keyboard"

	"github.com/veandco/go-sdl2/sdl"
)

func initSDL() (window *sdl.Window, err error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, err
	}
	window, err = sdl.CreateWindow("chip8", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 320, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}
	return window, nil

}
func handleSDLEvents(event sdl.Event, running *bool) {
	switch t := event.(type) {
	case sdl.QuitEvent:
		*running = false
		break
	case sdl.KeyboardEvent:
		keyboard.HandleSDLInputKeys(t.Keysym.Sym)
		break
	}
}
func main() {
	window, err := initSDL()
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			handleSDLEvents(event, &running)
		}
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		renderer.SetDrawColor(255, 255, 255, 255)

		renderer.Present()
		sdl.Delay(16)
	}
}
