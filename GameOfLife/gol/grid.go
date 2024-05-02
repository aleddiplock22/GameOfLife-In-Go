package gameoflife

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Tile struct {
	x int
	y int
}

type Grid struct {
	tiles []*Tile
}

func NewGrid() *Grid {
	var tiles []*Tile
	tiles = append(tiles, &Tile{WIDTH / 2, HEIGHT / 2})
	grid := &Grid{
		tiles: tiles,
	}
	return grid
}

func (g *Grid) Draw(gridImage *ebiten.Image) {
	DrawGridLines(WIDTH, HEIGHT, gridImage)
	for _, tile := range g.tiles {
		vector.DrawFilledRect(gridImage, float32(tile.x), float32(tile.y), 10, 10, color.RGBA{255, 0, 0, 1}, false)
	}
}

func DrawGridLines(width int, height int, gridImage *ebiten.Image) {
	var N float32 = 10
	var line_width float32 = 10.0
	// vertical
	var i float32 = 0
	for i < N {
		vector.DrawFilledRect(
			gridImage,
			i*(float32(height)/N),
			0,
			line_width,
			float32(height),
			color.White,
			false)
		i++
	}
	// horizontal
	i = 0
	for i < N {
		vector.DrawFilledRect(
			gridImage,
			0,
			i*(float32(height)/N),
			float32(width),
			line_width,
			color.White,
			false)
		i++
	}
}

func (g *Grid) Update() error {
	for _, tile := range g.tiles {
		tile.x++
		tile.y++
		fmt.Println(tile)
	}
	return nil
}
