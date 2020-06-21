package main

import (
	// "fmt"
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
	LY byte // keep track of the current line in the frame and index part of tile correctly
	LYC byte
	LX byte
	Cycles int
	BGP byte
}

func (ppu *PPU) tick() {
	ppu.Cycles++
	
	switch ppu.LCDCSTAT & 0x03 {
		case 0:
			// HBlank
			if ppu.Cycles >= 204 {
				ppu.Cycles = 0
				ppu.drawScanLine()
				ppu.LY++
				ppu.LX = 0
				
				
				if ppu.LY == 144 {
					ppu.LCDCSTAT |= 0x01
					
				}
			}
		case 1:
			// VBlank
			if ppu.Cycles >= 456 {
				ppu.Cycles = 0
				ppu.LY++
				
			} // implement as correct interrupt later
			if ppu.LY >= 153 {
				ppu.LCDCSTAT |= 0x10
				ppu.LY = 0
				// ppu.LX = 0
				
			}
		case 2: 
			// search OAM still got to implement
			if ppu.Cycles >= 80 {
				ppu.Cycles = 0
				ppu.LCDCSTAT |= 0x11
			}
		case 3:
			// read scanline from VRAM and put in framebuffer
			if ppu.Cycles >= 172 {
				// draw scanline
				ppu.LCDCSTAT |= 0x00
				ppu.Cycles = 0

			}
	}
}



// trying to draw line by line
func (ppu *PPU) drawScanLine() {
	
	// strategy: get starting bg map address
	// using SCX and SCY find the tile to start on and iterate over the 360 tiles
	startAddr := ppu.getBGMapAddr() + uint16((ppu.SCY + ppu.LY) * 32) + uint16(ppu.SCX)
	for i := 0; i < 20; i++ {
		if ppu.VRAM[startAddr + uint16(i) - 0x8000] > 0 {


			// fmt.Println(ppu.VRAM[startAddr + uint16(i) - 0x8000])
		}
		// ppu.drawBGLine()
	}

	// for i := 0; i < 8; i++ {
	// 	ppu.frameBuffer[ppu.yIndex][ppu.xIndex] = (((1 << i) & left) >> i | ((1 << i) & right) >> i)
		
		

		
	// }
	

}
// take 2 bytes and add to framebuffer
func (ppu *PPU) drawBGLine(left byte, ) {
	
	
	

	for i := 0; i < 8; i++ {
		
		
		// ppu.frameBuffer[ppu.yIndex][ppu.xIndex] = (((1 << i) & left) >> i | ((1 << i) & right) >> i)
		
		

		
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

