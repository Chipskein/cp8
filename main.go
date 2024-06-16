package main

import (
	"chip8/internal/cpu"
	"chip8/internal/keyboard"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var width int32 = 640
var height int32 = 320
var pixel_size int32 = 10
var scale int32 = 1
var title = "one day"
var fps int32 = 60

/*UpdateDisplay function is used to set a display buffer to the screen.*/
func UpdateDisplay(DisplayBuffer [32][64]int) {
	for y_index, y := range DisplayBuffer {
		for x_index, x := range y {
			if x == 1 {
				rl.DrawRectangle(scale*pixel_size*int32(x_index), scale*pixel_size*int32(y_index), pixel_size, pixel_size, rl.White)
			} else {
				rl.DrawRectangle(scale*pixel_size*int32(x_index), scale*pixel_size*int32(y_index), pixel_size, pixel_size, rl.Black)
			}
		}
	}
}

func main() {
	var path = "roms/test_opcode.ch8"
	rl.InitWindow(width, height, title)
	defer rl.CloseWindow()
	rl.SetTargetFPS(fps)
	c, err := cpu.NewCPU(path)
	if err != nil {
		panic(err)
	}
	for !rl.WindowShouldClose() {
		if c.PC > (c.StartPC + int32(c.RomSize)) {
			rl.CloseWindow()
			os.Exit(0)
		}
		keyboard.HandleInput(c)
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		c.Cycle()
		UpdateDisplay(c.GraphicsBuffer)
		rl.EndDrawing()
		rl.WaitTime(0.01667)
		if c.DT > 0 {
			c.DT--
		}
		if c.ST > 0 {
			c.ST--
		}

	}
}
