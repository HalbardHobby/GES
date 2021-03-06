package mos6502

import (
	"testing"
)

type instructionTest struct {
	name       string
	cpu        *CPU
	flags      flagSet
	testValue  uint8
	address    uint16
	addressing addressMode
}

var _testAddr uint16

func (tt *instructionTest) setupInstructionTest() {
	read, write := getTestMemory()
	tt.cpu.ReadBus = read
	tt.cpu.WriteBus = write

	tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address))
	tt.cpu.WriteBus(tt.cpu.programCounter+1, uint8(tt.address>>8))

	tt.cpu.WriteBus(tt.address, tt.testValue)
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
		{name: "base acc",
			cpu:        &CPU{accumulator: 0x02},
			addressing: acc},
		{name: "negative acc",
			cpu:        &CPU{accumulator: 0x5A},
			addressing: acc, flags: negative},
		{name: "zero acc",
			cpu:        &CPU{accumulator: 0x0},
			addressing: acc, flags: zero},
		{name: "carry acc",
			cpu:        &CPU{accumulator: 0xAA},
			addressing: acc, flags: carry},
		{name: "base mem",
			cpu:        &CPU{},
			addressing: abs,
			address:    0xAF9D, testValue: 0x02},
		{name: "negative mem",
			cpu:        &CPU{},
			addressing: abs, flags: negative,
			address: 0xAF9D, testValue: 0x5A},
		{name: "zero mem",
			cpu:        &CPU{},
			addressing: abs, flags: zero,
			address: 0xAF9D, testValue: 0x00},
		{name: "carry mem",
			cpu:        &CPU{},
			addressing: abs, flags: carry,
			address: 0xAF9D, testValue: 0xAA},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupInstructionTest()
			expected := tt.cpu.accumulator << 1
			if tt.testValue != 0 {
				expected = tt.testValue << 1
			}
			asl := asl(tt.cpu, tt.addressing)
			asl()

			var answer uint8
			if _, impl, _ := tt.addressing(tt.cpu); impl {
				answer = tt.cpu.accumulator
			} else {
				answer = tt.cpu.ReadBus(tt.address)
			}

			if answer != expected {
				t.Errorf("asl() got 0x%X, wanted 0x%X", answer, expected)
			}
			if tt.cpu.processorStatus != tt.flags {
				t.Errorf("processor status = %b, wanted %b",
					tt.cpu.processorStatus, tt.flags)
			}
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
		{name: "base",
			cpu: &CPU{}, flags: zero,
			testValue: 0x20, address: 0xAFED},
		{name: "base",
			cpu: &CPU{accumulator: 0x40}, flags: overflow,
			testValue: 0x40, address: 0xAFED},
		{name: "base",
			cpu: &CPU{accumulator: 0x80}, flags: negative,
			testValue: 0x80, address: 0xAFED},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupInstructionTest()
			bit := bit(tt.cpu, abs)
			bit()

			if tt.cpu.processorStatus != tt.flags {
				t.Errorf("processor status = %b, wanted %b",
					tt.cpu.processorStatus, tt.flags)
			}
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
		{name: "carry", },
		{name: "zero", },
		{name: "negative", },
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
		{name: "base",
			cpu: &CPU{}, address: 0xAFDE,
			testValue: 0x05},
		{name: "negative",
			cpu: &CPU{}, address: 0xAFDE,
			testValue: 0x00, flags: negative},
		{name: "zero",
			cpu: &CPU{}, address: 0xAFDE,
			testValue: 0x01, flags: zero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupInstructionTest()
			dec := dec(tt.cpu, abs)
			dec()

			if tt.cpu.ReadBus(tt.address) != tt.testValue-1 {
				t.Errorf("Address 0x%X is not 0x%X, got 0x%X",
					tt.address, tt.testValue-1, tt.cpu.ReadBus(tt.address))
			}
			if tt.cpu.processorStatus != tt.flags {
				t.Errorf("processor status = %b, wanted %b",
					tt.cpu.processorStatus, tt.flags)
			}
		})
	}
}

