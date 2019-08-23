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
	_ = iota
	mode01
	mode02
	mode03
	mode04
	mode05
	mode06
	mode07
	mode08
	mode09
	mode10
	mode11
	mode12
	mode13
)

// CPU struct
type CPU struct {
	pCounter uint16          // Program counter
	sPointer byte            // Stack Pointer
	acc      byte            // Accumulator
	regX     byte            // Index Register X
	regY     byte            // Index Register Y
	status   processorStatus // Proccessor status
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
