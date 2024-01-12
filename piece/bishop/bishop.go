package bishop

import (
	"fmt"
	"github.com/chess/board"
	"github.com/chess/board/grid"
	"github.com/chess/piece"
)

type Bishop struct{}

func (bs Bishop) GetLegalMoves(p piece.Piece, b board.Board) []grid.Cell {
	var moves []grid.Cell
	for _, diagonalDirection := range grid.DiagonalDirections {
		move, found := grid.DiagonalMovements[diagonalDirection]
		if !found {
			panic(fmt.Sprintf("No movement method for specified diagonal direction: %s", diagonalDirection))
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
