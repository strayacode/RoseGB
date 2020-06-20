package main

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
	// Cycles int
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

