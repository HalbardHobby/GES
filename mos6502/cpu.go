package mos6502

// Interrupt types
type Interrupt uint8

const (
	unknownInterrupt Interrupt = iota
	interruptNone              // No interruption
	interruptNMI               // Non maskable interrupts
	interruptIRQ               // maskable interrupts
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

func (cpu *CPU) clock() {

}

// Forces the CPU to a known state.
func (cpu *CPU) reset() {

}

func (cpu *CPU) interruptIRQ() {
	if !cpu.processorStatus.has(interruptDisable) {

	}
}

func (cpu *CPU) interruptNMI() {

}
