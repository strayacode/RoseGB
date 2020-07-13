package main

type Keypad struct {
	P1 byte
	button [4]bool // a, b, select, start
	direction [4]bool // right, left, up, down
	
}

func (keypad *Keypad) getP15() bool {
	// check if buttons are allowed
	if (keypad.P1 & 0x20) >> 5 == 0 {
		return true
	}
	return false
}

func (keypad *Keypad) getP14() bool {
	// check if direction keys are allowed
	if (keypad.P1 & 0x10) >> 4 == 0 {
		return true
	}
	return false
}

func (keypad *Keypad) setUp() {
	keypad.P1 &= 0xFB
}

func (keypad *Keypad) setSelect() {
	keypad.P1 &= 0xFB
}

func (keypad *Keypad) setDown() {
	keypad.P1 &= 0xF7
}

func (keypad *Keypad) setStart() {
	keypad.P1 &= 0xF7
}

func (keypad *Keypad) setLeft() {
	keypad.P1 &= 0xFD
}

func (keypad *Keypad) setRight() {
	keypad.P1 &= 0xFE
}

func (keypad *Keypad) setA() {
	keypad.P1 &= 0xFE
}

func (keypad *Keypad) setB() {
	keypad.P1 &= 0xFD
}

// yes, for example, P14 is associated with (in order from P10), right, left, up, down

// if I press any combination of these buttons when P14 is set to 1 (which is the default state most of the game), reading from rP1 will yield the lower nybble as being %xxxx1111, all 1s which means no button pressed

// if I press right while P14 is 0, and read from rP1, I'll get back %xxxx1110, right+left yields %xxxx1100



