package grid

import (
	"github.com/samber/lo"
)

type Grid struct {
	cells      []Cell
	Dimensions int
}

func NewGrid(dimensions int) Grid {
	grid := Grid{
		cells:      make([]Cell, dimensions*dimensions),
		Dimensions: dimensions,
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

		if (i+1)%g.Dimensions == 0 {
			x++
			y = 0
		}
	}
}

func (g Grid) GetRow(index int) []Cell {
	return lo.Filter(g.cells, func(cell Cell, _ int) bool {
		return cell.Coordinates.X == index
	})
}

func (g Grid) GetColumn(index int) []Cell {
	return lo.Filter(g.cells, func(cell Cell, _ int) bool {
		return cell.Coordinates.Y == index
	})
}

func (g Grid) GetDiagonals(c Cell) []Cell {
	var diagonalCells []Cell
	nextCell := c
	for {
		nextCell = nextCell.GoUp().GoRight()

		if g.IsValidCell(nextCell) {
			diagonalCells = append(diagonalCells, nextCell)
		} else {
			break
		}
	}

	nextCell = c
	for {
		nextCell = nextCell.GoDown().GoLeft()

		if g.IsValidCell(nextCell) {
			diagonalCells = append(diagonalCells, nextCell)
		} else {
			break
		}
	}

	nextCell = c
	for {
		nextCell = nextCell.GoUp().GoLeft()

		if g.IsValidCell(nextCell) {
			diagonalCells = append(diagonalCells, nextCell)
		} else {
			break
		}
	}

	nextCell = c
	for {
		nextCell = nextCell.GoDown().GoRight()

		if g.IsValidCell(nextCell) {
			diagonalCells = append(diagonalCells, nextCell)
		} else {
			break
		}
	}

	return diagonalCells
}

func (g Grid) IsValidCell(c Cell) bool {
	return c.X >= 0 && c.Y >= 0 && c.X < g.Dimensions && c.Y < g.Dimensions
}
