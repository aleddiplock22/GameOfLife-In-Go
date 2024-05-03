package gameoflife

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const WIDTH = 410
const HEIGHT = 410

type Game struct {
	grid *Grid
}

func (g *Game) Update() error {
	err := g.grid.Update()
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.grid.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func NewGame() *Game {
	return &Game{NewGrid()}
}

func (g *Game) DrawRect() error {
	return nil
}
