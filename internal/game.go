package internal

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game interface {
	Start()
}

type game struct {
	Map      WorldMap
	Player   Player
	Renderer Renderer
}

func NewGame(width, height int32, title string) Game {
	fmt.Println("Initializing window")

	rl.SetTraceLog(rl.LogNone)
	rl.InitWindow(width, height, title)
	rl.SetTargetFPS(60)

	fmt.Println("Initializing audio")
	rl.InitAudioDevice()

	return &game{
		Map:      NewWorldMap(60),
		Player:   NewPlayer(200, 200, 10, 50, 100, 100),
		Renderer: NewRenderer(480, 0, 520, 500),
	}
}

func (g *game) Start() {
	g.start()
}

func (g *game) start() {
	for !rl.WindowShouldClose() {
		g.Player.UpdatePlayer(g.Map)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawRectangle(480, 0, 520, 240, rl.Blue)
		rl.DrawRectangle(480, 241, 520, 240, rl.Green)
		rl.DrawRectangle(480, 0, 10, 480, rl.Black)

		g.Map.RenderMap()
		g.Player.RenderPlayer(g.Map)

		g.Renderer.Render(g.Map, g.Player)
		rl.EndDrawing()
	}
}
