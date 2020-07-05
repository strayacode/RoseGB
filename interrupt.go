package main

type Interrupt struct {
	IME byte
	IF byte
	IE byte
}

func (interrupt *Interrupt) requestVBlank() {
	
}