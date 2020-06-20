package main

import (
	"fmt"
	// "strconv"
)

type PPU struct {
	frameBuffer [144][160]byte
	VRAM [0x2000]byte
	OAM [0xA0]byte
	LCDC byte
	LCDCSTAT byte
	SCX byte
	SCY byte
	LY byte
	LYC byte
	Cycles int
	Index int
	VRAMIndex uint16
	BGP byte
}

func (ppu *PPU) tick() {
	
	// if ppu.Cycles == 0 {
	// 	switch LCDCSTAT & 0x03:
	// 		case 0:
	// 			// HBlank
	// 			ppu.Cycles = 204
	// 		case 1:
	// 			// VBlank
	// 			cpu.
	// }
}

func (ppu *PPU) drawFramebuffer() {
	for i := ppu.getBGMapAddr(); i < ppu.getBGMapAddr() + 1024; i++ {
		// string := ""
		for j := 0; j < 16; j++ {
			index := ppu.VRAM[ppu.getBGMapAddr() - 0x8000]
			fmt.Println(index)
			// if index > 0 {


				// fmt.Println(uint16(j) + uint16((index * 16)), index)
			// }
			// fmt.Println(uint16(j) + ppu.getBGStartAddr() - 0x8000 + (ppu.getBGStartAddr() - 0x8000 + uint16(index * 16) + uint16(j)) * 16)
			
			// string += strconv.FormatUint(uint64(ppu.VRAM[ppu.getBGStartAddr() - 0x8000 + uint16(index * 16) + uint16(j)]), 16)
		}
		// fmt.Println(string)
	}
	// ppu.getBGStartAddr()
}

// WARNING: THIS IS THE VERY FIRST PPU IMPLEMENTATION FOR DRAWING: ONLY TO DRAW TILES TO SCREEN
func (ppu *PPU) drawTile(j int) {
	for i := 0; i < 8; i++ {

	}
}

func (ppu *PPU) getBGStartAddr() uint16 {
	if ppu.LCDC & 0x10 >> 4 == 1 {
		return 0x8000
	}
	return 0x8800
}

func (ppu *PPU) getBGMapAddr() uint16 {
	if ppu.LCDC & 0x08 >> 3 == 1 {
		return 0x9C00
	}
	return 0x9800
}

