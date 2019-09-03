package mos6502

// OPCODES IMPLEMENTATION //

type opcode func()

// Add memory to accumulator with carry
// A + M + C -> A, C
func adc(c *CPU, operand addressMode) opcode {
	// TODO
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
	// TODO
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
	bcs := func() {
		val, _, _ := operand(c)
		if c.processorStatus.has(carry) {
			c.programCounter = val
		}
	}
	return bcs
}

// branch on result zero
func beq(c *CPU, operand addressMode) opcode {
	beq := func() {
		val, _, _ := operand(c)
		if c.processorStatus.has(zero) {
			c.programCounter = val
		}
	}
	return beq
}

func bit(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Branch on result minus
func bmi(c *CPU, operand addressMode) opcode {
	bmi := func() {
		val, _, _ := operand(c)
		if c.processorStatus.has(negative) {
			c.programCounter = val
		}
	}
	return bmi
}

// Branch on result not zero
func bne(c *CPU, operand addressMode) opcode {
	bne := func() {
		val, _, _ := operand(c)
		if !c.processorStatus.has(zero) {
			c.programCounter = val
		}
	}
	return bne
}

// Branch on result plus
func bpl(c *CPU, operand addressMode) opcode {
	bpl := func() {
		val, _, _ := operand(c)
		if !c.processorStatus.has(negative) {
			c.programCounter = val
		}
	}
	return bpl
}

// Force Break
func brk(c *CPU, operand addressMode) opcode {
	return func() {}
}

// Branch on overflow clear
func bvc(c *CPU, operand addressMode) opcode {
	bvc := func() {
		val, _, _ := operand(c)
		if !c.processorStatus.has(overflow) {
			c.programCounter = val
		}
	}
	return bvc
}

// Branch on overflow set
func bvs(c *CPU, operand addressMode) opcode {
	bvs := func() {
		val, _, _ := operand(c)
		if c.processorStatus.has(overflow) {
			c.programCounter = val
		}
	}
	return bvs
}

// Clear carry flag
func clc(c *CPU, operand addressMode) opcode {
	return func() {
		operand(c)
		c.processorStatus.clear(carry)
	}
}

// Clear decimal mode
func cld(c *CPU, operand addressMode) opcode {
	return func() {
		operand(c)
		c.processorStatus.clear(decimal)
	}
}

// Clear interrupt disable
func cli(c *CPU, operand addressMode) opcode {
	return func() {
		operand(c)
		c.processorStatus.clear(interruptDisable)
	}
}

// Clear overflow flag
func clv(c *CPU, operand addressMode) opcode {
	return func() {
		operand(c)
		c.processorStatus.clear(overflow)
	}
}

// Compare Memory with Accumulator
func cmp(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Compare Memory with Index X
func cpx(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Compare Memory with Index Y
func cpy(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Decrement memory by one
// M - 1 -> M
func dec(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Decrement X by One
// X - 1-> X
func dex(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Decrement Y by One
// Y - 1 -> Y
func dey(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Exclusive-Or Memory with accumulator
// A XOR M -> A
func eor(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Increment Memory by one
// M + 1 -> M
func inc(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Increment X by one
// X + 1 -> X
func inx(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Increment Y by One
// Y + 1 -> Y
func iny(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Jump to new location
func jmp(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

func jsr(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Load Accumulator with Memory
// M -> A
func lda(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Load X with Memory
// M -> X
func ldx(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Load Y with memory
// M -> Y
func ldy(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

func lsr(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

func nop(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Or Memory with Acumulator
// A OR M -> A
func ora(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

func pha(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

func php(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

func pla(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

func plp(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

func rol(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

func ror(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

func rti(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

func rts(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Substract memory from accumulator with Borrow
//  A - M - C -> A
func sbc(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Set Carry flag
// 1 -> C
func sec(c *CPU, operand addressMode) opcode {
	return func() {
		operand(c)
		c.processorStatus.set(carry)
	}
}

// Set Decimal flag
// 1 -> D
func sed(c *CPU, operand addressMode) opcode {
	return func() {
		operand(c)
		c.processorStatus.set(decimal)
	}
}

// Set interrupt disable
// 1 -> I
func sei(c *CPU, operand addressMode) opcode {
	return func() {
		operand(c)
		c.processorStatus.set(interruptDisable)
	}
}

// Store Accumulator to Memory
// A -> M
func sta(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Store X to Memory
// X -> M
func stx(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

// Store Y to Memory
// Y -> M
func sty(c *CPU, operand addressMode) opcode {
	// TODO
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
	// TODO
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
