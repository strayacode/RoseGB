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

func (cpu *CPU) PPUTick() {
	cpu.bus.ppu.Cycles++
	
	switch cpu.bus.ppu.LCDCSTAT & 0x03 {
		case 0:
			// HBlank
			if cpu.bus.ppu.Cycles >= 204 {
				cpu.bus.ppu.Cycles = 0
				cpu.bus.ppu.LX = 0
				cpu.bus.ppu.drawScanLine()
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
			
			if cpu.bus.ppu.Cycles >= 80 {
				
				cpu.bus.ppu.Cycles = 0
				// set to mode 3
				cpu.bus.ppu.LCDCSTAT |= 0x03
				cpu.bus.ppu.cpuVRAMAccess = false
			}
		case 3:
			// read scanline from VRAM and put in framebuffer
			
			if cpu.bus.ppu.Cycles >= 172 {
				
				// set to hblank
				cpu.bus.ppu.LCDCSTAT &= 0xFC
				cpu.bus.ppu.Cycles = 0
				cpu.bus.ppu.cpuVRAMAccess = true
				cpu.bus.ppu.setHBlankInterrupt()
				cpu.bus.interrupt.requestLCDCSTAT()

			}
	}
}



// trying to draw line by line
func (ppu *PPU) drawScanLine() {
	
	tileY := math.Floor(float64(ppu.LY + ppu.SCY) / 8)
	tileX := math.Floor(float64(ppu.SCX + ppu.LX) / 8)
	
	startAddr := ppu.getBGMapAddr() + uint16(tileY * 32) + uint16(tileX)
	// iterate through tiles in a scanline
	for i := 0; i < 20; i++ {
		if ppu.getBGStartAddr() == 0x8000 {
			start := (uint16(ppu.getBGStartAddr()) + uint16(ppu.VRAM[startAddr + uint16(i) - 0x8000]) * 16) - 0x8000
			tileOffset := uint16((ppu.SCY + ppu.LY) % 8) * 2
			
			ppu.drawBGLine(ppu.VRAM[start + tileOffset], ppu.VRAM[start + tileOffset + 1])
		} else {
			start := (uint16(ppu.getBGStartAddr()) + uint16(ppu.VRAM[uint16(int(startAddr) + 128) + uint16(i) - 0x8000]) * 16) - 0x8000
			tileOffset := uint16((ppu.SCY + ppu.LY) % 8) * 2
			
			ppu.drawBGLine(ppu.VRAM[start + tileOffset], ppu.VRAM[start + tileOffset + 1])
		}
	}
}
// take 2 bytes and add to framebuffer
func (ppu *PPU) drawBGLine(left byte, right byte) {
	for i := 0; i < 8; i++ {
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

func (ppu *PPU) resetPPU() {
	for i := 0; i < 144; i++ {
		for j := 0; j < 160; j++ {
			ppu.frameBuffer[i][j] = 0
		}
	}
	for i := 0; i < 0x2000; i++ {
		ppu.VRAM[i] = 0
	}
	for i := 0; i < 0xA0; i++ {
		ppu.OAM[i] = 0
	}
	ppu.LCDC = 0
	ppu.LCDCSTAT = 0
	ppu.SCX = 0
	ppu.SCY = 0
	ppu.LY = 0
	ppu.LYC = 0
	ppu.LX = 0
	ppu.WY = 0
	ppu.WX = 0
	ppu.Cycles = 0
	ppu.BGP = 0
	ppu.OBP0 = 0
	ppu.OBP1 = 0
	ppu.cpuVRAMAccess = false
	ppu.cpuOAMAccess = false
	ppu.DMA = 0

}




