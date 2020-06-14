package main

import (
	"fmt"
	"strconv"
)

func (cpu *CPU) debug() {
	fmt.Println("A: ", strconv.FormatUint(uint64(cpu.A), 16), "F ", strconv.FormatUint(uint64(cpu.F), 16))
	fmt.Println("B: ", strconv.FormatUint(uint64(cpu.B), 16), "C: ", strconv.FormatUint(uint64(cpu.C), 16))
	fmt.Println("D: ", strconv.FormatUint(uint64(cpu.D), 16), "E: ", strconv.FormatUint(uint64(cpu.E), 16))
	fmt.Println("H: ", strconv.FormatUint(uint64(cpu.H), 16), "L: ", strconv.FormatUint(uint64(cpu.L), 16))
	fmt.Println("PC: ", strconv.FormatUint(uint64(cpu.PC), 16), "SP: ", strconv.FormatUint(uint64(cpu.SP), 16))
}