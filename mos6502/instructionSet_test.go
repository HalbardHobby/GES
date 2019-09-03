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

var _testAddr uint16

func (tt *instructionTest) setupInstructionTest() {
	read, write := getTestMemory()
	tt.cpu.ReadBus = read
	tt.cpu.WriteBus = write
	return
}

func Test_adc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adc(tt.cpu, imm)
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
			tt.setupInstructionTest()

			and := and(tt.cpu, imm)
			tt.cpu.WriteBus(tt.cpu.programCounter, tt.testValue)
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
			asl(tt.cpu, imm)
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
			expected := tt.cpu.programCounter + 1
			tt.setupInstructionTest()
			if !tt.cpu.processorStatus.has(carry) {
				var movement = uint16(tt.testValue)
				if tt.testValue&0x80 != 0 {
					movement |= 0xFF00
				}
				expected += movement
			}
			tt.cpu.WriteBus(tt.cpu.programCounter, tt.testValue)
			bcc := bcc(tt.cpu, rel)
			bcc()

			if tt.cpu.programCounter != expected {
				t.Errorf("bcc() = 0x%X, wanted 0x%X", tt.cpu.programCounter, expected)
			}
		})
	}
}

func Test_bcs(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu:       &CPU{programCounter: 0x8000},
			testValue: 0x20},
		{name: "standard",
			cpu:       &CPU{programCounter: 0x8000, processorStatus: carry},
			testValue: 0x20},
		{name: "negative",
			cpu:       &CPU{programCounter: 0x8000, processorStatus: carry},
			testValue: 0x9F},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := tt.cpu.programCounter + 1
			tt.setupInstructionTest()
			if tt.cpu.processorStatus.has(carry) {
				var movement = uint16(tt.testValue)
				if tt.testValue&0x80 != 0 {
					movement |= 0xFF00
				}
				expected += movement
			}
			tt.cpu.WriteBus(tt.cpu.programCounter, tt.testValue)
			bcs := bcs(tt.cpu, rel)
			bcs()

			if tt.cpu.programCounter != expected {
				t.Errorf("bcs() = 0x%X, wanted 0x%X", tt.cpu.programCounter, expected)
			}
		})
	}
}

func Test_beq(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu:       &CPU{programCounter: 0x8000},
			testValue: 0x20},
		{name: "standard",
			cpu:       &CPU{programCounter: 0x8000, processorStatus: zero},
			testValue: 0x20},
		{name: "negative",
			cpu:       &CPU{programCounter: 0x8000, processorStatus: zero},
			testValue: 0x9F},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := tt.cpu.programCounter + 1
			tt.setupInstructionTest()
			if tt.cpu.processorStatus.has(zero) {
				var movement = uint16(tt.testValue)
				if tt.testValue&0x80 != 0 {
					movement |= 0xFF00
				}
				expected += movement
			}
			tt.cpu.WriteBus(tt.cpu.programCounter, tt.testValue)
			beq := beq(tt.cpu, rel)
			beq()

			if tt.cpu.programCounter != expected {
				t.Errorf("beq() = 0x%X, wanted 0x%X", tt.cpu.programCounter, expected)
			}
		})
	}
}

func Test_bit(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bit(tt.cpu, imm)
		})
	}
}

func Test_bmi(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu:       &CPU{programCounter: 0x8000},
			testValue: 0x20},
		{name: "standard",
			cpu:       &CPU{programCounter: 0x8000, processorStatus: negative},
			testValue: 0x20},
		{name: "negative",
			cpu:       &CPU{programCounter: 0x8000, processorStatus: negative},
			testValue: 0x9F},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := tt.cpu.programCounter + 1
			tt.setupInstructionTest()
			if tt.cpu.processorStatus.has(negative) {
				var movement = uint16(tt.testValue)
				if tt.testValue&0x80 != 0 {
					movement |= 0xFF00
				}
				expected += movement
			}
			tt.cpu.WriteBus(tt.cpu.programCounter, tt.testValue)
			bmi := bmi(tt.cpu, rel)
			bmi()

			if tt.cpu.programCounter != expected {
				t.Errorf("bmi() = 0x%X, wanted 0x%X", tt.cpu.programCounter, expected)
			}
		})
	}
}

func Test_bne(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu:       &CPU{programCounter: 0x8000, processorStatus: zero},
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
			expected := tt.cpu.programCounter + 1
			tt.setupInstructionTest()
			if !tt.cpu.processorStatus.has(zero) {
				var movement = uint16(tt.testValue)
				if tt.testValue&0x80 != 0 {
					movement |= 0xFF00
				}
				expected += movement
			}
			tt.cpu.WriteBus(tt.cpu.programCounter, tt.testValue)
			bne := bne(tt.cpu, rel)
			bne()

			if tt.cpu.programCounter != expected {
				t.Errorf("bmi() = 0x%X, wanted 0x%X", tt.cpu.programCounter, expected)
			}
		})
	}
}

