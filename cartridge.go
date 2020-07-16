package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"flag"
	"strings"
)

type Cartridge struct {
	ROM [0x4000]byte
	header Header
	rombank ROMBank
	rambank RAMBank
}

type Header struct {
	title [16]byte // title of cartridge in UPPER CASE ASCII
	cartridgeType byte // type of cartridge
	ROMSize byte // specify size of ROM (could have multiple banks)
	RAMSize byte // specify size of ERAM (could have multiple banks)
	isCGB bool

}

type ROMBank struct {
	bankptr uint16 // 0-256
	bank [512][0x4000]byte
}

type RAMBank struct {
	bankptr byte // 0-16
	bank [16][0x2000]byte
}

func (cartridge *Cartridge) loadBootROM() {
	cartridge.loadCartridge()
	_, err := os.Stat("bios.rom")
	if os.IsNotExist(err) {
		fmt.Println("no bios file detected!")
		os.Exit(3)
	}
	file, err := ioutil.ReadFile("bios.rom")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(file); i++ {
		cartridge.ROM[i] = file[i]
		cartridge.rombank.bank[0][i] = file[i]
	}

}

func (cartridge *Cartridge) loadCartridge() {
	rom := flag.Args()[0]
	_, err := os.Stat(rom)
	if os.IsNotExist(err) {
		fmt.Println(rom, "does not exist!")
		os.Exit(3)
	}
	if strings.HasSuffix(rom, ".gb") == false {
		fmt.Println("a .gb file is required!")
		os.Exit(3)
	}

	file, err := ioutil.ReadFile(rom)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 16; i++ {
		cartridge.header.title[i] = file[i + 0x134]
	}
	cartridge.header.cartridgeType = file[0x147]

	cartridge.header.ROMSize = file[0x148]
	cartridge.header.RAMSize = file[0x149]
	if file[0x143] == 0x80 {
		cartridge.header.isCGB = true
	} else {
		cartridge.header.isCGB = false
	}

	for i := 0; i < 0x4000; i++ {
		cartridge.ROM[i] = file[i]
		cartridge.rombank.bank[0][i] = file[i]
	}
	for i := 1; i < 512; i++ {
		for j := 0; j < 0x4000; j++ {
			if (i * 0x4000) + j < len(file) {
				cartridge.rombank.bank[i][j] = file[(i * 0x4000) + j]
			}
		}
	}
}

func (cartridge *Cartridge) unmapBootROM() {
	rom := flag.Args()[0]
	file, err := ioutil.ReadFile(rom)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 0xFF; i++ {
		cartridge.ROM[i] = file[i]
		cartridge.rombank.bank[0][i] = file[i]
	}

	

}