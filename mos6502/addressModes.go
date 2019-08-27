package mos6502

// Addressing modes
type addressMode func(*CPU) uint8

// ADDRESSING MODES IMPLEMENTATION //
// All addressing modes return the expected value

func (c *CPU) a() (value uint8) {
	return
}

func (c *CPU) abs() (value uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	hi := uint16(c.ReadBus(c.programCounter))
	c.programCounter++

	value = c.ReadBus(lo | (hi << 8))
	return
}

func (c *CPU) absX() (value uint8) {
	return
}

func (c *CPU) absY() (value uint8) {
	return
}

func (c *CPU) imm() (value uint8) {
	value = c.ReadBus(c.programCounter)
	c.programCounter++
	return
}

func (c *CPU) impl() (value uint8) {
	return
}

func (c *CPU) ind() (value uint8) {
	return
}

func (c *CPU) xInd() (value uint8) {
	return
}

func (c *CPU) indY() (value uint8) {
	return
}

func (c *CPU) rel() (value uint8) {
	return
}

func (c *CPU) zpg() (value uint8) {
	return
}

func (c *CPU) zpgX() (value uint8) {
	return
}

func (c *CPU) zpgY() (value uint8) {
	return
}
