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
var cycles_per_frame int = 10

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
	var path = "roms/games/Wall [David Winter].ch8"
	rl.InitWindow(width, height, title)
	defer rl.CloseWindow()
	rl.SetTargetFPS(fps)
	rl.InitAudioDevice() // Initialize audio device
	defer rl.CloseAudioDevice()
	beep := rl.LoadSound("resources/audio/beep.wav") // Load beep sound
	rl.SetMasterVolume(0.3)
	defer rl.UnloadSound(beep)
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
		for i := 0; i < cycles_per_frame; i++ {
			c.Cycle()
		}
		UpdateDisplay(c.GraphicsBuffer)
		rl.EndDrawing()
		rl.WaitTime(0.01667)
		if c.DT > 0 {
			c.DT--
		}
		if c.ST > 0 {
			if c.ST == 1 { // Play sound when ST is set
				rl.PlaySound(beep)
			}
			c.ST--
		}

	}
}
