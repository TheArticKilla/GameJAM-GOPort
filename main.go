package main

import (
	"log"

	"github.com/TheArticKilla/GameJAM-GOPort/Game"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	game, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Game name")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
