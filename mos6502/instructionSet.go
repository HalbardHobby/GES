package mos6502

// OPCODES IMPLEMENTATION //

func (c *cpu) adc(operand addressMode) {

}

func (c *cpu) and(operand addressMode) {

}

func (c *cpu) asl(operand addressMode) {

}

func (c *cpu) bcc(operand addressMode) {

}

func (c *cpu) bcs(operand addressMode) {

}

func (c *cpu) beq(operand addressMode) {

}

func (c *cpu) bit(operand addressMode) {

}

func (c *cpu) bmi(operand addressMode) {

}

func (c *cpu) bne(operand addressMode) {

}

func (c *cpu) bpl(operand addressMode) {

}

func (c *cpu) brk(operand addressMode) {

}

func (c *cpu) bvc(operand addressMode) {

}

func (c *cpu) bvs(operand addressMode) {

}

func (c *cpu) clc(operand addressMode) {
	c.processorStatus.clear(carry)
}

func (c *cpu) cld(operand addressMode) {
	c.processorStatus.clear(decimal)
}

func (c *cpu) cli(operand addressMode) {
	c.processorStatus.clear(interruptDisable)
}

func (c *cpu) clv(operand addressMode) {
	c.processorStatus.clear(overflow)
}

func (c *cpu) cmp(operand addressMode) {

}

func (c *cpu) cpx(operand addressMode) {

}

func (c *cpu) cpy(operand addressMode) {

}

func (c *cpu) dec(operand addressMode) {

}

func (c *cpu) dex(operand addressMode) {

}

func (c *cpu) dey(operand addressMode) {

}

func (c *cpu) eor(operand addressMode) {

}

func (c *cpu) inc(operand addressMode) {

}

func (c *cpu) inx(operand addressMode) {

}

func (c *cpu) iny(operand addressMode) {

}

func (c *cpu) jmp(operand addressMode) {

}

func (c *cpu) jsr(operand addressMode) {

}

func (c *cpu) lda(operand addressMode) {

}

func (c *cpu) ldx(operand addressMode) {

}

func (c *cpu) ldy(operand addressMode) {

}

func (c *cpu) lsr(operand addressMode) {

}

func (c *cpu) nop(operand addressMode) {
	return
}

func (c *cpu) ora(operand addressMode) {

}

func (c *cpu) pha(operand addressMode) {

}

func (c *cpu) php(operand addressMode) {

}

func (c *cpu) pla(operand addressMode) {

}

func (c *cpu) plp(operand addressMode) {

}

func (c *cpu) rol(operand addressMode) {

}

func (c *cpu) ror(operand addressMode) {

}

func (c *cpu) rti(operand addressMode) {

}

func (c *cpu) rts(operand addressMode) {

}

func (c *cpu) sbc(operand addressMode) {

}

func (c *cpu) sec(operand addressMode) {

}

func (c *cpu) sed(operand addressMode) {

}

func (c *cpu) sei(operand addressMode) {

}

func (c *cpu) sta(operand addressMode) {

}

func (c *cpu) stx(operand addressMode) {

}

func (c *cpu) sty(operand addressMode) {

}

func (c *cpu) tax(operand addressMode) {
	c.indexX = c.accumulator
	if c.indexX == 0x00 {
		c.processorStatus.set(zero)
	}
	if c.indexX&0x80 == 0x80 {
		c.processorStatus.set(negative)
	}
}

func (c *cpu) tay(operand addressMode) {

}

func (c *cpu) tsx(operand addressMode) {

}

func (c *cpu) txa(operand addressMode) {

}

func (c *cpu) txs(operand addressMode) {

}

func (c *cpu) tya(operand addressMode) {

}
