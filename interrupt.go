package main

type Interrupt struct {
	IME byte
	IF byte
	IE byte
}

func (interrupt *Interrupt) requestVBlank() {
	interrupt.IF |= 1
}

func (interrupt *Interrupt) requestLCDCSTAT() {
	interrupt.IF |= (1 << 1)
}

func (interrupt *Interrupt) requestTimer() {
	interrupt.IF |= (1 << 2)
}

func (interrupt *Interrupt) requestSerial() {
	interrupt.IF |= (1 << 3)
}

func (interrupt *Interrupt) requestJoypad() {
	interrupt.IF |= (1 << 4)
}



// EI:
//   # Enables IME, however there is a delay
//   ime_delay = true

// DI:
//   # DI immediately disables IME
//   ime = false
//   ime_delay = false

// your tick function:
//   Handle interrupts
  
//   if ime_delay:
//     ime_delay = false
//     ime = true

//   Get opcode
//   Execute opcode

// Interrupt handler:
//   IF = memory[0xFF0F]
//   IE = memory[0xFFFF]
//   potential interrupts = IF & IE & 0x1F
  
//   # None of the last five bits in IF or IE match with the other  
//   if !potential interrupts:
//     return
  
//   # IME is disabled
//   if !ime:
//     return

//   # Five different kinds of interrupts
//   for bit in range(5):
//     if !(potential interrupts & (1 << bit)):
//       continue; # Check the next bit

//     # At this point, we now have a requested interrupt that is enabled, so we should run it
    
//     # Turn off the requested interrupt bit
//     memory[0xFF0F] &= ~(1 << bit)

//     # IME is disabled after every serviced interrupt
//     ime = false

//     # Here you should determine which address to jump to based on which interrupt you are servicing.
//     # For example, if IF and IE both had bit 0 set, we would be doing a VBlank interrupt, so we would jump to $0040.

//     Push the current PC to the stack
//     Set PC to the new interrupt address