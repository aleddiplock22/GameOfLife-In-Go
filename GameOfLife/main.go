package main

import (
	gameoflife "gameoflife/gol"

	"github.com/hajimehoshi/ebiten/v2"
)

/*
Cellular Automata Simulator

Defined by starting position and rule_set
Currently only have rule_set of Conway's Game of Life
Starting positions defined in ./gol/starting_positions.go

Grid size and Number of cells can be set in ./gol/game.go and ./gol/grid.go respectively
Adjust framerate in main func below [SetTPS]

NOTE: ADJUST gameoflife.NUM_LINES = 100 in grid.go for GOSPER_GLIDER_GUN
*/

func main() {
	game := gameoflife.NewGame(gameoflife.GLIDER, gameoflife.ConwaysGameOfLife)
	ebiten.SetWindowSize(gameoflife.WIDTH, gameoflife.HEIGHT)
	ebiten.SetWindowTitle("Game Of Life Simulator")
	ebiten.SetTPS(10) // num ticks per second
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
