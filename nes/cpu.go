package nes

// Interrupt types
const (
	_             = iota
	interruptNone // No interruption
	interruptNMI  // Non maskable interrupts
	interruptIRQ  // maskable interrupts
)

// Addressing modes
const (
	_              = iota
	zeroPageIndexX // PEEK((arg + X) % 256)
	zeroPageIndexY // PEEK((arg + Y) % 256)
	absoluteIndexX // PEEK(arg + X)
	absoluteIndexY // PEEK(arg + Y)
	indirectindexX // PEEK(PEEK((arg + X) % 256) + PEEK((arg + X + 1) % 256) * 256)
	indirectIndexY // PEEK(arg + X) % 256
	implicit       // PEEK(PEEK(arg) + PEEK((arg + 1) % 256) * 256 + Y)
	acumulator     // Instructions with implied destination
	immediate      // Instructions that operate on the Accumulator
	zeroPage       // fetches 8-bit value on zero page
	absolute       // fetches value from 16-bit address anywhere in memory
	relative       // Branch inbstructions
	indirect
)

// CPU struct
type CPU struct {
	programCounter uint16          // Program counter
	stackPointer   byte            // Stack Pointer
	accumulator    byte            // Accumulator
	indexX         byte            // Index Register X
	indexY         byte            // Index Register Y
	statusRegister processorStatus // Proccessor status
}

type processorStatus struct {
	carry            bool // Flag set when instruction resulted in overflow
	zero             bool // Flag set if result is 0
	interruptDisable bool // Flag set if system need to ignore interruptions
	decimalMode      bool // Sets the processor to decimal mode, ignored in 2A03
	breakCommand     bool // Flag set if BRK was executed, causing IRQ
	overflow         bool // Set if sign changes because of overflow eg: 64 + 64 = -128
	negative         bool // Set if number is negative
}
