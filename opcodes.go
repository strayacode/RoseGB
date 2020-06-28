package main

import (
	"fmt"
	"os"
	"strconv"
)

type Opcode struct {
	Cycles int
	Exec func(*CPU)
}

func op_null(cpu *CPU) {
	fmt.Println("unimplemented opcode!", "0x" + strconv.FormatUint(uint64(cpu.Opcode), 16))
	cpu.debugCPU()
	os.Exit(3)
}

func cbop_null(cpu *CPU) {
	fmt.Println("unimplemented cb opcode!", "0x" + strconv.FormatUint(uint64(cpu.Opcode), 16))
	cpu.debugCPU()
	os.Exit(3)
}

// NOP DONE
func op_0x00(cpu *CPU) {

}

// LD BC, u16 DONE
func op_0x01(cpu *CPU) {
	cpu.B = cpu.bus.read(cpu.PC + 1)
	cpu.C = cpu.bus.read(cpu.PC)
	cpu.PC += 2
}

// INC B DONE
func op_0x04(cpu *CPU) {
	cpu.B++
	if cpu.B == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	if ((cpu.B & 0xF) + ((cpu.B - 1) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
}

// DEC B DONE
func op_0x05(cpu *CPU) {
	cpu.B--;
	if cpu.B == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.B & 0xF) + ((cpu.B + 1) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}


}

// LD B, u8 DONE
func op_0x06(cpu *CPU) {
	cpu.B = cpu.bus.read(cpu.PC)
	cpu.PC++
}

// INC C DONE
func op_0x0C(cpu *CPU) {
	cpu.C++
	if cpu.C == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	if ((cpu.C & 0xF) + ((cpu.C - 1) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
}

// DEC C DONE
func op_0x0D(cpu *CPU) {
	cpu.C--;
	if cpu.C == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.C & 0xF) + ((cpu.C + 1) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
}

// LD C, u8 DONE
func op_0x0E(cpu *CPU) {
	cpu.C = cpu.bus.read(cpu.PC)
	cpu.PC++
}

// LD DE, u16 DONE
func op_0x11(cpu *CPU) {
	cpu.D = cpu.bus.read(cpu.PC + 1)
	cpu.E = cpu.bus.read(cpu.PC)
	cpu.PC += 2
}

// LD (DE), A DONE
func op_0x12(cpu *CPU) {
	cpu.bus.write((uint16(cpu.D) << 8 | uint16(cpu.E)), cpu.A)
}

// INC DE DONE
func op_0x13(cpu *CPU) {
	cpu.E++
	if cpu.E == 0 {
		cpu.D++
	}
}

// INC D DONE
func op_0x14(cpu *CPU) {
	cpu.D++
	if cpu.D == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	if ((cpu.D & 0xF) + ((cpu.D - 1) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
}

// DEC D DONE
func op_0x15(cpu *CPU) {
	cpu.D--;
	if cpu.D == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.D & 0xF) + ((cpu.D + 1) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
}

// LD D, u8 DONE
func op_0x16(cpu *CPU) {
	cpu.D = cpu.bus.read(cpu.PC)
	cpu.PC++
}

// RLA DONE
func op_0x17(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	newcarry := cpu.A >> 7
	cpu.A = (cpu.A << 1 | tempcarry)
	cpu.CFlag(newcarry)
	cpu.ZFlag(0)
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// JR i8 DONE
func op_0x18(cpu *CPU) {
	offset := int8(cpu.bus.read(cpu.PC) + 1)
	cpu.PC += uint16(offset)
}

// LD A, (DE) DONE
func op_0x1A(cpu *CPU) {
	cpu.A = cpu.bus.read(uint16(cpu.D) << 8 | uint16(cpu.E))
}

// INC E DONE
func op_0x1C(cpu *CPU) {
	cpu.E++
	if cpu.E == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	if ((cpu.E & 0xF) + ((cpu.E - 1) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
}

// DEC E DONE
func op_0x1D(cpu *CPU) {
	cpu.E--;
	if cpu.E == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.E & 0xF) + ((cpu.E + 1) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}

}

// LD E, u8 DONE
func op_0x1E(cpu *CPU) {
	cpu.E = cpu.bus.read(cpu.PC)
	cpu.PC++
}

// RRA DONE
func op_0x1F(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	newcarry := (cpu.A & 0x01)
	cpu.CFlag(newcarry)
	cpu.A = (tempcarry | cpu.A >> 1)
	cpu.ZFlag(0)
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// JR NZ, i8 DONE
func op_0x20(cpu *CPU) {
	if cpu.getZFlag() == 0 {
		offset := int8(cpu.bus.read(cpu.PC) + 1)
		cpu.PC += uint16(offset)
		cpu.Cycles += 4
	} else {
		cpu.PC++
	}

}

// LD HL, u16 DONE
func op_0x21(cpu *CPU) {
	cpu.H = cpu.bus.read(cpu.PC + 1)
	cpu.L = cpu.bus.read(cpu.PC)
	cpu.PC += 2
}

// LD (HL+), A DONE
func op_0x22(cpu *CPU) {
	cpu.bus.write((uint16(cpu.H) << 8 | uint16(cpu.L)), cpu.A)
	cpu.L++
	if cpu.L == 0 {
		cpu.H++
	}
}

// INC HL DONE
func op_0x23(cpu *CPU) {
	cpu.L++
	if cpu.L == 0 {
		cpu.H++
	}
}

// INC H DONE
func op_0x24(cpu *CPU) {
	cpu.H++
	if cpu.H == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	if ((cpu.H & 0xF) + ((cpu.H - 1) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
}

// JR Z, i8 DONE
func op_0x28(cpu *CPU) {
	if cpu.getZFlag() == 1 {
		offset := int8(cpu.bus.read(cpu.PC) + 1)
		cpu.PC += uint16(offset)
		cpu.Cycles += 4
	} else {
		cpu.PC++
	}

}

// LD A, (HL+) DONE
func op_0x2A(cpu *CPU) {
	cpu.A = cpu.bus.read((uint16(cpu.H) << 8 | uint16(cpu.L)))
	cpu.L++
	if cpu.L == 0 {
		cpu.H++
	}

}

// LD L, u8
func op_0x2E(cpu *CPU) {
	cpu.L = cpu.bus.read(cpu.PC)
	cpu.PC++

}

// JR NC, i8 DONE
func op_0x30(cpu *CPU) {
	if cpu.getCFlag() == 0 {
		offset := int8(cpu.bus.read(cpu.PC) + 1)
		cpu.PC += uint16(offset)
		cpu.Cycles += 4
	} else {
		cpu.PC++
	}

}

// LD SP, u16 DONE
func op_0x31(cpu *CPU) {
	cpu.SP = cpu.bus.read16(cpu.PC)
	cpu.PC += 2
}

// LD (HL-), A DONE
func op_0x32(cpu *CPU) {
	cpu.bus.write((uint16(cpu.H) << 8 | uint16(cpu.L)), cpu.A)
	cpu.L--
	if cpu.L == 0xFF {
		cpu.H--
	}
}

// DEC A DONE
func op_0x3D(cpu *CPU) {
	cpu.A--;
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.A & 0xF) + ((cpu.A + 1) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}

}

// LD A, u8 DONE
func op_0x3E(cpu *CPU) {
	cpu.A = cpu.bus.read(cpu.PC)
	cpu.PC++
}

// LD B, A DONE
func op_0x47(cpu *CPU) {
	cpu.B = cpu.A
}

// LD C, A DONE
func op_0x4F(cpu *CPU) {
	cpu.C = cpu.A
}

// LD D, B DONE
func op_0x50(cpu *CPU) {
	cpu.D = cpu.B
}

// LD D, A DONE
func op_0x57(cpu *CPU) {
	cpu.D = cpu.A
}

// LD H, A DONE
func op_0x67(cpu *CPU) {
	cpu.H = cpu.A
}

// LD (HL), A DONE
func op_0x77(cpu *CPU) {
	cpu.bus.write((uint16(cpu.H) << 8 | uint16(cpu.L)), cpu.A)
}

// LD A, B DONE
func op_0x78(cpu *CPU) {
	cpu.A = cpu.B
}

// LD A, E DONE
func op_0x7B(cpu *CPU) {
	cpu.A = cpu.E
}

// LD A, H DONE
func op_0x7C(cpu *CPU) {
	cpu.A = cpu.H
}

// LD A, L DONE
func op_0x7D(cpu *CPU) {
	cpu.A = cpu.L
}

// ADD A, (HL)
func op_0x86(cpu *CPU) {
	cpu.A += cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	if ((cpu.A - cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))) & 0x08) != (cpu.A & 0x08) {
		cpu.HFlag(1)
	}
	if cpu.A & 0x80 != (cpu.A - cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))) & 0x80 {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// SUB A, B DONE
func op_0x90(cpu *CPU) {
	cpu.A -= cpu.B
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.A & 0xF) + ((cpu.A + cpu.B) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.B > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}

}

// XOR A, A DONE
func op_0xAF(cpu *CPU) {
	cpu.A ^= cpu.A
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)

}

// CP A, (HL) DONE
func op_0xBE(cpu *CPU) {
	result := cpu.A - cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((result & 0xF) + ((cpu.A) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L)) > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// POP BC DONE
func op_0xC1(cpu *CPU) {
	cpu.C = cpu.bus.read(cpu.SP)
	cpu.SP++
	cpu.B = cpu.bus.read(cpu.SP)
	cpu.SP++
}

// JP u16
func op_0xC3(cpu *CPU) {
	cpu.PC = cpu.bus.read16(cpu.PC)
}

// PUSH BC DONE
func op_0xC5(cpu *CPU) {
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.B)
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.C)
}

// RET DONE
func op_0xC9(cpu *CPU) {
	cpu.PC = cpu.bus.read16(cpu.SP)
	cpu.SP += 2
}

// PREFIX CB DONE
func op_0xCB(cpu *CPU) {
	cpu.Opcode = cpu.bus.read(cpu.PC)
	cbopcodes[cpu.Opcode].Exec(cpu)
	cpu.PC++
}

// CALL u16 DONE
func op_0xCD(cpu *CPU) {
	hi := byte((cpu.PC + 2) >> 8)
	lo := byte(((cpu.PC + 2) & 0xFF))
	cpu.SP--
	cpu.bus.write(cpu.SP, hi)
	cpu.SP--
	cpu.bus.write(cpu.SP, lo)
	cpu.PC = cpu.bus.read16(cpu.PC)
}

// ADC A, u8 DONE
func op_0xCE(cpu *CPU) {
	bitOverflow := uint16(cpu.A) + uint16(cpu.getCFlag()) + uint16(cpu.bus.read(cpu.PC))
	cpu.A += (cpu.getCFlag() + cpu.bus.read(cpu.PC))
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	if ((cpu.A & 0xF) + ((cpu.A - cpu.getCFlag() - cpu.bus.read(cpu.PC)) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}

	if bitOverflow > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}



}

// SUB A, u8 DONE
func op_0xD6(cpu *CPU) {
	cpu.A -= cpu.bus.read(cpu.PC)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.A & 0xF) + ((cpu.A + cpu.bus.read(cpu.PC)) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	}
	if cpu.bus.read(cpu.PC) > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}

}

// LD (FF00 + u8), A DONE
func op_0xE0(cpu *CPU) {
	cpu.bus.write((0xFF00 + uint16(cpu.bus.read(cpu.PC))), cpu.A)
	cpu.PC++
}

// POP HL DONE
func op_0xE1(cpu *CPU) {
	cpu.L = cpu.bus.read(cpu.SP)
	cpu.SP++
	cpu.H = cpu.bus.read(cpu.SP)
	cpu.SP++
}

// LD (FF00 + C), A DONE
func op_0xE2(cpu *CPU) {
	cpu.bus.write((0xFF00 + uint16(cpu.C)), cpu.A)
}

// PUSH HL DONE
func op_0xE5(cpu *CPU) {
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.H)
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.L)
}

