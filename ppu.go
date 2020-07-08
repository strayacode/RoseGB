package main

import (
	// "fmt"
	"math"
	"os"
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
	WY byte
	WX byte
	Cycles int
	BGP byte
	OBP0 byte
	OBP1 byte
	cpuVRAMAccess bool
	cpuOAMAccess bool
	DMA byte
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
				if ppu.LY == 144 {
					// execute VBlank
					ppu.LCDCSTAT |= 0x01
					ppu.LCDCSTAT &= 0xFD
				} else {
					// execute OAM search (mode 2)
					
					ppu.LCDCSTAT |= 0x02
					ppu.LCDCSTAT &= 0xFE
					
				}
			}
		case 1:
			// VBlank
			if ppu.Cycles >= 4560 {
				ppu.Cycles = 0

				ppu.LY = 0
				ppu.LX = 0
				// set to mode 2
				ppu.LCDCSTAT |= 0x02
				ppu.LCDCSTAT &= 0xFE
				
			} // implement as correct interrupt later
		case 2: 
			if ppu.Cycles >= 80 {
				
				ppu.Cycles = 0
				// set to mode 3
				ppu.LCDCSTAT |= 0x03
				ppu.cpuVRAMAccess = false
			}
		case 3:
			// read scanline from VRAM and put in framebuffer
			
			if ppu.Cycles >= 172 {
				
				// set to mode 0
				ppu.LCDCSTAT &= 0xFC
				ppu.Cycles = 0
				ppu.cpuVRAMAccess = true

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
		if ppu.getBGStartAddr() == 0x8000 {
			start := (uint16(ppu.getBGStartAddr()) + uint16(ppu.VRAM[startAddr + uint16(i) - 0x8000]) * 16) - 0x8000
			tileOffset := uint16((ppu.SCY + ppu.LY) % 8) * 2
			
			ppu.drawBGLine(ppu.VRAM[start + tileOffset], ppu.VRAM[start + tileOffset + 1])
		}
		//  else {
		// 	// fmt.Println("8800 addressing")
		// 	start := (uint16(ppu.getBGStartAddr()) + uint16(ppu.VRAM[uint16(int(startAddr) + 128) + uint16(i) - 0x8000]) * 16) - 0x8000
		// 	tileOffset := uint16((ppu.SCY + ppu.LY) % 8) * 2
			
		// 	ppu.drawBGLine(ppu.VRAM[start + tileOffset], ppu.VRAM[start + tileOffset + 1])
		// }
	}
}
// take 2 bytes and add to framebuffer
func (ppu *PPU) drawBGLine(left byte, right byte) {
	for i := 0; i < 8; i++ {
		// works fine
		ppu.frameBuffer[ppu.LY][ppu.LX] = ((((1 << (7 - i)) & left)) >> (7 - i)) << 1 | ((1 << (7 - i)) & right) >> (7 - i)
		ppu.LX++
	}
	

}

func (ppu *PPU) getBGStartAddr() uint16 {
	if ppu.LCDC & 0x10 >> 4 == 1 {
		return 0x8000
	}
	return 0x9000
}

func (ppu *PPU) getBGMapAddr() uint16 {
	if ppu.LCDC & 0x08 >> 3 == 1 {
		return 0x9C00
	}
	return 0x9800
}

func (ppu *PPU) writeDump(data string) {
	f, err := os.OpenFile("dump.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
	    panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(data); err != nil {
	    panic(err)
	}
}