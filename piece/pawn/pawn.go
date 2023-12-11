package pawn

import (
	"github.com/chess/grid"
	"github.com/chess/piece"
)

type Pawn struct {
	color      piece.Color
	position   grid.Cell
	LegalMoves []grid.Cell
}

func New(startPosition grid.Coordinates, color piece.Color) Pawn {
	return Pawn{
		color:      color,
		position:   grid.Cell{Coordinates: startPosition},
		LegalMoves: nil,
	}
}

func (p Pawn) GetLegalMoves(g grid.Grid) (moves []grid.Cell) {
	moves = append(
		moves,
		grid.Cell{
			Coordinates: grid.Coordinates{
				X: p.position.X + 1,
				Y: p.position.Y,
			},
		},
		grid.Cell{
			Coordinates: grid.Coordinates{
				X: p.position.X + 1,
				Y: p.position.Y - 1,
			},
		},
		grid.Cell{
			Coordinates: grid.Coordinates{
				X: p.position.X + 1,
				Y: p.position.Y + 1,
			},
		},
	)

	if p.position.X == 1 || p.position.X == g.GetDimensions()-1 {
		moves = append(moves, grid.Cell{
			Coordinates: grid.Coordinates{
				X: p.position.X + 2,
				Y: p.position.Y,
			},
		})
	}

	return moves
}

func (p Pawn) GetPosition() grid.Cell {
	return p.position
}

func (p Pawn) GetColor() piece.Color {
	return p.color
}
