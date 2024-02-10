package screen

import (
	"chip8/internal/enum"
	"chip8/internal/keyboard"

	"github.com/veandco/go-sdl2/sdl"
)

const WIDTH = 640
const HEIGHT = 320
const SCALE = 10

var SCREEN [HEIGHT][WIDTH]int

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

func scale_my_buffer(buf *[32][64]int) {
	scaleX := float64(len(SCREEN)) / float64(len(buf))
	scaleY := float64(len(SCREEN[0])) / float64(len(buf[0]))
	for i := range SCREEN {
		for j := range SCREEN[i] {
			origX := int(float64(i) / scaleX)
			origY := int(float64(j) / scaleY)
			SCREEN[i][j] = buf[origX][origY]
		}
	}
}

func draw(renderer *sdl.Renderer) {
	for y, y_value := range SCREEN {
		for x, x_value := range y_value {
			if x_value == 1 {
				renderer.SetDrawColor(255, 255, 255, 255)
			} else {
				renderer.SetDrawColor(0, 0, 0, 255)
			}
			rect := &sdl.FRect{W: 1, H: 1}
			rect.X = float32(x)
			rect.Y = float32(y)
			renderer.FillRectF(rect)
			renderer.DrawRectF(rect)

		}
	}
}
func Update(renderer *sdl.Renderer, display *[32][64]int) {
	renderer.Clear()
	scale_my_buffer(display)
	draw(renderer)
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
