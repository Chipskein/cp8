package cpu

func (c *CPU) Cls() {
	//00E0 - CLS
}
func (c *CPU) Ret() {
	//00EE - RET
}
func (c *CPU) Sys(mem_addr uint16) {
	//0nnn - SYS addr
}
func (c *CPU) Jp(mem_addr uint16) {
	//1nnn - JP addr
	//Bnnn - JP V0, addr
}
func (c *CPU) Call(mem_addr uint16) {
	//2nnn - CALL addr
}
func (c *CPU) Se(v_index uint16, byt uint8) {
	//3xkk - SE Vx, byte
	//5xy0 - SE Vx, Vy
}
func (c *CPU) Sne(v_index uint16, byt uint8) {
	//4xkk - SNE Vx, byte
	//9xy0 - SNE Vx, Vy
}
func (c *CPU) Ld(v_index uint16, byt uint8) {
	//6xkk - LD Vx, byte
	//8xy0 - LD Vx, Vy
	//Annn - LD I, addr
	//Fx07 - LD Vx, DT
	//Fx0A - LD Vx, K
	//Fx15 - LD DT, Vx
	//Fx18 - LD ST, Vx
	//Fx29 - LD F, Vx
	//Fx33 - LD B, Vx
	//Fx55 - LD [I], Vx
	//Fx65 - LD Vx, [I]
	//Fx30 - LD HF, Vx
	//Fx75 - LD R, Vx
	//Fx85 - LD Vx, R
}
func (c *CPU) Add(v_index uint16, byt uint8) {
	//7xkk - ADD Vx, byte
	//8xy4 - ADD Vx, Vy
	//Fx1E - ADD I, Vx
}
func (c *CPU) Or(v_index uint16, v_index2 uint16) {
	// 8xy1 - OR Vx, Vy
}
func (c *CPU) And(v_index uint16, v_index2 uint16) {
	//8xy2 - AND Vx, Vy
}
func (c *CPU) Xor(v_index uint16, v_index2 uint16) {
	//8xy3 - XOR Vx, Vy
}
func (c *CPU) Sub(v_index uint16, v_index2 uint16) {
	//8xy5 - SUB Vx, Vy
}
func (c *CPU) Shr(v_index uint16, t interface{}) {
	//8xy6 - SHR Vx {, Vy}
}
func (c *CPU) Subn(v_index uint16, v_index2 uint16) {
	//8xy7 - SUBN Vx, Vy
}
func (c *CPU) Shl(v_index uint16, t interface{}) {
	//8xyE - SHL Vx {, Vy}
}
func (c *CPU) Rnd(v_index uint16, byt uint8) {
	// /   Cxkk - RND Vx, byte
}
func (c *CPU) Drw(v_index uint16, v_index2 uint16, byt uint8) {

	//Dxyn - DRW Vx, Vy, nibble
	//Dxy0 - DRW Vx, Vy, 0
}
func (c *CPU) Skp(v_index uint16) {
	//Ex9E - SKP Vx
}
func (c *CPU) Skpn(v_index uint16) {
	//ExA1 - SKNP Vx
}
func (c *CPU) Scd(nibble uint16) {
	//00Cn - SCD nibble

}
func (c *CPU) Scr() {
	//00FB - SCR

}
func (c *CPU) Scl() {
	//00FC - SCL
}
func (c *CPU) Exit() {
	//00FD - EXIT
}
func (c *CPU) Low() {
	//00FE - LOW
}
func (c *CPU) High() {
	//00FF - HIGH
}
