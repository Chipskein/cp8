package keyboard

import (
	"chip8/config"
	"chip8/internal/cpu"
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/*
HandleInput function is used to handle the input from the keyboard.
*/
/*
	Key bindings for the Chip8 Keyboard
	1 2 3 4      1 2 3 C
	Q W E R  ->  4 5 6 D
	A S D F      7 8 9 E
	Z X C V      A 0 B F
	Esc -> Close the window
*/
func HandleInput(c *cpu.CPU) {
	if rl.IsKeyDown(rl.KeyEscape) {
		rl.CloseWindow()
		os.Exit(0)
	}
	if rl.IsKeyReleased(rl.KeyOne) || rl.IsKeyReleased(rl.KeyTwo) || rl.IsKeyReleased(rl.KeyThree) || rl.IsKeyReleased(rl.KeyFour) ||
		rl.IsKeyReleased(rl.KeyQ) || rl.IsKeyReleased(rl.KeyW) || rl.IsKeyReleased(rl.KeyE) || rl.IsKeyReleased(rl.KeyR) ||
		rl.IsKeyReleased(rl.KeyA) || rl.IsKeyReleased(rl.KeyS) || rl.IsKeyReleased(rl.KeyD) || rl.IsKeyReleased(rl.KeyF) ||
		rl.IsKeyReleased(rl.KeyZ) || rl.IsKeyReleased(rl.KeyX) || rl.IsKeyReleased(rl.KeyC) || rl.IsKeyReleased(rl.KeyV) {
		c.Key = -1
	}
	//line 1
	if rl.IsKeyDown(rl.KeyOne) {
		if config.DEBUG {
			fmt.Println("Key 1 pressed")
		}
		c.Key = 0x1
	}
	if rl.IsKeyDown(rl.KeyTwo) {
		if config.DEBUG {
			fmt.Println("Key 2 pressed")
		}
		c.Key = 0x2
	}
	if rl.IsKeyDown(rl.KeyThree) {
		if config.DEBUG {
			fmt.Println("Key 3 pressed")
		}
		c.Key = 0x3
	}
	if rl.IsKeyDown(rl.KeyFour) {
		if config.DEBUG {
			fmt.Println("Key C pressed")
		}
		c.Key = 0xc
	}
	//line 2
	if rl.IsKeyDown(rl.KeyQ) {
		if config.DEBUG {
			fmt.Println("Key 4 pressed")
		}
		c.Key = 0x4
	}
	if rl.IsKeyDown(rl.KeyW) {
		if config.DEBUG {
			fmt.Println("Key 5 pressed")
		}
		c.Key = 0x5
	}
	if rl.IsKeyDown(rl.KeyE) {
		if config.DEBUG {
			fmt.Println("Key 6 pressed")
		}
		c.Key = 0x6
	}
	if rl.IsKeyDown(rl.KeyR) {
		if config.DEBUG {
			fmt.Println("Key D pressed")
		}
		c.Key = 0xd
	}
	//line3
	if rl.IsKeyDown(rl.KeyA) {
		if config.DEBUG {
			fmt.Println("Key 7 pressed")
		}
		c.Key = 0x7
	}
	if rl.IsKeyDown(rl.KeyS) {
		if config.DEBUG {
			fmt.Println("Key 8 pressed")
		}
		c.Key = 0x8
	}
	if rl.IsKeyDown(rl.KeyD) {
		if config.DEBUG {
			fmt.Println("Key 9 pressed")
		}
		c.Key = 0x9
	}
	if rl.IsKeyDown(rl.KeyF) {
		if config.DEBUG {
			fmt.Println("Key E pressed")
		}
		c.Key = 0xe
	}
	//line 4
	if rl.IsKeyDown(rl.KeyZ) {
		if config.DEBUG {
			fmt.Println("Key A pressed")
		}
		c.Key = 0xa
	}
	if rl.IsKeyDown(rl.KeyX) {
		if config.DEBUG {
			fmt.Println("Key 0 pressed")
		}
		c.Key = 0x0
	}
	if rl.IsKeyDown(rl.KeyC) {
		if config.DEBUG {
			fmt.Println("Key B pressed")
		}
		c.Key = 0xb
	}
	if rl.IsKeyDown(rl.KeyV) {
		if config.DEBUG {
			fmt.Println("Key F pressed")
		}
		c.Key = 0xf
	}
}
