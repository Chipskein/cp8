package cpu_test

import (
	chip8_cpu "chip8/internal/cpu"
	"log"
	"testing"
)

func TestFetchFunction(t *testing.T) {
	var c = &chip8_cpu.CPU{}
	var bytes = []uint8{0, 224, 162, 42, 96, 12}
	for index, opcode := range bytes {
		c.Memory[0x200+index] = opcode
	}
	c.PC = 0x200
	ins := c.Fetch()
	log.Printf("%x", ins)

	ins = c.Fetch()
	log.Printf("%x", ins)

	ins = c.Fetch()
	log.Printf("%x", ins)
}
