package mos6502

// OPCODES IMPLEMENTATION //

func (c *CPU) adc(operand addressMode) {

}

func (c *CPU) and(operand addressMode) {

}

func (c *CPU) asl(operand addressMode) {

}

func (c *CPU) bcc(operand addressMode) {

}

func (c *CPU) bcs(operand addressMode) {

}

func (c *CPU) beq(operand addressMode) {

}

func (c *CPU) bit(operand addressMode) {

}

func (c *CPU) bmi(operand addressMode) {

}

func (c *CPU) bne(operand addressMode) {

}

func (c *CPU) bpl(operand addressMode) {

}

func (c *CPU) brk(operand addressMode) {

}

func (c *CPU) bvc(operand addressMode) {

}

func (c *CPU) bvs(operand addressMode) {

}

func (c *CPU) clc(operand addressMode) {
	c.processorStatus.clear(carry)
}

func (c *CPU) cld(operand addressMode) {
	c.processorStatus.clear(decimal)
}

func (c *CPU) cli(operand addressMode) {
	c.processorStatus.clear(interruptDisable)
}

func (c *CPU) clv(operand addressMode) {
	c.processorStatus.clear(overflow)
}

func (c *CPU) cmp(operand addressMode) {

}

func (c *CPU) cpx(operand addressMode) {

}

func (c *CPU) cpy(operand addressMode) {

}

func (c *CPU) dec(operand addressMode) {

}

func (c *CPU) dex(operand addressMode) {

}

func (c *CPU) dey(operand addressMode) {

}

func (c *CPU) eor(operand addressMode) {

}

func (c *CPU) inc(operand addressMode) {

}

func (c *CPU) inx(operand addressMode) {

}

func (c *CPU) iny(operand addressMode) {

}

func (c *CPU) jmp(operand addressMode) {

}

func (c *CPU) jsr(operand addressMode) {

}

func (c *CPU) lda(operand addressMode) {

}

func (c *CPU) ldx(operand addressMode) {

}

func (c *CPU) ldy(operand addressMode) {

}

func (c *CPU) lsr(operand addressMode) {

}

func (c *CPU) nop(operand addressMode) {
	return
}

func (c *CPU) ora(operand addressMode) {

}

func (c *CPU) pha(operand addressMode) {

}

func (c *CPU) php(operand addressMode) {

}

func (c *CPU) pla(operand addressMode) {

}

func (c *CPU) plp(operand addressMode) {

}

func (c *CPU) rol(operand addressMode) {

}

func (c *CPU) ror(operand addressMode) {

}

func (c *CPU) rti(operand addressMode) {

}

func (c *CPU) rts(operand addressMode) {

}

func (c *CPU) sbc(operand addressMode) {

}

func (c *CPU) sec(operand addressMode) {

}

func (c *CPU) sed(operand addressMode) {

}

func (c *CPU) sei(operand addressMode) {

}

func (c *CPU) sta(operand addressMode) {

}

func (c *CPU) stx(operand addressMode) {

}

func (c *CPU) sty(operand addressMode) {

}

func (c *CPU) tax(operand addressMode) {
	c.indexX = c.accumulator
	if c.indexX == 0x00 {
		c.processorStatus.set(zero)
	}
	if c.indexX&0x80 == 0x80 {
		c.processorStatus.set(negative)
	}
}

func (c *CPU) tay(operand addressMode) {

}

func (c *CPU) tsx(operand addressMode) {

}

func (c *CPU) txa(operand addressMode) {

}

func (c *CPU) txs(operand addressMode) {

}

func (c *CPU) tya(operand addressMode) {

}
