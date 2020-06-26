package main

import (
	// "fmt"
	"math"
)

type PPU struct {
	frameBuffer [144][160]byte
	VRAM [0x2000]byte
	OAM [0xA0]byte
	LCDC byte
	LCDCSTAT byte
	SCX byte
	SCY byte
	LY byte // keep track of the current line in the frame and index part of tile correctly
	LYC byte
	LX byte
	Cycles int
	BGP byte
}

func (ppu *PPU) tick() {
	ppu.Cycles++
	// fmt.Println(ppu.LCDCSTAT & 0x03)
	switch ppu.LCDCSTAT & 0x03 {
		case 0:
			// HBlank
			if ppu.Cycles >= 204 {
				ppu.Cycles = 0
				ppu.LX = 0
				ppu.drawScanLine()
				ppu.LY++
				
				// fmt.Println(ppu.LCDCSTAT & 0x03)
				
				if ppu.LY == 144 {
					// execute VBlank
					ppu.LCDCSTAT |= 0x01
					ppu.LCDCSTAT &= 0xFD
				} else {
					// execute OAM search
					
					ppu.LCDCSTAT |= 0x02
					ppu.LCDCSTAT &= 0xFE
					
				}
			}
		case 1:
			// VBlank
			if ppu.Cycles >= 4560 {
				// fmt.Println(ppu.LCDCSTAT & 0x03)
				// ppu.clearFramebuffer()
				ppu.Cycles = 0

				ppu.LY = 0
				ppu.LX = 0
				ppu.LCDCSTAT &= 0xFE
				
			} // implement as correct interrupt later
		case 2: 
			// search OAM still got to implement
			// fmt.Println(ppu.LCDCSTAT & 0x03)
			if ppu.Cycles >= 80 {
				
				ppu.Cycles = 0
				ppu.LCDCSTAT |= 0x03
			}
		case 3:
			// read scanline from VRAM and put in framebuffer

			if ppu.Cycles >= 172 {
				
				// draw scanline
				ppu.LCDCSTAT &= 0xFC
				ppu.Cycles = 0

			}
	}
}



// trying to draw line by line
func (ppu *PPU) drawScanLine() {
	
	
	tileY := math.Floor(float64(int(ppu.LY) + int(ppu.SCY)) / 8)
	tileX := math.Floor(float64(int(ppu.SCX) + int(ppu.LX) / 8))
	
	startAddr := ppu.getBGMapAddr() + uint16(tileY * 32) + uint16(tileX)
	// iterate through tiles in a scanline
	for i := 0; i < 20; i++ {
		start := (uint16(ppu.getBGStartAddr()) + uint16(ppu.VRAM[startAddr + uint16(i) - 0x8000])) - 0x8000
		tileIndex := uint16(math.Floor((float64(ppu.SCY % 8)) / 2)) 
		// fmt.Println(start, tileIndex)
		ppu.drawBGLine(ppu.VRAM[start + (2 * tileIndex)], ppu.VRAM[start + (2 * tileIndex) + 1])
		
		

		
		
	}
}
// take 2 bytes and add to framebuffer
func (ppu *PPU) drawBGLine(left byte, right byte) {
	
	// fmt.Println(left, right)
	

	for i := 0; i < 8; i++ {
		ppu.frameBuffer[ppu.LY][ppu.LX] = (((1 << i) & left) >> i | ((1 << i) & right) >> i)
		ppu.LX++
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

func (ppu *PPU) clearFramebuffer() {
	for i := 0; i < 160; i++ {
		for j := 0; j < 144; j++ {
			ppu.frameBuffer[j][i] = 0
		}
	}
}
