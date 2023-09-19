package screen

import (
	"chip8/internal/enum"
	"chip8/internal/keyboard"

	"github.com/veandco/go-sdl2/sdl"
)

func InitSDL() (renderer *sdl.Renderer, err error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, err
	}
	window, err := sdl.CreateWindow("chip8", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 64, 32, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}
	return renderer, nil
}
func Update(renderer *sdl.Renderer, display *[64 * 32]bool) {
	renderer.Clear()
	for pixel_index, pixel_set := range display {
		if pixel_set {
			renderer.SetDrawColor(255, 255, 255, 255)
		} else {
			renderer.SetDrawColor(0, 0, 0, 255)
		}
		rect := &sdl.FRect{W: 10, H: 10}
		rect.X = float32(pixel_index%640) * 10
		rect.Y = float32(pixel_index/640) * 10
		renderer.FillRectF(rect)
		renderer.DrawRectF(rect)
	}
	renderer.Present()
}
func HandleSDLEvents(event sdl.Event, state *enum.Machine_state) {
	switch t := event.(type) {
	case sdl.QuitEvent:
		*state = enum.Stop
	case sdl.KeyboardEvent:
		if t.GetType() == sdl.KEYDOWN {
			keyboard.HandleSDLInputKeys(t.Keysym.Sym, state)
		}
	}
}
