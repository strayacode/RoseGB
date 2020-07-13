package main

type Keypad struct {
	P1 byte
	button [4]byte // a, b, select, start
	direction [4]byte // right, left, up, down
	
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

func (keypad *Keypad) setUp(value byte) {
	keypad.direction[2] = value
}

func (keypad *Keypad) setSelect(value byte) {
	keypad.button[2] = value
}

func (keypad *Keypad) setDown(value byte) {
	keypad.direction[3] = value
}

func (keypad *Keypad) setStart(value byte) {
	keypad.button[3] = value
}

func (keypad *Keypad) setLeft(value byte) {
	keypad.direction[1] = value
}

func (keypad *Keypad) setRight(value byte) {
	keypad.direction[0] = value
}

func (keypad *Keypad) setA(value byte) {
	keypad.button[0] = value
}

func (keypad *Keypad) setB(value byte) {
	keypad.button[1] = value
}

// yes, for example, P14 is associated with (in order from P10), right, left, up, down

// if I press any combination of these buttons when P14 is set to 1 (which is the default state most of the game), reading from rP1 will yield the lower nybble as being %xxxx1111, all 1s which means no button pressed

// if I press right while P14 is 0, and read from rP1, I'll get back %xxxx1110, right+left yields %xxxx1100



