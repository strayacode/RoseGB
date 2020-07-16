package main

import (
	g "github.com/AllenDang/giu"
	"time"
	"image"
	"image/color"
	"strconv"
	"os"
	"flag"
	// "fmt"
)

var (
	cartridgetypes = [28]string{
		"ROM ONLY", "MBC1", "MBC1+RAM", "MBC1+RAM+BATTERY", "MBC2", "MBC2+BATTERY", "ROM+RAM", 
		"ROM+RAM+BATTERY", "MMM01", "MMM01+RAM", "MMM01+RAM+BATTERY", "MBC3+TIMER+BATTERY", "MBC3+TIMER+RAM+BATTERY", "MBC3", 
		"MBC3+RAM", "MBC3+RAM+BATTERY", "MBC5", "MBC5+RAM", "MBC5+RAM+BATTERY", "MBC5+RUMBLE", "MBC5+RUMBLE+RAM", 
		"MBC5+RUMBLE+RAM+BATTERY", "MBC6", "MBC7+SENSOR+RUMBLE+RAM+BATTERY", "POCKET CAMERA", "BANDAI TAMA5", "HuC3", "HuC1+RAM+BATTERY",
	}

	ramsize = [6]string{
		"None", "2KB", "8KB", "32KB", "128KB", "64KB",
	}
	romsize = ""
	upLeft = image.Point{0, 0}
	lowRight = image.Point{160, 144}
	texture *g.Texture
	tileTexture *g.Texture
	cycles = 70224
	cpu = CPU{}
)

func main() {
	// init cpu
	cpu.init()
	if checkBootromSkip() {
		cpu.skipBootROM()
	} else {
		cpu.bus.cartridge.loadBootROM()
	}
	switch cpu.bus.cartridge.header.ROMSize {
	case 0x00:
		romsize = "32KB"
	case 0x01:
		romsize = "64KB"
	case 0x02:
		romsize = "128KB"
	case 0x03:
		romsize = "256KB"
	case 0x04:
		romsize = "512KB"
	case 0x05:
		romsize = "1MB"
	case 0x06:
		romsize = "2MB"
	case 0x07:
		romsize = "4MB"
	case 0x08:
		romsize = "8MB"
	case 0x52:
		romsize = "1.1MB"
	case 0x53:
		romsize = "1.2MB"
	case 0x54:
		romsize = "1.5MB"
	}
	title := "RoseGB - " + string(cpu.bus.cartridge.header.title[:])
	wnd := g.NewMasterWindow(title, 1000, 500, g.MasterWindowFlagsNotResizable, nil)
	cpu.bus.keypad.P1 = 0xFF
    go refresh()

    wnd.Main(loop)
}

