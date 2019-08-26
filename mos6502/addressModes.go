package mos6502

// Addressing modes
type addressMode func(*cpu) uint8

// ADDRESSING MODES IMPLEMENTATION //
// All addressing modes return the expected value

func (c *cpu) a() (value uint8) {
	return
}

func (c *cpu) abs() (value uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	hi := uint16(c.ReadBus(c.programCounter))
	c.programCounter++

	value = c.ReadBus((hi << 8) | lo)
	return
}

func (c *cpu) absX() (value uint8) {
	return
}

func (c *cpu) absY() (value uint8) {
	return
}

func (c *cpu) imm() (value uint8) {
	value = c.ReadBus(c.programCounter)
	c.programCounter++
	return
}

func (c *cpu) impl() (value uint8) {
	return
}

func (c *cpu) ind() (value uint8) {
	return
}

func (c *cpu) xInd() (value uint8) {
	return
}

func (c *cpu) indY() (value uint8) {
	return
}

func (c *cpu) rel() (value uint8) {
	return
}

func (c *cpu) zpg() (value uint8) {
	return
}

func (c *cpu) zpgX() (value uint8) {
	return
}

func (c *cpu) zpgY() (value uint8) {
	return
}
