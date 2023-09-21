package cpu_test

import (
	chip8_cpu "chip8/internal/cpu"
	"log"
	"testing"
)

func TestInstruction_0x0000_CleanScreen(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	for index := range c.Display {
		c.Display[index] = true
	}
	var inst = &chip8_cpu.Instruction{Opcode: 0x0000, Kk: 0x00E0}
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Fail()
	}
	for _, is_pixel_up := range c.Display {
		if is_pixel_up {
			t.Fail()
		}
	}
}
func TestInstruction_0x0000_Return(t *testing.T) {
	var c = &chip8_cpu.CPU{SP: 1}
	var inst = &chip8_cpu.Instruction{Opcode: 0x0000, Kk: 0x00EE}
	c.DecodeExec(inst)
	if c.PC != c.Stack[c.SP] || c.SP != 0 {
		t.Log("c.PC!=c.Stack[c.SP]")
		t.Fail()
	}
	if c.SP != 0 {
		t.Log("c.SP != 0")
		t.Fail()
	}
}
func TestInstruction_0x1000_JumpToNNN(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x1000, Nnn: 0x3C}
	c.DecodeExec(inst)
	if c.PC != inst.Nnn {
		t.Log("c.PC != inst.Nnn")
		t.Fail()
	}
}

func TestInstruction_0x2000_CallAddressNNN(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x2000, Nnn: 0x3C}
	var before = c.PC
	c.DecodeExec(inst)
	if c.PC != inst.Nnn {
		t.Log("c.PC != inst.Nnn")
		t.Fail()
	}
	if c.SP != 1 {
		t.Log("c.SP != 1")
		t.Fail()
	}
	if c.Stack[c.SP-1] != (before + 2) {
		t.Log("c.Stack[c.SP-1] != (before + 2)")
		t.Fail()
	}
}

func TestInstruction_0x3xkk_SkipIfTrue(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x3000, X: 0xF, Kk: 12}
	c.V[inst.X] = inst.Kk
	c.DecodeExec(inst)
	if c.PC != 4 {
		t.Log("c.PC != 4")
		t.Fail()
	}
	c = &chip8_cpu.CPU{}
	inst = &chip8_cpu.Instruction{Opcode: 0x3000, X: 0xF, Kk: 12}
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("c.PC != 2")
		t.Fail()
	}
}
func TestInstruction_0x4xkk_SkipIfFalse(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x4000, X: 0xF, Kk: 12}
	c.DecodeExec(inst)
	if c.PC != 4 {
		t.Log("c.PC != 4")
		t.Fail()
	}
	c = &chip8_cpu.CPU{}
	inst = &chip8_cpu.Instruction{Opcode: 0x4000, X: 0xF, Kk: 12}
	c.V[inst.X] = inst.Kk
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("c.PC != 2")
		t.Fail()
	}
}
func TestInstruction_0x5xy0_SkipIfTrue(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x5000, X: 0xF, Y: 0xA}
	c.V[inst.X] = 0x9
	c.V[inst.X] = c.V[inst.Y]
	c.DecodeExec(inst)
	if c.PC != 4 {
		t.Log("c.PC != 4")
		t.Fail()
	}

	var c2 = &chip8_cpu.CPU{}
	var inst2 = &chip8_cpu.Instruction{Opcode: 0x5000, X: 0xF, Y: 0xA}
	c2.V[inst2.X] = 0x9
	c2.V[inst2.X] = 0xA
	c2.DecodeExec(inst2)
	if c2.PC != 2 {
		t.Log("c.PC != 2")
		t.Fail()
	}
}

func TestInstruction_0x6xkk_SetRegisterXtoKK(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x6000, X: 0xF, Kk: 0xA}
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("c.PC != 2")
		t.Fail()
	}
	if c.V[inst.X] != inst.Kk {
		t.Log("c.V[inst.X] != inst.Kk")
		t.Fail()
	}
}
func TestInstruction_0x7xkk_IncrementRegisterXWithKK(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x7000, X: 0xF, Kk: 0xA}
	var register_value_before = 0x1
	c.V[inst.X] = uint8(register_value_before)
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("c.PC != 2")
		t.Fail()
	}
	if c.V[inst.X] != uint8(register_value_before)+inst.Kk {
		t.Log("c.V[inst.X] != uint8(register_value_before)+inst.Kk")
		t.Fail()
	}

}

