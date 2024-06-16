package cpu

import (
	"cp8/config"
	"cp8/internal/utils"
	"fmt"
	"math/rand"
)

type Instruction struct {
	Opcode uint16
	X      uint16
	Y      uint16
	N      uint16
	KK     uint16
	NNN    uint16
}

// 00E0 - CLS: Clear the display.
func _00E0(c *CPU) {
	c.GraphicsBuffer = [32][64]int{}
	if config.DEBUG {
		fmt.Println("Clearing display")
	}
}

// 00EE - RET: Return from a subroutine. The interpreter sets the program counter to the address at the top of the stack, then subtracts 1 from the stack pointer.
func _00EE(c *CPU) {
	if config.DEBUG {
		fmt.Println("Returning from subroutine")
		fmt.Println("Stack pointer:", c.SP)
		fmt.Println("Stack:", c.Stack)
	}
	c.PC = c.Stack[c.SP]
	c.SP--
}

// 1nnn - JP addr: Jump to location nnn. The interpreter sets the program counter to nnn.
func _1nnn(c *CPU) {
	if config.DEBUG {
		fmt.Println("Current PC: ", c.PC)
		fmt.Println("Current NNN: ", c.CurrentInstruction.NNN)
		fmt.Println("Jumping to location nnn")
	}
	c.PC = int32(c.CurrentInstruction.NNN)
}

// 2nnn - CALL addr: Call subroutine at nnn. The interpreter increments the stack pointer, then puts the current PC on the top of the stack. The PC is then set to nnn.
func _2nnn(c *CPU) {
	if config.DEBUG {
		fmt.Println("Calling subroutine at nnn")
		fmt.Println("Current PC: ", c.PC)
		fmt.Println("Current NNN: ", c.CurrentInstruction.NNN)
		fmt.Println("Stack pointer:", c.SP)
	}
	c.SP++
	c.Stack[c.SP] = c.PC
	c.PC = int32(c.CurrentInstruction.NNN)
}

// 3xkk - SE Vx, byte: Skip next instruction if Vx = kk. The interpreter compares register Vx to kk, and if they are equal, increments the program counter by 2.
func _3xkk(c *CPU) {
	if config.DEBUG {
		fmt.Println("Skip next instruction if Vx = kk")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("kk:", c.CurrentInstruction.KK)
	}
	if c.V[c.CurrentInstruction.X] == int32(c.CurrentInstruction.KK) {
		c.PC += 2
	}
}

// 4xkk - SNE Vx, byte: Skip next instruction if Vx != kk. The interpreter compares register Vx to kk, and if they are not equal, increments the program counter by 2.
func _4xkk(c *CPU) {
	if config.DEBUG {
		fmt.Println("Skip next instruction if Vx != kk")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("kk:", c.CurrentInstruction.KK)
	}
	if c.V[c.CurrentInstruction.X] != int32(c.CurrentInstruction.KK) {
		c.PC += 2
	}
}

// 5xy0 - SE Vx, Vy: Skip next instruction if Vx = Vy. The interpreter compares register Vx to register Vy, and if they are equal, increments the program counter by 2.
func _5xy0(c *CPU) {
	if config.DEBUG {
		fmt.Println("Skip next instruction if Vx = Vy")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("Vy:", c.V[c.CurrentInstruction.Y])
	}
	if c.V[c.CurrentInstruction.X] == c.V[c.CurrentInstruction.Y] {
		c.PC += 2
	}
}

// 6xkk - LD Vx, byte: Set Vx = kk. The interpreter puts the value kk into register Vx.
func _6xkk(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set Vx = kk")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("kk:", c.CurrentInstruction.KK)
	}
	c.V[c.CurrentInstruction.X] = int32(c.CurrentInstruction.KK)
}

// 7xkk - ADD Vx, byte: Set Vx = Vx + kk. Adds the value kk to the value of register Vx, then stores the result in Vx.
func _7xkk(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set Vx = Vx + kk")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("kk:", c.CurrentInstruction.KK)
	}
	c.V[c.CurrentInstruction.X] += int32(c.CurrentInstruction.KK)
}

