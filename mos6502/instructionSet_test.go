package mos6502

import (
	"testing"
)

type instructionTest struct {
	name      string
	cpu       *CPU
	flags     flagSet
	testValue uint8
}

// Test address mode
func testOperand(c *CPU) uint8 {
	return _testVal
}

var _testVal uint8

func Test_adc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adc(tt.cpu, testOperand)
		})
	}
}

func Test_and(t *testing.T) {

	tests := []instructionTest{
		{name: "base",
			cpu: &CPU{accumulator: 0x41}, flags: 0, testValue: 0x7B},
		{name: "negative",
			cpu: &CPU{accumulator: 0x8F}, flags: negative, testValue: 0x9F},
		{name: "zero",
			cpu: &CPU{accumulator: 0x41}, flags: zero, testValue: 0x24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := tt.cpu.accumulator & tt.testValue

			_testVal = tt.testValue
			and := and(tt.cpu, testOperand)
			and()

			if tt.cpu.accumulator != expected {
				t.Errorf("and() = %b, wanted %b", tt.cpu.accumulator, expected)
			}
			if tt.cpu.processorStatus != tt.flags {
				t.Errorf("processor status = %b, wanted %b", tt.cpu.processorStatus, tt.flags)
			}
		})
	}
}

func Test_asl(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asl(tt.cpu, testOperand)
		})
	}
}

func Test_bcc(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu:       &CPU{programCounter: 0x8000, processorStatus: carry},
			testValue: 0x20},
		{name: "standard",
			cpu:       &CPU{programCounter: 0x8000},
			testValue: 0x20},
		{name: "negative",
			cpu:       &CPU{programCounter: 0x8000},
			testValue: 0x9F},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := tt.cpu.programCounter

			if !tt.cpu.processorStatus.has(carry) {
				var movement = uint16(tt.testValue)
				if tt.testValue&0x80 != 0 {
					movement |= 0xFF00
				}
				expected += movement
			}

			_testVal = tt.testValue
			bcc := bcc(tt.cpu, testOperand)
			bcc()

			if tt.cpu.programCounter != expected {
				t.Errorf("bcc() = 0x%X, wanted 0x%X", tt.cpu.programCounter, expected)
			}
		})
	}
}

func Test_bcs(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bcs(tt.cpu, testOperand)
		})
	}
}

func Test_beq(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			beq(tt.cpu, testOperand)
		})
	}
}

func Test_bit(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bit(tt.cpu, testOperand)
		})
	}
}

func Test_bmi(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bmi(tt.cpu, testOperand)
		})
	}
}

func Test_bne(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bne(tt.cpu, testOperand)
		})
	}
}

func Test_bpl(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bpl(tt.cpu, testOperand)
		})
	}
}

func Test_brk(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			brk(tt.cpu, testOperand)
		})
	}
}

func Test_bvc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bvc(tt.cpu, testOperand)
		})
	}
}

func Test_bvs(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bvs(tt.cpu, testOperand)
		})
	}
}

func Test_clc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clc(tt.cpu, testOperand)
		})
	}
}

func Test_cld(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cld(tt.cpu, testOperand)
		})
	}
}

func Test_cli(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli(tt.cpu, testOperand)
		})
	}
}

func Test_clv(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clv(tt.cpu, testOperand)
		})
	}
}

func Test_cmp(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmp(tt.cpu, testOperand)
		})
	}
}

func Test_cpx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpx(tt.cpu, testOperand)
		})
	}
}

func Test_cpy(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpy(tt.cpu, testOperand)
		})
	}
}

func Test_dec(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dec(tt.cpu, testOperand)
		})
	}
}

func Test_dex(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dex(tt.cpu, testOperand)
		})
	}
}

func Test_dey(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dey(tt.cpu, testOperand)
		})
	}
}

func Test_eor(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eor(tt.cpu, testOperand)
		})
	}
}

func Test_inc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inc(tt.cpu, testOperand)
		})
	}
}

func Test_inx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inx(tt.cpu, testOperand)
		})
	}
}

func Test_iny(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iny(tt.cpu, testOperand)
		})
	}
}

func Test_jmp(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jmp(tt.cpu, testOperand)
		})
	}
}

func Test_jsr(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsr(tt.cpu, testOperand)
		})
	}
}

func Test_lda(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lda(tt.cpu, testOperand)
		})
	}
}

func Test_ldx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ldx(tt.cpu, testOperand)
		})
	}
}

func Test_ldy(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ldy(tt.cpu, testOperand)
		})
	}
}

func Test_lsr(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lsr(tt.cpu, testOperand)
		})
	}
}

func Test_nop(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nop(tt.cpu, testOperand)
		})
	}
}

func Test_ora(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ora(tt.cpu, testOperand)
		})
	}
}

func Test_pha(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pha(tt.cpu, testOperand)
		})
	}
}

func Test_php(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			php(tt.cpu, testOperand)
		})
	}
}

func Test_pla(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pla(tt.cpu, testOperand)
		})
	}
}

func Test_plp(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plp(tt.cpu, testOperand)
		})
	}
}

func Test_rol(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rol(tt.cpu, testOperand)
		})
	}
}

func Test_ror(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ror(tt.cpu, testOperand)
		})
	}
}

func Test_rti(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rti(tt.cpu, testOperand)
		})
	}
}

func Test_rts(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rts(tt.cpu, testOperand)
		})
	}
}

func Test_sbc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sbc(tt.cpu, testOperand)
		})
	}
}

func Test_sec(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sec(tt.cpu, testOperand)
		})
	}
}

func Test_sed(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sed(tt.cpu, testOperand)
		})
	}
}

func Test_sei(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sei(tt.cpu, testOperand)
		})
	}
}

func Test_sta(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sta(tt.cpu, testOperand)
		})
	}
}

func Test_stx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stx(tt.cpu, testOperand)
		})
	}
}

func Test_sty(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sty(tt.cpu, testOperand)
		})
	}
}

func Test_tax(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tax(tt.cpu, testOperand)
		})
	}
}

func Test_tay(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tay(tt.cpu, testOperand)
		})
	}
}

func Test_tsx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tsx(tt.cpu, testOperand)
		})
	}
}

func Test_txa(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txa(tt.cpu, testOperand)
		})
	}
}

func Test_txs(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txs(tt.cpu, testOperand)
		})
	}
}

func Test_tya(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tya(tt.cpu, testOperand)
		})
	}
}