func TestInstruction_0x8xyn_ExecMathOpNWithXandY(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	//Op 0x0
	var inst = &chip8_cpu.Instruction{Opcode: 0x8000, X: 0x2, Y: 0x4, N: 0}
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("Op 0x0 c.PC != 2")
		t.Fail()
	}
	if c.V[inst.X] != c.V[inst.Y] {
		t.Log("Op 0x0 c.V[inst.X] != c.V[inst.Y]")
		t.Fail()
	}

	//Op 0x1
	c = &chip8_cpu.CPU{}
	inst = &chip8_cpu.Instruction{Opcode: 0x8000, X: 0x2, Y: 0x4, N: 1}
	c.V[inst.X] = 1
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("Op 0x1 c.PC != 2")
		t.Fail()
	}
	if c.V[inst.X] != (c.V[inst.X] | c.V[inst.Y]) {
		t.Log("Op 0x1 c.V[inst.X] != (c.V[inst.X] | c.V[inst.Y])")
		t.Fail()
	}

	//Op 0x2
	c = &chip8_cpu.CPU{}
	inst = &chip8_cpu.Instruction{Opcode: 0x8000, X: 0x2, Y: 0x4, N: 2}
	c.V[inst.X] = 1
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("Op 0x2 c.PC != 2")
		t.Fail()
	}
	if c.V[inst.X] != (c.V[inst.X] & c.V[inst.Y]) {
		t.Log("Op 0x2 c.V[inst.X] != (c.V[inst.X] & c.V[inst.Y])")
		t.Fail()
	}
	//Op 0x3
	c = &chip8_cpu.CPU{}
	inst = &chip8_cpu.Instruction{Opcode: 0x8000, X: 0x2, Y: 0x4, N: 3}
	c.V[inst.X] = 1
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("Op 0x3 c.PC != 2")
		t.Fail()
	}
	if c.V[inst.X] != (c.V[inst.X] ^ c.V[inst.Y]) {
		t.Log("Op 0x3 c.V[inst.X] != (c.V[inst.X]^c.V[inst.Y])")
		t.Fail()
	}
	//Op 0x4
	c = &chip8_cpu.CPU{}
	inst = &chip8_cpu.Instruction{Opcode: 0x8000, X: 0x2, Y: 0x4, N: 4}
	var value_in_X = 10
	var value_in_Y = 20
	c.V[inst.X] = uint8(value_in_X)
	c.V[inst.Y] = uint8(value_in_Y)
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("Op 0x4 c.PC != 2")
		t.Fail()
	}
	if c.V[0xF] == 1 {
		t.Log("Op 0x4 c.V[0xF] == 1")
		t.Fail()
	}
	if c.V[inst.X] != (uint8(value_in_X) + uint8(value_in_Y)) {
		t.Log("Op 0x4 c.V[inst.X] != (uint8(value_in_X)+uint8(value_in_Y))")
		t.Fail()
	}
	//Op 0x5
	c = &chip8_cpu.CPU{}
	inst = &chip8_cpu.Instruction{Opcode: 0x8000, X: 0x2, Y: 0x4, N: 5}
	value_in_X = 10
	value_in_Y = 20
	c.V[inst.X] = uint8(value_in_X)
	c.V[inst.Y] = uint8(value_in_Y)
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("Op 0x5 c.PC != 2")
		t.Fail()
	}
	if c.V[0xF] == 1 {
		t.Log("Op 0x5 c.V[0xF] == 1")
		t.Fail()
	}
	if c.V[inst.X] != (uint8(value_in_X) - uint8(value_in_Y)) {
		t.Log("Op 0x5 c.V[inst.X] != (uint8(value_in_X)-uint8(value_in_Y))")
		t.Fail()
	}
	//Op 0x6
	c = &chip8_cpu.CPU{}
	inst = &chip8_cpu.Instruction{Opcode: 0x8000, X: 0x2, Y: 0x4, N: 6}
	value_in_X = 10
	c.V[inst.X] = uint8(value_in_X)
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("Op 0x6 c.PC != 2")
		t.Fail()
	}
	if c.V[0xF] != (uint8(value_in_X) & 0x1) {
		t.Log("Op 0x6 c.V[0xF] != (uint8(value_in_X) & 0x1)")
		t.Fail()
	}
	if c.V[inst.X] != (uint8(value_in_X) >> 0x1) {
		t.Log("Op 0x6 c.V[inst.X] != (uint8(value_in_X) >> 0x1)")
		t.Fail()
	}
	//Op 0x7
	c = &chip8_cpu.CPU{}
	inst = &chip8_cpu.Instruction{Opcode: 0x8000, X: 0x2, Y: 0x4, N: 7}
	value_in_X = 10
	value_in_Y = 20
	c.V[inst.X] = uint8(value_in_X)
	c.V[inst.Y] = uint8(value_in_Y)
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("Op 0x7 c.PC != 2")
		t.Fail()
	}
	if c.V[0xF] == 0 {
		t.Log("Op 0x7 c.V[0xF] == 0")
		t.Fail()
	}
	if c.V[inst.X] != (uint8(value_in_Y) - uint8(value_in_X)) {
		t.Log("Op 0x7 c.V[inst.X] != (uint8(value_in_Y) - uint8(value_in_X))")
		t.Fail()
	}

	//Op 0xE
	c = &chip8_cpu.CPU{}
	inst = &chip8_cpu.Instruction{Opcode: 0x8000, X: 0x2, Y: 0x4, N: 0xE}
	value_in_X = 10
	c.V[inst.X] = uint8(value_in_X)
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("0xE c.PC != 2")
		t.Fail()
	}
	if c.V[0xF] != ((uint8(value_in_X) >> 7) & 0x1) {
		t.Log("0xE c.V[0xF] != ((uint8(value_in_X) >> 7) & 0x1)")
		t.Fail()
	}
	if c.V[inst.X] != (uint8(value_in_X) << 1) {
		t.Log("0xE c.V[inst.X] != (uint8(value_in_X) << 1)")
		t.Fail()
	}

}