// 8xy0 - LD Vx, Vy: Set Vx = Vy. Stores the value of register Vy in register Vx.
func _8xy0(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set Vx = Vy")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
	}
	c.V[c.CurrentInstruction.X] = c.V[c.CurrentInstruction.Y]

}

// 8xy1 - OR Vx, Vy: Set Vx = Vx OR Vy. Performs a bitwise OR on the values of Vx and Vy, then stores the result in Vx. A bitwise OR compares the corresponding bits from two values, and if either bit is 1, then the same bit in the result is also 1. Otherwise, it is 0.
func _8xy1(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set Vx = Vx OR Vy")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("Vy:", c.V[c.CurrentInstruction.Y])
	}
	c.V[c.CurrentInstruction.X] |= c.V[c.CurrentInstruction.Y]
}

// 8xy2 - AND Vx, Vy: Set Vx = Vx AND Vy. Performs a bitwise AND on the values of Vx and Vy, then stores the result in Vx. A bitwise AND compares the corresponding bits from two values, and if both bits are 1, then the same bit in the result is also 1. Otherwise, it is 0.
func _8xy2(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set Vx = Vx AND Vy")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("Vy:", c.V[c.CurrentInstruction.Y])
	}
	c.V[c.CurrentInstruction.X] &= c.V[c.CurrentInstruction.Y]
}

// 8xy3 - XOR Vx, Vy: Set Vx = Vx XOR Vy. Performs a bitwise exclusive OR on the values of Vx and Vy, then stores the result in Vx. An exclusive OR compares the corresponding bits from two values, and if the bits are not both the same, then the corresponding bit in the result is set to 1. Otherwise, it is 0.
func _8xy3(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set Vx = Vx XOR Vy")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("Vy:", c.V[c.CurrentInstruction.Y])
	}
	c.V[c.CurrentInstruction.X] ^= c.V[c.CurrentInstruction.Y]
}

// 8xy4 - ADD Vx, Vy: Set Vx = Vx + Vy, set VF = carry. The values of Vx and Vy are added together. If the result is greater than 8 bits (i.e., > 255,) VF is set to 1, otherwise 0. Only the lowest 8 bits of the result are kept, and stored in Vx.
func _8xy4(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set Vx = Vx + Vy")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("Vy:", c.V[c.CurrentInstruction.Y])
	}
	sum := c.V[c.CurrentInstruction.X] + c.V[c.CurrentInstruction.Y]
	if sum > 255 {
		c.V[0xF] = 1
	} else {
		c.V[0xF] = 0
	}
	c.V[c.CurrentInstruction.X] = sum & 0xFF

}

// 8xy5 - SUB Vx, Vy: Set Vx = Vx - Vy, set VF = NOT borrow. If Vx > Vy, then VF is set to 1, otherwise 0. Then Vy is subtracted from Vx, and the results stored in Vx.
func _8xy5(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set Vx = Vx - Vy")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("Vy:", c.V[c.CurrentInstruction.Y])
	}
	if c.V[c.CurrentInstruction.X] > c.V[c.CurrentInstruction.Y] {
		c.V[0xF] = 1
	} else {
		c.V[0xF] = 0
	}
	c.V[c.CurrentInstruction.X] -= c.V[c.CurrentInstruction.Y]

}

// 8xy6 - SHR Vx {, Vy}: Set Vx = Vx SHR 1. If the least-significant bit of Vx is 1, then VF is set to 1, otherwise 0. Then Vx is divided by 2.
func _8xy6(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set Vx = Vx SHR 1")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
	}

	if c.V[c.CurrentInstruction.X]&0x1 == 1 {
		c.V[0xF] = 1
	} else {
		c.V[0xF] = 0
	}
	c.V[c.CurrentInstruction.X] >>= 1

}