func Test_bpl(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu:       &CPU{programCounter: 0x8000, processorStatus: negative},
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
			expected := tt.cpu.programCounter + 1
			tt.setupInstructionTest()
			if !tt.cpu.processorStatus.has(negative) {
				var movement = uint16(tt.testValue)
				if tt.testValue&0x80 != 0 {
					movement |= 0xFF00
				}
				expected += movement
			}
			tt.cpu.WriteBus(tt.cpu.programCounter, tt.testValue)
			bpl := bpl(tt.cpu, rel)
			bpl()

			if tt.cpu.programCounter != expected {
				t.Errorf("bpl() = 0x%X, wanted 0x%X", tt.cpu.programCounter, expected)
			}
		})
	}
}

func Test_brk(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			brk(tt.cpu, imm)
		})
	}
}

func Test_bvc(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu:       &CPU{programCounter: 0x8000, processorStatus: overflow},
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
			expected := tt.cpu.programCounter + 1
			tt.setupInstructionTest()
			if !tt.cpu.processorStatus.has(overflow) {
				var movement = uint16(tt.testValue)
				if tt.testValue&0x80 != 0 {
					movement |= 0xFF00
				}
				expected += movement
			}
			tt.cpu.WriteBus(tt.cpu.programCounter, tt.testValue)
			bvc := bvc(tt.cpu, rel)
			bvc()

			if tt.cpu.programCounter != expected {
				t.Errorf("bvc() = 0x%X, wanted 0x%X", tt.cpu.programCounter, expected)
			}
		})
	}
}

func Test_bvs(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu:       &CPU{programCounter: 0x8000, processorStatus: overflow},
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
			expected := tt.cpu.programCounter + 1
			tt.setupInstructionTest()
			if tt.cpu.processorStatus.has(overflow) {
				var movement = uint16(tt.testValue)
				if tt.testValue&0x80 != 0 {
					movement |= 0xFF00
				}
				expected += movement
			}
			tt.cpu.WriteBus(tt.cpu.programCounter, tt.testValue)
			bvs := bvs(tt.cpu, rel)
			bvs()

			if tt.cpu.programCounter != expected {
				t.Errorf("bvs() = 0x%X, wanted 0x%X", tt.cpu.programCounter, expected)
			}
		})
	}
}

func Test_clc(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu: &CPU{programCounter: 0x8000, processorStatus: carry}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clc := clc(tt.cpu, impl)
			clc()

			if tt.cpu.processorStatus.has(carry) {
				t.Errorf("clc() didn't clear carry flag")
			}
		})
	}
}

func Test_cld(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu: &CPU{programCounter: 0x8000, processorStatus: decimal}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cld := cld(tt.cpu, impl)
			cld()

			if tt.cpu.processorStatus.has(decimal) {
				t.Errorf("cld() didn't clear decimal flag")
			}
		})
	}
}

func Test_cli(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu: &CPU{programCounter: 0x8000, processorStatus: interruptDisable}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := cli(tt.cpu, impl)
			cli()

			if tt.cpu.processorStatus.has(interruptDisable) {
				t.Errorf("cli() didn't clear interrupt disable flag")
			}
		})
	}
}

func Test_clv(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu: &CPU{programCounter: 0x8000, processorStatus: overflow}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clv := clv(tt.cpu, impl)
			clv()

			if tt.cpu.processorStatus.has(overflow) {
				t.Errorf("clv() didn't clear overflow flag")
			}
		})
	}
}

func Test_cmp(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmp(tt.cpu, imm)
		})
	}
}

func Test_cpx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpx(tt.cpu, imm)
		})
	}
}

func Test_cpy(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpy(tt.cpu, imm)
		})
	}
}

func Test_dec(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dec(tt.cpu, imm)
		})
	}
}

func Test_dex(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dex(tt.cpu, imm)
		})
	}
}

func Test_dey(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dey(tt.cpu, imm)
		})
	}
}

func Test_eor(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eor(tt.cpu, imm)
		})
	}
}

func Test_inc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inc(tt.cpu, imm)
		})
	}
}

func Test_inx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inx(tt.cpu, imm)
		})
	}
}

func Test_iny(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iny(tt.cpu, imm)
		})
	}
}

func Test_jmp(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jmp(tt.cpu, imm)
		})
	}
}

func Test_jsr(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsr(tt.cpu, imm)
		})
	}
}

func Test_lda(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lda(tt.cpu, imm)
		})
	}
}

func Test_ldx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ldx(tt.cpu, imm)
		})
	}
}

func Test_ldy(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ldy(tt.cpu, imm)
		})
	}
}

