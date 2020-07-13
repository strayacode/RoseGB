package main

type Timer struct {
	TIMA byte
	TMA byte
	TAC byte
	DIV byte
	elaspedCycles int
	timerInterrupt bool
	divCycles int
}

func (timer *Timer) tick() {
	// DIV check increment
	if timer.divCycles >= 256 {
		timer.DIV++
		timer.divCycles = 0
	}
	if timer.readTimerEnable() == true {
		threshold := timer.readFrequency()
		if timer.elaspedCycles >= threshold {
			timer.TIMA++
			if timer.TIMA == 0x00 {
				timer.TIMA = timer.TMA
				timer.timerInterrupt = true
			}
			timer.elaspedCycles -= threshold
			
		}
	}
}

func (timer *Timer) readFrequency() int {
	result := timer.TAC & 0x03
	switch result {
	case 0:
		return 1024
	case 1:
		return 16
	case 2:
		return 64
	case 3:
		return 256
	default:
		return 0
	}
}

func (timer *Timer) readTimerEnable() bool {
	return (((timer.TAC & (1 << 2)) >> 2) == 1)
}

func (timer *Timer) resetTimer() {
	timer.DIV, timer.TIMA, timer.TMA, timer.TAC = 0, 0, 0, 0
}