// 8xy7 - SUBN Vx, Vy: Set Vx = Vy - Vx, set VF = NOT borrow. If Vy > Vx, then VF is set to 1, otherwise 0. Then Vx is subtracted from Vy, and the results stored in Vx.
func _8xy7(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set Vx = Vy - Vx")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("Vy:", c.V[c.CurrentInstruction.Y])
	}
	if c.V[c.CurrentInstruction.Y] > c.V[c.CurrentInstruction.X] {
		c.V[0xF] = 1
	} else {
		c.V[0xF] = 0
	}
	c.V[c.CurrentInstruction.X] = c.V[c.CurrentInstruction.Y] - c.V[c.CurrentInstruction.X]

}

// 8xyE - SHL Vx {, Vy}: Set Vx = Vx SHL 1. If the most-significant bit of Vx is 1, then VF is set to 1, otherwise to 0. Then Vx is multiplied by 2.
func _8xyE(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set Vx = Vx SHL 1")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
	}
	if c.V[c.CurrentInstruction.X]&0x80 == 1 {
		c.V[0xF] = 1
	} else {
		c.V[0xF] = 0
	}
	c.V[c.CurrentInstruction.X] <<= 1

}

// 9xy0 - SNE Vx, Vy: Skip next instruction if Vx != Vy. The values of Vx and Vy are compared, and if they are not equal, the program counter is increased by 2.
func _9xy0(c *CPU) {
	if config.DEBUG {
		fmt.Println("Skip next instruction if Vx != Vy")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("Vy:", c.V[c.CurrentInstruction.Y])
	}
	if c.V[c.CurrentInstruction.X] != c.V[c.CurrentInstruction.Y] {
		c.PC += 2
	}
}

// Annn - LD I, addr: Set I = nnn. The value of register I is set to nnn.
func _Annn(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set I = nnn")
		fmt.Println("I:", c.I)
		fmt.Println("nnn:", c.CurrentInstruction.NNN)
	}
	c.I = int32(c.CurrentInstruction.NNN)

}

// Bnnn - JP V0, addr: Jump to location nnn + V0. The program counter is set to nnn plus the value of V0.
func _Bnnn(c *CPU) {
	if config.DEBUG {
		fmt.Println("Jump to location nnn + V0")
		fmt.Println("nnn:", c.CurrentInstruction.NNN)
		fmt.Println("V0:", c.V[0])

	}
	c.PC = int32(c.CurrentInstruction.NNN) + c.V[0]
}

// Cxkk - RND Vx, byte: Set Vx = random byte AND kk. The interpreter generates a random number from 0 to 255, which is then ANDed with the value kk. The results are stored in Vx. See instruction 8xy2 for more information on AND.
func _Cxkk(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set Vx = random byte AND kk")
		fmt.Println("x:", c.CurrentInstruction.X)
		fmt.Println("kk:", c.CurrentInstruction.KK)
	}
	r := int32(rand.Intn(255))
	c.V[c.CurrentInstruction.X] = r & int32(c.CurrentInstruction.KK)
}

