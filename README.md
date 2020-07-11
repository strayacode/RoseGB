# RoseGB
reformed gameboy emulator in golang with revised structure

**Test Pass Checklist:**
- [x] 01-special.gb
- [ ] 02-interrupts.gb
- [x] 03-op sp, hl.gb
- [x] 04-op r, imm.gb
- [x] 05-op rp.gb
- [x] 06-ld r, r.gb
- [x] 07-jr, jp, call, ret, rst.gb
- [x] 08-misc instrs.gb
- [x] 09-op r,r .gb
- [x] 10-bit ops.gb
- [x] 11-op a, (hl).gb

**TODO:**
- [x] Interrupts
- [ ] Timers
- [ ] More interactive VRAM viewer
- [ ] play/pause/step functionality

**Usage:**

to compile RoseGB in the terminal type ```go build rosegb.go apu.go bus.go cartridge.go cpu.go debug.go interrupt.go opcodes.go ppu.go timer.go keypad.go``` and use ```./rosegb path-to-file``` to run