func Test_dex(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu: &CPU{indexX: 0x05}},
		{name: "negative",
			cpu: &CPU{indexX: 0x00}, flags: negative},
		{name: "zero",
			cpu: &CPU{indexX: 0x01}, flags: zero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupInstructionTest()
			expected := tt.cpu.indexX - 1
			dex := dex(tt.cpu, impl)
			dex()

			if tt.cpu.indexX != expected {
				t.Errorf("Index X is not 0x%X, got 0x%X",
					expected, tt.cpu.indexX)
			}
			if tt.cpu.processorStatus != tt.flags {
				t.Errorf("processor status = %b, wanted %b",
					tt.cpu.processorStatus, tt.flags)
			}
		})
	}
}

func Test_dey(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu: &CPU{indexY: 0x05}},
		{name: "negative",
			cpu: &CPU{indexY: 0x00}, flags: negative},
		{name: "zero",
			cpu: &CPU{indexY: 0x01}, flags: zero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupInstructionTest()
			expected := tt.cpu.indexY - 1
			dey := dey(tt.cpu, impl)
			dey()

			if tt.cpu.indexY != expected {
				t.Errorf("Index Y is not 0x%X, got 0x%X",
					expected, tt.cpu.indexY)
			}
			if tt.cpu.processorStatus != tt.flags {
				t.Errorf("processor status = %b, wanted %b",
					tt.cpu.processorStatus, tt.flags)
			}
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
		{name: "base",
			cpu: &CPU{}, address: 0xAFDE,
			testValue: 0x05},
		{name: "negative",
			cpu: &CPU{}, address: 0xAFDE,
			testValue: 0x7F, flags: negative},
		{name: "zero",
			cpu: &CPU{}, address: 0xAFDE,
			testValue: 0xFF, flags: zero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupInstructionTest()
			inc := inc(tt.cpu, abs)
			inc()

			if tt.cpu.ReadBus(tt.address) != tt.testValue+1 {
				t.Errorf("Address 0x%X is not 0x%X, got 0x%X",
					tt.address, tt.testValue+1, tt.cpu.ReadBus(tt.address))
			}
			if tt.cpu.processorStatus != tt.flags {
				t.Errorf("processor status = %b, wanted %b",
					tt.cpu.processorStatus, tt.flags)
			}
		})
	}
}

func Test_inx(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu: &CPU{indexX: 0x05}},
		{name: "negative",
			cpu: &CPU{indexX: 0x7F}, flags: negative},
		{name: "zero",
			cpu: &CPU{indexX: 0xFF}, flags: zero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupInstructionTest()
			expected := tt.cpu.indexX + 1
			inx := inx(tt.cpu, impl)
			inx()

			if tt.cpu.indexX != expected {
				t.Errorf("Index X is not 0x%X, got 0x%X",
					expected, tt.cpu.indexX)
			}
			if tt.cpu.processorStatus != tt.flags {
				t.Errorf("processor status = %b, wanted %b",
					tt.cpu.processorStatus, tt.flags)
			}
		})
	}
}

func Test_iny(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu: &CPU{indexY: 0x05}},
		{name: "negative",
			cpu: &CPU{indexY: 0x7F}, flags: negative},
		{name: "zero",
			cpu: &CPU{indexY: 0xFF}, flags: zero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupInstructionTest()
			expected := tt.cpu.indexY + 1
			iny := iny(tt.cpu, impl)
			iny()

			if tt.cpu.indexY != expected {
				t.Errorf("Index Y is not 0x%X, got 0x%X",
					expected, tt.cpu.indexY)
			}
			if tt.cpu.processorStatus != tt.flags {
				t.Errorf("processor status = %b, wanted %b",
					tt.cpu.processorStatus, tt.flags)
			}
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
		{name: "base",
			cpu: &CPU{}, address: 0xAFDE,
			testValue: 0x05},
		{name: "negative",
			cpu: &CPU{}, address: 0xAFDE,
			testValue: 0xAA, flags: negative},
		{name: "zero",
			cpu: &CPU{}, address: 0xAFDE,
			testValue: 0x00, flags: zero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupInstructionTest()
			lda := lda(tt.cpu, abs)
			lda()

			if tt.cpu.accumulator != tt.testValue {
				t.Errorf("Accumulator did not load data from memory")
			}
			if tt.cpu.processorStatus != tt.flags {
				t.Errorf("processor status = %b, wanted %b", tt.cpu.processorStatus, tt.flags)
			}
		})
	}
}

