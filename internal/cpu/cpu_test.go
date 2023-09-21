package cpu_test

import (
	chip8_cpu "chip8/internal/cpu"
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
		t.Failed()
	}
	for _, is_pixel_up := range c.Display {
		if is_pixel_up {
			t.Failed()
		}
	}
}
func TestInstruction_0x0000_Return(t *testing.T) {
	var c = &chip8_cpu.CPU{SP: 1}
	var inst = &chip8_cpu.Instruction{Opcode: 0x0000, Kk: 0x00EE}
	c.DecodeExec(inst)
	if c.PC != c.Stack[c.SP] || c.SP != 0 {
		t.Failed()
	}
}
func TestInstruction_0x1000_JumpToNNN(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x1000, Nnn: 0x3C}
	c.DecodeExec(inst)
	if c.PC != inst.Nnn {
		t.Failed()
	}
}

func TestInstruction_0x2000_CallAddressNNN(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x2000, Nnn: 0x3C}
	var before = c.PC
	c.DecodeExec(inst)
	if c.PC != inst.Nnn || c.SP != 1 || c.Stack[c.SP-1] != (before+2) {
		t.Failed()
	}
}

func TestInstruction_0x3xkk_SkipIfTrue(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x3000, X: 0xF, Kk: 12}
	c.V[inst.X] = inst.Kk
	c.DecodeExec(inst)
	if c.PC != 4 {
		t.Failed()
	}

	var c2 = &chip8_cpu.CPU{}
	var inst2 = &chip8_cpu.Instruction{Opcode: 0x3000, X: 0xF, Kk: 12}
	c2.DecodeExec(inst2)
	if c2.PC != 2 {
		t.Failed()
	}
}
func TestInstruction_0x4xkk_SkipIfFalse(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x4000, X: 0xF, Kk: 12}
	c.DecodeExec(inst)
	if c.PC != 4 {
		t.Failed()
	}
	var c2 = &chip8_cpu.CPU{}
	var inst2 = &chip8_cpu.Instruction{Opcode: 0x4000, X: 0xF, Kk: 12}
	c2.V[inst.X] = inst2.Kk
	c2.DecodeExec(inst2)
	if c2.PC != 2 {
		t.Failed()
	}
}
func TestInstruction_0x5xy0_SkipIfTrue(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x5000, X: 0xF, Y: 0xA}
	c.V[inst.X] = 0x9
	c.V[inst.X] = c.V[inst.Y]
	c.DecodeExec(inst)
	if c.PC != 4 {
		t.Failed()
	}

	var c2 = &chip8_cpu.CPU{}
	var inst2 = &chip8_cpu.Instruction{Opcode: 0x5000, X: 0xF, Y: 0xA}
	c2.V[inst2.X] = 0x9
	c2.V[inst2.X] = 0xA
	c2.DecodeExec(inst2)
	if c2.PC != 2 {
		t.Failed()
	}
}

func TestInstruction_0x6xkk_SetRegisterXtoKK(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x6000, X: 0xF, Kk: 0xA}
	c.DecodeExec(inst)
	if c.PC != 2 || c.V[inst.X] != inst.Kk {
		t.Failed()
	}
}
func TestInstruction_0x7xkk_IncrementRegisterXWithKK(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x7000, X: 0xF, Kk: 0xA}
	var register_value_before = 0x1
	c.V[inst.X] = uint8(register_value_before)
	c.DecodeExec(inst)
	if c.PC != 2 || c.V[inst.X] != uint8(register_value_before)+inst.Kk {
		t.Failed()
	}
}

/*
func TestInstruction_0x8xyn_DoNWithXandY(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x8000, X: 0xD, Y: 0x1}
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Failed()
	}
}
*/

func TestInstruction_0x9000_SkipIfXDifferentThanY(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0x9000, X: 0xD, Y: 0x1}
	c.DecodeExec(inst)
	if c.PC != 4 {
		t.Failed()
	}
	var c2 = &chip8_cpu.CPU{}
	var inst2 = &chip8_cpu.Instruction{Opcode: 0x9000, X: 0xD, Y: 0xD}
	c2.DecodeExec(inst2)
	if c2.PC != 2 {
		t.Failed()
	}
}
func TestInstruction_0xAnnn_SetMemoryIndexToNNN(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0xA000, Nnn: 0x204}
	c.DecodeExec(inst)
	if c.PC != 2 || c.I != inst.Nnn {
		t.Failed()
	}
}
func TestInstruction_0xBnnn_SetMemoryIndexToNNN(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0xB000, Nnn: 0x204}
	c.DecodeExec(inst)
	if c.PC != 2 {
		t.Failed()
	}
}
func TestInstruction_0xCxkk_SetXRandomByteAndKK(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var inst = &chip8_cpu.Instruction{Opcode: 0xC000, X: 0xC, Kk: 0xd}
	c.V[inst.X] = 0
	c.DecodeExec(inst)
	if c.PC != 2 || c.V[inst.X] == 0 {
		t.Failed()
	}
}
