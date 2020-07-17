package main

// TODO: MBC1, 128 ROM banks (max 2mb), 4 RAM banks (max 32kb)


import (
	"fmt"
	"os"
)

type Bus struct {
	cartridge Cartridge
	WRAM [0x2000]byte
	ppu PPU
	apu APU
	HRAM [0x80]byte
	timer Timer
	interrupt Interrupt
	keypad Keypad
	SB byte
	SC byte
	KEY1 byte
	enableERAM bool
	bankingMode byte
}


func (bus *Bus) read(addr uint16) byte {
	switch {
	case addr >= 0x0000 && addr <= 0x3FFF:
		return bus.cartridge.ROM[addr]
	case addr >= 0x4000 && addr <= 0x7FFF:
		return bus.cartridge.rombank.bank[bus.cartridge.rombank.bankptr][addr - 0x4000]
	case addr >= 0x8000 && addr <= 0x9FFF:
		return bus.ppu.VRAM[addr - 0x8000]
	case addr >= 0xA000 && addr <= 0xBFFF:
		switch bus.cartridge.header.cartridgeType {
		case 0:
			return bus.cartridge.rambank.bank[bus.cartridge.rambank.bankptr][addr - 0xA000]
		case 1, 2, 3, 0x11, 0x12, 0x13, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E:
			if bus.enableERAM == true {
				return bus.cartridge.rambank.bank[bus.cartridge.rambank.bankptr][addr - 0xA000]
			} else {
				return 0xFF
			}
		default:
			return 0
		}
	case addr >= 0xC000 && addr <= 0xDFFF:
		return bus.WRAM[addr - 0xC000]
	case addr >= 0xE000 && addr <= 0xFDFF:
		return bus.WRAM[addr - 0xE000]
	case addr >= 0xFE00 && addr <= 0xFE9F:
		if bus.ppu.cpuOAMAccess {
			return bus.ppu.OAM[addr - 0xFE00]
		}
	case addr >= 0xFF00 && addr <= 0xFF7F:
		return bus.readIO(addr)
	case addr >= 0xFF80 && addr <= 0xFFFE:
		return bus.HRAM[addr - 0xFF80]
	case addr == 0xFFFF:
		return bus.interrupt.IE
	default:
		fmt.Println("DEBUG: non-readable memory location!", addr)
		return 0
	}
	return 0
}

func (bus *Bus) read16(addr uint16) uint16 {
	switch {
	case addr >= 0x0000 && addr <= 0x3FFF:
		return uint16(bus.cartridge.ROM[addr + 1]) << 8 | uint16(bus.cartridge.ROM[addr])
	case addr >= 0x4000 && addr <= 0x7FFF:
		return uint16(bus.cartridge.rombank.bank[bus.cartridge.rombank.bankptr][addr + 1 - 0x4000]) << 8 | uint16(bus.cartridge.rombank.bank[bus.cartridge.rombank.bankptr][addr - 0x4000])
	case addr >= 0x8000 && addr <= 0x9FFF:
		return uint16(bus.ppu.VRAM[addr + 1 - 0x8000]) << 8 | uint16(bus.ppu.VRAM[addr - 0x8000])
	case addr >= 0xA000 && addr <= 0xBFFF:
		return uint16(bus.cartridge.rambank.bank[bus.cartridge.rambank.bankptr][addr + 1 - 0xA000]) << 8 | uint16(bus.cartridge.rambank.bank[bus.cartridge.rambank.bankptr][addr - 0xA000])
	case addr >= 0xC000 && addr <= 0xDFFF:
		return uint16(bus.WRAM[addr + 1 - 0xC000]) << 8 | uint16(bus.WRAM[addr - 0xC000])
	case addr >= 0xE000 && addr <= 0xFDFF:
		return uint16(bus.WRAM[addr + 1 - 0xE000]) << 8 | uint16(bus.WRAM[addr - 0xE000])
	case addr >= 0xFF80 && addr <= 0xFFFE:
		return uint16(bus.HRAM[addr + 1 - 0xFF80]) << 8 | uint16(bus.HRAM[addr - 0xFF80])
	default:
		fmt.Println("DEBUG: non-16bit readable memory location!", addr)
		os.Exit(3)
		return 0
	}
}

