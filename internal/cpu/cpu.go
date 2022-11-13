package cpu

import (
	"log"
	"math/rand"
	"time"
)

type CPU struct {
	rom_size    uint16
	Stack       [16]uint16 // (0x200 or 0x600)-0XFFF avaliable Memory for run programms ; 0X000 - 0x1FF avaliable for chip8 interpreter
	Memory      [0xFFF]uint8
	V           [0xF + 1]uint8 //Registers V0-VF
	I           uint16         //used as memory index store
	SP          uint8          //stack pointer
	PC          uint16         //pc counter
	delay_timer uint8
	sound_timer uint8
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
func (c *CPU) fetch() (opcode uint16, x uint8, y uint8, n uint8, kk uint8, nnn uint16) {
	var addr1 = c.Memory[c.PC]
	var addr2 = c.Memory[c.PC+1]
	c.PC += 2
	opcode = uint16(addr1)<<8 | uint16(addr2)
	x = uint8((opcode >> 8) & 0x000F) // the lower 4 bits of the high byte
	y = uint8((opcode >> 4) & 0x000F) // the upper 4 bits of the low byte
	n = uint8(opcode & 0x000F)        // the lowest 4 bits
	kk = uint8(opcode & 0x00FF)       // the lowest 8 bits
	nnn = opcode & 0x0FFF             // the lowest 12 bits
	return opcode, x, y, n, kk, nnn
}
func (c *CPU) decodeExec(opcode uint16, x uint8, y uint8, n uint8, kk uint8, nnn uint16) {
	log.Printf("Opcode:%x\n x:%x\n y:%x\n n:%x\n kk:%x\n nnn:%x\n", opcode, x, y, n, kk, nnn)
	switch opcode & 0xF000 {
	case 0x0000:
		switch kk {
		case 0x00E0: // clear the screen
			log.Printf("Clear the screen\n")
			//c.PC += 2
			break
		case 0x00EE: // ret
			log.Printf("return\n")
			c.SP--
			c.PC = c.Stack[c.SP]
			break
		}
		break
	case 0x1000: // 1nnn: jump to address nnn
		log.Printf("Jump to address 0x%x\n", nnn)
		c.PC = nnn
		break
	case 0x2000: // 2nnn: call address nnn
		log.Printf("Call address 0x%x\n", nnn)
		var index = c.SP
		c.Stack[index] = c.PC + 2
		c.SP++
		c.PC = nnn
		break
	case 0x3000: // 3xkk: skip next instr if V[x] = kk
		log.Printf("Skip next instruction if 0x%x == 0x%x\n", c.V[x], kk)
		if c.V[x] == kk {
			c.PC += 2
			break
		}
		//c.PC += 2
		break
	case 0x4000: // 4xkk: skip next instr if V[x] != kk
		log.Printf("Skip next instruction if 0x%x != 0x%x\n", c.V[x], kk)
		if c.V[x] != kk {
			c.PC += 2
			break
		}
		//c.PC += 2
		break
	case 0x5000: // 5xy0: skip next instr if V[x] == V[y]
		log.Printf("Skip next instruction if 0x%x == 0x%x\n", c.V[x], c.V[y])
		if c.V[x] == c.V[y] {
			c.PC += 2
			break
		}
		//c.PC += 2
		break
	case 0x6000: // 6xkk: set V[x] = kk
		log.Printf("Set V[0x%x] to 0x%x\n", x, kk)
		c.V[x] = kk
		//c.PC += 2
		break
	case 0x7000: // 7xkk: set V[x] = V[x] + kk
		log.Printf("Set V[0x%d] to V[0x%d] + 0x%x\n", x, x, kk)
		c.V[x] += kk
		//c.PC += 2
		break
	case 0x8000: // 8xyn: Arithmetic stuff
		switch n {
		case 0x0:
			log.Printf("V[0x%x] = V[0x%x] = 0x%x\n", x, y, c.V[y])
			c.V[x] = c.V[y]
			break
		case 0x1:
			log.Printf("V[0x%x] |= V[0x%x] = 0x%x\n", x, y, c.V[y])
			c.V[x] = c.V[x] | c.V[y]
			break
		case 0x2:
			log.Printf("V[0x%x] &= V[0x%x] = 0x%x\n", x, y, c.V[y])
			c.V[x] = c.V[x] & c.V[y]
			break
		case 0x3:
			log.Printf("V[0x%x] ^= V[0x%x] = 0x%x\n", x, y, c.V[y])
			c.V[x] = c.V[x] ^ c.V[y]
			break
		case 0x4:
			log.Printf("V[0x%x] = V[0x%x] + V[0x%x] = 0x%x + 0x%x\n", x, x, y, c.V[x], c.V[y])
			if int(c.V[x])+int(c.V[y]) > 255 {
				c.V[0xF] = 1
			} else {
				c.V[0xF] = 0
			}
			c.V[x] = c.V[x] + c.V[y]
			break
		case 0x5:
			log.Printf("V[0x%x] = V[0x%x] - V[0x%x] = 0x%x - 0x%x\n", x, x, y, c.V[x], c.V[y])
			if c.V[x] > c.V[y] {
				c.V[0xF] = 1
			} else {
				c.V[0xF] = 0
			}
			c.V[x] = c.V[x] - c.V[y]
			break
		case 0x6:
			log.Printf("V[0x%x] = V[0x%x] >> 1 = 0x%x >> 1\n", x, x, c.V[x])
			c.V[0xF] = c.V[x] & 0x1
			c.V[x] = (c.V[x] >> 1)
			break
		case 0x7:
			log.Printf("V[0x%x] = V[0x%x] - V[0x%x] = 0x%x - 0x%x\n", x, y, x, c.V[y], c.V[x])
			if c.V[y] > c.V[x] {
				c.V[0xF] = 1
			} else {
				c.V[0xF] = 0
			}
			c.V[x] = c.V[y] - c.V[x]
			break
		case 0xE:
			log.Printf("V[0x%x] = V[0x%x] << 1 = 0x%x << 1\n", x, x, c.V[x])
			c.V[0xF] = (c.V[x] >> 7) & 0x1
			c.V[x] = (c.V[x] << 1)
			break
		}
		//PC += 2;
		break
	case 0x9000: // 9xy0: skip instruction if Vx != Vy
		switch n {
		case 0x0:
			log.Printf("Skip next instruction if 0x%x != 0x%x\n", c.V[x], c.V[y])
			if c.V[x] != c.V[y] {
				c.PC += 2
				break
			}
			break
		}
		break
	case 0xA000: // Annn: set I to address nnn
		log.Printf("Set I to 0x%x\n", nnn)
		c.I = nnn
		//PC += 2;
		break
	case 0xB000: // Bnnn: jump to location nnn + V[0]
		log.Printf("Jump to 0x%x + V[0] (0x%x)\n", nnn, c.V[0])
		c.PC = nnn + uint16(c.V[0])
		break
	case 0xC000: // Cxkk: V[x] = random byte AND kk
		log.Printf("V[0x%x] = random byte\n", x)
		c.V[x] = uint8(rand.Uint64()%256) & kk
		//PC += 2;
		break
	case 0xD000: // Dxyn: Display an n-byte sprite starting at memory
		// location I at (Vx, Vy) on the screen, VF = collision
		log.Printf("Draw sprite at (V[0x%x], V[0x%x]) = (0x%x, 0x%x) of height %d", x, y, c.V[x], c.V[y], n)
		//draw_sprite(V[x], V[y], n);
		//PC += 2;
		//chip8_draw_flag = true;
		break
	case 0xE000: // key-pressed events
		switch kk {
		case 0x9E: // skip next instr if key[Vx] is pressed
			log.Printf("Skip next instruction if key[%d] is pressed\n", x)
			//PC += (key[V[x]]) ? 4 : 2;
			break
		case 0xA1: // skip next instr if key[Vx] is not pressed
			log.Printf("Skip next instruction if key[%d] is NOT pressed\n", x)
			//PC += (!key[V[x]]) ? 4 : 2;
			break

		}
		break
	case 0xF000: // misc
		switch kk {
		case 0x07:
			log.Printf("V[0x%x] = delay timer = %d\n", x, c.delay_timer)
			c.V[x] = c.delay_timer
			//PC += 2;
			break
		case 0x0A:
			//printf("Wait for key instruction\n")
			//PC += 2;
			break
		case 0x15:
			log.Printf("delay timer = V[0x%x] = %d\n", x, c.V[x])
			c.delay_timer = c.V[x]
			break
		case 0x18:
			log.Printf("sound timer = V[0x%x] = %d\n", x, c.V[x])
			c.sound_timer = c.V[x]
			//PC += 2;
			break
		case 0x1E:
			log.Printf("I = I + V[0x%x] = 0x%x + 0x%x\n", x, c.I, c.V[x])
			if c.I+uint16(c.V[x]) > uint16(0xfff) {
				c.V[0xF] = 1
			} else {
				c.V[0xF] = 0
			}
			c.I = c.I + uint16(c.V[x])
			//PC += 2;
			break
		case 0x29:
			log.Printf("I = location of font for character V[0x%x] = 0x%x\n", x, c.V[x])
			c.I = 5 * uint16(c.V[x])
			break
		case 0x33:
			log.Printf("Store BCD for %d starting at address 0x%x\n", c.V[x], c.I)
			c.Memory[c.I] = uint8((uint16(c.V[x]) % 1000) / 100) // hundred's digit
			c.Memory[c.I+1] = (c.V[x] % 100) / 10                // ten's digit
			c.Memory[c.I+2] = (c.V[x] % 10)                      // one's digit
			//PC += 2;
			break
		case 0x55:
			log.Printf("Copy sprite from registers 0 to 0x%x into memory at address 0x%x\n", x, c.I)
			//for(i = 0; i <= x; i++) {
			//	c.Memory[c.I + i] = c.V[i];
			//}
			c.I += uint16(x) + 1
			break
		case 0x65:
			log.Printf("Copy sprite from memory at address 0x%x into registers 0 to 0x%x\n", x, c.I)
			//for(i = 0; i <= x; i++) { c.V[i] = c.Memory[c.I + i]; }
			c.I += uint16(x) + 1
			//PC += 2;
			break
		}
		break
	}
}

func (c *CPU) cycle() {
	opcode, x, y, n, kk, nnn := c.fetch()
	c.decodeExec(opcode, x, y, n, kk, nnn)
}
func (c *CPU) loadROM(rom []byte) {
	c.rom_size = uint16(len(rom))
	for addr, Bytes := range rom {
		c.Memory[0x200+addr] = Bytes
	}
	c.PC = 0x200
}
func (c *CPU) run() {
	for {
		if c.PC <= (c.rom_size)+0x200 {
			c.cycle()
			time.Sleep(time.Millisecond * 500)
		}
		if c.rom_size+0x200 < c.PC {
			break
		}
	}
}
func Init(rom []byte) {
	var cpu = &CPU{}
	cpu.loadFontData()
	cpu.loadROM(rom)
	cpu.run()
}
