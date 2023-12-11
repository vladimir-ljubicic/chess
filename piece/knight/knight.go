package knight

import (
	"github.com/chess/grid"
	"github.com/chess/piece"
	"github.com/samber/lo"
)

type Knight struct {
	Color      piece.Color
	Position   grid.Cell
	LegalMoves []grid.Cell
}

func New(startPosition grid.Coordinates, color piece.Color) Knight {
	return Knight{
		Color:      color,
		Position:   grid.Cell{Coordinates: startPosition},
		LegalMoves: nil,
	}
}

func (k Knight) GetLegalMoves(g grid.Grid) (moves []grid.Cell) {

	moves = append(moves,
		k.Position.GoUp().GoUp().GoLeft(),
		k.Position.GoUp().GoUp().GoRight(),
		k.Position.GoDown().GoDown().GoLeft(),
		k.Position.GoDown().GoDown().GoRight(),
		k.Position.GoLeft().GoLeft().GoUp(),
		k.Position.GoLeft().GoLeft().GoDown(),
		k.Position.GoRight().GoRight().GoUp(),
		k.Position.GoRight().GoRight().GoDown(),
	)

	//	Filter out invalid cells
	moves = lo.Filter(moves, func(c grid.Cell, _ int) bool {
		return g.IsValidCell(c)
	})

	return moves
}
