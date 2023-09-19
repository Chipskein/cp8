package cpu

import (
	"chip8/internal/enum"
	"chip8/internal/screen"
	"log"
	"math/rand"

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
	Display             [64 * 32]bool
	Current_instruction *Instruction
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
	var memAddress = 0x050
	for _, value := range font_set {
		c.Memory[memAddress] = value
		memAddress++
	}
}
func (c *CPU) fetch() *Instruction {
	var addr1 = c.Memory[c.PC]
	var addr2 = c.Memory[c.PC+1]
	c.PC += 2
	var opcode = uint16(addr1)<<8 | uint16(addr2)
	var x = uint8((opcode >> 8) & 0x000F) // the lower 4 bits of the high byte
	var y = uint8((opcode >> 4) & 0x000F) // the upper 4 bits of the low byte
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
				for i := range c.Display {
					c.Display[i] = false
				}
				c.PC += 2
			case 0x00EE: // ret
				c.SP--
				c.PC = c.Stack[c.SP]
			}
		}
	case 0x1000: // 1nnn: jump to address nnn
		c.PC = nnn
	case 0x2000: // 2nnn: call address nnn
		var index = c.SP
		c.Stack[index] = c.PC + 2
		c.SP++
		c.PC = nnn
	case 0x3000: // 3xkk: skip next instr if V[x] = kk
		if c.V[x] == kk {
			c.PC += 2
			break
		}
		c.PC += 2
	case 0x4000: // 4xkk: skip next instr if V[x] != kk
		log.Printf("Skip next instruction if 0x%x != 0x%x\n", c.V[x], kk)
		if c.V[x] != kk {
			c.PC += 2
			break
		}
		c.PC += 2
	case 0x5000: // 5xy0: skip next instr if V[x] == V[y]
		log.Printf("Skip next instruction if 0x%x == 0x%x\n", c.V[x], c.V[y])
		if c.V[x] == c.V[y] {
			c.PC += 2
			break
		}
		c.PC += 2
	case 0x6000: // 6xkk: set V[x] = kk
		log.Printf("Set V[0x%x] to 0x%x\n", x, kk)
		c.V[x] = kk
		c.PC += 2
	case 0x7000: // 7xkk: set V[x] = V[x] + kk
		log.Printf("Set V[0x%d] to V[0x%d] + 0x%x\n", x, x, kk)
		c.V[x] += kk
		c.PC += 2
	case 0x8000: // 8xyn: Arithmetic stuff
		switch n {
		case 0x0:
			log.Printf("V[0x%x] = V[0x%x] = 0x%x\n", x, y, c.V[y])
			c.V[x] = c.V[y]
		case 0x1:
			log.Printf("V[0x%x] |= V[0x%x] = 0x%x\n", x, y, c.V[y])
			c.V[x] = c.V[x] | c.V[y]
		case 0x2:
			log.Printf("V[0x%x] &= V[0x%x] = 0x%x\n", x, y, c.V[y])
			c.V[x] = c.V[x] & c.V[y]
		case 0x3:
			log.Printf("V[0x%x] ^= V[0x%x] = 0x%x\n", x, y, c.V[y])
			c.V[x] = c.V[x] ^ c.V[y]
		case 0x4:
			log.Printf("V[0x%x] = V[0x%x] + V[0x%x] = 0x%x + 0x%x\n", x, x, y, c.V[x], c.V[y])
			if int(c.V[x])+int(c.V[y]) > 255 {
				c.V[0xF] = 1
			} else {
				c.V[0xF] = 0
			}
			c.V[x] = c.V[x] + c.V[y]
		case 0x5:
			log.Printf("V[0x%x] = V[0x%x] - V[0x%x] = 0x%x - 0x%x\n", x, x, y, c.V[x], c.V[y])
			if c.V[x] > c.V[y] {
				c.V[0xF] = 1
			} else {
				c.V[0xF] = 0
			}
			c.V[x] = c.V[x] - c.V[y]
		case 0x6:
			log.Printf("V[0x%x] = V[0x%x] >> 1 = 0x%x >> 1\n", x, x, c.V[x])
			c.V[0xF] = c.V[x] & 0x1
			c.V[x] = (c.V[x] >> 1)
		case 0x7:
			log.Printf("V[0x%x] = V[0x%x] - V[0x%x] = 0x%x - 0x%x\n", x, y, x, c.V[y], c.V[x])
			if c.V[y] > c.V[x] {
				c.V[0xF] = 1
			} else {
				c.V[0xF] = 0
			}
			c.V[x] = c.V[y] - c.V[x]
		case 0xE:
			log.Printf("V[0x%x] = V[0x%x] << 1 = 0x%x << 1\n", x, x, c.V[x])
			c.V[0xF] = (c.V[x] >> 7) & 0x1
			c.V[x] = (c.V[x] << 1)
		}
		c.PC += 2
	case 0x9000: // 9xy0: skip instruction if Vx != Vy
		switch n {
		case 0x0:
			log.Printf("Skip next instruction if 0x%x != 0x%x\n", c.V[x], c.V[y])
			if c.V[x] != c.V[y] {
				c.PC += 2
				break
			}
		}
	case 0xA000: // Annn: set I to address nnn
		log.Printf("Set I to 0x%x\n", nnn)
		c.I = nnn
		c.PC += 2
	case 0xB000: // Bnnn: jump to location nnn + V[0]
		log.Printf("Jump to 0x%x + V[0] (0x%x)\n", nnn, c.V[0])
		c.PC = nnn + uint16(c.V[0])
	case 0xC000: // Cxkk: V[x] = random byte AND kk
		log.Printf("V[0x%x] = random byte\n", x)
		c.V[x] = uint8(rand.Uint64()%256) & kk
		c.PC += 2
	case 0xD000: // Dxyn: Display an n-byte sprite starting at memory
		// location I at (Vx, Vy) on the screen, VF = collision
		log.Printf("Draw sprite at (V[0x%x], V[0x%x]) = (0x%x, 0x%x) of height %d", x, y, c.V[x], c.V[y], n)
		var x_coord = c.V[x] % 64
		var y_coord = c.V[y] % 32
		var original_x = x_coord
		c.V[0xF] = 0
		var i uint8 = 0
		for i < n {
			var sprite_data = c.Memory[uint8(c.I)+i]
			var j uint8 = 7
			for j > 0 {
				x_coord = original_x
				var pixel = &c.Display[y_coord*64+x_coord]
				var pixel_set_uint8 uint8

				if *pixel {
					pixel_set_uint8 = 1
				} else {
					pixel_set_uint8 = 0
				}

				var sprite_bit uint8 = sprite_data & (1 << j)
				var bitset bool
				if sprite_bit == 1 {
					bitset = true
				} else {
					bitset = false
				}
				if bitset && *pixel {
					c.V[0xF] = 1
				}
				pixel_set_uint8 ^= sprite_bit

				if pixel_set_uint8 == 1 {
					*pixel = true
				} else {
					*pixel = false
				}
				if (x_coord + 1) >= 64 {
					break
				}
				j--

			}

			if (y_coord + 1) >= 32 {
				break
			}

			i++
		}
	case 0xE000: // key-pressed events
		switch kk {
		case 0x9E: // skip next instr if key[Vx] is pressed
			log.Printf("Skip next instruction if key[%d] is pressed\n", x)
			//PC += (key[V[x]]) ? 4 : 2;
		case 0xA1: // skip next instr if key[Vx] is not pressed
			log.Printf("Skip next instruction if key[%d] is NOT pressed\n", x)
			//PC += (!key[V[x]]) ? 4 : 2;

		}
	case 0xF000: // misc
		switch kk {
		case 0x07:
			log.Printf("V[0x%x] = delay timer = %d\n", x, c.Delay_timer)
			c.V[x] = c.Delay_timer
			c.PC += 2
		case 0x0A:
			log.Printf("Wait for key instruction\n")
			c.PC += 2
		case 0x15:
			log.Printf("delay timer = V[0x%x] = %d\n", x, c.V[x])
			c.Delay_timer = c.V[x]
		case 0x18:
			log.Printf("sound timer = V[0x%x] = %d\n", x, c.V[x])
			c.Sound_timer = c.V[x]
			c.PC += 2
		case 0x1E:
			log.Printf("I = I + V[0x%x] = 0x%x + 0x%x\n", x, c.I, c.V[x])
			if c.I+uint16(c.V[x]) > uint16(0xfff) {
				c.V[0xF] = 1
			} else {
				c.V[0xF] = 0
			}
			c.I = c.I + uint16(c.V[x])
			c.PC += 2
		case 0x29:
			log.Printf("I = location of font for character V[0x%x] = 0x%x\n", x, c.V[x])
			c.I = 5 * uint16(c.V[x])
		case 0x33:
			log.Printf("Store BCD for %d starting at address 0x%x\n", c.V[x], c.I)
			c.Memory[c.I] = uint8((uint16(c.V[x]) % 1000) / 100) // hundred's digit
			c.Memory[c.I+1] = (c.V[x] % 100) / 10                // ten's digit
			c.Memory[c.I+2] = (c.V[x] % 10)                      // one's digit
			c.PC += 2
		case 0x55:
			log.Printf("Copy sprite from registers 0 to 0x%x into memory at address 0x%x\n", x, c.I)

			for i := 0; i <= int(x); i++ {
				c.Memory[c.I+uint16(i)] = c.V[i]
			}
			c.I += uint16(x) + 1
		case 0x65:
			log.Printf("Copy sprite from memory at address 0x%x into registers 0 to 0x%x\n", x, c.I)
			for i := 0; i <= int(x); i++ {
				c.V[i] = c.Memory[c.I+uint16(i)]
			}
			c.I += uint16(x) + 1
			c.PC += 2
		}
	}

}

func (c *CPU) cycle() {
	if c.Status != enum.Paused {
		insc := c.fetch()
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
		if c.PC <= (c.Rom_size)+0x200 {
			c.cycle()
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				screen.HandleSDLEvents(event, &c.Status)
			}
			screen.Update(renderer, &c.Display)
			sdl.Delay(100)
		}

	}
}
func Init(rom []byte) {
	var cpu = &CPU{}
	cpu.loadFontData()
	cpu.loadROM(rom)
	cpu.run()
}
