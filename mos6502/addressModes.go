package mos6502

// Addressing modes
type addressMode func(*CPU) uint8

// ADDRESSING MODES IMPLEMENTATION //
// All addressing modes return the expected value

func acc(c *CPU) (value uint8) {
	return c.accumulator
}

func abs(c *CPU) (value uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	hi := uint16(c.ReadBus(c.programCounter))
	c.programCounter++

	value = c.ReadBus(lo | (hi << 8))
	return
}

func absX(c *CPU) (value uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	hi := uint16(c.ReadBus(c.programCounter))
	c.programCounter++

	value = c.ReadBus((lo | (hi << 8)) + uint16(c.indexX))
	return
}

func absY(c *CPU) (value uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	hi := uint16(c.ReadBus(c.programCounter))
	c.programCounter++

	value = c.ReadBus((lo | (hi << 8)) + uint16(c.indexY))
	return
}

func imm(c *CPU) (value uint8) {
	value = c.ReadBus(c.programCounter)
	c.programCounter++
	return
}

func impl(c *CPU) (value uint8) {
	// implicit returns 0 since no operands are required
	return
}

func ind(c *CPU) (value uint8) {
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

func xInd(c *CPU) (value uint8) {
	// Calculate address
	address := c.ReadBus(c.programCounter) + c.indexX
	c.programCounter++
	// Obtener dirección efectiva
	lo := uint16(c.ReadBus(uint16(address)))
	hi := uint16(c.ReadBus(uint16(address + 1)))

	// Retornar valor
	return c.ReadBus(lo | (hi << 8))
}

func indY(c *CPU) (value uint8) {
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
func rel(c *CPU) (value uint8) {
	value = c.ReadBus(c.programCounter)
	c.programCounter++
	return
}

func zpg(c *CPU) (value uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	return c.ReadBus(lo & 0x00FF)
}

func zpgX(c *CPU) (value uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	return c.ReadBus((lo + uint16(c.indexX)) & 0x00FF)
}

func zpgY(c *CPU) (value uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	return c.ReadBus((lo + uint16(c.indexY)) & 0x00FF)
}