func Test_ldx(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu: &CPU{}, address: 0xAFDE,
			testValue: 0x05},
		{name: "negative",
			cpu: &CPU{}, address: 0xAFDE,
			testValue: 0xAA, flags: negative},
		{name: "zero",
			cpu: &CPU{}, address: 0xAFDE,
			testValue: 0x00, flags: zero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupInstructionTest()
			ldx := ldx(tt.cpu, abs)
			ldx()

			if tt.cpu.indexX != tt.testValue {
				t.Errorf("X register did not load data")
			}
			if tt.cpu.processorStatus != tt.flags {
				t.Errorf("processor status = %b, wanted %b", tt.cpu.processorStatus, tt.flags)
			}
		})
	}
}

func Test_ldy(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu: &CPU{}, address: 0xAFDE,
			testValue: 0x05},
		{name: "negative",
			cpu: &CPU{}, address: 0xAFDE,
			testValue: 0xAA, flags: negative},
		{name: "zero",
			cpu: &CPU{}, address: 0xAFDE,
			testValue: 0x00, flags: zero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupInstructionTest()
			ldy := ldy(tt.cpu, abs)
			ldy()

			if tt.cpu.indexY != tt.testValue {
				t.Errorf("Y register did not load data")
			}
			if tt.cpu.processorStatus != tt.flags {
				t.Errorf("processor status = %b, wanted %b", tt.cpu.processorStatus, tt.flags)
			}
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
		{name: "base",
			cpu:     &CPU{indexX: 0xAA},
			address: 0x00DE},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupInstructionTest()
			sta := sta(tt.cpu, zpg)
			sta()

			if tt.cpu.ReadBus(tt.address) != tt.cpu.accumulator {
				t.Errorf("0x%X address != 0x%X, got 0x%X",
					tt.address, tt.cpu.accumulator, tt.cpu.ReadBus(tt.address))
			}
		})
	}
}

func Test_stx(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu:     &CPU{indexX: 0xAA},
			address: 0x00DE},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupInstructionTest()
			stx := stx(tt.cpu, zpg)
			stx()

			if tt.cpu.ReadBus(tt.address) != tt.cpu.indexX {
				t.Errorf("0x%X address != 0x%X, got 0x%X",
					tt.address, tt.cpu.indexX, tt.cpu.ReadBus(tt.address))
			}
		})
	}
}

func Test_sty(t *testing.T) {
	tests := []instructionTest{
		{name: "base",
			cpu:     &CPU{indexY: 0xAA},
			address: 0x00DE},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupInstructionTest()
			sty := sty(tt.cpu, zpg)
			sty()

			if tt.cpu.ReadBus(tt.address) != tt.cpu.indexY {
				t.Errorf("0x%X address != 0x%X, got 0x%X",
					tt.address, tt.cpu.indexY, tt.cpu.ReadBus(tt.address))
			}
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
		{name: "base",
			cpu:   &CPU{accumulator: 0x41},
			flags: 0},
		{name: "negative",
			cpu:   &CPU{accumulator: 0x8F},
			flags: negative},
		{name: "zero",
			cpu:   &CPU{accumulator: 0x00},
			flags: zero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tay := tay(tt.cpu, imm)
			tay()

			if tt.cpu.accumulator != tt.cpu.indexY {
				t.Errorf("tay() = 0x%X wanted 0x%X", tt.cpu.indexY, tt.cpu.accumulator)
			}
			if tt.cpu.processorStatus != tt.flags {
				t.Errorf("processor status = %b, wanted %b", tt.cpu.processorStatus, tt.flags)
			}
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