// Dxyn - DRW Vx, Vy, nibble: Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision. The interpreter reads n bytes from memory, starting at the address stored in I. These bytes are then displayed as sprites on screen at coordinates (Vx, Vy). Sprites are XORed onto the existing screen. If this causes any pixels to be erased, VF is set to 1, otherwise it is set to 0. If the sprite is positioned so part of it is outside the coordinates of the display, it wraps around to the opposite side of the screen. See instruction 8xy3 for more information on XOR, and section 2.4, Display, for more information on the Chip-8 screen and sprites.
func _Dxyn(c *CPU) {
	if config.DEBUG {
		fmt.Println("Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("Vy:", c.V[c.CurrentInstruction.Y])
		fmt.Println("n:", c.CurrentInstruction.N)
	}
	c.V[0xF] = 0 //reset collision flag
	//For each row of the sprite
	for y := uint16(0); y < c.CurrentInstruction.N; y++ {
		spriteBytes := c.Memory[uint16(c.I)+y] //read sprite byte
		for x := uint16(0); x < 8; x++ {
			if (spriteBytes & (0x80 >> x)) != 0 { // Check if pixel should be flipped
				screenX := (int(c.V[c.CurrentInstruction.X]) + int(x)) % 64
				screenY := (int(c.V[c.CurrentInstruction.Y]) + int(y)) % 32
				if c.GraphicsBuffer[screenY][screenX] == 1 {
					c.V[0xF] = 1 // Set VF if pixel was erased
				}
				c.GraphicsBuffer[screenY][screenX] ^= 1 // XOR the pixel
			}
		}
	}

}

// Keyboard
// Ex9E - SKP Vx: Skip next instruction if key with the value of Vx is pressed. Checks the keyboard, and if the key corresponding to the value of Vx is currently in the down position, PC is increased by 2.
func _Ex9E(c *CPU) {
	if config.DEBUG {
		fmt.Println("Skip next instruction if key with the value of Vx is pressed")
	}
	if c.Key == c.V[c.CurrentInstruction.X] {
		c.PC += 2
	}
}

// ExA1 - SKNP Vx: Skip next instruction if key with the value of Vx is not pressed. Checks the keyboard, and if the key corresponding to the value of Vx is currently in the up position, PC is increased by 2.
func _ExA1(c *CPU) {
	if config.DEBUG {
		fmt.Println("Skip next instruction if key with the value of Vx is not pressed")
	}
	if c.Key != c.V[c.CurrentInstruction.X] {
		c.PC += 2
	}
}

// Fx07 - LD Vx, DT: Set Vx = delay timer value. The value of DT is placed into Vx.
func _Fx07(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set Vx = delay timer value")
	}
	c.V[c.CurrentInstruction.X] = c.DT
}

// Fx0A - LD Vx, K: Wait for a key press, store the value of the key in Vx. All execution stops until a key is pressed, then the value of that key is stored in Vx.
func _Fx0A(c *CPU) {
	if config.DEBUG {
		fmt.Println("Wait for a key press, store the value of the key in Vx")
		fmt.Println("Key:", c.Key)
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
	}
	if c.Key == -1 {
		c.PC -= 2
	} else {
		c.V[c.CurrentInstruction.X] = c.Key
	}

}

// Fx15 - LD DT, Vx: Set delay timer = Vx. DT is set equal to the value of Vx.
func _Fx15(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set delay timer = Vx")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("DT:", c.DT)
	}
	c.DT = c.V[c.CurrentInstruction.X]
}

// Fx18 - LD ST, Vx: Set sound timer = Vx. ST is set equal to the value of Vx.
func _Fx18(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set sound timer = Vx")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("ST:", c.ST)
	}
	c.ST = c.V[c.CurrentInstruction.X]
}

// Fx1E - ADD I, Vx: Set I = I + Vx. The values of I and Vx are added, and the results are stored in I.
func _Fx1E(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set I = I + Vx")
		fmt.Println("I:", c.I)
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
	}
	c.I += c.V[c.CurrentInstruction.X]

}

// Fx29 - LD F, Vx: Set I = location of sprite for digit Vx. The value of I is set to the location for the hexadecimal sprite corresponding to the value of Vx. See section 2.4, Display, for more information on the Chip-8 hexadecimal font.
func _Fx29(c *CPU) {
	if config.DEBUG {
		fmt.Println("Set I = location of sprite for digit Vx")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
	}
	c.I = c.V[c.CurrentInstruction.X]

}

// Fx33 - LD B, Vx: Store Binary Code Decimal representation of Vx in memory locations I, I+1, and I+2. The interpreter takes the decimal value of Vx, and places the hundreds digit in memory at location in I, the tens digit at location I+1, and the ones digit at location I+2.
func _Fx33(c *CPU) {
	if config.DEBUG {
		fmt.Println("Store BCD representation of Vx in memory locations I, I+1, and I+2")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
	}
	h, d, o, err := utils.ExtractDecimalHouses(c.V[c.CurrentInstruction.X])
	if err != nil {
		fmt.Println("error extracting decimal houses at Fx33 instruction")
	}
	c.Memory[c.I] = uint16(h)
	c.Memory[c.I+1] = uint16(d)
	c.Memory[c.I+2] = uint16(o)

}

