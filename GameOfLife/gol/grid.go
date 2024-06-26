package gameoflife

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const LINE_WIDTH = ADJ
const WIDTH_adj = WIDTH - LINE_WIDTH
const HEIGHT_adj = HEIGHT - LINE_WIDTH
const NUM_LINES = 50 // use 50 for Gosper Glider Gun
const BOX_SIDE_LEN = WIDTH_adj/NUM_LINES - LINE_WIDTH

var COLOR_RED = color.RGBA{255, 0, 0, 1}

type Tile struct {
	x int
	y int
}

type Grid struct {
	tiles []*Tile
}

func NewGrid(starting_positions [][]int) *Grid {
	var tiles []*Tile
	centering_adjustment := (NUM_LINES / 2) - int(NUM_LINES/10)
	for _, position := range starting_positions {
		x := position[0] + centering_adjustment
		y := position[1] + centering_adjustment
		tiles = append(tiles, &Tile{x, y})
	}
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
	vector.DrawFilledRect(gridImage, x_coord, y_coord, BOX_SIDE_LEN, BOX_SIDE_LEN, COLOR_RED, false)
}

func DrawGridLines(width int, height int, gridImage *ebiten.Image) {
	var N float32 = float32(NUM_LINES)
	// vertical
	var i float32 = 0
	for i <= N {
		vector.DrawFilledRect(
			gridImage,
			i*(float32(width)/N),
			0,
			LINE_WIDTH,
			float32(height)+LINE_WIDTH,
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
			float32(width)+LINE_WIDTH,
			LINE_WIDTH,
			color.White,
			false)
		i++
	}
}

func getXYCoordFromGridRef(x_coord int, y_coord int) (float32, float32) {
	// where we have a grid (x_coord=0, y_coord=0) -> (x_coord=N-1, y_coord=N-1)
	x := LINE_WIDTH + x_coord*(BOX_SIDE_LEN+LINE_WIDTH)
	y := LINE_WIDTH + y_coord*(BOX_SIDE_LEN+LINE_WIDTH)
	return float32(x), float32(y)
}
