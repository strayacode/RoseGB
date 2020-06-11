type Cartridge struct {
	ROM [0x7FFF]byte
	ERAM [0x2000]byte

}

type Header struct {
	Title [16]byte // title of cartridge in UPPER CASE ASCII
	CartridgeType byte // type of cartridge
	ROMSize byte // specify size of ROM (could have multiple banks)
	RAMSize byte // specify size of ERAM (could have multiple banks)

}