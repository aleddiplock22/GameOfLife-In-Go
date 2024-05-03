package gameoflife

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

const ADJ = 2
const WIDTH = 500 + ADJ
const HEIGHT = 500 + ADJ

type Game struct {
	grid  *Grid
	state *State
	rules RuleSet
}

func (g *Game) Update() error {
	g.state.PerformGeneration(g.rules)
	var new_tiles []*Tile
	for x, col := range *g.state {
		for y, alive := range col {
			if alive {
				fmt.Println("FOUND ALIVE: ", x, y)
				new_tiles = append(new_tiles, &Tile{x, y})
			}
		}
	}

	g.grid.tiles = new_tiles
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.grid.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func NewGame(starting_positions [][]int, rule RuleSet) *Game {
	new_grid := NewGrid(starting_positions)
	return &Game{new_grid, DefineState(new_grid.tiles), rule}
}

func (g *Game) DrawRect() error {
	return nil
}
