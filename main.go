package main

import "fmt"
import "github.com/HalbardHobby/GES/mos6502"

func main() {

	ram := make([]byte, 0xFFFF)
	read := func(adr uint16) uint8 {
		return ram[adr]
	}

	write := func(adr uint16, val uint8) {
		ram[adr] = val
	}

	var cpu mos6502.CPU
	cpu.ReadBus = read
	cpu.WriteBus = write

	fmt.Println("Hello, GES")
}
