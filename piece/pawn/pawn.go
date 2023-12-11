package pawn

import (
	"github.com/chess/grid"
	"github.com/chess/piece"
)

type Pawn struct {
	Color      piece.Color
	Position   grid.Cell
	LegalMoves []grid.Cell
}

func New(startPosition grid.Coordinates, color piece.Color) Pawn {
	return Pawn{
		Color:      color,
		Position:   grid.Cell{Coordinates: startPosition},
		LegalMoves: nil,
	}
}

func (p Pawn) GetLegalMoves(g grid.Grid) (moves []grid.Cell) {
	moves = append(
		moves,
		grid.Cell{
			Coordinates: grid.Coordinates{
				X: p.Position.X + 1,
				Y: p.Position.Y,
			},
		},
		grid.Cell{
			Coordinates: grid.Coordinates{
				X: p.Position.X + 1,
				Y: p.Position.Y - 1,
			},
		},
		grid.Cell{
			Coordinates: grid.Coordinates{
				X: p.Position.X + 1,
				Y: p.Position.Y + 1,
			},
		},
	)

	if p.Position.X == 1 || p.Position.X == g.Dimensions-1 {
		moves = append(moves, grid.Cell{
			Coordinates: grid.Coordinates{
				X: p.Position.X + 2,
				Y: p.Position.Y,
			},
		})
	}

	return moves
}
