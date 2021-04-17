package nes

const (
	MemoryStart = 0x0000
	EndOfMemory = 0xFFFF
)

type Memory struct {
	RAM [EndOfMemory]byte
}

func (m *Memory) Read(address uint16, readOnly bool) byte {
	switch {
	case address >= MemoryStart && address < EndOfMemory:
		return m.RAM[address]
	default:
		// TODO
		// Memory out of bounds
	}
	return 0x00
}

func (m *Memory) Write(address uint16, value uint8) {
	switch {
	case address >= MemoryStart && address < EndOfMemory:
		m.RAM[address] = value
	default:
		// TODO
		// Memory out of bounds
	}
}
