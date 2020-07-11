package main

type Timer struct {
	TIMA byte
	TMA byte
	TAC byte
	DIV byte
	elaspedCycles int
	timerInterrupt bool
}

func (timer *Timer) tick() {
	if timer.readTimerEnable() == true {
		threshold := timer.readFrequency()
		if timer.elaspedCycles >= threshold {
			if timer.TIMA == 0xFF {
				timer.TIMA = timer.TMA
				// request timer interrupt
				timer.timerInterrupt = true
			} else {
				timer.TIMA++
				timer.elaspedCycles -= threshold
			}
		}
	}
}

func (timer *Timer) readFrequency() int {
	result := timer.TAC & 0x03
	switch result {
	case 0:
		return 4096
	case 1:
		return 262144
	case 2:
		return 65536
	case 3:
		return 16384
	default:
		return 0
	}
}

func (timer *Timer) readTimerEnable() bool {
	return (((timer.TAC & (1 << 2)) >> 2) == 1)
}