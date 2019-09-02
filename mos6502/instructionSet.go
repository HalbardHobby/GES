package mos6502

// OPCODES IMPLEMENTATION //

type opcode func()

// Add memory to accumulator with carry
// A + M + C -> A, C
func adc(c *CPU, operand addressMode) opcode {
	return func() {}
}

// AND Memory with accumulator
// A and M -> A
func and(c *CPU, operand addressMode) opcode {
	and := func() {
		address, _, _ := operand(c)
		c.accumulator &= c.ReadBus(address)

		c.processorStatus.setValue(negative, c.accumulator&0x80 != 0)
		c.processorStatus.setValue(zero, c.accumulator == 0)
	}
	return and
}

func asl(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Branch on carry clear
func bcc(c *CPU, operand addressMode) opcode {
	bcc := func() {
		val, _, _ := operand(c)
		if !c.processorStatus.has(carry) {
			c.programCounter = val
		}
	}
	return bcc
}

// branch on carry set
func bcs(c *CPU, operand addressMode) opcode {
	return func() {}
}

// branch on result zero
func beq(c *CPU, operand addressMode) opcode {
	return func() {}
}

func bit(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Branch on result minus
func bmi(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Branch on result not zero
func bne(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Branch on result plus
func bpl(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Force Break
func brk(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Branch on overflow clear
func bvc(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Branch on overflow set
func bvs(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Clear carry flag
func clc(c *CPU, operand addressMode) opcode {
	c.processorStatus.clear(carry)
	return func() {}
}

// Clear decimal mode
func cld(c *CPU, operand addressMode) opcode {
	c.processorStatus.clear(decimal)
	return func() {}
}

// Clear interrupt disable
func cli(c *CPU, operand addressMode) opcode {
	c.processorStatus.clear(interruptDisable)
	return func() {}
}

// Clear overflow flag
func clv(c *CPU, operand addressMode) opcode {
	c.processorStatus.clear(overflow)
	return func() {}
}

// Compare Memory with Accumulator
func cmp(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Compare Memory with Index X
func cpx(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Compare Memory with Index Y
func cpy(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Decrement memory by one
// M - 1 -> M
func dec(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Decrement X by One
// X - 1-> X
func dex(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Decrement Y by One
// Y - 1 -> Y
func dey(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Exclusive-Or Memory with accumulator
// A XOR M -> A
func eor(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Increment Memory by one
// M + 1 -> M
func inc(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Increment X by one
// X + 1 -> X
func inx(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Increment Y by One
// Y + 1 -> Y
func iny(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Jump to new location
func jmp(c *CPU, operand addressMode) opcode {
	return func() {}
}

func jsr(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Load Accumulator with Memory
// M -> A
func lda(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Load X with Memory
// M -> X
func ldx(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Load Y with memory
// M -> Y
func ldy(c *CPU, operand addressMode) opcode {
	return func() {}
}

func lsr(c *CPU, operand addressMode) opcode {
	return func() {}
}

func nop(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Or Memory with Acumulator
// A OR M -> A
func ora(c *CPU, operand addressMode) opcode {
	return func() {}
}

func pha(c *CPU, operand addressMode) opcode {
	return func() {}
}

func php(c *CPU, operand addressMode) opcode {
	return func() {}
}

func pla(c *CPU, operand addressMode) opcode {
	return func() {}
}

func plp(c *CPU, operand addressMode) opcode {
	return func() {}
}

func rol(c *CPU, operand addressMode) opcode {
	return func() {}
}

func ror(c *CPU, operand addressMode) opcode {
	return func() {}
}

func rti(c *CPU, operand addressMode) opcode {
	return func() {}
}

func rts(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Substract memory from accumulator with Borrow
//  A - M - C -> A
func sbc(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Set Carry flag
// 1 -> C
func sec(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Set Decimal flag
// 1 -> D
func sed(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Set interrupt disable
// 1 -> I
func sei(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Store Accumulator to Memory
// A -> M
func sta(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Store X to Memory
// X -> M
func stx(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Store Y to Memory
// Y -> M
func sty(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Transfer Accumulator to X
// A -> X
func tax(c *CPU, operand addressMode) opcode {
	c.indexX = c.accumulator
	if c.indexX == 0x00 {
		c.processorStatus.set(zero)
	}
	if c.indexX&0x80 == 0x80 {
		c.processorStatus.set(negative)
	}
	return func() {}
}

// Transfer Accumulator to Y
// A -> Y
func tay(c *CPU, operand addressMode) opcode {
	return func() {}
}

func tsx(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Transfer X to Accumulator
// X -> A
func txa(c *CPU, operand addressMode) opcode {
	return func() {
		operand(c)
		c.accumulator = c.indexX
		if c.accumulator&0x80 != 0 {
			c.processorStatus.set(negative)
		}
		if c.accumulator == 0 {
			c.processorStatus.set(zero)
		}
	}
}

// X -> ST
func txs(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Transfer Y to Acumulator
// Y -> A
func tya(c *CPU, operand addressMode) opcode {
	return func() {
		operand(c)
		c.accumulator = c.indexY
		if c.accumulator&0x80 != 0 {
			c.processorStatus.set(negative)
		}
		if c.accumulator == 0 {
			c.processorStatus.set(zero)
		}
	}
}
