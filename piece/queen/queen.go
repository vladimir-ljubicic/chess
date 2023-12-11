package queen

import (
	"github.com/chess/grid"
	"github.com/chess/piece"
	"github.com/samber/lo"
)

type Queen struct {
	Color      piece.Color
	Position   grid.Cell
	LegalMoves []grid.Cell
}

func New(startPosition grid.Coordinates, color piece.Color) Queen {
	return Queen{
		Color:      color,
		Position:   grid.Cell{Coordinates: startPosition},
		LegalMoves: nil,
	}
}

func (q Queen) GetLegalMoves(g grid.Grid) (moves []grid.Cell) {
	row := g.GetRow(q.Position.X)
	moves = append(moves, row...)

	column := g.GetColumn(q.Position.Y)
	moves = append(moves, column...)

	diagonals := g.GetDiagonals(q.Position)
	moves = append(moves, diagonals...)

	//	Filter out current grid position
	moves = lo.Filter(moves, func(c grid.Cell, _ int) bool {
		return c.Coordinates != q.Position.Coordinates
	})

	return moves
}
