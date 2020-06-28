package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"flag"
)


type Window struct {
	width int
	height int
	skipBootROM bool
}
func main() {

	// init window with properties
	window := Window{}
	window.checkFlags()

	// init cpu
	cpu := CPU{}


	// check skip bootrom flag
	if window.skipBootROM == false {
		cpu.bus.cartridge.loadBootROM()
	} else {
		cpu.skipBootROM()
	}

	title := "RoseGB - " + string(cpu.bus.cartridge.header.title[:])
	rl.InitWindow(int32(window.width), int32(window.height), title)

	rl.SetTargetFPS(60)

	// mainloop which is executed 60 times per second
	for !rl.WindowShouldClose() {
		for i := 0; i < 17556; i++ {
			cpu.tick()
			cpu.bus.ppu.tick()
			
		}
		cpu.drawFramebuffer()
		// cpu.debugWRAM()
		
	}

	rl.CloseWindow()
}

func (window *Window) checkFlags() {
	// flags to add
	// window size
	// emulation speed
	var intvar int
	var boolvar bool
	flag.IntVar(&intvar, "window-size", 1, "specifies the window size as a multiple")
	flag.BoolVar(&boolvar, "skip-bootrom", false, "user specifies whether the bootrom should be loaded or not")
	window.width = 160 * intvar
	window.height = 144 * intvar
	window.skipBootROM = boolvar
	flag.Parse()
	
}

func (cpu *CPU) skipBootROM() {
	cpu.A = 0x01
	cpu.B = 0x00
	cpu.C = 0x13
	cpu.D = 0x00
	cpu.E = 0xD8
	cpu.F = 0xB0
	cpu.H = 0x01
	cpu.L = 0x4D
	cpu.SP = 0xFFFE
	cpu.PC = 0x100
	cpu.bus.write(0xFF05, 0x00)
	cpu.bus.write(0xFF06, 0x00)
  	cpu.bus.write(0xFF07, 0x00)
  	cpu.bus.write(0xFF10, 0x80)
  	cpu.bus.write(0xFF11, 0xBF)
 	cpu.bus.write(0xFF12, 0xF3)
 	cpu.bus.write(0xFF14, 0xBF)
 	cpu.bus.write(0xFF16, 0x3F)
 	cpu.bus.write(0xFF17, 0x00)
 	cpu.bus.write(0xFF19, 0xBF)
 	cpu.bus.write(0xFF1A, 0x7F)
 	cpu.bus.write(0xFF1B, 0xFF)
 	cpu.bus.write(0xFF1C, 0x9F)
 	cpu.bus.write(0xFF1E, 0xBF)
 	cpu.bus.write(0xFF20, 0xFF)
 	cpu.bus.write(0xFF21, 0x00)
 	cpu.bus.write(0xFF22, 0x00)
 	cpu.bus.write(0xFF23, 0xBF)
 	cpu.bus.write(0xFF24, 0x77)
 	cpu.bus.write(0xFF25, 0xF3)
 	cpu.bus.write(0xFF26, 0xF1) // F0 on super gameboy
 	cpu.bus.write(0xFF40, 0x91)
  	cpu.bus.write(0xFF42, 0x00)
  	cpu.bus.write(0xFF43, 0x00)
  	cpu.bus.write(0xFF45, 0x00)
  	cpu.bus.write(0xFF47, 0xFC)
  	cpu.bus.write(0xFF48, 0xFF)
  	cpu.bus.write(0xFF49, 0xFF)
  	cpu.bus.write(0xFF4A, 0x00)
  	cpu.bus.write(0xFF4B, 0x00)
  	cpu.bus.write(0xFFFF, 0x00)
  	cpu.bus.cartridge.loadCartridge()
}

type Colour struct {
	R byte
	G byte
	B byte
	A byte
}

func (cpu *CPU) drawFramebuffer() {
	var colours [4]Colour = [4]Colour{
		Colour{255, 255, 255, 255}, Colour{192, 192, 192, 255}, Colour{96, 96, 96, 255}, Colour{0, 0, 0, 255},
	} 
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)
	for i := 0; i < 144; i++ {
		for j := 0; j < 160; j++ {
			// get palette
			
			
			tileColour := cpu.bus.read(0xFF47)
			if cpu.bus.ppu.frameBuffer[i][j] == 0 {
				tileColour &= 0x03
				
			} else if cpu.bus.ppu.frameBuffer[i][j] == 1 {
				tileColour = (tileColour & 0xC) >> 2
				
			} else if cpu.bus.ppu.frameBuffer[i][j] == 2 {
				tileColour = (tileColour & 0x30) >> 4
				
			} else if cpu.bus.ppu.frameBuffer[i][j] == 3 {
				tileColour = (tileColour & 0xC0) >> 6
				
			}

			
				
			
			rl.DrawRectangle(int32(j), int32(i), 1, 1, rl.NewColor(colours[tileColour].R, colours[tileColour].G, colours[tileColour].B, 255))
		}
	}
	
	rl.EndDrawing()
}