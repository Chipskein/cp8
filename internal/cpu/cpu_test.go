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
	if c.PC != 2 {
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
	if c.PC != 2 {
		t.Failed()
	}
}
