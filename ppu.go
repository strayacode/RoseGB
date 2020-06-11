type PPU struct {
	frameBuffer [144][160]byte
	VRAM [0x2000]byte
}