func TestInstruction_0x9000_SkipIfXDifferentThanY(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x9000, X: 0xD, Y: 0x1}
	c.V[inst.X] = inst.X
	c.V[inst.Y] = inst.Y
	c.DecodeExec(inst)
	if c.PC != 4 {
		t.Log("c.PC != 4")
		t.Fail()
	}
	c = &chip8_cpu.CPU{}
	inst = &chip8_cpu.Instruction{Opcode: 0x9000, X: 0xD, Y: 0xD}
	c.V[inst.X] = inst.X
	c.V[inst.Y] = inst.Y
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("c.PC != 2")
		t.Fail()
	}
}
func TestInstruction_0xAnnn_SetMemoryIndexToNNN(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0xA000, Nnn: 0x204}
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("c.PC != 2")
		t.Fail()
	}
	if c.I != inst.Nnn {
		t.Log("c.I != inst.Nnn")
		t.Fail()
	}
}
func TestInstruction_0xBnnn_JumpToNNNPlusV0(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0xB000, Nnn: 0x204}
	c.V[0] = 12
	c.DecodeExec(inst)
	if c.PC != inst.Nnn+uint16(c.V[0]) {
		t.Log("c.PC != inst.Nnn+uint16(c.V[0])")
		t.Fail()
	}
}
func TestInstruction_0xCxkk_SetXRandomByteAndKK(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0xC000, X: 0xC, Kk: 0xd}
	c.V[inst.X] = 0
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("c.PC != 2")
		t.Fail()
	}
	if c.V[inst.X] == 0 {
		t.Log("c.V[inst.X] == 0")
		t.Fail()
	}
}

func TestInstruction_0xDxyn_DisplayAnNByteAtMemory(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0xD000, X: 0x4, Y: 0xA, N: 0xA}
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Log("c.PC != 2")
		t.Fail()
	}
	pixel_on_counter := 0
	for pixel_index, pixel := range c.Display {
		if pixel {
			pixel_on_counter++
			log.Printf("Pixel %d On\n", pixel_index)
		}
	}
	if pixel_on_counter == 0 {
		t.Log("pixel_on_counter == 0")
		t.Fail()
	}
}
