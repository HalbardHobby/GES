package mos6502

// Addressing modes
type addressMode uint8

const (
	unknownIndex   addressMode = iota
	zeroPageIndexX             // PEEK((arg + X) % 256)
	zeroPageIndexY             // PEEK((arg + Y) % 256)
	absoluteIndexX             // PEEK(arg + X)
	absoluteIndexY             // PEEK(arg + Y)
	indirectindexX             // PEEK(PEEK((arg + X) % 256) + PEEK((arg + X + 1) % 256)*256)
	indirectIndexY             // PEEK(arg + X) % 256
	implicit                   // PEEK(PEEK(arg) + PEEK((arg + 1) % 256) * 256 + Y)
	acumulator                 // Instructions with implied destination
	immediate                  // Instructions that operate on the Accumulator
	zeroPage                   // fetches 8-bit value on zero page
	absolute                   // fetches value from 16-bit address anywhere in memory
	relative                   // Branch instructions
	indirect
)

// ADDRESSING MODES IMPLEMENTATION //

func (cpu *CPU) a() {

}

func (cpu *CPU) abs() {

}

func (cpu *CPU) absX() {

}

func (cpu *CPU) absY() {

}

func (cpu *CPU) imm() {

}

func (cpu *CPU) impl() {

}

func (cpu *CPU) ind() {

}

func (cpu *CPU) xInd() {

}

func (cpu *CPU) indY() {

}

func (cpu *CPU) rel() {

}

func (cpu *CPU) zpg() {

}

func (cpu *CPU) zpgX() {

}

func (cpu *CPU) zpgY() {

}
