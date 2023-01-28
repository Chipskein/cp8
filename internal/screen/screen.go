package screen

import (
	"chip8/internal/enum"
	"chip8/internal/keyboard"

	"github.com/veandco/go-sdl2/sdl"
)

func InitSDL() (window *sdl.Window, renderer *sdl.Renderer, err error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, nil, err
	}
	window, err = sdl.CreateWindow("chip8", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 320, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, nil, err
	}
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, nil, err
	}

	return window, renderer, nil
}
func HandleSDLEvents(event sdl.Event, state *enum.Machine_state) {
	switch t := event.(type) {
	case sdl.QuitEvent:
		*state = enum.Stop
		break
	case sdl.KeyboardEvent:
		if t.GetType() == sdl.KEYDOWN {
			keyboard.HandleSDLInputKeys(t.Keysym.Sym, state)
		}
		break
	}
}
