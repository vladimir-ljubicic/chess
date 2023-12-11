package rook

import (
	"github.com/chess/grid"
	"github.com/chess/piece"
	"github.com/samber/lo"
)

type Rook struct {
	Color      piece.Color
	Position   grid.Cell
	LegalMoves []grid.Cell
}

func New(startPosition grid.Coordinates, color piece.Color) Rook {
	return Rook{
		Color: color,
		Position: grid.Cell{
			Coordinates: startPosition,
		},
		LegalMoves: nil,
	}
}

func (r Rook) GetLegalMoves(g grid.Grid) (moves []grid.Cell) {
	row := g.GetRow(r.Position.X)
	moves = append(moves, row...)

	column := g.GetColumn(r.Position.Y)
	moves = append(moves, column...)

	//	Filter out current grid position
	moves = lo.Filter(moves, func(c grid.Cell, _ int) bool {
		return c.Coordinates != r.Position.Coordinates
	})

	return moves
}