func loop() {
	g.MainMenuBar(g.Layout{
		g.Menu("File", g.Layout{
			// g.MenuItem("Reset", reset),
			g.MenuItem("Exit", exit),
			
		}),
		g.Menu("Options", g.Layout {
						g.Menu("Emulation Speed", g.Layout {
						g.MenuItem("1/2x", func() {cycles = 35112}),
						g.MenuItem("1x", func() {cycles = 70224}),
						g.MenuItem("2x", func() {cycles = 140448}),
						g.MenuItem("4x", func() {cycles = 280896}),
					},
					),
			},
			),
	}).Build()

	g.Window("RoseGB", 10, 30, 180, 180, g.Layout{
		g.Custom(func() {
			canvas := g.GetCanvas()
			pos := g.GetCursorScreenPos()
			if texture != nil {
				canvas.AddImage(texture, pos.Add(image.Pt(0, 0)), pos.Add(image.Pt(160, 144)))
			}
		}),
	})
	g.Window("Debugger", 200, 30, 300, 400, g.Layout{
		g.Label("A: 0x" + strconv.FormatUint(uint64(cpu.A), 16)),
		g.Label("B: 0x" + strconv.FormatUint(uint64(cpu.B), 16)),
		g.Label("C: 0x" + strconv.FormatUint(uint64(cpu.C), 16)),
		g.Label("D: 0x" + strconv.FormatUint(uint64(cpu.D), 16)),
		g.Label("E: 0x" + strconv.FormatUint(uint64(cpu.E), 16)),
		g.Label("F: 0x" + strconv.FormatUint(uint64(cpu.F), 16)),
		g.Label("H: 0x" + strconv.FormatUint(uint64(cpu.H), 16)),
		g.Label("L: 0x" + strconv.FormatUint(uint64(cpu.L), 16)),
		g.Label("LCDC: 0x" + strconv.FormatUint(uint64(cpu.bus.ppu.LCDC), 16)),
		g.Label("LCDCSTAT: 0x" + strconv.FormatUint(uint64(cpu.bus.ppu.LCDCSTAT), 16)),
		g.Label("Opcode: 0x" + strconv.FormatUint(uint64(cpu.Opcode), 16)),
		g.Label("IME: 0x" + strconv.FormatUint(uint64(cpu.bus.interrupt.IME), 16) + " IF: 0x" + strconv.FormatUint(uint64(cpu.bus.interrupt.IF), 16) + " IE: 0x" + strconv.FormatUint(uint64(cpu.bus.interrupt.IE), 16)),
		g.Label("DIV: 0x" + strconv.FormatUint(uint64(cpu.bus.timer.DIV), 16) + " TIMA: 0x" + strconv.FormatUint(uint64(cpu.bus.timer.TIMA), 16) + " TMA: 0x" + strconv.FormatUint(uint64(cpu.bus.timer.TMA), 16) + " TAC: 0x" + strconv.FormatUint(uint64(cpu.bus.timer.TAC), 16)),
		g.Label("halt: " + strconv.FormatBool(cpu.halt)),
		g.Label("P1: " + strconv.FormatUint(uint64(cpu.bus.keypad.P1), 2)),
		g.Label("Direction: " + strconv.FormatUint(uint64(cpu.bus.keypad.direction[3] << 3 | cpu.bus.keypad.direction[2] << 2 | cpu.bus.keypad.direction[1] << 1 | cpu.bus.keypad.direction[0]), 2)),
		g.Label("Button: " + strconv.FormatUint(uint64(cpu.bus.keypad.button[3] << 3 | cpu.bus.keypad.button[2] << 2 | cpu.bus.keypad.button[1] << 1 | cpu.bus.keypad.button[0]), 2)),
	})

	g.Window("Tile Viewer", 510, 30, 200, 300, g.Layout{
		g.Custom(func() {
			canvas := g.GetCanvas()
			pos := g.GetCursorScreenPos()
			if texture != nil {
				canvas.AddImage(tileTexture, pos.Add(image.Pt(0, 0)), pos.Add(image.Pt(128, 192)))
			}
		}),
	})
	
	g.Window("Cartridge", 720, 30, 250, 300, g.Layout{
		g.Label("Cartridge Type: " + cartridgetypes[cpu.bus.cartridge.header.cartridgeType]),
		g.Label("ROM Size: " + romsize),
		g.Label("RAM Size: " + ramsize[cpu.bus.cartridge.header.RAMSize]),
		g.Label("Banking Mode: " + strconv.Itoa(int(cpu.bus.bankingMode))),
		g.Label("ROMBankptr: " + strconv.Itoa(int(cpu.bus.cartridge.rombank.bankptr))),
		g.Label("RAMBankptr: " + strconv.Itoa(int(cpu.bus.cartridge.rambank.bankptr))),
		g.Label("isCBG: " + strconv.FormatBool(cpu.bus.cartridge.header.isCGB)),
	})

	
}

func exit() {
	os.Exit(3)
}

func checkBootromSkip() bool {
	boolPtr := flag.Bool("skip-bootrom", false, "user specifies whether they want to skip the bootrom or not")
	flag.Parse()
	return *boolPtr
}

