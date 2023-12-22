package grid

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
			Coordinates: Coordinates{
				X: x,
				Y: y,
			},
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
