package nes

// CPU struct
type CPU struct {
	pCounter uint16 // Program counter
	sPointer byte   // Stack Pointer
	acc      byte   // Accumulator
	regX     byte   // Index Register X
	regY     byte   // Index Register Y
	pStatus  byte   // Proccessor status
}

type processorStatus struct {
	carry            bool
	zero             bool
	interruptDisable bool
	decimalMode      bool
	breakCommand     bool
	overflow         bool
	negative         bool
}