func (bus *Bus) write(addr uint16, data byte) {
	switch {
	case addr >= 0x0000 && addr <= 0x1FFF:
		// MBC1
		switch bus.cartridge.header.cartridgeType {
		case 1, 2, 3, 0x11, 0x12, 0x13, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E:
			if (data & 0xF) == 0xA {
				bus.enableERAM = true
			} else {
				bus.enableERAM = false
			}
		}
	case addr >= 0x2000 && addr <= 0x3FFF:
		// MBC1
		switch bus.cartridge.header.cartridgeType {
		case 0:
		case 1:
			// if cart is more than 32kb use the secondary register to address ptr with more than 5 bits
			if bus.cartridge.header.ROMSize > 0x04 {
				if data == 0x00 {
					bus.cartridge.rombank.bankptr = 0x01
				} else {
					bus.cartridge.rombank.bankptr = (uint16(bus.cartridge.rambank.bankptr) << 5 | (uint16(data) & 0x1F))
				}
			} else {
				if data == 0x00 {
					bus.cartridge.rombank.bankptr = 0x01
				} else {
					bus.cartridge.rombank.bankptr = (uint16(data) & 0x1F)
				}
			}
		case 0x11, 0x12, 0x13:
			if data == 0x00 {
				bus.cartridge.rombank.bankptr = 0x01
			} else {
				bus.cartridge.rombank.bankptr = uint16(data) & 0x7F
			}
		}
	case addr >= 0x2000 && addr <= 0x2FFF:
		// MBC1
		switch bus.cartridge.header.cartridgeType {
		
		case 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E:
			bus.cartridge.rombank.bankptr = (bus.cartridge.rombank.bankptr & 0x100) | uint16(data)
		}
	case addr >= 0x3000 && addr <= 0x3FFF:
		// MBC1
		switch bus.cartridge.header.cartridgeType {
		
		case 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E:
			bus.cartridge.rombank.bankptr = (bus.cartridge.rombank.bankptr & 0xFF) | (uint16(data) << 8)
		}
	case addr >= 0x4000 && addr <= 0x5FFF:
		// MBC1
		switch bus.cartridge.header.cartridgeType {
			case 1, 2, 3:
				// ram banking only applies for 32kb ram
				if bus.cartridge.header.RAMSize == 0x03 {
					bus.cartridge.rambank.bankptr = (data & 0x03)
				}
			case 0x11, 0x12, 0x13:
				bus.cartridge.rambank.bankptr = (data & 0x03)
			case 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E:
				bus.cartridge.rambank.bankptr = (data & 0xF)
			default:

		}
	case addr >= 0x6000 && addr <= 0x7FFF:
		// MBC1
		if bus.cartridge.header.cartridgeType == 1 {
			bus.bankingMode = (data & 0x1)
		}
	case addr >= 0x8000 && addr <= 0x9FFF:
		bus.ppu.VRAM[addr - 0x8000] = data

	case addr >= 0xA000 && addr <= 0xBFFF:
		// MBC1
		switch bus.cartridge.header.cartridgeType {
		case 0:
			bus.cartridge.rambank.bank[bus.cartridge.rambank.bankptr][addr - 0xA000] = data
		case 1, 2, 3, 0x11, 0x12, 0x13, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E:

			if bus.enableERAM == true {
				bus.cartridge.rambank.bank[bus.cartridge.rambank.bankptr][addr - 0xA000] = data
			}
		}

	case addr >= 0xC000 && addr <= 0xDFFF:
		bus.WRAM[addr - 0xC000] = data
	case addr >= 0xE000 && addr <= 0xFDFF:
		bus.WRAM[addr - 0xE000] = data
	case addr >= 0xFE00 && addr <= 0xFE9F:
		if bus.ppu.cpuOAMAccess {
			bus.ppu.OAM[addr - 0xFE00] = data
		}
	case addr >= 0xFF00 && addr <= 0xFF7F:
		bus.writeIO(addr, data)
	case addr >= 0xFF80 && addr <= 0xFFFE:
		bus.HRAM[addr - 0xFF80] = data
	case addr == 0xFFFF:
		bus.interrupt.IE = data
	default:
	}
}

