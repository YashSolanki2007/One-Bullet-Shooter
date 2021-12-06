package main 

import ("github.com/gen2brain/raylib-go/raylib")

type Bullet struct {
	pos_x  int32 
	pos_y  int32
	y_veolcity int32
	radius float32
	Draw bool 
	Color rl.Color
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(600)
	var shouldShoot bool = false
	var hasShotten bool = false

	// Creating the main window 
	rl.InitWindow(screenWidth, screenHeight, "One Bullet Shooter")

	// Creating the player
	player := rl.LoadImage("player.png")
	var player_x int32 = screenWidth / 2
	var player_y int32 = screenHeight - 100
	var player_x_change int32 = 0
	player_texture := rl.LoadTextureFromImage(player)
	Bullets := Bullet{}

	// Create the main game loop 
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawTexture(player_texture, player_x, player_y, rl.White)
		
		if (shouldShoot) {
			rl.DrawCircle(Bullets.pos_x + 32, Bullets.pos_y, Bullets.radius, Bullets.Color)
			Bullets.pos_y += Bullets.y_veolcity
		}

		// KeyPresses
		if rl.IsKeyDown(rl.KeyRight) {
			player_x_change = 1
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			player_x_change = -1
		}

		if (rl.IsKeyDown(rl.KeySpace) && !hasShotten) {
			current_bullet := Bullet{int32(player_x), int32(player_y), int32(-2), float32(10), true, rl.Black}
			shouldShoot = true
			Bullets = current_bullet
			hasShotten = true
		}

		if rl.IsKeyUp(rl.KeyRight) && rl.IsKeyUp(rl.KeyLeft) {
			player_x_change = 0
		}

		// Border checking 
		if (player_x < 0) {
			player_x = 0
		}
		if (player_x >= screenWidth - 70) {
			player_x = screenWidth - 70
		}

		player_x += player_x_change
		rl.EndDrawing()
	}
	rl.CloseWindow()
}