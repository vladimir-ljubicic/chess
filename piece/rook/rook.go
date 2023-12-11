package rook

import (
	"github.com/chess/board"
	"github.com/chess/grid"
	"github.com/chess/piece"
	"github.com/samber/lo"
)

type Rook struct {
	color      piece.Color
	position   grid.Cell
	LegalMoves []grid.Cell
}

func New(startPosition grid.Coordinates, color piece.Color) Rook {
	return Rook{
		color: color,
		position: grid.Cell{
			Coordinates: startPosition,
		},
		LegalMoves: nil,
	}
}

func (r Rook) GetLegalMoves(g grid.Grid) (moves []grid.Cell) {
	row := g.GetRow(r.Position.X)
	row := g.GetRow(r.position.X)
	moves = append(moves, row...)

	column := g.GetColumn(r.position.Y)
	moves = append(moves, column...)

	//	Filter out current grid position
	moves = lo.Filter(moves, func(c grid.Cell, _ int) bool {
		return c.Coordinates != r.position.Coordinates
	})

	return moves
}

func (r Rook) GetPosition() grid.Cell {
	return r.position
}

func (r Rook) GetColor() piece.Color {
	return r.color
}
