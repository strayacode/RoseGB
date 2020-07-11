package main

import (
	"fmt"
)

type Interrupt struct {
	IME byte
	IF byte
	IE byte
	IMEDelay bool
}

func (interrupt *Interrupt) requestVBlank() {
	interrupt.IF |= 1
}

func (interrupt *Interrupt) resetVBlank() {
	interrupt.IF &= 0xFE
	// fmt.Println(interrupt.IF & 1)
}

func (interrupt *Interrupt) requestLCDCSTAT() {
	interrupt.IF |= (1 << 1)
}

func (interrupt *Interrupt) resetLCDCSTAT() {
	interrupt.IF &= 0xFD
}

func (interrupt *Interrupt) requestTimer() {
	interrupt.IF |= (1 << 2)
}

func (interrupt *Interrupt) resetTimer() {
	interrupt.IF &= 0xFB
}

func (interrupt *Interrupt) requestSerial() {
	interrupt.IF |= (1 << 3)
}

func (interrupt *Interrupt) resetSerial() {
	interrupt.IF &= 0xF7
}

func (interrupt *Interrupt) requestJoypad() {
	interrupt.IF |= (1 << 4)
}

func (interrupt *Interrupt) resetJoypad() {
	interrupt.IF &= 0xEF
}


func (interrupt *Interrupt) executeVBlank() {
	hi := byte((cpu.PC) >> 8)
	lo := byte(((cpu.PC) & 0xFF))
	cpu.SP--
	cpu.bus.write(cpu.SP, hi)
	cpu.SP--
	cpu.bus.write(cpu.SP, lo)
	cpu.PC = 0x0040
	// fmt.Println("good")
}

func (interrupt *Interrupt) executeLCDCSTAT() {
	hi := byte((cpu.PC) >> 8)
	lo := byte(((cpu.PC) & 0xFF))
	cpu.SP--
	cpu.bus.write(cpu.SP, hi)
	cpu.SP--
	cpu.bus.write(cpu.SP, lo)
	cpu.PC = 0x0048
}

func (interrupt *Interrupt) executeTimer() {
	hi := byte((cpu.PC) >> 8)
	lo := byte(((cpu.PC) & 0xFF))
	cpu.SP--
	cpu.bus.write(cpu.SP, hi)
	cpu.SP--
	cpu.bus.write(cpu.SP, lo)
	cpu.PC = 0x0050
}

func (interrupt *Interrupt) executeSerial() {
	hi := byte((cpu.PC) >> 8)
	lo := byte(((cpu.PC) & 0xFF))
	cpu.SP--
	cpu.bus.write(cpu.SP, hi)
	cpu.SP--
	cpu.bus.write(cpu.SP, lo)
	cpu.PC = 0x0058
}

func (interrupt *Interrupt) executeJoypad() {
	hi := byte((cpu.PC) >> 8)
	lo := byte(((cpu.PC) & 0xFF))
	cpu.SP--
	cpu.bus.write(cpu.SP, hi)
	cpu.SP--
	cpu.bus.write(cpu.SP, lo)
	cpu.PC = 0x0060
	fmt.Println("GOOD!")
	
}


func (interrupt *Interrupt) handleInterrupts() {
	// only handle interrupts if IME is set
	if interrupt.IME == 1 {
		for i := 0; i < 5; i++ {
			// check if bit is a potential interrupt
			if ((1 << i) & interrupt.IE) >> i == 1 && ((1 << i) & interrupt.IF) >> i == 1 {
				// turn off corresponding bit in IF
				switch i {
				case 0:
					interrupt.executeVBlank()
				case 1:
					interrupt.executeLCDCSTAT()
				case 2:
					interrupt.executeTimer()
				case 3:
					interrupt.executeSerial()
				case 4:
					interrupt.executeJoypad()
				}

				switch i {
				case 0:
					interrupt.resetVBlank()
				case 1:
					interrupt.resetLCDCSTAT()
				case 2:
					interrupt.resetTimer()
				case 3:
					interrupt.resetSerial()
				case 4:
					interrupt.resetJoypad()
				}
				
				interrupt.IME = 0
			}
		}
	}
}