func Test_lsr(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lsr(tt.cpu, imm)
		})
	}
}

func Test_nop(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nop(tt.cpu, imm)
		})
	}
}

func Test_ora(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ora(tt.cpu, imm)
		})
	}
}

func Test_pha(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pha(tt.cpu, imm)
		})
	}
}

func Test_php(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			php(tt.cpu, imm)
		})
	}
}

func Test_pla(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pla(tt.cpu, imm)
		})
	}
}

func Test_plp(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plp(tt.cpu, imm)
		})
	}
}

func Test_rol(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rol(tt.cpu, imm)
		})
	}
}

func Test_ror(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ror(tt.cpu, imm)
		})
	}
}

func Test_rti(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rti(tt.cpu, imm)
		})
	}
}

func Test_rts(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rts(tt.cpu, imm)
		})
	}
}

func Test_sbc(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sbc(tt.cpu, imm)
		})
	}
}

func Test_sec(t *testing.T) {
	tests := []instructionTest{
		{name: "base", cpu: &CPU{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sec := sec(tt.cpu, impl)
			sec()
			if !tt.cpu.processorStatus.has(carry) {
				t.Errorf("sec() didn't set carry flag")
			}
		})
	}
}

func Test_sed(t *testing.T) {
	tests := []instructionTest{
		{name: "base", cpu: &CPU{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sed := sed(tt.cpu, impl)
			sed()
			if !tt.cpu.processorStatus.has(decimal) {
				t.Errorf("sed() didn't set decimal flag")
			}
		})
	}
}

func Test_sei(t *testing.T) {
	tests := []instructionTest{
		{name: "base", cpu: &CPU{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sei := sei(tt.cpu, impl)
			sei()
			if !tt.cpu.processorStatus.has(interruptDisable) {
				t.Errorf("sei() didn't set interrupt flag")
			}
		})
	}
}

func Test_sta(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sta(tt.cpu, imm)
		})
	}
}

func Test_stx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stx(tt.cpu, imm)
		})
	}
}

func Test_sty(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sty(tt.cpu, zpg)
		})
	}
}

func Test_tax(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu:   &CPU{accumulator: 0x41, indexX: 0x05},
			flags: 0},
		{name: "negative",
			cpu:   &CPU{accumulator: 0xA0, indexX: 0x95},
			flags: negative},
		{name: "zero",
			cpu:   &CPU{accumulator: 0x0, indexX: 0x00},
			flags: zero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupInstructionTest()
			tax := tax(tt.cpu, impl)
			tax()

			if tt.cpu.accumulator != tt.cpu.indexX {
				t.Errorf("tax() = 0x%X wanted 0x%X", tt.cpu.accumulator, tt.cpu.indexX)
			}
			if tt.cpu.processorStatus != tt.flags {
				t.Errorf("processor status = %b, wanted %b", tt.cpu.processorStatus, tt.flags)
			}
		})
	}
}

func Test_tay(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tay(tt.cpu, imm)
		})
	}
}

func Test_tsx(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tsx(tt.cpu, imm)
		})
	}
}

func Test_txa(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu:   &CPU{accumulator: 0x41, indexX: 0x05},
			flags: 0},
		{name: "negative",
			cpu:   &CPU{accumulator: 0x8F, indexX: 0x95},
			flags: negative},
		{name: "zero",
			cpu:   &CPU{accumulator: 0x41, indexX: 0x00},
			flags: zero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txa := txa(tt.cpu, imm)
			txa()

			if tt.cpu.accumulator != tt.cpu.indexX {
				t.Errorf("txa() = 0x%X wanted 0x%X", tt.cpu.accumulator, tt.cpu.indexX)
			}
			if tt.cpu.processorStatus != tt.flags {
				t.Errorf("processor status = %b, wanted %b", tt.cpu.processorStatus, tt.flags)
			}
		})
	}
}

func Test_txs(t *testing.T) {
	tests := []instructionTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txs(tt.cpu, imm)
		})
	}
}

func Test_tya(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu:   &CPU{accumulator: 0x41, indexY: 0x05},
			flags: 0},
		{name: "negative",
			cpu:   &CPU{accumulator: 0x8F, indexY: 0x95},
			flags: negative},
		{name: "zero",
			cpu:   &CPU{accumulator: 0x41, indexY: 0x00},
			flags: zero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tya := tya(tt.cpu, impl)
			tya()
		})
		if tt.cpu.accumulator != tt.cpu.indexY {
			t.Errorf("tya() = 0x%X wanted 0x%X", tt.cpu.accumulator, tt.cpu.indexY)
		}
		if tt.cpu.processorStatus != tt.flags {
			t.Errorf("processor status = %b, wanted %b", tt.cpu.processorStatus, tt.flags)
		}
	}
}
