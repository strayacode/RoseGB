package main

import (
	"fmt"
	"strconv"
	"os"
)

func (cpu *CPU) debugCPU() {
	fmt.Println("A: ", strconv.FormatUint(uint64(cpu.A), 16), "F ", strconv.FormatUint(uint64(cpu.F), 16))
	fmt.Println("B: ", strconv.FormatUint(uint64(cpu.B), 16), "C: ", strconv.FormatUint(uint64(cpu.C), 16))
	fmt.Println("D: ", strconv.FormatUint(uint64(cpu.D), 16), "E: ", strconv.FormatUint(uint64(cpu.E), 16))
	fmt.Println("H: ", strconv.FormatUint(uint64(cpu.H), 16), "L: ", strconv.FormatUint(uint64(cpu.L), 16))
	fmt.Println("PC: ", strconv.FormatUint(uint64(cpu.PC), 16), "SP: ", strconv.FormatUint(uint64(cpu.SP), 16))
}

func (cpu *CPU) debugPPU() {
	
	// if cpu.bus.ppu.LCDC & 0x10 == 1 {
	// 	cpu.debugVRAM(0x8000, 0x8FFF)
	// } else {
	// 	cpu.debugVRAM(0x8800, 0x97FF)
	// }
	fmt.Println("LCDC: ", cpu.bus.ppu.LCDC)
	fmt.Println("LCDCSTAT", cpu.bus.ppu.LCDCSTAT)
}

func (cpu *CPU) setPCBreakpoint(PC uint16) {
	if cpu.PC == PC {
		cpu.debugCPU()
		// cpu.debugPPU()
		// cpu.bus.ppu.drawFramebuffer()
		os.Exit(3)
	}

}

func (cpu *CPU) debugVRAM() {
	string := ""
	for i := 0x8000; i < 0x9FFF; i++ {
		string += " " + strconv.FormatUint(uint64(cpu.bus.ppu.VRAM[i - 0x8000]), 16)
	}
	fmt.Println(string)
}
