package nes

// Interrupt types
type Interrupt uint8

const (
	unknownInterrupt Interrupt = iota
	interruptNone              // No interruption
	interruptNMI               // Non maskable interrupts
	interruptIRQ               // maskable interrupts
)

// Addressing modes
type AddressMode uint8

const (
	unknownIndex   AddressMode = iota
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

// FlagMask are the flags for processor status
type flagSet byte

const (
	carry            flagSet = 1 << iota // Flag set when instruction resulted in overflow
	zero                                 // Flag set if result is 0
	interruptDisable                     // Flag set if system need to ignore interruptions
	decimalMode                          // Sets the processor to decimal mode, ignored in 2A03
	breakCommand                         // Flag set if BRK was executed, causing IRQ
	_                                    // unused flag
	overflow                             // Set if sign changes because of overflow eg: 64 + 64 = -128
	negative                             // Set if number is negative
)

// CPU struct
type CPU struct {
	programCounter  uint16  // Program counter
	stackPointer    byte    // Stack Pointer
	accumulator     byte    // Accumulator
	indexX          byte    // Index Register X
	indexY          byte    // Index Register Y
	processorStatus flagSet // Proccessor status
}

type instruction struct {
	mode   AddressMode // Addressing mode for the function
	size   byte        // Instruction size
	cycles byte        // Number of cycles used by the instruction
	name   [3]byte     // instruction name
}

func (f *flagSet) set(bits flagSet) {
	*f = *f | bits
}
func (f *flagSet) clear(bits flagSet) {
	*f = *f &^ bits
}
func (f *flagSet) toggle(bits flagSet) {
	*f = *f ^ bits
}
func (f flagSet) has(bits flagSet) bool {
	return f&bits != 0
}
