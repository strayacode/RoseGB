package main


type Memory struct {
	bootROM [0xFF]byte
	
	
	
	WRAM [0x2000]byte
	IO []
	HRAM []
	


}

func (memory *Memory) read(addr uint16) byte {
	switch {
		case addr >= 0x0000 && addr <= 0x7FFF:
			return memory.ROM[addr]
		case addr >= 0x8000 && addr <= 
	}
}