package main

import (
	g "github.com/AllenDang/giu"
	"time"
	"image"
	"image/color"
	"strconv"
	"os"
	"flag"
	"fmt"
)

var (
	upLeft = image.Point{0, 0}
	lowRight = image.Point{160, 144}
	texture *g.Texture
	tileTexture *g.Texture
	cpu = CPU{}
)

func main() {
	// init cpu
	if checkBootromSkip() {
		cpu.skipBootROM()
	} else {
		cpu.bus.cartridge.loadBootROM()
	}
	title := "RoseGB - " + string(cpu.bus.cartridge.header.title[:])
	wnd := g.NewMasterWindow(title, 1000, 500, g.MasterWindowFlagsNotResizable, nil)

    go refresh()

    wnd.Main(loop)
}

func loop() {
	g.MainMenuBar(g.Layout{
		g.Menu("File", g.Layout{
			g.MenuItem("Exit", exit),
		}),
		
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
	g.Window("Debugger", 200, 30, 200, 300, g.Layout{
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
	})

	g.Window("Tile Viewer", 410, 30, 200, 300, g.Layout{
		g.Custom(func() {
			canvas := g.GetCanvas()
			pos := g.GetCursorScreenPos()
			if texture != nil {
				canvas.AddImage(tileTexture, pos.Add(image.Pt(0, 0)), pos.Add(image.Pt(128, 192)))
			}
		}),
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
    	for i := 0; i < 17556; i++ {
    		cpu.tick()
    		cpu.PPUTick()
    	}
    	cpu.drawTileViewer()
    	cpu.drawFramebuffer()
    	cpu.checkInput()
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
		Colour{255, 255, 255, 255}, Colour{192, 192, 192, 255}, Colour{96, 96, 96, 255}, Colour{0, 0, 0, 255},
	} 
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
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
			colour := color.RGBA{colours[tileColour].R, colours[tileColour].G, colours[tileColour].B, colours[tileColour].A}
			img.Set(j, i, colour)
		}
	}
	texture, _ = g.NewTextureFromRgba(img)
}

func (cpu *CPU) drawTileViewer() {
	var colours [4]Colour = [4]Colour{
		Colour{255, 255, 255, 255}, Colour{192, 192, 192, 255}, Colour{96, 96, 96, 255}, Colour{0, 0, 0, 255},
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
	if g.IsKeyPressed(90) == true && cpu.bus.keypad.getP15() == true {
		// fmt.Println("A!")
		cpu.bus.keypad.setA()
		cpu.bus.interrupt.requestJoypad()
	}
	// x, B
	// g.IsKeyPressed(88)
	// up, up
	// fmt.Println(g.IsKeyPressed(38), g.IsKeyPressed(90))
	if g.IsKeyPressed(38) {
		fmt.Println("up!")
		cpu.bus.keypad.setUp()
		cpu.bus.interrupt.requestJoypad()
	}
	// // down, down
	// g.IsKeyPressed(40)
	// // left, left
	// g.IsKeyPressed(37)
	// // right, right
	// g.IsKeyPressed(39)
}