func (bus *Bus) readIO(addr uint16) byte {
	switch addr {
	case 0xFF00:
		// fmt.Println(bus.keypad.P1)
		// override the buttons
		if bus.keypad.getP14() {
			// fmt.Println("good")
			bus.keypad.P1 &= (0xF0 | bus.keypad.direction[3] << 3 | bus.keypad.direction[2] << 2 | bus.keypad.direction[1] << 1 | bus.keypad.direction[0])
		}
		if bus.keypad.getP15() {
			bus.keypad.P1 &= (0xF0 | bus.keypad.button[3] << 3 | bus.keypad.button[2] << 2 | bus.keypad.button[1] << 1 | bus.keypad.button[0])
		} 

		return bus.keypad.P1
	case 0xFF01:
		return bus.SB
	case 0xFF02:
		return bus.SC
	case 0xFF04:
		return bus.timer.DIV
	case 0xFF0F:
		return bus.interrupt.IF
	case 0xFF10:
		return bus.apu.NR10 
	case 0xFF11:
		return bus.apu.NR11 
	case 0xFF12:
		return bus.apu.NR12 
	case 0xFF13:
		return bus.apu.NR13 
	case 0xFF14:
		return bus.apu.NR14 
	case 0xFF16:
		return bus.apu.NR21 
	case 0xFF17:
		return bus.apu.NR22 
	case 0xFF19:
		return bus.apu.NR24 
	case 0xFF1A:
		return bus.apu.NR30 
	case 0xFF1B:
		return bus.apu.NR31 
	case 0xFF1C:
		return bus.apu.NR32 
	case 0xFF1D:
		return bus.apu.NR33
	case 0xFF1E:
		return bus.apu.NR34 
	case 0xFF20:
		return bus.apu.NR41 
	case 0xFF21:
		return bus.apu.NR42 
	case 0xFF22:
		return bus.apu.NR43 
	case 0xFF23:
		return bus.apu.NR44 
	case 0xFF24:
		return bus.apu.NR50 
	case 0xFF25:
		return bus.apu.NR51 
	case 0xFF26:
		return bus.apu.NR52 
	case 0xFF40:
		return bus.ppu.LCDC
	case 0xFF41:
		return bus.ppu.LCDCSTAT
	case 0xFF42:
		return bus.ppu.SCY
	case 0xFF43:
		return bus.ppu.SCX
	case 0xFF44:
		return bus.ppu.LY
	case 0xFF45:
		return bus.ppu.LYC
	case 0xFF47:
		return bus.ppu.BGP
	case 0xFF48:
		return bus.ppu.OBP0
	case 0xFF49:
		return bus.ppu.OBP1
	case 0xFF4A:
		return bus.ppu.WY
	case 0xFF4B:
		return bus.ppu.WX
	case 0xFF4D:
		return bus.KEY1
	default:
		fmt.Println(addr, "IO read not implemented yet!")
		// os.Exit(3)
		return 0xFF
	}
}

