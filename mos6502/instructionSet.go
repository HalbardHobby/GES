package mos6502

// OPCODES IMPLEMENTATION //

func (cpu *CPU) adc(operand addressMode) {

}

func (cpu *CPU) and(operand addressMode) {

}

func (cpu *CPU) asl(operand addressMode) {

}

func (cpu *CPU) bcc(operand addressMode) {

}

func (cpu *CPU) bcs(operand addressMode) {

}

func (cpu *CPU) beq(operand addressMode) {

}

func (cpu *CPU) bit(operand addressMode) {

}

func (cpu *CPU) bmi(operand addressMode) {

}

func (cpu *CPU) bne(operand addressMode) {

}

func (cpu *CPU) bpl(operand addressMode) {

}

func (cpu *CPU) brk(operand addressMode) {

}

func (cpu *CPU) bvc(operand addressMode) {

}

func (cpu *CPU) bvs(operand addressMode) {

}

func (cpu *CPU) clc(operand addressMode) {
	cpu.processorStatus.clear(carry)
}

func (cpu *CPU) cld(operand addressMode) {
	cpu.processorStatus.clear(decimal)
}

func (cpu *CPU) cli(operand addressMode) {
	cpu.processorStatus.clear(interruptDisable)
}

func (cpu *CPU) clv(operand addressMode) {
	cpu.processorStatus.clear(overflow)
}

func (cpu *CPU) cmp(operand addressMode) {

}

func (cpu *CPU) cpx(operand addressMode) {

}

func (cpu *CPU) cpy(operand addressMode) {

}

func (cpu *CPU) dec(operand addressMode) {

}

func (cpu *CPU) dex(operand addressMode) {

}

func (cpu *CPU) dey(operand addressMode) {

}

func (cpu *CPU) eor(operand addressMode) {

}

func (cpu *CPU) inc(operand addressMode) {

}

func (cpu *CPU) inx(operand addressMode) {

}

func (cpu *CPU) iny(operand addressMode) {

}

func (cpu *CPU) jmp(operand addressMode) {

}

func (cpu *CPU) jsr(operand addressMode) {

}

func (cpu *CPU) lda(operand addressMode) {

}

func (cpu *CPU) ldx(operand addressMode) {

}

func (cpu *CPU) ldy(operand addressMode) {

}

func (cpu *CPU) lsr(operand addressMode) {

}

func (cpu *CPU) nop(operand addressMode) {
	return
}

func (cpu *CPU) ora(operand addressMode) {

}

func (cpu *CPU) pha(operand addressMode) {

}

func (cpu *CPU) php(operand addressMode) {

}

func (cpu *CPU) pla(operand addressMode) {

}

func (cpu *CPU) plp(operand addressMode) {

}

func (cpu *CPU) rol(operand addressMode) {

}

func (cpu *CPU) ror(operand addressMode) {

}

func (cpu *CPU) rti(operand addressMode) {

}

func (cpu *CPU) rts(operand addressMode) {

}

func (cpu *CPU) sbc(operand addressMode) {

}

func (cpu *CPU) sec(operand addressMode) {

}

func (cpu *CPU) sed(operand addressMode) {

}

func (cpu *CPU) sei(operand addressMode) {

}

func (cpu *CPU) sta(operand addressMode) {

}

func (cpu *CPU) stx(operand addressMode) {

}

func (cpu *CPU) sty(operand addressMode) {

}

func (cpu *CPU) tax(operand addressMode) {
	cpu.indexX = cpu.accumulator
	if cpu.indexX == 0x00 {
		cpu.processorStatus.set(zero)
	}
	if cpu.indexX&0x80 == 0x80 {
		cpu.processorStatus.set(negative)
	}
}

func (cpu *CPU) tay(operand addressMode) {

}

func (cpu *CPU) tsx(operand addressMode) {

}

func (cpu *CPU) txa(operand addressMode) {

}

func (cpu *CPU) txs(operand addressMode) {

}

func (cpu *CPU) tya(operand addressMode) {

}
