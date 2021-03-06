package main

import (
	// "fmt"
	// "strconv"
)

type CPU struct {
	A byte
	B byte
	C byte
	D byte
	E byte
	F byte // used for storing z, n, h, c flags | ordered like znhc0000
	H byte
	L byte
	PC uint16
	SP uint16
	Cycles int
	Opcode byte
	bus Bus
	halt bool
}

func (cpu *CPU) tick() {
	// run instructions when cpu is not halted
	if !cpu.halt {
		if cpu.Cycles == 0 {
			
			if cpu.bus.interrupt.IMEDelay {
				cpu.bus.interrupt.IMEDelay = false
				cpu.bus.interrupt.IME = 1
			}
			cpu.Opcode = cpu.bus.read(cpu.PC)
			cpu.PC++
			cpu.Cycles = opcodes[cpu.Opcode].Cycles
			opcodes[cpu.Opcode].Exec(cpu)
			if cpu.bus.timer.readTimerEnable() {
				cpu.bus.timer.elaspedCycles += opcodes[cpu.Opcode].Cycles
			}
		}
			
	} else {
		// once interrupt has been serviced
		if cpu.bus.interrupt.IME == 0 {
			pendingInterrupts := (cpu.bus.interrupt.IE & cpu.bus.interrupt.IF) != 0
			if pendingInterrupts {
				cpu.halt = false
			}
		}
		if cpu.bus.timer.readTimerEnable() {
				cpu.bus.timer.elaspedCycles += 4
		}
		cpu.Cycles += 4
	}
	cpu.Cycles--
}



func (cpu *CPU) ZFlag(value byte) {
	if value == 1 {
		cpu.F |= (1 << 7)
	} else {
		cpu.F &= 0x7F
	}
}

func (cpu *CPU) NFlag(value byte) {
	if value == 1 {
		cpu.F |= (1 << 6)
	} else {
		cpu.F &= 0xBF
	}
}
func (cpu *CPU) HFlag(value byte) {
	if value == 1 {
		cpu.F |= (1 << 5)
	} else {
		cpu.F &= 0xDF
	}
}
func (cpu *CPU) CFlag(value byte) {
	if value == 1 {
		cpu.F |= (1 << 4)
	} else {
		cpu.F &= 0xEF
	}
}

func (cpu *CPU) getZFlag() byte {
	return (cpu.F & 0x80) >> 7
}

func (cpu *CPU) getNFlag() byte {
	return (cpu.F & 0x40) >> 6
}

func (cpu *CPU) getHFlag() byte {
	return (cpu.F & 0x20) >> 5
}

func (cpu *CPU) getCFlag() byte {
	return (cpu.F & 0x10) >> 4
}

