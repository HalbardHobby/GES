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
	memory, read, write := getTestMemory()
	tt.cpu.ReadBus = read
	tt.cpu.WriteBus = write
	expectedCounter = tt.cpu.programCounter + tt.pcMovement
	memory[tt.cpu.programCounter] = uint8(tt.address)
	memory[tt.address] = tt.wantValue
	return
}

// This function creates a sample memory array with its read
// and write buses for testing purposes.
func getTestMemory() (mem memory, read readBus, write writeBus) {
	mem = make([]byte, 0xFFFF)
	read = func(adr uint16) uint8 {
		return mem[adr]
	}
	write = func(adr uint16, val uint8) {
		mem[adr] = val
	}
	return
}

func TestCPU_a(t *testing.T) {
	tests := []addressTest{
		{"base", &CPU{}, 0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCounter := tt.setupAddressTest()
			if gotValue := tt.cpu.a(); gotValue != tt.wantValue {
				t.Errorf("CPU.a() = %v, want %v", gotValue, tt.wantValue)
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
		{"base", GetCPU(), 0x0FDE, 2, 0x15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedPC := tt.setupAddressTest()
			tt.cpu.WriteBus(tt.cpu.programCounter, uint8(tt.address&0xFF))
			tt.cpu.WriteBus(tt.cpu.programCounter+1, uint8(tt.address>>8))
			if gotValue := tt.cpu.abs(); gotValue != tt.wantValue {

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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.absX(); gotValue != tt.wantValue {
				t.Errorf("CPU.absX() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_absY(t *testing.T) {
	tests := []addressTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.absY(); gotValue != tt.wantValue {
				t.Errorf("CPU.absY() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_imm(t *testing.T) {
	tests := []addressTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.imm(); gotValue != tt.wantValue {
				t.Errorf("CPU.imm() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_impl(t *testing.T) {
	tests := []addressTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.impl(); gotValue != tt.wantValue {
				t.Errorf("CPU.impl() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_ind(t *testing.T) {
	tests := []addressTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.ind(); gotValue != tt.wantValue {
				t.Errorf("CPU.ind() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_xInd(t *testing.T) {
	tests := []addressTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.xInd(); gotValue != tt.wantValue {
				t.Errorf("CPU.xInd() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_indY(t *testing.T) {
	tests := []addressTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.indY(); gotValue != tt.wantValue {
				t.Errorf("CPU.indY() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_rel(t *testing.T) {
	tests := []addressTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.rel(); gotValue != tt.wantValue {
				t.Errorf("CPU.rel() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_zpg(t *testing.T) {
	c := GetCPU()
	tests := []addressTest{
		{"general", c, 0x00CD, 1, 0x0F},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memory, read, write := getTestMemory()
			tt.cpu.ReadBus = read
			tt.cpu.WriteBus = write
			expectedCounter := tt.cpu.programCounter + 1
			memory[tt.cpu.programCounter] = uint8(tt.address)
			memory[tt.address] = tt.wantValue

			if gotValue := tt.cpu.zpg(); gotValue != tt.wantValue {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.zpgX(); gotValue != tt.wantValue {
				t.Errorf("CPU.zpgX() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_zpgY(t *testing.T) {
	tests := []addressTest{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.zpgY(); gotValue != tt.wantValue {
				t.Errorf("CPU.zpgY() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