func refresh() {
    ticker := time.NewTicker(time.Second / 60)
    for {
    	for i := 0; i < cycles; i++ {
    		cpu.bus.interrupt.handleInterrupts()
    		cpu.tick()
    		cpu.PPUTick()
    		cpu.bus.timer.tick()
    		if cpu.bus.timer.timerInterrupt == true {
    			cpu.bus.interrupt.requestTimer()
    			cpu.bus.timer.timerInterrupt = false
    		}
    	}
    	cpu.drawTileViewer()
    	cpu.drawFramebuffer()
    	cpu.checkInput()
    	// fmt.Println(cpu.bus.ppu.OAM)
        g.Update()
        <-ticker.C
    } 
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
		Colour{202, 220, 159, 255}, Colour{155, 188, 15, 255}, Colour{48, 98, 48, 255}, Colour{15, 56, 15, 255},
	} 
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for i := 0; i < 144; i++ {
		for j := 0; j < 160; j++ {
			// get palette
			tileColour := cpu.bus.ppu.frameBuffer[i][j]
			colour := color.RGBA{colours[tileColour].R, colours[tileColour].G, colours[tileColour].B, colours[tileColour].A}
			img.Set(j, i, colour)
		}
	}
	texture, _ = g.NewTextureFromRgba(img)
}

func (cpu *CPU) drawTileViewer() {
	var colours [4]Colour = [4]Colour{
		Colour{202, 220, 159, 255}, Colour{155, 188, 15, 255}, Colour{48, 98, 48, 255}, Colour{15, 56, 15, 255},
	} 
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{128, 192}})
	x := 0
	y := 0
	for i := 0; i < 0x1800; i += 2 {
		for j := 0; j < 8; j++ {
			// works fine
			tileColour := cpu.bus.read(0xFF47)
			value := ((((1 << (7 - j)) & cpu.bus.ppu.VRAM[i])) >> (7 - j)) << 1 | ((1 << (7 - j)) & cpu.bus.ppu.VRAM[i + 1]) >> (7 - j)
			if value == 0 {
				tileColour &= 0x03
				
			} else if value == 1 {
				tileColour = (tileColour & 0xC) >> 2
				
			} else if value == 2 {
				tileColour = (tileColour & 0x30) >> 4
				
			} else if value == 3 {
				tileColour = (tileColour & 0xC0) >> 6	
			}
			colour := color.RGBA{colours[tileColour].R, colours[tileColour].G, colours[tileColour].B, colours[tileColour].A}
			
			img.Set(x, y, colour)
			x++
			if x % 8 == 0 {
				x -= 8
			}
		}
		y++
		if y % 8 == 0 {
			y -= 8
			x += 8
			if x == 128 {
				y += 8
				x = 0
			}
		}	
	}
	tileTexture, _ = g.NewTextureFromRgba(img)
}

func (cpu *CPU) checkInput() {
	// z, A
	if g.IsKeyDown(90) {
		cpu.bus.keypad.setA(0)
	} else {
		cpu.bus.keypad.setA(1)
	}
		
	// x, B
	if g.IsKeyDown(88) {
		cpu.bus.keypad.setB(0)
	} else {
		cpu.bus.keypad.setB(1)
	}

	// return, start
	if g.IsKeyDown(257)  {
		cpu.bus.keypad.setStart(0)
	} else {
		cpu.bus.keypad.setStart(1)
	}

	// backspace, select
	if g.IsKeyDown(259) {
		cpu.bus.keypad.setSelect(0)
	} else {
		cpu.bus.keypad.setSelect(1)
	}
	
	// right key
	if g.IsKeyDown(262) {
		cpu.bus.keypad.setRight(0)
	} else {
		cpu.bus.keypad.setRight(1)
	}

	// left key
	if g.IsKeyDown(263) {
		cpu.bus.keypad.setLeft(0)
	} else {
		cpu.bus.keypad.setLeft(1)
	}

	// up key
	if g.IsKeyDown(265) {
		cpu.bus.keypad.setUp(0)
	} else {
		cpu.bus.keypad.setUp(1)
	}

	// down key
	if g.IsKeyDown(264) {
		cpu.bus.keypad.setDown(0)
	} else {
		cpu.bus.keypad.setDown(1)
	}



	
}

func reset() { // still experimental
	// reset state
	
	cpu = CPU{}
	cpu.bus.cartridge.loadBootROM()
}

func (cpu *CPU) init() {

	cpu.bus.cartridge.rombank.bankptr = 0x01
}
