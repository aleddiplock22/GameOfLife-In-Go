package gameoflife

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const WIDTH = 400
const HEIGHT = 400

type Game struct {
	grid *Grid
}

func (g *Game) Update() error {
	err := g.grid.Update()
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "GAME OF LIFE!")
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
