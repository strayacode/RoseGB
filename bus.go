package main

import (
	"fmt"
)

type Bus struct {
	cartridge Cartridge
	WRAM [0x2000]byte
	ppu PPU
	IOReg [0x80]byte
	HRAM [0x80]byte
	IE byte
}


func (bus *Bus) read(addr uint16) byte {
	switch {
		case addr >= 0x0000 && addr <= 0x7FFF:
			return bus.cartridge.ROM[addr]
		case addr >= 0x8000 && addr <= 0x9FFF:
			return bus.ppu.VRAM[addr - 0x8000]
		case addr >= 0xA000 && addr <= 0xBFFF:
			return bus.cartridge.ERAM[addr - 0xA000]
		case addr >= 0xC000 && addr <= 0xDFFF:
			return bus.WRAM[addr - 0xC000]
		case addr >= 0xFF00 && addr <= 0xFF7F:
			return bus.IOReg[addr - 0xFF00]
		case addr >= 0xFF80 && addr <= 0xFFFE:
			return bus.HRAM[addr - 0xFF80]
		default:
			return 0
	}
}

func (bus *Bus) write(addr uint16, data byte) {
	switch {
		case addr >= 0x8000 && addr <= 0x9FFF:
			bus.ppu.VRAM[addr - 0x8000] = data
		case addr >= 0xA000 && addr <= 0xBFFF:
			bus.cartridge.ERAM[addr - 0xA000] = data
		case addr >= 0xC000 && addr <= 0xDFFF:
			bus.WRAM[addr - 0xC000] = data
		case addr >= 0xFF00 && addr <= 0xFF7F:
			bus.IOReg[addr - 0xFF00] = data
		case addr == 0xFFFF:
			bus.IE = data
		default:
			fmt.Println("DEBUG: non-writeable memory location!", addr)
	}
}