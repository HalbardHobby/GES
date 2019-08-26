package mos6502

// Addressing modes
type addressMode func(*CPU) uint8

// ADDRESSING MODES IMPLEMENTATION //
// All addressing modes return the expected value

func (cpu *CPU) a() (value uint8) {
	return
}

func (cpu *CPU) abs() (value uint8) {
	lo := uint16(cpu.ReadBus(cpu.programCounter))
	cpu.programCounter++
	hi := uint16(cpu.ReadBus(cpu.programCounter))
	cpu.programCounter++

	value = cpu.ReadBus((hi << 8) | lo)
	return
}

func (cpu *CPU) absX() (value uint8) {
	return
}

func (cpu *CPU) absY() (value uint8) {
	return
}

func (cpu *CPU) imm() (value uint8) {
	value = cpu.ReadBus(cpu.programCounter)
	cpu.programCounter++
	return
}

func (cpu *CPU) impl() (value uint8) {
	return
}

func (cpu *CPU) ind() (value uint8) {
	return
}

func (cpu *CPU) xInd() (value uint8) {
	return
}

func (cpu *CPU) indY() (value uint8) {
	return
}

func (cpu *CPU) rel() (value uint8) {
	return
}

func (cpu *CPU) zpg() (value uint8) {
	return
}

func (cpu *CPU) zpgX() (value uint8) {
	return
}

func (cpu *CPU) zpgY() (value uint8) {
	return
}
