package knight

import (
	"github.com/chess/grid"
	"github.com/chess/piece"
	"github.com/samber/lo"
)

type Knight struct {
	color      piece.Color
	position   grid.Cell
	LegalMoves []grid.Cell
}

func New(startPosition grid.Coordinates, color piece.Color) Knight {
	return Knight{
		color:      color,
		position:   grid.Cell{Coordinates: startPosition},
		LegalMoves: nil,
	}
}

func (k Knight) GetLegalMoves(g grid.Grid) (moves []grid.Cell) {

	moves = append(moves,
		k.position.Up().Up().Left(),
		k.position.Up().Up().Right(),
		k.position.Down().Down().Left(),
		k.position.Down().Down().Right(),
		k.position.Left().Left().Up(),
		k.position.Left().Left().Down(),
		k.position.Right().Right().Up(),
		k.position.Right().Right().Down(),
	)

	//	Filter out invalid cells
	moves = lo.Filter(moves, func(c grid.Cell, _ int) bool {
		return g.IsValidCell(c)
	})

	return moves
}

func (k Knight) GetPosition() grid.Cell {
	return k.position
}

func (k Knight) GetColor() piece.Color {
	return k.color
}
