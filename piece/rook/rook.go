package rook

import (
	"fmt"
	"github.com/chess/board"
	"github.com/chess/board/grid"
	"github.com/chess/piece"
)

type Rook struct{}

func (r Rook) GetLegalMoves(p piece.Piece, b board.Board) []grid.Cell {
	var moves []grid.Cell
	for _, direction := range grid.Directions {
		move, found := grid.Movements[direction]
		if !found {
			panic(fmt.Sprintf("No movement method for specified direction: %s", direction))
		}

		x := p.Position
		for {
			x = move(x)

			if !b.Grid.IsValidCell(x) {
				break
			}

			if encountered := b.GetPieceOn(x); encountered != nil {
				if (*encountered).Color != p.Color {
					moves = append(moves, x)
				}
				break
			}

			moves = append(moves, x)
		}
	}

	return moves
}
