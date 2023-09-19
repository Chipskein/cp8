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
