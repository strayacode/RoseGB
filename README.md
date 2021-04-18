# RoseGB
reformed gameboy emulator in golang with revised structure

**Screenshots:**

![](https://github.com/strayacode/RoseGB/blob/master/screenshots/pic1.png)

**Test Pass Checklist:**

**Blargg:**
- [x] 01-special.gb
- [x] 02-interrupts.gb
- [x] 03-op sp, hl.gb
- [x] 04-op r, imm.gb
- [x] 05-op rp.gb
- [x] 06-ld r, r.gb
- [x] 07-jr, jp, call, ret, rst.gb
- [x] 08-misc instrs.gb
- [x] 09-op r,r .gb
- [x] 10-bit ops.gb
- [x] 11-op a, (hl).gb
- [x] cpu_instrs.gb
- [ ] instr_timing.gb

**tests:**
- [x] boot_regs-dmgABC.gb

**instr:**
- [x] daa.gb

**MBC1:**
- [x] bits_bank1.gb
- [x] bits_bank2.gb
- [ ] bits_mode.gb
- [x] bits-ramg.gb
- [ ] multicart_rom_8Mb.gb
- [ ] ram_64kb.gb
- [ ] ram_256kb.gb
- [ ] rom_1Mb.gb
- [ ] rom_2Mb.gb
- [ ] rom_4Mb.gb
- [ ] rom_8Mb.gb
- [ ] rom_16Mb.gb
- [ ] ram_512kb.gb

**TODO:**
- [x] Interrupts
- [x] Timers
- [x] Proper joypad functionality
- [x] OAM + DMA
- [ ] finish off MBC1
- [ ] More interactive VRAM viewer
- [ ] play/pause/step functionality
- [ ] add emulator window that can be "popped out" of the main window to play the emulator on any screen
- [ ] be able to easily change the size of the game window
- [ ] pixel fifo
- [ ] full restructure of emulator for optimization
- [ ] audio
- [ ] add more debugging features, live preview of background/window with frame, sprite properties etc. 

**Usage:**

To compile RoseGB in the terminal, type `go build` and use `./rosegb path-to-file` to run

use the flag `-skip-bootrom` if you don't own a copy of the dmg bootrom
