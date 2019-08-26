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
	decimal                              // Sets the processor to decimal mode, ignored in 2A03
	breakCommand                         // Flag set if BRK was executed, causing IRQ
	unused                               // unused flag
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
	opcode          byte    // current operation
	processorStatus flagSet // Proccessor status

	ReadBus  func(uint16) uint8  // 16-bit bus read chanel
	WriteBus func(uint16, uint8) // 16-bit write bus connection
}

var processor *CPU

// GetCPU returns the global processor instance. Should be preferred over
func GetCPU() *CPU {
	if processor == nil {
		processor = new(CPU)
	}
	return processor
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

func (c *CPU) clock() {
	c.opcode = c.ReadBus(c.programCounter)
	c.programCounter++

}

// Forces the CPU to a known state.
func (c *CPU) reset() {

}

func (c *CPU) interruptIRQ() {
	if !c.processorStatus.has(interruptDisable) {

	}
}

func (c *CPU) interruptNMI() {

}
