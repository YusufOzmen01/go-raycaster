package main

import "raycaster/internal"

func main() {
	game := internal.NewGame(1000, 480, "Raycaster")
	game.Start()
}
