package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {

	var playerX float32 = 25
	var playerY float32 = 100
	var playerSpeed float32 = 100
	var playerSize float32 = 50

	rl.InitWindow(800, 450, "Game!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(120)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		//draw player rectangle
		rl.DrawRectangle(int32(playerX), int32(playerY), int32(playerSize), int32(playerSize), rl.White)

		//keyboard input
		//top of screen
		if rl.IsKeyDown(rl.KeyW) && playerY > 0 {
			playerY -= playerSpeed * rl.GetFrameTime()
		}
		//bottom of screen
		if rl.IsKeyDown(rl.KeyS) && playerY < 450-playerSize {
			playerY += playerSpeed * rl.GetFrameTime()
		}
		//left screen block
		if rl.IsKeyDown(rl.KeyA) && playerX > 0 {
			playerX -= playerSpeed * rl.GetFrameTime()
		}
		//right screen block
		if rl.IsKeyDown(rl.KeyD) && playerX < 800-playerSize {
			playerX += playerSpeed * rl.GetFrameTime()
		}

		rl.EndDrawing()
	}
}