func (bus *Bus) writeIO(addr uint16, data byte) {
	switch addr {
	case 0xFF00:
		bus.keypad.P1 = (data & 0xF0) | 0xF
	case 0xFF01:
		bus.SB = data
	case 0xFF02:
		bus.SC = data
	case 0xFF04:
		bus.timer.DIV = 0
	case 0xFF05:
		bus.timer.TIMA = data
	case 0xFF06:
		bus.timer.TMA = data
	case 0xFF07:
		bus.timer.TAC = data
	case 0xFF0F:
		bus.interrupt.IF = data
	case 0xFF10:
		bus.apu.NR10 = data
	case 0xFF11:
		bus.apu.NR11 = data
	case 0xFF12:
		bus.apu.NR12 = data
	case 0xFF13:
		bus.apu.NR13 = data
	case 0xFF14:
		bus.apu.NR14 = data
	case 0xFF16:
		bus.apu.NR21 = data
	case 0xFF17:
		bus.apu.NR22 = data
	case 0xFF19:
		bus.apu.NR24 = data
	case 0xFF1A:
		bus.apu.NR30 = data
	case 0xFF1B:
		bus.apu.NR31 = data
	case 0xFF1C:
		bus.apu.NR32 = data
	case 0xFF1D:
		bus.apu.NR33 = data
	case 0xFF1E:
		bus.apu.NR34 = data
	case 0xFF20:
		bus.apu.NR41 = data
	case 0xFF21:
		bus.apu.NR42 = data
	case 0xFF22:
		bus.apu.NR43 = data
	case 0xFF23:
		bus.apu.NR44 = data
	case 0xFF24:
		bus.apu.NR50 = data
	case 0xFF25:
		bus.apu.NR51 = data
	case 0xFF26:
		bus.apu.NR52 = data
	case 0xFF30:
		bus.apu.WAVEPATTERNRAM[0] = data
	case 0xFF31:
		bus.apu.WAVEPATTERNRAM[1] = data
	case 0xFF32:
		bus.apu.WAVEPATTERNRAM[2] = data
	case 0xFF33:
		bus.apu.WAVEPATTERNRAM[3] = data
	case 0xFF34:
		bus.apu.WAVEPATTERNRAM[4] = data
	case 0xFF35:
		bus.apu.WAVEPATTERNRAM[5] = data
	case 0xFF36:
		bus.apu.WAVEPATTERNRAM[6] = data
	case 0xFF37:
		bus.apu.WAVEPATTERNRAM[7] = data
	case 0xFF38:
		bus.apu.WAVEPATTERNRAM[8] = data
	case 0xFF39:
		bus.apu.WAVEPATTERNRAM[9] = data
	case 0xFF3A:
		bus.apu.WAVEPATTERNRAM[10] = data
	case 0xFF3B:
		bus.apu.WAVEPATTERNRAM[11] = data
	case 0xFF3C:
		bus.apu.WAVEPATTERNRAM[12] = data
	case 0xFF3D:
		bus.apu.WAVEPATTERNRAM[13] = data
	case 0xFF3E:
		bus.apu.WAVEPATTERNRAM[14] = data
	case 0xFF3F:
		bus.apu.WAVEPATTERNRAM[15] = data
	case 0xFF40:
		bus.ppu.LCDC = data
	case 0xFF41:
		bus.ppu.LCDCSTAT = (data & 0xF8) | bus.ppu.LCDCSTAT

	case 0xFF42:
		bus.ppu.SCY = data
	case 0xFF43:
		bus.ppu.SCX = data
	case 0xFF44:
		bus.ppu.LY = data
	case 0xFF45:
		bus.ppu.LYC = data
	case 0xFF46:
		bus.ppu.DMA = data
		bus.ppu.dmaTransfer()
		// fmt.Println(data, "dma")
	case 0xFF47:
		bus.ppu.BGP = data
	case 0xFF48:
		bus.ppu.OBP0 = data
	case 0xFF49:
		bus.ppu.OBP1 = data
	case 0xFF4A:
		bus.ppu.WY = data
	case 0xFF4B:
		bus.ppu.WX = data
	case 0xFF4D:
		bus.KEY1 = data
	case 0xFF50:
		bus.cartridge.unmapBootROM()
	default:
	}
	
}