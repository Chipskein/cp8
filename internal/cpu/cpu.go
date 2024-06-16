package cpu

import (
	"chip8/config"
	"fmt"
	"os"
)

type CPU struct {
	StartPC            int32 //512 at most programs ,but 1536 in ETI 660 Chip8-Programs
	GraphicsBuffer     [32][64]int
	I                  int32
	V                  [16]int32
	Stack              [16]int32
	SP                 int32
	PC                 int32
	Memory             [4096]uint16
	RomSize            int
	CurrentInstruction *Instruction
	Key                int32
	DT                 int32
	ST                 int32
}

func (c *CPU) LoadFont() {
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
		c.Memory[addr] = uint16(b)
	}

}
func (c *CPU) LoadROM(filePath string) (int, error) {
	var start = int(c.StartPC)
	romBytes, err := os.ReadFile(filePath)
	if err != nil {
		return 0, fmt.Errorf("error loading rom %s:%s", filePath, err.Error())
	}
	for i, b := range romBytes {
		c.Memory[start+i] = uint16(b)
		if config.DEBUG {
			fmt.Printf("Memory[%d]=%x\n", start+i, c.Memory[start+i])
		}
	}
	return len(romBytes), nil

}
func (c *CPU) Fetch() {
	var addr1 = c.Memory[c.PC]
	var addr2 = c.Memory[c.PC+1]
	c.PC += 2
	var opcode = addr1<<8 | addr2
	var x = (opcode >> 8) & 0x0F
	var y = (opcode >> 4) & 0x00F // the upper 4 bits of the low byte
	var n = opcode & 0x000F       // the lowest 4 bits
	var kk = opcode & 0x00FF      // the lowest 8 bits
	var nnn = opcode & 0x0FFF
	if config.DEBUG {
		fmt.Printf("############################################\n")
		fmt.Printf("Opcode:%x\n", opcode)
		fmt.Printf("	x:%x\n", x)
		fmt.Printf("	y:%x\n", y)
		fmt.Printf("	n:%x\n", n)
		fmt.Printf("	kk:%x\n", kk)
		fmt.Printf("	nnn:%x\n", nnn)
		fmt.Printf("############################################\n")
	}
	c.CurrentInstruction = &Instruction{
		Opcode: opcode,
		X:      x,
		Y:      y,
		N:      n,
		KK:     kk,
		NNN:    nnn,
	}
}

func (c *CPU) Cycle() {
	c.Fetch()
	Exec(c)
}
func NewCPU(filePath string) (*CPU, error) {
	//Setup Flag Case is a ETI 660 Chip8
	var c = CPU{
		StartPC: 512,
		Key:     -1,
	}
	c.PC = c.StartPC
	c.LoadFont()
	size, err := c.LoadROM(filePath)
	if err != nil {
		return nil, fmt.Errorf("error loading rom %s:%s", filePath, err.Error())
	}
	c.RomSize = size
	return &c, nil
}
