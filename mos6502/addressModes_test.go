package mos6502

import "testing"

type memory []byte
type readBus func(uint16) uint8
type writeBus func(uint16, uint8)

// Struct type used to model test cases
type addressTest struct {
	name       string
	cpu        *CPU
	address    uint16
	pcMovement uint16
	wantValue  uint8
}

// Utility function that sets up the testcase
func (tt *addressTest) setupAddressTest() (expectedCounter uint16) {
	read, write := getTestMemory()
	tt.cpu.ReadBus = read
	tt.cpu.WriteBus = write
	expectedCounter = tt.cpu.programCounter + tt.pcMovement
	return
}

// This function creates a sample memory array with its read
// and write buses for testing purposes.
func getTestMemory() (read readBus, write writeBus) {
	mem := make([]byte, 0xFFFF)
	read = func(adr uint16) uint8 {
		return mem[adr]
	}
	write = func(adr uint16, val uint8) {
		mem[adr] = val
	}
	return
}

func TestCPU_acc(t *testing.T) {
	tests := []addressTest{
		{name: "base", cpu: &CPU{accumulator: 0x80}, wantValue: 0x80},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCounter := tt.setupAddressTest()
			if gotValue := acc(tt.cpu); gotValue != tt.wantValue {
				t.Errorf("CPU.acc() = %v, want %v", gotValue, tt.wantValue)
			}
			if tt.cpu.programCounter != expectedCounter {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedCounter)
			}
		})
	}
}

