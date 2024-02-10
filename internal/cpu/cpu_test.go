package cpu_test

import (
	chip8_cpu "chip8/internal/cpu"
	"testing"
)

var expected_instruction = []uint16{0x00E0, 0xA22A, 0x600C, 0x6108, 0xD01F, 0x7009, 0xA239, 0xD01F, 0xA248, 0x7008, 0xD01F, 0x7004, 0xA257, 0xD01F, 0x7008, 0xA266, 0xD01F, 0x7008, 0xA275, 0xD01F, 0x1228, 0xFF00, 0xFF00, 0x3C00, 0x3C00, 0x3C00, 0x3C00, 0xFF00, 0xFFFF, 0x00FF, 0x0038, 0x003F, 0x003F, 0x0038, 0x00FF, 0x00FF, 0x8000, 0xE000, 0xE000, 0x8000, 0x8000, 0xE000, 0xE000, 0x80F8, 0x00FC, 0x003E, 0x003F, 0x003B, 0x0039, 0x00F8, 0x00F8, 0x0300, 0x0700, 0x0F00, 0xBF00, 0xFB00, 0xF300, 0xE300, 0x43E0, 0x00E0, 0x0080, 0x0080, 0x0080, 0x0080, 0x00E0, 0x00E0}
var ibm_logo_bytes = []uint8{0, 224, 162, 42, 96, 12, 97, 8, 208, 31, 112, 9, 162, 57, 208, 31, 162, 72, 112, 8, 208, 31, 112, 4, 162, 87, 208, 31, 112, 8, 162, 102, 208, 31, 112, 8, 162, 117, 208, 31, 18, 40, 255, 0, 255, 0, 60, 0, 60, 0, 60, 0, 60, 0, 255, 0, 255, 255, 0, 255, 0, 56, 0, 63, 0, 63, 0, 56, 0, 255, 0, 255, 128, 0, 224, 0, 224, 0, 128, 0, 128, 0, 224, 0, 224, 0, 128, 248, 0, 252, 0, 62, 0, 63, 0, 59, 0, 57, 0, 248, 0, 248, 3, 0, 7, 0, 15, 0, 191, 0, 251, 0, 243, 0, 227, 0, 67, 224, 0, 224, 0, 128, 0, 128, 0, 128, 0, 128, 0, 224, 0, 224}

func TestFetchFunction(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	for index, opcode := range ibm_logo_bytes {
		c.Memory[0x200+index] = opcode
	}
	c.PC = 0x200
	var instruction_counter = 0
	var rom_size = uint16(len(ibm_logo_bytes))
	for c.PC < rom_size+0x200 {
		var addr1 = uint16(c.Memory[c.PC])
		var addr2 = uint16(c.Memory[c.PC+1])
		ins := c.Fetch()
		if ins.Opcode != expected_instruction[instruction_counter] {
			t.Fatalf("[FAIL]\nintruction counter:%d\n\naddr1:%d\naddr2:%d\nOpcode: %x!=%x\n", instruction_counter, addr1, addr2, ins.Opcode, expected_instruction[instruction_counter])
		}
		//t.Logf("%x=%x\n", ins.Opcode, expected_instruction[instruction_counter])
		instruction_counter++
	}
}
func TestInstructionClearScreenFunction(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	for index, opcode := range ibm_logo_bytes {
		c.Memory[0x200+index] = opcode
	}
	c.PC = 0x200
	ins := c.Fetch()
	c.DecodeExec(ins)
}
func TestInstructionSetIToNNNunction(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	for index, opcode := range ibm_logo_bytes {
		c.Memory[0x200+index] = opcode
	}
	var initI = c.I
	c.PC = 0x200 + 2
	ins := c.Fetch()
	c.DecodeExec(ins)
	if initI == c.I {
		t.Fatalf("[FAIL] Opcode:%x\n nnn:%x\n", ins.Opcode, ins.Nnn)
	}
}
func TestInstructionSetVxToKKfunction(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	for index, opcode := range ibm_logo_bytes {
		c.Memory[0x200+index] = opcode
	}
	c.PC = 0x200 + 4
	ins := c.Fetch()
	var initVx = c.V[ins.X]
	c.DecodeExec(ins)
	if initVx == c.V[ins.X] || c.V[ins.X] != ins.Kk {
		t.Fatalf("[FAIL] Opcode:%x\n X:%x\n V[%x]=%x\n kk:%x\n", ins.Opcode, ins.X, ins.X, c.V[ins.X], ins.Kk)
	}

	c.PC = 0x200 + 6
	ins = c.Fetch()
	initVx = c.V[ins.X]
	c.DecodeExec(ins)
	if initVx == c.V[ins.X] || c.V[ins.X] != ins.Kk {
		t.Fatalf("[FAIL] Opcode:%x\n X:%x\n V[%x]=%x\n kk:%x\n", ins.Opcode, ins.X, ins.X, c.V[ins.X], ins.Kk)
	}
}
func TestInstructionDrawDxynfunction(t *testing.T) {
	var c = &chip8_cpu.CPU{}

	for index, opcode := range ibm_logo_bytes {
		c.Memory[0x200+index] = opcode
	}
	c.PC = 0x200 + 8
	t.Fatalf("[FAIL] NÃO IMPLEMENTADO\n")
}
