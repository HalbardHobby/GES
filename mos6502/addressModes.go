package mos6502

// Addressing modes
type addressMode func(*CPU) uint8

// ADDRESSING MODES IMPLEMENTATION //
// All addressing modes return the expected value

func (c *CPU) acc() (value uint8) {
	return c.accumulator
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
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	hi := uint16(c.ReadBus(c.programCounter))
	c.programCounter++

	value = c.ReadBus((lo | (hi << 8)) + uint16(c.indexX))
	return
}

func (c *CPU) absY() (value uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	hi := uint16(c.ReadBus(c.programCounter))
	c.programCounter++

	value = c.ReadBus((lo | (hi << 8)) + uint16(c.indexY))
	return
}

func (c *CPU) imm() (value uint8) {
	value = c.ReadBus(c.programCounter)
	c.programCounter++
	return
}

func (c *CPU) impl() (value uint8) {
	// implicit returns 0 since no operands are required
	return
}

func (c *CPU) ind() (value uint8) {
	// retireve address
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	hi := uint16(c.ReadBus(c.programCounter))
	c.programCounter++

	// Create address
	tempAddr := lo | (hi << 8)
	// Retrieve expected address
	addrLo := uint16(c.ReadBus(tempAddr))
	addrHi := uint16(c.ReadBus(tempAddr + 1))

	return c.ReadBus(addrLo | (addrHi << 8))
}

func (c *CPU) xInd() (value uint8) {
	// Calculate address
	address := c.ReadBus(c.programCounter) + c.indexX
	c.programCounter++
	// Obtener dirección efectiva
	lo := uint16(c.ReadBus(uint16(address)))
	hi := uint16(c.ReadBus(uint16(address + 1)))

	// Retornar valor
	return c.ReadBus(lo | (hi << 8))
}

func (c *CPU) indY() (value uint8) {
	// Calculate address
	address := c.ReadBus(c.programCounter)
	c.programCounter++
	// Obtener dirección efectiva
	lo := uint16(c.ReadBus(uint16(address)))
	hi := uint16(c.ReadBus(uint16(address + 1)))

	// Retornar valor
	return c.ReadBus((lo | (hi << 8)) + uint16(c.indexY))
}

// relative is a special adressing mode used by branching functions
func (c *CPU) rel() (value uint8) {
	value = c.ReadBus(c.programCounter)
	c.programCounter++
	return
}

func (c *CPU) zpg() (value uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	return c.ReadBus(lo & 0x00FF)
}

func (c *CPU) zpgX() (value uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	return c.ReadBus((lo + uint16(c.indexX)) & 0x00FF)
}

func (c *CPU) zpgY() (value uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	return c.ReadBus((lo + uint16(c.indexY)) & 0x00FF)
}
