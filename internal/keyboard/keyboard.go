package keyboard

import (
	"chip8/internal/enum"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func HandleSDLInputKeys(keyCode sdl.Keycode, state *enum.Machine_state) {
	switch keyCode {

	case 27:
		//esc
		*state = enum.Stop
	case 32:
		//space
		if *state == enum.Paused {
			*state = enum.Running
			return
		}
		*state = enum.Paused
		log.Printf("Machine Paused")
	case 48:
		//0
		log.Printf("Key 0 Pressed %d\n", keyCode)
	case 49:
		//1
		log.Printf("Key 1 Pressed %d\n", keyCode)
	case 50:
		//2
		log.Printf("Key 2 Pressed %d\n", keyCode)
	case 51:
		//3
		log.Printf("Key 3 Pressed %d\n", keyCode)
	case 52:
		//4
		log.Printf("Key 4 Pressed %d\n", keyCode)
	case 53:
		//5
		log.Printf("Key 5 Pressed %d\n", keyCode)
	case 54:
		//6
		log.Printf("Key 6 Pressed %d\n", keyCode)
	case 55:
		//7
		log.Printf("Key 7 Pressed %d\n", keyCode)
	case 56:
		//8
		log.Printf("Key 8 Pressed %d\n", keyCode)
	case 57:
		//9
		log.Printf("Key 9 Pressed %d\n", keyCode)
	case 97:
		//a
		log.Printf("Key A Pressed %d\n", keyCode)
	case 98:
		//b
		log.Printf("Key B Pressed %d\n", keyCode)
	case 99:
		//c
		log.Printf("Key C Pressed %d\n", keyCode)
	case 100:
		//d
		log.Printf("Key D Pressed %d\n", keyCode)
	case 101:
		//e
		log.Printf("Key E Pressed %d\n", keyCode)
	case 102:
		//f
		log.Printf("Key F Pressed %d\n", keyCode)
	}

}
