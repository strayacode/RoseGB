package main

type Keypad struct {
	P1 byte
	
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

func (keypad *Keypad) setDown() {
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


