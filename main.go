package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"fmt"
)

func main() {

	rl.InitWindow(160, 144, "RoseGB")

	rl.SetTargetFPS(60)

	

	
	// mainloop which is executed 60 times per second
	for !rl.WindowShouldClose() {
		
		// rl.BeginDrawing()
		// rl.ClearBackground(rl.RayWhite)
		// rl.DrawRectangle(1, 1, 1, 1, rl.Red)
		

		// rl.EndDrawing()
	}

	rl.CloseWindow()
}