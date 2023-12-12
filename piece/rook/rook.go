package rook

import (
	"fmt"
	"github.com/chess/board"
	"github.com/chess/grid"
	"github.com/chess/piece"
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

func (r Rook) GetLegalMoves(b board.Board) (moves []grid.Cell) {
	for _, direction := range grid.Directions {
		move, found := grid.Movements[direction]
		if !found {
			panic(fmt.Sprintf("No movement method for specified direction: %s", direction))
		}

		p := r.position
		for {
			p = move(p)

			if !b.Grid.IsValidCell(p) {
				break
			}

			if encountered := b.GetPieceOn(p); encountered != nil {
				if (*encountered).GetColor() != r.color {
					moves = append(moves, p)
				}
				break
			}

			moves = append(moves, p)
		}
	}

	return moves
}

func (r Rook) GetPosition() grid.Cell {
	return r.position
}

func (r Rook) GetColor() piece.Color {
	return r.color
}
