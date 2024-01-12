package grid

import "github.com/samber/lo"

type Grid struct {
	cells      []Cell
	dimensions int
}

func NewGrid(dimensions int) Grid {
	grid := Grid{
		cells:      make([]Cell, dimensions*dimensions),
		dimensions: dimensions,
	}

	grid.initCoordinates()

	return grid
}

func (g Grid) initCoordinates() {
	x, y := 0, 0

	for i, _ := range g.cells {
		g.cells[i] = Cell{
			X: x,
			Y: y,
		}
		y++

		if (i+1)%g.dimensions == 0 {
			x++
			y = 0
		}
	}
}

func (g Grid) GetDimensions() int {
	return g.dimensions
}

func (g Grid) IsValidCell(c Cell) bool {
	return c.X >= 0 && c.Y >= 0 && c.X < g.dimensions && c.Y < g.dimensions
}

func GetMedialCells(c1 Cell, c2 Cell) (medialCells []Cell) {
	if c1 == c2 {
		return
	}

	minByX := lo.MinBy([]Cell{c1, c2}, func(a Cell, b Cell) bool {
		return a.X < b.X
	})
	maxByX := lo.MaxBy([]Cell{c1, c2}, func(a Cell, b Cell) bool {
		return a.X > b.X
	})
	minByY := lo.MinBy([]Cell{c1, c2}, func(a Cell, b Cell) bool { return a.Y < b.Y })
	maxByY := lo.MaxBy([]Cell{c1, c2}, func(a Cell, b Cell) bool { return a.Y > b.Y })

	//	medial squares on the same column
	if c1.X == c2.X {
		c := minByY
		for c.Y < maxByY.Y-1 {
			c = c.Up()
			medialCells = append(medialCells, c)
		}
	}
	//	medial squares on the same row
	if c1.Y == c2.Y {
		c := minByX
		for c.X < maxByX.X-1 {
			c = c.Right()
			medialCells = append(medialCells, c)
		}
	}
	//	medial squares on a descending diagonal
	if c1.X+c1.Y == c2.X+c2.Y {
		c := minByX
		for c.X < maxByX.X-1 {
			c = c.Right().Down()
			medialCells = append(medialCells, c)
		}

	}

	//	medial squares on an ascending diagonal
	if c1.X-c2.X == c1.Y-c2.Y {
		c := minByX
		for c.X < maxByX.X-1 {
			c = c.Right().Up()
			medialCells = append(medialCells, c)
		}
	}

	return
}
