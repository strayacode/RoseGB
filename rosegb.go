package main

import (
	g "github.com/AllenDang/giu"
	"time"
	"image"
	"image/color"
	"strconv"
	"os"
	"flag"
)

var (
	upLeft = image.Point{0, 0}
	lowRight = image.Point{160, 144}
	texture *g.Texture
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
	wnd := g.NewMasterWindow(title, 800, 400, g.MasterWindowFlagsNotResizable, nil)

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
	g.Window("Debugger", 300, 30, 200, 300, g.Layout{
		g.Label("A: " + strconv.Itoa(int(cpu.A))),
		g.Label("B: " + strconv.Itoa(int(cpu.B))),
		g.Label("C: " + strconv.Itoa(int(cpu.C))),
		g.Label("D: " + strconv.Itoa(int(cpu.D))),
		g.Label("E: " + strconv.Itoa(int(cpu.E))),
		g.Label("F: " + strconv.Itoa(int(cpu.F))),
		g.Label("H: " + strconv.Itoa(int(cpu.H))),
		g.Label("L: " + strconv.Itoa(int(cpu.L))),
		g.Label("LCDC: " + strconv.Itoa(int(cpu.bus.ppu.LCDC))),
		g.Label("LCDCSTAT: " + strconv.Itoa(int(cpu.bus.ppu.LCDCSTAT))),
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
    		
    		// cpu.setPCBreakpoint(0x740)
    		cpu.bus.ppu.tick()
    	}
    	
    	cpu.drawFramebuffer()
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

