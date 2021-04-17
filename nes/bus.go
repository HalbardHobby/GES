package nes

// The bus only cares if a Device is capable of reading and writing
type Device interface {
	Read(uint16, bool) uint8
	Write(uint16, uint8)
}

type Bus struct {
	Devices []Device
	Memory  Memory
}