func TestCPU_abs(t *testing.T) {
	tests := []addressTest{
		{name: "base", cpu: &CPU{},
			address: 0x0FDE, pcMovement: 2, wantValue: 0x15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedPC := tt.setupAddressTest()
			tt.cpu.WriteBus(tt.address, tt.wantValue)
			tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address&0xFF))
			tt.cpu.WriteBus(tt.cpu.programCounter+1, uint8(tt.address>>8))
			if gotValue := abs(tt.cpu); gotValue != tt.wantValue {

				t.Errorf("CPU.abs() = %v, want %v", gotValue, tt.wantValue)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func TestCPU_absX(t *testing.T) {
	tests := []addressTest{
		{name: "base",
			cpu: &CPU{indexX: 0}, address: 0x0F00,
			pcMovement: 2, wantValue: 0x15},
		{name: "advance",
			cpu: &CPU{indexX: 0x25}, address: 0x0F00,
			pcMovement: 2, wantValue: 0x15},
		{name: "page overflow",
			cpu: &CPU{indexX: 0x40}, address: 0x0FDD,
			pcMovement: 2, wantValue: 0x15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedPC := tt.setupAddressTest()
			adr := tt.address + uint16(tt.cpu.indexX)
			tt.cpu.WriteBus(adr, tt.wantValue)
			tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address&0xFF))
			tt.cpu.WriteBus(tt.cpu.programCounter+1, uint8(tt.address>>8))

			if gotValue := absX(tt.cpu); gotValue != tt.wantValue {
				t.Errorf("CPU.absX() = 0x%X, want 0x%X", gotValue, tt.wantValue)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func TestCPU_absY(t *testing.T) {
	tests := []addressTest{
		{name: "base",
			cpu: &CPU{indexY: 0}, address: 0x0F00,
			pcMovement: 2, wantValue: 0x15},
		{name: "advance",
			cpu: &CPU{indexY: 0x25}, address: 0x0F00,
			pcMovement: 2, wantValue: 0x15},
		{name: "page overflow",
			cpu: &CPU{indexY: 0x40}, address: 0x0FDD,
			pcMovement: 2, wantValue: 0x15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedPC := tt.setupAddressTest()
			adr := tt.address + uint16(tt.cpu.indexY)
			tt.cpu.WriteBus(adr, tt.wantValue)
			tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address&0xFF))
			tt.cpu.WriteBus(tt.cpu.programCounter+1, uint8(tt.address>>8))

			if gotValue := absY(tt.cpu); gotValue != tt.wantValue {
				t.Errorf("CPU.absY() = 0x%X, want 0x%X", gotValue, tt.wantValue)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func TestCPU_imm(t *testing.T) {
	tests := []addressTest{
		{name: "base", cpu: &CPU{},
			pcMovement: 1, wantValue: 0x15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedPC := tt.setupAddressTest()
			tt.cpu.WriteBus(tt.cpu.programCounter, tt.wantValue)
			if gotValue := imm(tt.cpu); gotValue != tt.wantValue {
				t.Errorf("CPU.imm() = 0x%X, want 0x%X", gotValue, tt.wantValue)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func TestCPU_impl(t *testing.T) {
	tests := []addressTest{
		{name: "base", cpu: &CPU{},
			pcMovement: 0, wantValue: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedPC := tt.setupAddressTest()
			if gotValue := impl(tt.cpu); gotValue != tt.wantValue {
				t.Errorf("CPU.impl() = %v, want %v", gotValue, tt.wantValue)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func TestCPU_ind(t *testing.T) {
	tests := []addressTest{
		{name: "base", cpu: &CPU{},
			address: 0x0FDE, pcMovement: 2, wantValue: 0x15},
	}
	for _, tt := range tests {
		expectedPC := tt.setupAddressTest()
		intadr := uint16(0x0AEF)
		tt.cpu.WriteBus(intadr, uint8(tt.address&0xFF))
		tt.cpu.WriteBus(intadr+1, uint8(tt.address>>8))

		tt.cpu.WriteBus(tt.cpu.programCounter, uint8(intadr&0xFF))
		tt.cpu.WriteBus(tt.cpu.programCounter+1, uint8(intadr>>8))

		tt.cpu.WriteBus(tt.address, tt.wantValue)

		t.Run(tt.name, func(t *testing.T) {
			if gotValue := ind(tt.cpu); gotValue != tt.wantValue {
				t.Errorf("CPU.ind() = 0x%X, want 0x%X", gotValue, tt.wantValue)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func TestCPU_xInd(t *testing.T) {
	tests := []addressTest{
		{name: "static", cpu: &CPU{},
			address: 0x00DE, pcMovement: 1, wantValue: 0x15},
		{name: "standard", cpu: &CPU{indexX: 0x05},
			address: 0x00DE, pcMovement: 1, wantValue: 0x15},
		{name: "overflow", cpu: &CPU{indexX: 0x30},
			address: 0x00DE, pcMovement: 1, wantValue: 0x15},
	}
	for _, tt := range tests {
		expectedPC := tt.setupAddressTest()
		effAdr := uint16(0x5F2A)

		tt.cpu.WriteBus((tt.address+uint16(tt.cpu.indexX))%256,
			uint8(effAdr&0xFF))
		tt.cpu.WriteBus((tt.address+uint16(tt.cpu.indexX)+1)%256,
			uint8(effAdr>>8))

		tt.cpu.WriteBus(effAdr, tt.wantValue)
		tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address&0xFF))

		t.Run(tt.name, func(t *testing.T) {
			if gotValue := xInd(tt.cpu); gotValue != tt.wantValue {
				t.Errorf("CPU.xInd() = 0x%X, want 0x%X", gotValue, tt.wantValue)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func TestCPU_indY(t *testing.T) {
	tests := []addressTest{
		{name: "static", cpu: &CPU{},
			address: 0x00DE, pcMovement: 1, wantValue: 0x15},
		{name: "standard", cpu: &CPU{indexY: 0x05},
			address: 0x00DE, pcMovement: 1, wantValue: 0x15},
		{name: "overflow", cpu: &CPU{indexY: 0x30},
			address: 0x00DE, pcMovement: 1, wantValue: 0x15},
	}
	for _, tt := range tests {
		expectedPC := tt.setupAddressTest()
		effAdr := uint16(0x5F2A)

		tt.cpu.WriteBus((tt.address)%256,
			uint8(effAdr&0xFF))
		tt.cpu.WriteBus((tt.address+1)%256,
			uint8(effAdr>>8))

		tt.cpu.WriteBus(effAdr+uint16(tt.cpu.indexY), tt.wantValue)
		tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address&0xFF))

		t.Run(tt.name, func(t *testing.T) {
			if gotValue := indY(tt.cpu); gotValue != tt.wantValue {
				t.Errorf("CPU.indY() = 0x%X, want 0x%X", gotValue, tt.wantValue)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func TestCPU_rel(t *testing.T) {
	tests := []addressTest{
		{name: "static",
			cpu: &CPU{programCounter: 0x8000}, address: 0x0F00,
			pcMovement: 1, wantValue: 0x0},
		{name: "positive",
			cpu: &CPU{programCounter: 0x8000}, address: 0x0F00,
			pcMovement: 1, wantValue: 0x10},
		{name: "negative",
			cpu: &CPU{programCounter: 0x8000}, address: 0x0FDD,
			pcMovement: 1, wantValue: 0x90},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedPC := tt.setupAddressTest()
			tt.cpu.WriteBus(tt.cpu.programCounter, tt.wantValue)
			if gotValue := rel(tt.cpu); gotValue != tt.wantValue {
				t.Errorf("CPU.rel() = %v, want %v", gotValue, tt.wantValue)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func TestCPU_zpg(t *testing.T) {
	tests := []addressTest{
		{name: "general", cpu: GetCPU(),
			address: 0x00CD, pcMovement: 1, wantValue: 0x0F},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCounter := tt.setupAddressTest()
			tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address))
			tt.cpu.WriteBus(tt.address, tt.wantValue)

			if gotValue := zpg(tt.cpu); gotValue != tt.wantValue {
				t.Errorf("CPU.zpg() = 0x%X, want 0x%X", gotValue, tt.wantValue)
			}
			if tt.cpu.programCounter != expectedCounter {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedCounter)
			}
		})
	}
}

func TestCPU_zpgX(t *testing.T) {
	tests := []addressTest{
		{name: "base", cpu: &CPU{indexX: 0x02},
			address: 0x005F, pcMovement: 1, wantValue: 0x8F},
		{name: "page overflow", cpu: &CPU{indexX: 0xDD},
			address: 0x00DD, pcMovement: 1, wantValue: 0x8F},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCounter := tt.setupAddressTest()
			tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address))
			tt.cpu.WriteBus((tt.address+uint16(tt.cpu.indexX))&0x00FF,
				tt.wantValue)
			if gotValue := zpgX(tt.cpu); gotValue != tt.wantValue {
				t.Errorf("CPU.zpgX() = 0x%X, want 0x%X", gotValue, tt.wantValue)
			}
			if tt.cpu.programCounter != expectedCounter {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedCounter)
			}
		})
	}
}

func TestCPU_zpgY(t *testing.T) {
	tests := []addressTest{
		{name: "base", cpu: &CPU{indexY: 0x02},
			address: 0x005F, pcMovement: 1, wantValue: 0x8F},
		{name: "page overflow", cpu: &CPU{indexX: 0xDD},
			address: 0x00DD, pcMovement: 1, wantValue: 0x8F},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCounter := tt.setupAddressTest()
			tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address))
			tt.cpu.WriteBus((tt.address+uint16(tt.cpu.indexY))&0x00FF,
				tt.wantValue)
			if gotValue := zpgY(tt.cpu); gotValue != tt.wantValue {
				t.Errorf("CPU.zpgY() = 0x%X, want 0x%X", gotValue, tt.wantValue)
			}
			if tt.cpu.programCounter != expectedCounter {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedCounter)
			}
		})
	}
}
