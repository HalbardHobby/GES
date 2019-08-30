package mos6502

import (
	"testing"
)

type instructionTest struct {
	name        string
	cpu         *CPU
	flags       flagSet
	wantedValue uint8
}

func testOperand(c *CPU) uint8 {
	return 0
}

func TestCPU_adc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.adc(testOperand)
		})
	}
}

func TestCPU_and(t *testing.T) {

	tests := []instructionTest{
		{name: "base", cpu: &CPU{}, flags: 0},
		{name: "negative"},
		{name: "zero"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.and(testOperand)
		})
	}
}

func TestCPU_asl(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.asl(testOperand)
		})
	}
}

func TestCPU_bcc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.bcc(testOperand)
		})
	}
}

func TestCPU_bcs(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.bcs(testOperand)
		})
	}
}

func TestCPU_beq(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.beq(testOperand)
		})
	}
}

func TestCPU_bit(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.bit(testOperand)
		})
	}
}

func TestCPU_bmi(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.bmi(testOperand)
		})
	}
}

func TestCPU_bne(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.bne(testOperand)
		})
	}
}

func TestCPU_bpl(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.bpl(testOperand)
		})
	}
}

func TestCPU_brk(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.brk(testOperand)
		})
	}
}

func TestCPU_bvc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.bvc(testOperand)
		})
	}
}

func TestCPU_bvs(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.bvs(testOperand)
		})
	}
}

func TestCPU_clc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.clc(testOperand)
		})
	}
}

func TestCPU_cld(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.cld(testOperand)
		})
	}
}

func TestCPU_cli(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.cli(testOperand)
		})
	}
}

func TestCPU_clv(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.clv(testOperand)
		})
	}
}

func TestCPU_cmp(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.cmp(testOperand)
		})
	}
}

func TestCPU_cpx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.cpx(testOperand)
		})
	}
}

func TestCPU_cpy(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.cpy(testOperand)
		})
	}
}

func TestCPU_dec(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.dec(testOperand)
		})
	}
}

func TestCPU_dex(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.dex(testOperand)
		})
	}
}

func TestCPU_dey(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.dey(testOperand)
		})
	}
}

func TestCPU_eor(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.eor(testOperand)
		})
	}
}

func TestCPU_inc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.inc(testOperand)
		})
	}
}

func TestCPU_inx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.inx(testOperand)
		})
	}
}

func TestCPU_iny(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.iny(testOperand)
		})
	}
}

func TestCPU_jmp(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.jmp(testOperand)
		})
	}
}

func TestCPU_jsr(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.jsr(testOperand)
		})
	}
}

func TestCPU_lda(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.lda(testOperand)
		})
	}
}

func TestCPU_ldx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.ldx(testOperand)
		})
	}
}

func TestCPU_ldy(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.ldy(testOperand)
		})
	}
}

func TestCPU_lsr(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.lsr(testOperand)
		})
	}
}

func TestCPU_nop(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.nop(testOperand)
		})
	}
}

func TestCPU_ora(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.ora(testOperand)
		})
	}
}

func TestCPU_pha(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.pha(testOperand)
		})
	}
}

func TestCPU_php(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.php(testOperand)
		})
	}
}

func TestCPU_pla(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.pla(testOperand)
		})
	}
}

func TestCPU_plp(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.plp(testOperand)
		})
	}
}

func TestCPU_rol(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.rol(testOperand)
		})
	}
}

func TestCPU_ror(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.ror(testOperand)
		})
	}
}

func TestCPU_rti(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.rti(testOperand)
		})
	}
}

func TestCPU_rts(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.rts(testOperand)
		})
	}
}

func TestCPU_sbc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.sbc(testOperand)
		})
	}
}

func TestCPU_sec(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.sec(testOperand)
		})
	}
}

func TestCPU_sed(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.sed(testOperand)
		})
	}
}

func TestCPU_sei(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.sei(testOperand)
		})
	}
}

func TestCPU_sta(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.sta(testOperand)
		})
	}
}

func TestCPU_stx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.stx(testOperand)
		})
	}
}

func TestCPU_sty(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.sty(testOperand)
		})
	}
}

func TestCPU_tax(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.tax(testOperand)
		})
	}
}

func TestCPU_tay(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.tay(testOperand)
		})
	}
}

func TestCPU_tsx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.tsx(testOperand)
		})
	}
}

func TestCPU_txa(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.txa(testOperand)
		})
	}
}

func TestCPU_txs(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.txs(testOperand)
		})
	}
}

func TestCPU_tya(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.tya(testOperand)
		})
	}
}