// Fx55 - LD [I], Vx: Store registers V0 through Vx in memory starting at location I. The interpreter copies the values of registers V0 through Vx into memory, starting at the address in I.
func _Fx55(c *CPU) {
	if config.DEBUG {
		fmt.Println("Store registers V0 through Vx in memory starting at location I")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("I:", c.I)
	}
	for i := uint16(0); i <= c.CurrentInstruction.X; i++ {
		c.Memory[uint16(c.I)+i] = uint16(c.V[i])
	}

}

// Fx65 - LD Vx, [I]: Read registers V0 through Vx from memory starting at location I. The interpreter reads values from memory starting at location I into registers V0 through Vx.
func _Fx65(c *CPU) {
	if config.DEBUG {
		fmt.Println("Read registers V0 through Vx from memory starting at location I")
		fmt.Println("Vx:", c.V[c.CurrentInstruction.X])
		fmt.Println("I:", c.I)
	}
	for i := uint16(0); i <= c.CurrentInstruction.X; i++ {
		c.V[i] = int32(c.Memory[uint16(c.I)+i])
	}

}

func Exec(c *CPU) {
	opcode := c.CurrentInstruction.Opcode
	// Mask the opcode with 0xF000 to check the first nibble
	switch opcode & 0xF000 {
	case 0x0000:
		// Further check the last byte
		switch opcode & 0x00FF {
		case 0x00E0:
			_00E0(c)
		case 0x00EE:
			_00EE(c)
		default:
			fmt.Println("Unknown opcode" + fmt.Sprintf("%x", opcode))
		}
	case 0x1000:
		_1nnn(c)
	case 0x2000:
		_2nnn(c)
	case 0x3000:
		_3xkk(c)
	case 0x4000:
		_4xkk(c)
	case 0x5000:
		_5xy0(c)
	case 0x6000:
		_6xkk(c)
	case 0x7000:
		_7xkk(c)
	case 0x8000:
		// Further check the last nibble
		switch opcode & 0x000F {
		case 0x0000:
			_8xy0(c)
		case 0x0001:
			_8xy1(c)
		case 0x0002:
			_8xy2(c)
		case 0x0003:
			_8xy3(c)
		case 0x0004:
			_8xy4(c)
		case 0x0005:
			_8xy5(c)
		case 0x0006:
			_8xy6(c)
		case 0x0007:
			_8xy7(c)
		case 0x000E:
			_8xyE(c)
		default:
			fmt.Println("Unknown opcode" + fmt.Sprintf("%x", opcode))
		}
	case 0x9000:
		_9xy0(c)
	case 0xA000:
		_Annn(c)
	case 0xB000:
		_Bnnn(c)
	case 0xC000:
		_Cxkk(c)
	case 0xD000:
		_Dxyn(c)
	case 0xE000:
		// Further check the last byte
		switch opcode & 0x00FF {
		case 0x009E:
			_Ex9E(c)
		case 0x00A1:
			_ExA1(c)
		default:
			fmt.Println("Unknown opcode" + fmt.Sprintf("%x", opcode))
		}
	case 0xF000:
		// Further check the last byte
		switch opcode & 0x00FF {
		case 0x0007:
			_Fx07(c)
		case 0x000A:
			_Fx0A(c)
		case 0x0015:
			_Fx15(c)
		case 0x0018:
			_Fx18(c)
		case 0x001E:
			_Fx1E(c)
		case 0x0029:
			_Fx29(c)
		case 0x0033:
			_Fx33(c)
		case 0x0055:
			_Fx55(c)
		case 0x0065:
			_Fx65(c)
		default:
			fmt.Println("Unknown opcode" + fmt.Sprintf("%x", opcode))
		}
	default:
		fmt.Println("Unknown opcode" + fmt.Sprintf("%x", opcode))
	}
}
