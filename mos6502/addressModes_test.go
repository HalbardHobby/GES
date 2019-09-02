package mos6502

import "testing"

type memory []byte
type readBus func(uint16) uint8
type writeBus func(uint16, uint8)

// Struct type used to model test cases
type addressTest struct {
	name        string
	cpu         *CPU
	wantImplied bool
	address     uint16
	pcMovement  uint16
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

func Test_acc(t *testing.T) {
	tests := []addressTest{
		{name: "base", cpu: &CPU{accumulator: 0x80}, wantImplied: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCounter := tt.setupAddressTest()
			if _, gotImpl, _ := acc(tt.cpu); gotImpl != tt.wantImplied {
				t.Errorf("CPU.acc() = %v, want %v", gotImpl, tt.wantImplied)
			}
			if tt.cpu.programCounter != expectedCounter {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedCounter)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	tests := []addressTest{
		{name: "base", cpu: &CPU{},
			address: 0x0FDE, pcMovement: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedPC := tt.setupAddressTest()
			tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address&0xFF))
			tt.cpu.WriteBus(tt.cpu.programCounter+1, uint8(tt.address>>8))
			if gotAddress, _, _ := abs(tt.cpu); gotAddress != tt.address {
				t.Errorf("CPU.abs() = %v, want %v", gotAddress, tt.address)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func Test_absX(t *testing.T) {
	tests := []addressTest{
		{name: "base",
			cpu: &CPU{indexX: 0}, address: 0x0F00,
			pcMovement: 2},
		{name: "advance",
			cpu: &CPU{indexX: 0x25}, address: 0x0F00,
			pcMovement: 2},
		{name: "page overflow",
			cpu: &CPU{indexX: 0x40}, address: 0x0FDD,
			pcMovement: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedPC := tt.setupAddressTest()
			adr := tt.address + uint16(tt.cpu.indexX)
			tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address&0xFF))
			tt.cpu.WriteBus(tt.cpu.programCounter+1, uint8(tt.address>>8))

			if gotValue, _, _ := absX(tt.cpu); gotValue != adr {
				t.Errorf("CPU.absX() = 0x%X, want 0x%X", gotValue, adr)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func Test_absY(t *testing.T) {
	tests := []addressTest{
		{name: "base",
			cpu: &CPU{indexY: 0}, address: 0x0F00,
			pcMovement: 2},
		{name: "advance",
			cpu: &CPU{indexY: 0x25}, address: 0x0F00,
			pcMovement: 2},
		{name: "page overflow",
			cpu: &CPU{indexY: 0x40}, address: 0x0FDD,
			pcMovement: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedPC := tt.setupAddressTest()
			adr := tt.address + uint16(tt.cpu.indexY)
			tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address&0xFF))
			tt.cpu.WriteBus(tt.cpu.programCounter+1, uint8(tt.address>>8))

			if gotValue, _, _ := absY(tt.cpu); gotValue != adr {
				t.Errorf("CPU.absY() = 0x%X, want 0x%X", gotValue, adr)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func Test_imm(t *testing.T) {
	tests := []addressTest{
		{name: "base", cpu: &CPU{},
			pcMovement: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedPC := tt.setupAddressTest()
			if gotValue, _, _ := imm(tt.cpu); gotValue != tt.address {
				t.Errorf("CPU.imm() = 0x%X, want 0x%X", gotValue, tt.address)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func Test_impl(t *testing.T) {
	tests := []addressTest{
		{name: "base", cpu: &CPU{},
			wantImplied: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedPC := tt.setupAddressTest()
			if _, implied, _ := impl(tt.cpu); implied != tt.wantImplied {
				t.Errorf("CPU.impl() = %v, want %v", implied, tt.wantImplied)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func Test_ind(t *testing.T) {
	tests := []addressTest{
		{name: "base", cpu: &CPU{},
			address: 0x0FDE, pcMovement: 2},
	}
	for _, tt := range tests {
		expectedPC := tt.setupAddressTest()
		intadr := uint16(0x0AEF)
		tt.cpu.WriteBus(intadr, uint8(tt.address&0xFF))
		tt.cpu.WriteBus(intadr+1, uint8(tt.address>>8))

		tt.cpu.WriteBus(tt.cpu.programCounter, uint8(intadr&0xFF))
		tt.cpu.WriteBus(tt.cpu.programCounter+1, uint8(intadr>>8))

		t.Run(tt.name, func(t *testing.T) {
			if gotValue, _, _ := ind(tt.cpu); gotValue != tt.address {
				t.Errorf("CPU.ind() = 0x%X, want 0x%X", gotValue, tt.address)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func Test_xInd(t *testing.T) {
	tests := []addressTest{
		{name: "static", cpu: &CPU{},
			address: 0x00DE, pcMovement: 1},
		{name: "standard", cpu: &CPU{indexX: 0x05},
			address: 0x00DE, pcMovement: 1},
		{name: "overflow", cpu: &CPU{indexX: 0x30},
			address: 0x00DE, pcMovement: 1},
	}
	for _, tt := range tests {
		expectedPC := tt.setupAddressTest()
		effAdr := uint16(0x5F2A)

		tt.cpu.WriteBus((tt.address+uint16(tt.cpu.indexX))%256,
			uint8(effAdr&0xFF))
		tt.cpu.WriteBus((tt.address+uint16(tt.cpu.indexX)+1)%256,
			uint8(effAdr>>8))

		tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address&0xFF))

		t.Run(tt.name, func(t *testing.T) {
			if gotValue, _, _ := xInd(tt.cpu); gotValue != effAdr {
				t.Errorf("CPU.xInd() = 0x%X, want 0x%X", gotValue, effAdr)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func Test_indY(t *testing.T) {
	tests := []addressTest{
		{name: "static", cpu: &CPU{},
			address: 0x00DE, pcMovement: 1},
		{name: "standard", cpu: &CPU{indexY: 0x05},
			address: 0x00DE, pcMovement: 1},
		{name: "overflow", cpu: &CPU{indexY: 0x30},
			address: 0x00DE, pcMovement: 1},
	}
	for _, tt := range tests {
		expectedPC := tt.setupAddressTest()
		effAdr := uint16(0x5F2A)

		tt.cpu.WriteBus((tt.address)%256,
			uint8(effAdr&0xFF))
		tt.cpu.WriteBus((tt.address+1)%256,
			uint8(effAdr>>8))

		tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address&0xFF))

		t.Run(tt.name, func(t *testing.T) {
			if gotValue, _, _ := indY(tt.cpu); gotValue != effAdr+uint16(tt.cpu.indexY) {
				t.Errorf("CPU.indY() = 0x%X, want 0x%X", gotValue, effAdr+uint16(tt.cpu.indexY))
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func Test_rel(t *testing.T) {
	tests := []addressTest{
		{name: "static",
			cpu: &CPU{programCounter: 0x8000}, address: 0x0,
			pcMovement: 1},
		{name: "positive",
			cpu: &CPU{programCounter: 0x8000}, address: 0x10,
			pcMovement: 1},
		{name: "negative",
			cpu: &CPU{programCounter: 0x8000}, address: 0x90,
			pcMovement: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedPC := tt.setupAddressTest()
			addressMovement := tt.address
			if addressMovement&0x80 != 0 {
				addressMovement |= 0xFF00
			}

			tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address))
			if gotValue, _, _ := rel(tt.cpu); gotValue != tt.cpu.programCounter+addressMovement {
				t.Errorf("CPU.rel() = 0x%X, want 0x%X", gotValue, tt.cpu.programCounter+addressMovement)
			}
			if tt.cpu.programCounter != expectedPC {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedPC)
			}
		})
	}
}

func Test_zpg(t *testing.T) {
	tests := []addressTest{
		{name: "general", cpu: GetCPU(),
			address: 0x00CD, pcMovement: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCounter := tt.setupAddressTest()
			tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address))

			if gotValue, _, _ := zpg(tt.cpu); gotValue != tt.address {
				t.Errorf("CPU.zpg() = 0x%X, want 0x%X", gotValue, tt.address)
			}
			if tt.cpu.programCounter != expectedCounter {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedCounter)
			}
		})
	}
}

func Test_zpgX(t *testing.T) {
	tests := []addressTest{
		{name: "base", cpu: &CPU{indexX: 0x02},
			address: 0x005F, pcMovement: 1},
		{name: "page overflow", cpu: &CPU{indexX: 0xDD},
			address: 0x00DD, pcMovement: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCounter := tt.setupAddressTest()
			expectedAddress := (tt.address + uint16(tt.cpu.indexX)) & 0x00FF
			tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address))
			if gotValue, _, _ := zpgX(tt.cpu); gotValue != expectedAddress {
				t.Errorf("CPU.zpgX() = 0x%X, want 0x%X", gotValue, expectedAddress)
			}
			if tt.cpu.programCounter != expectedCounter {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedCounter)
			}
		})
	}
}

func Test_zpgY(t *testing.T) {
	tests := []addressTest{
		{name: "base", cpu: &CPU{indexY: 0x02},
			address: 0x005F, pcMovement: 1},
		{name: "page overflow", cpu: &CPU{indexX: 0xDD},
			address: 0x00DD, pcMovement: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCounter := tt.setupAddressTest()
			expectedAddress := (tt.address + uint16(tt.cpu.indexY)) & 0x00FF
			tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address))
			if gotValue, _, _ := zpgY(tt.cpu); gotValue != expectedAddress {
				t.Errorf("CPU.zpgY() = 0x%X, want 0x%X", gotValue, expectedAddress)
			}
			if tt.cpu.programCounter != expectedCounter {
				t.Errorf("counter is 0x%X, should be 0x%X",
					tt.cpu.programCounter, expectedCounter)
			}
		})
	}
}
