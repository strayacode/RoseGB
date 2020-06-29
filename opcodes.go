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

// INC BC DONE
func op_0x03(cpu *CPU) {
	cpu.C++
	if cpu.C == 0 {
		cpu.B++
	}
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

// RLCA DONE
func op_0x07(cpu *CPU) {
	tempcarry := cpu.A >> 7
	cpu.CFlag(tempcarry)
	cpu.A = (cpu.A << 1 | tempcarry)
	cpu.ZFlag(0)
	cpu.NFlag(0)
	cpu.HFlag(0)

}

// LD (u16), SP DONE
func op_0x08(cpu *CPU) {
	cpu.bus.write(cpu.bus.read16(cpu.PC), byte(cpu.SP & 0xFF))
	cpu.bus.write(cpu.bus.read16(cpu.PC) + 1, byte(cpu.SP >> 8))
	cpu.PC += 2
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

// DEC H DONE
func op_0x25(cpu *CPU) {
	cpu.H--;
	if cpu.H == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.H & 0xF) + ((cpu.H + 1) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
}

// LD H, u8
func op_0x26(cpu *CPU) {
	cpu.H = cpu.bus.read(cpu.PC)
	cpu.PC++
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

// ADD HL, HL WRONG
func op_0x29(cpu *CPU) {
	hl := uint32(uint16(cpu.H) << 8 | uint16(cpu.L))
	hl += uint32((uint16(cpu.H) << 8 | uint16(cpu.L)))
	if hl > 0xFFFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	if ((hl & 0xF0FF) + ((hl / 2) & 0xF0FF)) & 0xEFFF == 0xEFFF {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	cpu.NFlag(0)
	hl &= 0xFF
	cpu.H = (byte(hl) >> 8)
	cpu.L = (byte(hl) & 0xFF)
}

// LD A, (HL+) DONE
func op_0x2A(cpu *CPU) {
	cpu.A = cpu.bus.read((uint16(cpu.H) << 8 | uint16(cpu.L)))
	cpu.L++
	if cpu.L == 0 {
		cpu.H++
	}
}

// INC L DONE
func op_0x2C(cpu *CPU) {
	cpu.L++
	if cpu.L == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	if ((cpu.L & 0xF) + ((cpu.L - 1) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
}

// DEC L
func op_0x2D(cpu *CPU) {
	cpu.L--;
	if cpu.L == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((cpu.L & 0xF) + ((cpu.L + 1) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
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

// DEC (HL) DONE
func op_0x35(cpu *CPU) {
	result := cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	result--
	result &= 0xFF
	cpu.bus.write(uint16(cpu.H) << 8 | uint16(cpu.L), result)
	if result == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(1)
	if ((result & 0xF) + ((result + 1) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}


}

// JR C, i8
func op_0x38(cpu *CPU) {
	if cpu.getCFlag() == 1 {
		offset := int8(cpu.bus.read(cpu.PC) + 1)
		cpu.PC += uint16(offset)
		cpu.Cycles += 4
	} else {
		cpu.PC++
	}
}

// INC A DONE
func op_0x3C(cpu *CPU) {
	cpu.A++
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	if ((cpu.A & 0xF) + ((cpu.A - 1) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
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

// LD B, B DONE
func op_0x40(cpu *CPU) {
	cpu.B = cpu.B
}

// LD B, C DONE
func op_0x41(cpu *CPU) {
	cpu.B = cpu.C
}

// LD B, D DONE
func op_0x42(cpu *CPU) {
	cpu.B = cpu.D
}

// LD B, E DONE
func op_0x43(cpu *CPU) {
	cpu.B = cpu.E
}

// LD B, H DONE
func op_0x44(cpu *CPU) {
	cpu.B = cpu.H
}

// LD B, L DONE
func op_0x45(cpu *CPU) {
	cpu.B = cpu.L
}

// LD B, (HL) DONE
func op_0x46(cpu *CPU) {
	cpu.B = cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
}

// LD B, A DONE
func op_0x47(cpu *CPU) {
	cpu.B = cpu.A
}

// LD C, B DONE
func op_0x48(cpu *CPU) {
	cpu.C = cpu.B
}

// LD C, C DONE
func op_0x49(cpu *CPU) {
	cpu.C = cpu.C
}

// LD C, D DONE
func op_0x4A(cpu *CPU) {
	cpu.C = cpu.D
}

// LD C, E DONE
func op_0x4B(cpu *CPU) {
	cpu.C = cpu.E
}

// LD C, H DONE
func op_0x4C(cpu *CPU) {
	cpu.C = cpu.H
}

// LD C, L DONE
func op_0x4D(cpu *CPU) {
	cpu.C = cpu.L
}

// LD C, (HL) DONE
func op_0x4E(cpu *CPU) {
	cpu.C = cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
}

// LD C, A DONE
func op_0x4F(cpu *CPU) {
	cpu.C = cpu.A
}

// LD D, B DONE
func op_0x50(cpu *CPU) {
	cpu.D = cpu.B
}

// LD D, C DONE
func op_0x51(cpu *CPU) {
	cpu.D = cpu.C
}

// LD D, D DONE
func op_0x52(cpu *CPU) {
	cpu.D = cpu.D
}

// LD D, E DONE
func op_0x53(cpu *CPU) {
	cpu.D = cpu.E
}

// LD D, H DONE
func op_0x54(cpu *CPU) {
	cpu.D = cpu.H
}

// LD D, L DONE
func op_0x55(cpu *CPU) {
	cpu.D = cpu.L
}

// LD D, (HL) DONE
func op_0x56(cpu *CPU) {
	cpu.D = cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
}

// LD D, A DONE
func op_0x57(cpu *CPU) {
	cpu.D = cpu.A
}

// LD E, B DONE
func op_0x58(cpu *CPU) {
	cpu.E = cpu.B
}

// LD E, C DONE
func op_0x59(cpu *CPU) {
	cpu.E = cpu.C
}

// LD E, D DONE
func op_0x5A(cpu *CPU) {
	cpu.E = cpu.D
}

// LD E, E DONE
func op_0x5B(cpu *CPU) {
	cpu.E = cpu.E
}

// LD E, H DONE
func op_0x5C(cpu *CPU) {
	cpu.E = cpu.H
}

// LD E, L DONE
func op_0x5D(cpu *CPU) {
	cpu.E = cpu.L
}

// LD E, (HL) DONE
func op_0x5E(cpu *CPU) {
	cpu.D = cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
}

// LD E, A DONE
func op_0x5F(cpu *CPU) {
	cpu.E = cpu.A
}

// LD H, B DONE
func op_0x60(cpu *CPU) {
	cpu.H = cpu.B
}
// LD H, C DONE
func op_0x61(cpu *CPU) {
	cpu.H = cpu.C
}
// LD H, D DONE
func op_0x62(cpu *CPU) {
	cpu.H = cpu.D
}
// LD H, E DONE
func op_0x63(cpu *CPU) {
	cpu.H = cpu.E
}
// LD H, H DONE
func op_0x64(cpu *CPU) {
	cpu.H = cpu.H
}
// LD H, L
func op_0x65(cpu *CPU) {
	cpu.H = cpu.L
}
// LD H, (HL) DONE
func op_0x66(cpu *CPU) {
	cpu.H = cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
}

// LD H, A DONE
func op_0x67(cpu *CPU) {
	cpu.H = cpu.A
}
// LD L, B DONE
func op_0x68(cpu *CPU) {
	cpu.L = cpu.B
}
// LD L, C DONE
func op_0x69(cpu *CPU) {
	cpu.L = cpu.C
}	

// LD L, D DONE
func op_0x6A(cpu *CPU) {
	cpu.L = cpu.D
}
// LD L, E DONE
func op_0x6B(cpu *CPU) {
	cpu.L = cpu.E
}
// LD L, H DONE
func op_0x6C(cpu *CPU) {
	cpu.L = cpu.H
}

// LD L, L DONE
func op_0x6D(cpu *CPU) {
	cpu.L = cpu.L
}

// LD L, (HL) DONE
func op_0x6E(cpu *CPU) {
	cpu.L = cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
}

// LD L, A DONE
func op_0x6F(cpu *CPU) {
	cpu.L = cpu.A
}

// LD (HL), B DONE
func op_0x70(cpu *CPU) {
	cpu.bus.write((uint16(cpu.H) << 8 | uint16(cpu.L)), cpu.B)
}

// LD (HL), C DONE
func op_0x71(cpu *CPU) {
	cpu.bus.write((uint16(cpu.H) << 8 | uint16(cpu.L)), cpu.C)
}

// LD (HL), D DONE
func op_0x72(cpu *CPU) {
	cpu.bus.write((uint16(cpu.H) << 8 | uint16(cpu.L)), cpu.D)
}

// LD (HL), A DONE
func op_0x77(cpu *CPU) {
	cpu.bus.write((uint16(cpu.H) << 8 | uint16(cpu.L)), cpu.A)
}

// LD A, B DONE
func op_0x78(cpu *CPU) {
	cpu.A = cpu.B
}

// LD A, C DONE
func op_0x79(cpu *CPU) {
	cpu.A = cpu.C
}

// LD A, D DONE
func op_0x7A(cpu *CPU) {
	cpu.A = cpu.D
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

// ADD A, D
func op_0x82(cpu *CPU) {
	result := uint16(cpu.A) + uint16(cpu.D)
	cpu.A += cpu.D
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	if ((cpu.A & 0xF) + ((cpu.A - cpu.D) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if result > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADD A, (HL)
func op_0x86(cpu *CPU) {
	result := uint16(cpu.A) + uint16(cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L)))
	cpu.A += cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	if ((cpu.A & 0xF) + ((cpu.A - cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L)) & 0xF))) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if result > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// ADC A, D DONE
func op_0x8A(cpu *CPU) {
	bitOverflow := uint16(cpu.A) + uint16(cpu.getCFlag()) + uint16(cpu.D)
	cpu.A += (cpu.getCFlag() + cpu.D)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	if ((cpu.A & 0xF) + ((cpu.A - cpu.getCFlag() - cpu.D) & 0xF)) & 0x10 == 0x10 {
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

// ADC A, L DONE
func op_0x8D(cpu *CPU) {
	bitOverflow := uint16(cpu.A) + uint16(cpu.getCFlag()) + uint16(cpu.L)
	cpu.A += (cpu.getCFlag() + cpu.L)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	if ((cpu.A & 0xF) + ((cpu.A - cpu.getCFlag() - cpu.L) & 0xF)) & 0x10 == 0x10 {
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

// SBC A, D DONE
func op_0x99(cpu *CPU) {
	cpu.A -= (cpu.getCFlag() + cpu.D)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	if ((cpu.A & 0xF) + ((cpu.A + cpu.getCFlag() + cpu.D) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}

	if cpu.D + cpu.getCFlag() > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// AND A, A DONE
func op_0xA7(cpu *CPU) {
	cpu.A &= cpu.A
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.CFlag(0)
}

// XOR A, C DONE
func op_0xA9(cpu *CPU) {
	cpu.A ^= cpu.C
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)

}

// XOR A, (HL) DONE
func op_0xAE(cpu *CPU) {
	cpu.A ^= cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)

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

// OR A, B
func op_0xB0(cpu *CPU) {
	cpu.A |= cpu.B
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// OR A, C
func op_0xB1(cpu *CPU) {
	cpu.A |= cpu.C
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// OR A, D
func op_0xB2(cpu *CPU) {
	cpu.A |= cpu.D
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// OR A, E
func op_0xB3(cpu *CPU) {
	cpu.A |= cpu.E
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// OR A, H
func op_0xB4(cpu *CPU) {
	cpu.A |= cpu.H
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// OR A, L
func op_0xB5(cpu *CPU) {
	cpu.A |= cpu.L
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// OR A, (HL)
func op_0xB6(cpu *CPU) {
	cpu.A |= cpu.bus.read(uint16(cpu.H) << 8 | uint16(cpu.L))
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// OR A, A
func op_0xB7(cpu *CPU) {
	cpu.A |= cpu.A
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
}

// CP A, D DONE
func op_0xBA(cpu *CPU) {
	result := cpu.A - cpu.D
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
	if cpu.D > cpu.A {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
	cpu.PC++
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

// RET NZ DONE
func op_0xC0(cpu *CPU) {
	// if Z Flag is reset condition
	if cpu.getZFlag() == 0 {
		cpu.PC = cpu.bus.read16(cpu.SP)
		cpu.SP += 2
		cpu.Cycles += 12
	} else {
		cpu.PC++
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

// CALL NZ, u16 
func op_0xC4(cpu *CPU) {
	// if z flag is not set
	if cpu.getZFlag() == 0 {
		hi := byte((cpu.PC + 2) >> 8)
		lo := byte(((cpu.PC + 2) & 0xFF))
		cpu.SP--
		cpu.bus.write(cpu.SP, hi)
		cpu.SP--
		cpu.bus.write(cpu.SP, lo)
		cpu.PC = cpu.bus.read16(cpu.PC)
		cpu.Cycles += 12
	}
}

// PUSH BC DONE
func op_0xC5(cpu *CPU) {
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.B)
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.C)
}

// ADD A, u8
func op_0xC6(cpu *CPU) {
	result := uint16(cpu.A + cpu.bus.read(cpu.PC))
	cpu.A += cpu.bus.read(cpu.PC)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	if ((cpu.A & 0xF) + ((cpu.A - cpu.bus.read(cpu.PC)) & 0xF)) & 0x10 == 0x10 {
		cpu.HFlag(1)
	} else {
		cpu.HFlag(0)
	}
	if result > 0xFF {
		cpu.CFlag(1)
	} else {
		cpu.CFlag(0)
	}
}

// RST 0x00 NOT DONE
func op_0xC7(cpu *CPU) {
	hi := byte((cpu.PC + 2) >> 8)
	lo := byte(((cpu.PC + 2) & 0xFF))
	cpu.SP--
	cpu.bus.write(cpu.SP, hi)
	cpu.SP--
	cpu.bus.write(cpu.SP, lo)
	cpu.PC = 0x0000
}

// RET Z DONE
func op_0xC8(cpu *CPU) {
	if cpu.getZFlag() == 1 {
		cpu.PC = cpu.bus.read16(cpu.SP)
		cpu.SP += 2
		cpu.Cycles += 12
	} else {
		cpu.PC++
	}
}

// RET DONE
func op_0xC9(cpu *CPU) {
	cpu.PC = cpu.bus.read16(cpu.SP)
	cpu.SP += 2
}

// JP Z, u16
func op_0xCA(cpu *CPU) {
	if cpu.getZFlag() == 1 {
		cpu.PC = cpu.bus.read16(cpu.PC)
		cpu.Cycles += 4
	} else {
		cpu.PC += 2
	}
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
	cpu.PC++



}

// POP DE DONE
func op_0xD1(cpu *CPU) {
	cpu.E = cpu.bus.read(cpu.SP)
	cpu.SP++
	cpu.D = cpu.bus.read(cpu.SP)
	cpu.SP++
}


// PUSH DE DONE
func op_0xD5(cpu *CPU) {
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.D)
	cpu.SP--
	cpu.bus.write(cpu.SP, cpu.E)
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
	cpu.PC++
}

// RET C DONE
func op_0xD8(cpu *CPU) {
	// if C Flag is set condition
	if cpu.getCFlag() == 1 {
		cpu.PC = cpu.bus.read16(cpu.SP)
		cpu.SP += 2
		cpu.Cycles += 12
	} else {
		cpu.PC++
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

// AND A, u8 DONE
func op_0xE6(cpu *CPU) {
	cpu.A &= cpu.bus.read(cpu.PC)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(1)
	cpu.CFlag(0)
	cpu.PC++
	// make sure to check other u8 instructions later
}

// LD (u16), A DONE
func op_0xEA(cpu *CPU) {
	cpu.bus.write((uint16(cpu.bus.read(cpu.PC + 1))) << 8 | uint16((cpu.bus.read(cpu.PC))), cpu.A)
	cpu.PC += 2
}

// XOR A, u8
func op_0xEE(cpu *CPU) {
	cpu.A ^= cpu.bus.read(cpu.PC)
	if cpu.A == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
	cpu.CFlag(0)
	cpu.PC++
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

// LD A, (u16) DONE
func op_0xFA(cpu *CPU) {
	addr := cpu.bus.read16(cpu.PC)
	cpu.A = cpu.bus.read(addr)
	cpu.PC += 2
}

// EI DONE
func op_0xFB(cpu *CPU) {
	cpu.bus.IME = 1
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

// RL B DONE
func cbop_0x10(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	newcarry := cpu.B >> 7
	cpu.B = (cpu.B << 1 | tempcarry)
	cpu.CFlag(newcarry)
	if cpu.B == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)
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

// RR C DONE
func cbop_0x19(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	cpu.CFlag(cpu.C & 0x1)
	cpu.C = (tempcarry << 7 | cpu.C >> 1)
	if cpu.C == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)


}

// RR D DONE
func cbop_0x1A(cpu *CPU) {
	tempcarry := cpu.getCFlag()
	cpu.CFlag(cpu.D & 0x1)
	cpu.D = (tempcarry << 7 | cpu.D >> 1)
	if cpu.D == 0 {
		cpu.ZFlag(1)
	} else {
		cpu.ZFlag(0)
	}
	cpu.NFlag(0)
	cpu.HFlag(0)


}

// SRL B 
func cbop_0x38(cpu *CPU) {
	cpu.CFlag(cpu.B & 0x1)
	cpu.B >>= 1
	if cpu.B == 0 {
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
	Opcode{4, op_0x00}, Opcode{12, op_0x01}, Opcode{8, op_null}, Opcode{8, op_0x03}, Opcode{4, op_0x04}, Opcode{4, op_0x05}, Opcode{8, op_0x06}, Opcode{4, op_0x07}, Opcode{20, op_0x08}, Opcode{8, op_null}, Opcode{8, op_null}, Opcode{8, op_null}, Opcode{4, op_0x0C}, Opcode{4, op_0x0D}, Opcode{8, op_0x0E}, Opcode{4, op_null},
	Opcode{4, op_null}, Opcode{12, op_0x11}, Opcode{8, op_0x12}, Opcode{8, op_0x13}, Opcode{4, op_0x14}, Opcode{4, op_0x15}, Opcode{8, op_0x16}, Opcode{4, op_0x17}, Opcode{12, op_0x18}, Opcode{8, op_null}, Opcode{8, op_0x1A}, Opcode{8, op_null}, Opcode{4, op_0x1C}, Opcode{4, op_0x1D}, Opcode{8, op_0x1E}, Opcode{4, op_0x1F},
	Opcode{8, op_0x20}, Opcode{12, op_0x21}, Opcode{8, op_0x22}, Opcode{8, op_0x23}, Opcode{4, op_0x24}, Opcode{4, op_0x25}, Opcode{8, op_0x26}, Opcode{4, op_null}, Opcode{8, op_0x28}, Opcode{8, op_0x29}, Opcode{8, op_0x2A}, Opcode{8, op_null}, Opcode{4, op_0x2C}, Opcode{4, op_0x2D}, Opcode{8, op_0x2E}, Opcode{4, op_null},
	Opcode{8, op_0x30}, Opcode{12, op_0x31}, Opcode{8, op_0x32}, Opcode{8, op_null}, Opcode{12, op_null}, Opcode{12, op_0x35}, Opcode{12, op_null}, Opcode{4, op_null}, Opcode{8, op_0x38}, Opcode{8, op_null}, Opcode{8, op_null}, Opcode{8, op_null}, Opcode{4, op_0x3C}, Opcode{4, op_0x3D}, Opcode{8, op_0x3E}, Opcode{4, op_null},
	Opcode{4, op_0x40}, Opcode{4, op_0x41}, Opcode{4, op_0x42}, Opcode{4, op_0x43}, Opcode{4, op_0x44}, Opcode{4, op_0x45}, Opcode{8, op_0x46}, Opcode{4, op_0x47}, Opcode{4, op_0x48}, Opcode{4, op_0x49}, Opcode{4, op_0x4A}, Opcode{4, op_0x4B}, Opcode{4, op_0x4C}, Opcode{4, op_0x4D}, Opcode{8, op_0x4E}, Opcode{4, op_0x4F},
	Opcode{4, op_0x50}, Opcode{4, op_0x51}, Opcode{4, op_0x52}, Opcode{4, op_0x53}, Opcode{4, op_0x54}, Opcode{4, op_0x55}, Opcode{8, op_0x56}, Opcode{4, op_0x57}, Opcode{4, op_0x58}, Opcode{4, op_0x59}, Opcode{4, op_0x5A}, Opcode{4, op_0x5B}, Opcode{4, op_0x5C}, Opcode{4, op_0x5D}, Opcode{8, op_0x5E}, Opcode{4, op_0x5F},
	Opcode{4, op_0x60}, Opcode{4, op_0x61}, Opcode{4, op_0x62}, Opcode{4, op_0x63}, Opcode{4, op_0x64}, Opcode{4, op_0x65}, Opcode{8, op_0x66}, Opcode{4, op_0x67}, Opcode{4, op_0x68}, Opcode{4, op_0x69}, Opcode{4, op_0x6A}, Opcode{4, op_0x6B}, Opcode{4, op_0x6C}, Opcode{4, op_0x6D}, Opcode{8, op_0x6E}, Opcode{4, op_0x6F},
	Opcode{8, op_0x70}, Opcode{8, op_0x71}, Opcode{8, op_0x72}, Opcode{8, op_null}, Opcode{8, op_null}, Opcode{8, op_null}, Opcode{4, op_null}, Opcode{8, op_0x77}, Opcode{4, op_0x78}, Opcode{4, op_0x79}, Opcode{4, op_0x7A}, Opcode{4, op_0x7B}, Opcode{4, op_0x7C}, Opcode{4, op_0x7D}, Opcode{8, op_null}, Opcode{4, op_null},
	Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_0x82}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_0x86}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_0x8A}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_0x8D}, Opcode{8, op_null}, Opcode{4, op_null},
	Opcode{4, op_0x90}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_0x99}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{4, op_null},
	Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{4, op_0xA7}, Opcode{4, op_null}, Opcode{4, op_0xA9}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_0xAE}, Opcode{4, op_0xAF},
	Opcode{4, op_0xB0}, Opcode{4, op_0xB1}, Opcode{4, op_0xB2}, Opcode{4, op_0xB3}, Opcode{4, op_0xB4}, Opcode{4, op_0xB5}, Opcode{8, op_0xB6}, Opcode{4, op_0xB7}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_0xBA}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_0xBE}, Opcode{4, op_null},
	Opcode{8, op_0xC0}, Opcode{12, op_0xC1}, Opcode{12, op_null}, Opcode{16, op_0xC3}, Opcode{12, op_0xC4}, Opcode{16, op_0xC5}, Opcode{8, op_0xC6}, Opcode{16, op_0xC7}, Opcode{8, op_0xC8}, Opcode{16, op_0xC9}, Opcode{12, op_0xCA}, Opcode{4, op_0xCB}, Opcode{12, op_null}, Opcode{24, op_0xCD}, Opcode{8, op_0xCE}, Opcode{16, op_null},
	Opcode{8, op_null}, Opcode{12, op_0xD1}, Opcode{12, op_null}, Opcode{8, op_null}, Opcode{12, op_null}, Opcode{16, op_0xD5}, Opcode{8, op_0xD6}, Opcode{16, op_null}, Opcode{8, op_0xD8}, Opcode{16, op_null}, Opcode{12, op_null}, Opcode{8, op_null}, Opcode{12, op_null}, Opcode{4, op_null}, Opcode{8, op_null}, Opcode{16, op_null},
	Opcode{12, op_0xE0}, Opcode{12, op_0xE1}, Opcode{8, op_0xE2}, Opcode{8, op_null}, Opcode{4, op_null}, Opcode{16, op_0xE5}, Opcode{8, op_0xE6}, Opcode{16, op_null}, Opcode{16, op_null}, Opcode{4, op_null}, Opcode{16, op_0xEA}, Opcode{8, op_null}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_0xEE}, Opcode{16, op_null},
	Opcode{12, op_0xF0}, Opcode{12, op_0xF1}, Opcode{8, op_null}, Opcode{4, op_0xF3}, Opcode{4, op_null}, Opcode{16, op_0xF5}, Opcode{8, op_null}, Opcode{16, op_null}, Opcode{12, op_null}, Opcode{8, op_null}, Opcode{16, op_0xFA}, Opcode{4, op_0xFB}, Opcode{4, op_null}, Opcode{4, op_null}, Opcode{8, op_0xFE}, Opcode{16, op_null},

}

var cbopcodes [256]Opcode = [256]Opcode {
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_0x10}, Opcode{8, cbop_0x11}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_0x19}, Opcode{8, cbop_0x1A}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null},
	Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_0x38}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{8, cbop_null}, Opcode{16, cbop_null}, Opcode{8, cbop_null},
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