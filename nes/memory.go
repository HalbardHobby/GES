package nes

const (
	internalRAMSize       = 0x0800
	internalRAMMirrorSize = 0x1800
	ppuRegisterSize       = 0x0008
	ppuRegisterMirrorSize = 0x1FF8
	apuIORegisterSize     = 0x0018
	apuIOTestSize         = 0x0008
	cartridgeSpaceSize    = 0xBFE0
)

const (
	internalMemoryStart    = 0x0000
	mirrorMemoryStart      = 0x0800
	ppuRegisterStart       = 0x2000
	mirrorPPURegisterStart = 0x2008
	apuIORegisterStart     = 0x4000
	apuIORegisterTestStart = 0x4018
	cartridgespaceStart    = 0x4020
	endOfMemory            = 0xFFFF
)

type Memory struct {
	internalRAM    [internalRAMSize]byte    // Internal 2KB RAM
	ppuRegisters   [ppuRegisterSize]byte    // PPU Registers
	apuIORegisters [apuIORegisterSize]byte  // APU & I/O registers
	apuIOTest      [apuIOTestSize]byte      // APU & I/O functionality for test purposes
	cartridgeSpace [cartridgeSpaceSize]byte // cartridge space
}

func (m *Memory) Read(address uint16) byte {
	switch {
	case address >= internalMemoryStart && address < mirrorMemoryStart:
		// 2KB internal RAM
	case address >= mirrorMemoryStart && address < ppuRegisterStart:
		// three Mirrors of $0000-$07FF
	case address >= ppuRegisterStart && address < mirrorPPURegisterStart:
		// PPU registers
	case address >= mirrorPPURegisterStart && address < apuIORegisterStart:
		// Mirrors of $2000-$2007, repeats every 8 bytes
	case address >= apuIORegisterStart && address < apuIORegisterTestStart:
		// APU & I/O registers
	case address >= apuIORegisterTestStart && address < cartridgespaceStart:
		// APU & I/O test functionality
	case address >= cartridgespaceStart && address < endOfMemory:
		// Cartridge space: PRG ROM, PRG RAM and mapper registers
	default:
		// TODO
		// Memory out of bounds
	}
	return 0
}

func (m *Memory) Write(address uint16, value uint8) {
	switch {
	case address >= internalMemoryStart && address < mirrorMemoryStart:
		// 2KB internal RAM
	case address >= mirrorMemoryStart && address < ppuRegisterStart:
		// three Mirrors of $0000-$07FF
	case address >= ppuRegisterStart && address < mirrorPPURegisterStart:
		// PPU registers
	case address >= mirrorPPURegisterStart && address < apuIORegisterStart:
		// Mirrors of $2000-$2007, repeats every 8 bytes
	case address >= apuIORegisterStart && address < apuIORegisterTestStart:
		// APU & I/O registers
	case address >= apuIORegisterTestStart && address < cartridgespaceStart:
		// APU & I/O test functionality
	case address >= cartridgespaceStart && address < endOfMemory:
		// Cartridge space: PRG ROM, PRG RAM and mapper registers
	default:
		// TODO
		// Memory out of bounds
	}
}
