package queen

import (
	"fmt"
	"github.com/chess/board"
	"github.com/chess/board/grid"
	"github.com/chess/piece"
)

type Queen struct{}

func (q Queen) UpdateLegalMoves(p *piece.Piece, b board.Board) {
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

	p.LegalMoves = moves
}
