package mos6502

// OPCODES IMPLEMENTATION //

type opcode func()

func adc(c *CPU, operand addressMode) opcode {
	return func() {}
}

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

func bcc(c *CPU, operand addressMode) opcode {
	bcc := func() {
		val, _, _ := operand(c)
		if !c.processorStatus.has(carry) {
			c.programCounter = val
		}
	}
	return bcc
}

func bcs(c *CPU, operand addressMode) opcode {
	return func() {}
}

func beq(c *CPU, operand addressMode) opcode {
	return func() {}
}

func bit(c *CPU, operand addressMode) opcode {
	return func() {}
}

func bmi(c *CPU, operand addressMode) opcode {
	return func() {}
}

func bne(c *CPU, operand addressMode) opcode {
	return func() {}
}

func bpl(c *CPU, operand addressMode) opcode {
	return func() {}
}

func brk(c *CPU, operand addressMode) opcode {
	return func() {}
}

func bvc(c *CPU, operand addressMode) opcode {
	return func() {}
}

func bvs(c *CPU, operand addressMode) opcode {
	return func() {}
}

func clc(c *CPU, operand addressMode) opcode {
	c.processorStatus.clear(carry)
	return func() {}
}

func cld(c *CPU, operand addressMode) opcode {
	c.processorStatus.clear(decimal)
	return func() {}
}

func cli(c *CPU, operand addressMode) opcode {
	c.processorStatus.clear(interruptDisable)
	return func() {}
}

func clv(c *CPU, operand addressMode) opcode {
	c.processorStatus.clear(overflow)
	return func() {}
}

func cmp(c *CPU, operand addressMode) opcode {
	return func() {}
}

func cpx(c *CPU, operand addressMode) opcode {
	return func() {}
}

func cpy(c *CPU, operand addressMode) opcode {
	return func() {}
}

func dec(c *CPU, operand addressMode) opcode {
	return func() {}
}

func dex(c *CPU, operand addressMode) opcode {
	return func() {}
}

func dey(c *CPU, operand addressMode) opcode {
	return func() {}
}

func eor(c *CPU, operand addressMode) opcode {
	return func() {}
}

func inc(c *CPU, operand addressMode) opcode {
	return func() {}
}

func inx(c *CPU, operand addressMode) opcode {
	return func() {}
}

func iny(c *CPU, operand addressMode) opcode {
	return func() {}
}

func jmp(c *CPU, operand addressMode) opcode {
	return func() {}
}

func jsr(c *CPU, operand addressMode) opcode {
	return func() {}
}

func lda(c *CPU, operand addressMode) opcode {
	return func() {}
}

func ldx(c *CPU, operand addressMode) opcode {
	return func() {}
}

func ldy(c *CPU, operand addressMode) opcode {
	return func() {}
}

func lsr(c *CPU, operand addressMode) opcode {
	return func() {}
}

func nop(c *CPU, operand addressMode) opcode {
	return func() {}
}

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

func sbc(c *CPU, operand addressMode) opcode {
	return func() {}
}

func sec(c *CPU, operand addressMode) opcode {
	return func() {}
}

func sed(c *CPU, operand addressMode) opcode {
	return func() {}
}

func sei(c *CPU, operand addressMode) opcode {
	return func() {}
}

func sta(c *CPU, operand addressMode) opcode {
	return func() {}
}

func stx(c *CPU, operand addressMode) opcode {
	return func() {}
}

func sty(c *CPU, operand addressMode) opcode {
	return func() {}
}

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

func tay(c *CPU, operand addressMode) opcode {
	return func() {}
}

func tsx(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

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

func txs(c *CPU, operand addressMode) opcode {
	// TODO
	return func() {}
}

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
