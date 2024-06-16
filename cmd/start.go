package cmd

import (
	"cp8/config"
	"cp8/internal/cpu"
	"cp8/internal/keyboard"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var FilePath string

func Start(filePath string) {
	rl.InitWindow(config.Width, config.Height, config.Title)
	defer rl.CloseWindow()
	rl.SetTargetFPS(config.Fps)
	rl.InitAudioDevice() // Initialize audio device
	defer rl.CloseAudioDevice()
	beep := rl.LoadSound("resources/audio/beep.wav") // Load beep sound
	rl.SetMasterVolume(0.3)
	defer rl.UnloadSound(beep)

	c, err := cpu.NewCPU(filePath)
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
		for i := 0; i < config.CyclesPerFrame; i++ {
			c.Cycle()
		}
		cpu.UpdateDisplay(c.GraphicsBuffer)
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