// LD (u16), A DONE
func op_0xEA(cpu *CPU) {
	cpu.bus.write((uint16(cpu.bus.read(cpu.PC + 1))) << 8 | uint16((cpu.bus.read(cpu.PC))), cpu.A)
	cpu.PC += 2
}

// LD A, (FF00 + u8) DONE
func op_0xF0(cpu *CPU) {
	cpu.A = cpu.bus.read(0xFF00 + uint16(cpu.bus.read(cpu.PC)))
	cpu.PC++

}

// POP AF DONE
func op_0xF1(cpu *CPU) {
	cpu.F = cpu.bus.read(cpu.SP)
	cpu.SP++
	cpu.A = cpu.bus.read(cpu.SP)
	cpu.SP++
}

// DI DONE
func op_0xF3(cpu *CPU) {
	cpu.IME = 0
}

// PUSH AF DONE
func op_0xF5(cpu *CPU) {
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.A)
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.F)
}

// CP A, u8 DONE
func op_0xFE(cpu *CPU) {
	result := cpu.A - cpu.bus.read(cpu.PC)
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((result & 0xF) + ((cpu.A) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if cpu.bus.read(cpu.PC) > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.PC++
}

// RL C DONE
func cbop_0x11(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	newcarry := cpu.C >> 7
	cpu.C = (cpu.C << 1 | tempcarry)
	cpu.CFlag(newcarry)
	if cpu.C == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
}

// BIT 7, H DONE
func cbop_0x7C(cpu *CPU) {
	bit := (cpu.H & 0x80) >> 7
	if bit == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
}

var opcodes [256]Opcode = [256]Opcode {
	Opcode{4, op_0x00}, Opcode{12, op_0x01}, Opcode{8, op_null}, Opcode{8, op_null}, Opcode{4, op_0x04}, Opcode{4, op_0x05}, Opcode{8, op_0x06}, Opcode{4, op_null}, Opcode{20, op_null}, Opcode{8, op_null}, Opcode{8, op_null}, Opcode{8, op_null}, Opcode{4, op_0x0C}, Opcode{4, op_0x0D}, Opcode{8, op_0x0E}, Opcode{4, op_null},
	Opcode{4, op_null}, Opcode{12, op_0x11}, Opcode{8, op_0x12}, Opcode{8, op_0x13}, Opcode{4, op_0x14}, Opcode{4, op_0x15}, Opcode{8, op_0x16}, Opcode{4, op_0x17}, Opcode{12, op_0x18}, Opcode{8, op_null}, Opcode{8, op_0x1A}, Opcode{8, op_null}, Opcode{4, op_0x1C}, Opcode{4, op_0x1D}, Opcode{8, op_0x1E}, Opcode{4, op_0x1F},
	Opcode{8, op_0x20}, Opcode{12, op_0x21}, Opcode{8, op_0x22}, Opcode{8, op_0x23}, Opcode{4, op_0x24}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{4, op_null}, Opcode{8, op_0x28}, Opcode{8, op_null}, Opcode{8, op_0x2A}, Opcode{8, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_0x2E}, Opcode{4, op_null},
	Opcode{8, op_0x30}, Opcode{12, op_0x31}, Opcode{8, op_0x32}, Opcode{8, op_null}, Opcode{12, op_null}, Opcode{12, op_null}, Opcode{12, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{8, op_null}, Opcode{8, op_null}, Opcode{8, op_null}, Opcode{4, op_null}, Opcode{4, op_0x3D}, Opcode{8, op_0x3E}, Opcode{4, op_null},
	Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{4, op_0x47}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{4, op_0x4F},
	Opcode{4, op_0x50}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{4, op_0x57}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{4, op_null},
	Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{4, op_0x67}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{4, op_null},
	Opcode{8, op_null}, Opcode{8, op_null}, Opcode{8, op_null}, Opcode{8, op_null}, Opcode{8, op_null}, Opcode{8, op_null}, Opcode{4, op_null}, Opcode{8, op_0x77}, Opcode{4, op_0x78}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_0x7B}, Opcode{4, op_0x7C}, Opcode{4, op_0x7D}, Opcode{8, op_null}, Opcode{4, op_null},
	Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_0x86}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{4, op_null},
	Opcode{4, op_0x90}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{4, op_null},
	Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{4, op_0xAF},
	Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_0xBE}, Opcode{4, op_null},
	Opcode{8, op_null}, Opcode{12, op_0xC1}, Opcode{12, op_null}, Opcode{16, op_0xC3}, Opcode{12, op_null}, Opcode{16, op_0xC5}, Opcode{8, op_null}, Opcode{16, op_null}, Opcode{8, op_null}, Opcode{16, op_0xC9}, Opcode{12, op_null}, Opcode{4, op_0xCB}, Opcode{12, op_null}, Opcode{24, op_0xCD}, Opcode{8, op_0xCE}, Opcode{16, op_null},
	Opcode{8, op_null}, Opcode{12, op_null}, Opcode{12, op_null}, Opcode{8, op_null}, Opcode{12, op_null}, Opcode{16, op_null}, Opcode{8, op_0xD6}, Opcode{16, op_null}, Opcode{8, op_null}, Opcode{16, op_null}, Opcode{12, op_null}, Opcode{8, op_null}, Opcode{12, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{16, op_null},
	Opcode{12, op_0xE0}, Opcode{12, op_0xE1}, Opcode{8, op_0xE2}, Opcode{8, op_null}, Opcode{4, op_null}, Opcode{16, op_0xE5}, Opcode{8, op_null}, Opcode{16, op_null}, Opcode{16, op_null}, Opcode{4, op_null}, Opcode{16, op_0xEA}, Opcode{8, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{16, op_null},
	Opcode{12, op_0xF0}, Opcode{12, op_0xF1}, Opcode{8, op_null}, Opcode{4, op_0xF3}, Opcode{4, op_null}, Opcode{16, op_0xF5}, Opcode{8, op_null}, Opcode{16, op_null}, Opcode{12, op_null}, Opcode{8, op_null}, Opcode{16, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_0xFE}, Opcode{16, op_null},

}

var cbopcodes [256]Opcode = [256]Opcode {
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_0x11}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{12, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{12, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{12, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{12, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{12, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{12, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{12, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_0x7C}, Opcode{8, cbop_null}, Opcode{12, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null},

}