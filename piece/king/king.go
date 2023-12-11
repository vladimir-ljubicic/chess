package king

import (
	"github.com/chess/grid"
	"github.com/chess/piece"
)

type King struct {
	Color      piece.Color
	Position   grid.Cell
	LegalMoves []grid.Cell
}

func New(startPosition grid.Coordinates, color piece.Color) King {
	return King{
		Color:      color,
		Position:   grid.Cell{Coordinates: startPosition},
		LegalMoves: nil,
	}
}

func (k King) GetLegalMoves(g grid.Grid) (moves []grid.Cell) {
	//	This will be by far the trickiest one. King cannot move into check, so we need
	//	to check if any other opposing piece has kings destination as a Legal move.

	return moves
}
