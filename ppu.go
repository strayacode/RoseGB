package main

import (
	"math"
	"os"
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
	visibleSprites [10]int
}

func (cpu *CPU) PPUTick() {
	cpu.bus.ppu.Cycles++
	
	switch cpu.bus.ppu.LCDCSTAT & 0x03 {
		case 0:
			// HBlank
			if cpu.bus.ppu.Cycles >= 204 {
				cpu.bus.ppu.Cycles = 0
				cpu.bus.ppu.LX = 0
				cpu.bus.ppu.LY++
				if cpu.bus.ppu.LYC == cpu.bus.ppu.LY {
					cpu.bus.ppu.setLYCInterrupt()
					cpu.bus.interrupt.requestLCDCSTAT()
				}
				if cpu.bus.ppu.LY == 144 {
					// execute VBlank
					cpu.bus.ppu.setVBlankInterrupt()
					cpu.bus.interrupt.requestVBlank()
					cpu.bus.ppu.LCDCSTAT |= 0x01
					cpu.bus.ppu.LCDCSTAT &= 0xFD
					
				} else {
					// execute OAM search (mode 2)
					cpu.bus.ppu.LCDCSTAT |= 0x02
					cpu.bus.ppu.LCDCSTAT &= 0xFE
				}
			}
		case 1:
			
			// VBlank
			if cpu.bus.ppu.Cycles % 456 == 0 {
				cpu.bus.ppu.LY++
				if cpu.bus.ppu.LYC == cpu.bus.ppu.LY {
					cpu.bus.ppu.setLYCInterrupt()
					cpu.bus.interrupt.requestLCDCSTAT()
				}
			}
			if cpu.bus.ppu.Cycles >= 4560 {
				cpu.bus.ppu.Cycles = 0
				cpu.bus.ppu.LY = 0
				if cpu.bus.ppu.LYC == cpu.bus.ppu.LY {
					cpu.bus.ppu.setLYCInterrupt()
					cpu.bus.interrupt.requestLCDCSTAT()
				}
				cpu.bus.ppu.LX = 0
				// set to mode 2
				cpu.bus.ppu.LCDCSTAT |= 0x02
				cpu.bus.ppu.LCDCSTAT &= 0xFE
				cpu.bus.ppu.setOAMInterrupt()
				cpu.bus.interrupt.requestLCDCSTAT()
				
			} 
		case 2: 
			// search oam for 10 sprites that can fit on the current line and place in an array
			if cpu.bus.ppu.Cycles >= 80 {
				cpu.bus.ppu.spriteSelect()
				cpu.bus.ppu.Cycles = 0
				// set to mode 3
				cpu.bus.ppu.LCDCSTAT |= 0x03
				cpu.bus.ppu.cpuVRAMAccess = false
			}
		case 3:
			// read scanline from VRAM and put in framebuffer
			
			if cpu.bus.ppu.Cycles >= 172 {
				if (cpu.bus.ppu.LCDC & (1 << 7)) >> 7 == 1 {
					cpu.bus.ppu.drawScanLine()
				}
				// set to hblank
				cpu.bus.ppu.LCDCSTAT &= 0xFC
				cpu.bus.ppu.Cycles = 0
				cpu.bus.ppu.cpuVRAMAccess = true
				cpu.bus.ppu.setHBlankInterrupt()
				cpu.bus.interrupt.requestLCDCSTAT()

			}
	}
}


// search oam for first 10 sprites that can go on a line 
func (ppu *PPU) spriteSelect() {
	// clear previous sprites selection first
	// fmt.Println(ppu.OAM)
	for i := 0; i < 10; i++ {
		ppu.visibleSprites[i] = -1
	}
	index := 0
	for j := 0; j < 0xA0; j += 4 {
		if index < 10 {
			// y coordinate
			y := ppu.OAM[j]
			if y > 0 && y < 160 {
				
				if y + 8 >= ppu.LY && y <= ppu.LY {

					// x coordinate
					x := ppu.OAM[j + 1]
					// fmt.Println(j)
					if x > 0 && x < 168 {
						ppu.visibleSprites[index] = j
					} else {
						ppu.visibleSprites[index] = -1 // if negative dont draw the sprite
					}
					index++
				}
			}
			
		}
	}
	

}


