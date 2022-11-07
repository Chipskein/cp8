package cpu

type CPU struct {
	rom_size uint16
	Stack    [16]uint16 // (0x200 or 0x600)-0XFFF avaliable Memory for run programms ; 0X000 - 0x1FF avaliable for chip8 interpreter
	Memory   [0xFFF]uint8
	V        [0xF]uint8 //Registers V0-VF
	I        uint16     //used as memory index store
	SP       uint8      //stack pointer
	PC       uint16     //pc counter
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
func (c *CPU) fetch() uint16 {
	var addr1 = c.Memory[c.PC]
	var addr2 = c.Memory[c.PC+1]
	c.PC += 2
	var opcode = uint16(addr1)<<8 | uint16(addr2)
	return opcode
}
func (c *CPU) decode(opcode uint16) {
}
func (c *CPU) execute() {

}
func (c *CPU) cycle() {
	var opcode = c.fetch()
	c.decode(opcode)
	c.execute()
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
