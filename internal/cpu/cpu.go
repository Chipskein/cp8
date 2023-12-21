package cpu

import (
	"chip8/internal/enum"
	"chip8/internal/screen"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type Instruction struct {
	Opcode uint16
	Nnn    uint16
	Kk     uint8
	N      uint8
	X      uint8
	Y      uint8
}
type CPU struct {
	Rom_path            string
	Status              enum.Machine_state
	Rom_size            uint16
	Stack               [16]uint16 // (0x200 or 0x600)-0XFFF avaliable Memory for run programms ; 0X000 - 0x1FF avaliable for chip8 interpreter
	Memory              [0xFFF]uint8
	V                   [0xF + 1]uint8 //Registers V0-VF
	I                   uint16         //used as memory index store
	SP                  uint8          //stack pointer
	PC                  uint16         //pc counter
	Delay_timer         uint8
	Sound_timer         uint8
	Display             [32][64]int
	Current_instruction *Instruction
	Current_key         uint8 //Key pressed
	Update_Screen       bool
}

func (c *CPU) loadFontData() {
	var font_set = []uint8{
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
	var memAddress = 0x000
	for _, value := range font_set {
		c.Memory[memAddress] = value
		memAddress++
	}
}
func (c *CPU) Fetch() *Instruction {
	var addr1 = c.Memory[c.PC]
	var addr2 = c.Memory[c.PC+1]
	c.PC += 2
	var opcode = uint16(addr1)<<8 | uint16(addr2)
	var x = uint8((opcode & 0x0F00) >> 8) // the lower 4 bits of the high byte
	var y = uint8((opcode & 0x00F0) >> 4) // the upper 4 bits of the low byte
	var n = uint8(opcode & 0x000F)        // the lowest 4 bits
	var kk = uint8(opcode & 0x00FF)       // the lowest 8 bits
	var nnn = opcode & 0x0FFF             // the lowest 12 bits
	return &Instruction{Opcode: opcode, X: x, Y: y, N: n, Kk: kk, Nnn: nnn}
}
func (c *CPU) DecodeExec(inst *Instruction) {
	var opcode = inst.Opcode
	var x = inst.X
	var y = inst.Y
	var n = inst.N
	var kk = inst.Kk
	var nnn = inst.Nnn

	switch opcode & 0xF000 {
	case 0x0000:
		{
			switch kk {
			case 0x00E0: // clear the screen
				for row_index, row := range c.Display {
					for column_index := range row {
						c.Display[row_index][column_index] = 0
					}
				}
				c.Update_Screen = true
				c.PC += 2
			}
		}
	case 0x1000: // 1nnn: jump to address nnn
		c.PC = nnn
	case 0x6000: // 6xkk: set V[x] = kk
		c.V[x] = kk
		c.PC += 2
	case 0x7000: // 7xkk: set V[x] = V[x] + kk
		c.V[x] += kk
		c.PC += 2
	case 0xA000: // Annn: set I to address nnn
		c.I = nnn
		c.PC += 2
	case 0xD000: // Dxyn: Display an n-byte sprite starting at memory
		// location I at (Vx, Vy) on the screen, VF = collision
		// Initialize collision flag to 0
		c.V[0xF] = 0
		startX := c.V[x]
		startY := c.V[y]
		log.Printf("%x\n", inst.Opcode)
		for y_sprite := uint16(0); y_sprite < uint16(n); y_sprite++ {
			spriteByte := c.Memory[c.I+y_sprite]
			for x_sprite := uint16(0); x_sprite < 8; x_sprite++ {
				displayX := (startX + uint8(x_sprite)) % 64
				displayY := (startY + uint8(y_sprite)) % 32

				currentPixel := c.Display[displayY][displayX]
				spriteBit := (spriteByte >> (7 - uint8(x_sprite))) & 1
				if currentPixel == 1 && spriteBit == 1 {
					c.V[0xF] = 1
					log.Printf("COLLISÃO:%x\n", inst.Opcode)
				}
				c.Display[displayY][displayX] ^= int(spriteBit)
			}
		}
		c.Update_Screen = true
		c.PC += 2
	default:
		log.Printf("\nIntruction\n\topcode:%x\n", opcode)
		log.Panicf("UNKNOW INSTRUCTION\n")
	}

}

func (c *CPU) cycle() {
	if c.Status != enum.Paused {
		insc := c.Fetch()
		c.Current_instruction = insc
		c.DecodeExec(insc)
	}
}
func (c *CPU) loadROM(rom []byte) {
	c.Rom_size = uint16(len(rom))
	for addr, Bytes := range rom {
		c.Memory[0x200+addr] = Bytes
	}
	c.PC = 0x200
}
func (c *CPU) run() {
	renderer, err := screen.InitSDL()
	if err != nil {
		panic(err)
	}
	c.Status = enum.Running
	for c.Status != enum.Stop {
		if c.PC < (c.Rom_size)+0x200 {
			c.cycle()
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				screen.HandleSDLEvents(event, &c.Status)
			}
			if c.Update_Screen {
				screen.Update(renderer, &c.Display)
				c.Update_Screen = false
			}
			sdl.Delay(300)
		}

	}
}
func Init(rom []byte) {
	var cpu = &CPU{}
	cpu.loadFontData()
	cpu.loadROM(rom)
	cpu.run()
}
