package gameoflife

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const LINE_WIDTH = 10
const WIDTH_adj = WIDTH - LINE_WIDTH
const HEIGHT_adj = HEIGHT - LINE_WIDTH
const NUM_LINES = 10
const BOX_SIDE_LEN = WIDTH_adj/NUM_LINES - LINE_WIDTH

var COLOR_RED = color.RGBA{255, 0, 0, 1}

type Tile struct {
	x int
	y int
}

type Grid struct {
	tiles []*Tile
}

func NewGrid() *Grid {
	var tiles []*Tile
	tiles = append(tiles, &Tile{5, 5})
	grid := &Grid{
		tiles: tiles,
	}
	return grid
}

func (g *Grid) Draw(gridImage *ebiten.Image) {
	DrawGridLines(WIDTH_adj, HEIGHT_adj, gridImage)
	for _, tile := range g.tiles {
		DrawTile(tile.x, tile.y, gridImage)
	}
}

func DrawTile(x, y int, gridImage *ebiten.Image) {
	x_coord, y_coord := getXYCoordFromGridRef(x, y)
	fmt.Println(x_coord, x, "|", y_coord, y)
	vector.DrawFilledRect(gridImage, x_coord, y_coord, BOX_SIDE_LEN, BOX_SIDE_LEN, COLOR_RED, false)
}

func DrawGridLines(width int, height int, gridImage *ebiten.Image) {
	var N float32 = float32(NUM_LINES)
	var line_width float32 = 10.0
	// vertical
	var i float32 = 0
	for i <= N {
		vector.DrawFilledRect(
			gridImage,
			i*(float32(width)/N),
			0,
			line_width,
			float32(height)+line_width,
			color.White,
			false)
		i++
	}
	// horizontal
	i = 0
	for i <= N {
		vector.DrawFilledRect(
			gridImage,
			0,
			i*(float32(height)/N),
			float32(width)+line_width,
			line_width,
			color.White,
			false)
		i++
	}
}

func (g *Grid) Update() error {
	g.tiles = append(g.tiles, &Tile{
		rand.Intn(NUM_LINES + 1), rand.Intn(NUM_LINES + 1),
	})
	return nil
}

func getXYCoordFromGridRef(x_coord int, y_coord int) (float32, float32) {
	// where we have a grid (x_coord=0, y_coord=0) -> (x_coord=N, y_coord=N)
	x := LINE_WIDTH + x_coord*(BOX_SIDE_LEN+LINE_WIDTH)
	y := LINE_WIDTH + y_coord*(BOX_SIDE_LEN+LINE_WIDTH)
	return float32(x), float32(y)
}
