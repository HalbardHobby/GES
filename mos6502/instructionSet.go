package mos6502

type instruction struct {
	mode   func() // Addressing mode for the function
	opcode func() // opcode to execute
	size   byte   // Instruction size
	cycles byte   // Number of cycles used by the instruction
	name   string // instruction name
}
