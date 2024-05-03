package main

import (
	gameoflife "gameoflife/gol"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := gameoflife.NewGame(gameoflife.BLINKER, gameoflife.ConwaysGameOfLife)
	ebiten.SetWindowSize(gameoflife.WIDTH, gameoflife.HEIGHT)
	ebiten.SetWindowTitle("Game Of Life Simulator")
	ebiten.SetTPS(5) // num ticks per second
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
