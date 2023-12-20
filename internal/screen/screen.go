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
	window, err := sdl.CreateWindow("chip8", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 320, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}
	renderer, err = sdl.CreateRenderer(window, 0, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}
	return renderer, nil
}
func Update(renderer *sdl.Renderer, display *[32][64]int) {
	renderer.Clear()
	for pixel_row_index, pixel_row := range display {
		for pixel_column_index, pixel_set := range pixel_row {
			if pixel_set == 1 {
				renderer.SetDrawColor(255, 255, 255, 255)
			} else {
				renderer.SetDrawColor(0, 0, 0, 255)
			}
			rect := &sdl.FRect{W: 1, H: 1}
			rect.X = float32(pixel_column_index)
			rect.Y = float32(pixel_row_index)
			renderer.FillRectF(rect)
			renderer.DrawRectF(rect)
		}
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
