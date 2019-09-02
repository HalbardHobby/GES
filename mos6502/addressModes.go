package mos6502

// Addressing modes
type addressMode func(*CPU) (uint16, bool, uint8)

// ADDRESSING MODES IMPLEMENTATION //
// All addressing modes return the following values:
// address uint16: the absolute address needed for the operation
// impled bool: true if the operation has an implied operand
// cycles uint8: true if the conditions for the "oops" extra cycle have been triggered

func acc(c *CPU) (address uint16, implied bool, cycles uint8) {
	implied = true
	return
}

func abs(c *CPU) (address uint16, implied bool, cycles uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	hi := uint16(c.ReadBus(c.programCounter))
	c.programCounter++

	address = (lo | (hi << 8))
	return
}

func absX(c *CPU) (address uint16, implied bool, cycles uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	hi := uint16(c.ReadBus(c.programCounter))
	c.programCounter++

	address = ((lo | (hi << 8)) + uint16(c.indexX))
	return
}

func absY(c *CPU) (address uint16, implied bool, cycles uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	hi := uint16(c.ReadBus(c.programCounter))
	c.programCounter++

	address = ((lo | (hi << 8)) + uint16(c.indexY))
	return
}

func imm(c *CPU) (address uint16, implied bool, cycles uint8) {
	address = (c.programCounter)
	c.programCounter++
	return
}

func impl(c *CPU) (address uint16, implied bool, cycles uint8) {
	// implicit returns 0 since no operands are required
	implied = true
	return
}

func ind(c *CPU) (address uint16, implied bool, cycles uint8) {
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

	address = (addrLo | (addrHi << 8))
	return
}

func xInd(c *CPU) (address uint16, implied bool, cycles uint8) {
	// Calculate address
	addr := c.ReadBus(c.programCounter) + c.indexX
	c.programCounter++
	// Obtener dirección efectiva
	lo := uint16(c.ReadBus(uint16(addr)))
	hi := uint16(c.ReadBus(uint16(addr + 1)))

	// Retornar valor
	address = lo | (hi << 8)
	return
}

func indY(c *CPU) (address uint16, implied bool, cycles uint8) {
	// Calculate address
	addr := c.ReadBus(c.programCounter)
	c.programCounter++
	// Obtener dirección efectiva
	lo := uint16(c.ReadBus(uint16(addr)))
	hi := uint16(c.ReadBus(uint16(addr + 1)))

	// Retornar valor
	address = (lo | (hi << 8) + uint16(c.indexY))
	return
}

// relative is a special adressing mode used by branching functions
func rel(c *CPU) (address uint16, implied bool, cycles uint8) {
	addr := uint16(c.ReadBus(c.programCounter))

	c.programCounter++
	if addr&0x80 != 0 {
		addr |= 0xFF00
	}

	address = c.programCounter + addr
	return
}

func zpg(c *CPU) (address uint16, implied bool, cycles uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	address = lo & 0x00FF
	return
}

func zpgX(c *CPU) (address uint16, implied bool, cycles uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	address = (lo + uint16(c.indexX)) & 0x00FF
	cycles = 4
	return
}

func zpgY(c *CPU) (address uint16, implied bool, cycles uint8) {
	lo := uint16(c.ReadBus(c.programCounter))
	c.programCounter++
	address = (lo + uint16(c.indexY)) & 0x00FF
	cycles = 4
	return
}
