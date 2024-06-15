package main

import (
	"chip8/config"
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var width int32 = 640
var height int32 = 320
var pixel_size int32 = 10
var scale int32 = 2
var title = "one day"
var fps int32 = 15

func HandleInput() {
	if rl.IsKeyDown(rl.KeyEscape) {
		rl.CloseWindow()
	}
}

/*
Function tha update pixel using DisplayBuffer as base
if pixel is on == 1
if is not then is off
*/
func UpdateDisplay(DisplayBuffer [32][64]int) {
	for y_index, y := range DisplayBuffer {
		for x_index, x := range y {
			if config.DEBUG {
				fmt.Printf("DisplayBuffer[%d][%d]=%d\n", y_index, x_index, x)
			}
			if x == 1 {
				rl.DrawRectangle(scale*pixel_size*int32(x_index), scale*pixel_size*int32(y_index), pixel_size, pixel_size, rl.White)
			} else {
				rl.DrawRectangle(scale*pixel_size*int32(x_index), scale*pixel_size*int32(y_index), pixel_size, pixel_size, rl.Black)
			}
		}
	}
}

var GraphicsBuffer [32][64]int
var I int32
var V [16]int32
var SP int32
var PC int32 = 512
var Memory [4096]uint16

func LoadFont() {
	var chip8Fontset = []byte{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4
		0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x10, 0x20, 0x40, 0x40, // 7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
		0xF0, 0x90, 0xF0, 0x90, 0x90, // A
		0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
		0xF0, 0x80, 0x80, 0x80, 0xF0, // C
		0xE0, 0x90, 0x90, 0x90, 0xE0, // D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
		0xF0, 0x80, 0xF0, 0x80, 0x80, // F
	}
	for addr, b := range chip8Fontset {
		Memory[addr] = uint16(b)
	}

}
func LoadROM(filePath string) (int, error) {
	var start = 512 //1536 in ETI 660 Chip8-Programs
	romBytes, err := os.ReadFile(filePath)
	if err != nil {
		return 0, fmt.Errorf("error loading rom %s:%s", filePath, err.Error())
	}
	for i, b := range romBytes {
		Memory[start+i] = uint16(b)
		if config.DEBUG {
			fmt.Printf("Memory[%d]=%x\n", start+i, Memory[start+i])
		}
	}
	return len(romBytes), nil

}
func cycle() {
	//fetch
	var addr1 = Memory[PC]
	var addr2 = Memory[PC+1]
	PC += 2
	var opcode = addr1<<8 | addr2
	var x = (opcode >> 8) & 0x0F
	var y = (opcode >> 4) & 0x00F // the upper 4 bits of the low byte
	var n = opcode & 0x000F       // the lowest 4 bits
	var kk = opcode & 0x00FF      // the lowest 8 bits
	var nnn = opcode & 0x0FFF
	if config.DEBUG {
		fmt.Printf("opcode:%x\n", opcode)
		fmt.Printf("	x:%x\n", x)
		fmt.Printf("	y:%x\n", y)
		fmt.Printf("	n:%x\n", n)
		fmt.Printf("	kk:%x\n", kk)
		fmt.Printf("	nnn:%x\n", nnn)
	}
	//exec
}
func main() {
	var path = "internal/__test_roms__/IBM Logo2.ch8"
	size, err := LoadROM(path)
	if err != nil {
		panic(err)
	}
	LoadFont()
	rl.InitWindow(width, height, title)
	defer rl.CloseWindow()
	rl.SetTargetFPS(fps)
	for !rl.WindowShouldClose() {
		if PC > (512 + int32(size)) {
			rl.CloseWindow()
		}
		HandleInput()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		cycle()
		//UpdateDisplay(GraphicsBuffer)
		rl.EndDrawing()
		rl.WaitTime(0.2)
	}
}