// trying to draw line by line
func (ppu *PPU) drawScanLine() {
	
	offsetY := math.Floor(float64(ppu.LY + ppu.SCY) / 8)
	offsetX := math.Floor(float64(ppu.SCX + ppu.LX) / 8)
	tileOffset := uint16((ppu.SCY + ppu.LY) % 8) * 2
	startAddr := ppu.getBGMapAddr() + uint16(offsetY * 32) + uint16(offsetX)
	
	// iterate through tiles in a scanline
	for i := 0; i < 20; i++ {
		if ppu.getBGStartAddr() == 0x8000 {
			start := (uint16(ppu.getBGStartAddr()) + uint16(ppu.VRAM[(startAddr + uint16(i) - 0x8000) & 0x1FFF]) * 16) - 0x8000
			ppu.drawBGLine(ppu.VRAM[start + tileOffset], ppu.VRAM[start + tileOffset + 1])
		} else {
			tileNumber := ppu.VRAM[(startAddr + uint16(i) - 0x8000) & 0x1FFF]
			
			start := (uint16(ppu.getBGStartAddr()) + uint16(ppu.VRAM[(startAddr + uint16(i) - 0x8000) & 0x1FFF]) * 16) - 0x8000 
			
			if tileNumber <= 127 {
				start = (uint16(ppu.getBGStartAddr()) + uint16(ppu.VRAM[(startAddr + uint16(i) - 0x8000) & 0x1FFF]) * 16) - 0x8000	
			} else if tileNumber >= 128 {
				start = uint16(tileNumber) * 16
			}
			ppu.drawBGLine(ppu.VRAM[start + tileOffset], ppu.VRAM[start + tileOffset + 1])
		}

		// draw sprites now
		for j := 0; j < 10; j++ {
			if (ppu.LCDC & (1 << 1)) >> 1 == 1 && ppu.visibleSprites[j] >= 0 {
				
				if ppu.OAM[ppu.visibleSprites[j] + 1] - 8 >= ppu.LX && ppu.OAM[ppu.visibleSprites[j] + 1] - 8 <= ppu.LX + 8 { 
					// get tile data
					start := (uint16(ppu.OAM[ppu.visibleSprites[j] + 2]) * 16)
					// draw sprite
					tileOffset := uint16((ppu.LY) % 8) * 2
					ppu.drawSPRLine(ppu.OAM[ppu.visibleSprites[j] + 3], ppu.VRAM[start + tileOffset], ppu.VRAM[start + tileOffset + 1])
				}
			}
		}


		ppu.LX += 8
	}

	
}

// take 2 bytes and add to framebuffer
func (ppu *PPU) drawSPRLine(attr byte, left byte, right byte) {
	// fmt.Println(left, right)
	for i := 0; i < 8; i++ {
		var selection byte = ((((1 << (7 - i)) & left)) >> (7 - i)) << 1 | ((1 << (7 - i)) & right) >> (7 - i)
		var colour byte = 0
		if ppu.spritePalette(attr) == 1 {
			colour = (ppu.OBP1 & (0x03 << (selection * 2))) >> (selection * 2)
		} else {
			colour = (ppu.OBP0 & (0x03 << (selection * 2))) >> (selection * 2)
		}
		ppu.frameBuffer[ppu.LY][ppu.LX + byte(i)] = colour

		
	}
}

// take 2 bytes and add to framebuffer
func (ppu *PPU) drawBGLine(left byte, right byte) {
	for i := 0; i < 8; i++ {
		selection := ((((1 << (7 - i)) & left)) >> (7 - i)) << 1 | ((1 << (7 - i)) & right) >> (7 - i)
		colour := (ppu.BGP & (0x03 << (selection * 2))) >> (selection * 2)
		ppu.frameBuffer[ppu.LY][ppu.LX + byte(i)] = colour
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

func (ppu *PPU) setLYCInterrupt() {
	ppu.LCDCSTAT |= (1 << 6)
}

func (ppu *PPU) setOAMInterrupt() {
	ppu.LCDCSTAT |= (1 << 5)
}

func (ppu *PPU) setVBlankInterrupt() {
	ppu.LCDCSTAT |= (1 << 4)
}

func (ppu *PPU) setHBlankInterrupt() {
	ppu.LCDCSTAT |= (1 << 3)
}

func (ppu *PPU) dmaTransfer() {
	for i := 0; i < 0x9F; i++ {
		cpu.bus.write(0xFE00 + uint16(i), cpu.bus.read(uint16(cpu.bus.ppu.DMA) * 256 + uint16(i)))
	}
}

func (ppu *PPU) spritePalette(attr byte) byte {
	return ((attr & 0x10) >> 4)
}




