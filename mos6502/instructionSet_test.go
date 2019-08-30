package mos6502

import (
	"testing"
)

type instructionTest struct {
	name    string
	c       *CPU
	operand addressMode
}

func TestCPU_adc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.adc(tt.operand)
		})
	}
}

func TestCPU_and(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.and(tt.operand)
		})
	}
}

func TestCPU_asl(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.asl(tt.operand)
		})
	}
}

func TestCPU_bcc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.bcc(tt.operand)
		})
	}
}

func TestCPU_bcs(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.bcs(tt.operand)
		})
	}
}

func TestCPU_beq(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.beq(tt.operand)
		})
	}
}

func TestCPU_bit(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.bit(tt.operand)
		})
	}
}

func TestCPU_bmi(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.bmi(tt.operand)
		})
	}
}

func TestCPU_bne(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.bne(tt.operand)
		})
	}
}

func TestCPU_bpl(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.bpl(tt.operand)
		})
	}
}

func TestCPU_brk(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.brk(tt.operand)
		})
	}
}

func TestCPU_bvc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.bvc(tt.operand)
		})
	}
}

func TestCPU_bvs(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.bvs(tt.operand)
		})
	}
}

func TestCPU_clc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.clc(tt.operand)
		})
	}
}

func TestCPU_cld(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.cld(tt.operand)
		})
	}
}

func TestCPU_cli(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.cli(tt.operand)
		})
	}
}

func TestCPU_clv(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.clv(tt.operand)
		})
	}
}

func TestCPU_cmp(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.cmp(tt.operand)
		})
	}
}

func TestCPU_cpx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.cpx(tt.operand)
		})
	}
}

func TestCPU_cpy(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.cpy(tt.operand)
		})
	}
}

func TestCPU_dec(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.dec(tt.operand)
		})
	}
}

func TestCPU_dex(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.dex(tt.operand)
		})
	}
}

func TestCPU_dey(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.dey(tt.operand)
		})
	}
}

func TestCPU_eor(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.eor(tt.operand)
		})
	}
}

func TestCPU_inc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.inc(tt.operand)
		})
	}
}

func TestCPU_inx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.inx(tt.operand)
		})
	}
}

func TestCPU_iny(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.iny(tt.operand)
		})
	}
}

func TestCPU_jmp(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.jmp(tt.operand)
		})
	}
}

func TestCPU_jsr(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.jsr(tt.operand)
		})
	}
}

func TestCPU_lda(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.lda(tt.operand)
		})
	}
}

func TestCPU_ldx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.ldx(tt.operand)
		})
	}
}

func TestCPU_ldy(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.ldy(tt.operand)
		})
	}
}

func TestCPU_lsr(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.lsr(tt.operand)
		})
	}
}

func TestCPU_nop(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.nop(tt.operand)
		})
	}
}

func TestCPU_ora(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.ora(tt.operand)
		})
	}
}

func TestCPU_pha(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.pha(tt.operand)
		})
	}
}

func TestCPU_php(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.php(tt.operand)
		})
	}
}

func TestCPU_pla(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.pla(tt.operand)
		})
	}
}

func TestCPU_plp(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.plp(tt.operand)
		})
	}
}

func TestCPU_rol(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.rol(tt.operand)
		})
	}
}

func TestCPU_ror(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.ror(tt.operand)
		})
	}
}

func TestCPU_rti(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.rti(tt.operand)
		})
	}
}

func TestCPU_rts(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.rts(tt.operand)
		})
	}
}

func TestCPU_sbc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.sbc(tt.operand)
		})
	}
}

func TestCPU_sec(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.sec(tt.operand)
		})
	}
}

func TestCPU_sed(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.sed(tt.operand)
		})
	}
}

func TestCPU_sei(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.sei(tt.operand)
		})
	}
}

func TestCPU_sta(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.sta(tt.operand)
		})
	}
}

func TestCPU_stx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.stx(tt.operand)
		})
	}
}

func TestCPU_sty(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.sty(tt.operand)
		})
	}
}

func TestCPU_tax(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.tax(tt.operand)
		})
	}
}

func TestCPU_tay(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.tay(tt.operand)
		})
	}
}

func TestCPU_tsx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.tsx(tt.operand)
		})
	}
}

func TestCPU_txa(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.txa(tt.operand)
		})
	}
}

func TestCPU_txs(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.txs(tt.operand)
		})
	}
}

func TestCPU_tya(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.tya(tt.operand)
		})
	}
}
