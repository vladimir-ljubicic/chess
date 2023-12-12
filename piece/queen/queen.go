package queen

import (
	"fmt"
	"github.com/chess/board"
	"github.com/chess/grid"
	"github.com/chess/piece"
)

type Queen struct {
	color      piece.Color
	position   grid.Cell
	LegalMoves []grid.Cell
}

func New(startPosition grid.Coordinates, color piece.Color) Queen {
	return Queen{
		color:      color,
		position:   grid.Cell{Coordinates: startPosition},
		LegalMoves: nil,
	}
}

func (q Queen) GetLegalMoves(b board.Board) (moves []grid.Cell) {
	for _, direction := range grid.Directions {
		move, found := grid.Movements[direction]
		if !found {
			panic(fmt.Sprintf("No movement method for specified direction: %s", direction))
		}

		p := q.position
		for {
			p = move(p)

			if !b.Grid.IsValidCell(p) {
				break
			}

			if encountered := b.GetPieceOn(p); encountered != nil {
				if (*encountered).GetColor() != q.color {
					moves = append(moves, p)
				}
				break
			}

			moves = append(moves, p)
		}
	}

	for _, diagonalDirection := range grid.DiagonalDirections {
		move, found := grid.DiagonalMovements[diagonalDirection]
		if !found {
			panic(fmt.Sprintf("No movement method for specified diagonal direction: %s", diagonalDirection))
		}

		p := q.position
		for {
			p = move(p)

			if !b.Grid.IsValidCell(p) {
				break
			}

			if encountered := b.GetPieceOn(p); encountered != nil {
				if (*encountered).GetColor() != q.color {
					moves = append(moves, p)
				}
				break
			}

			moves = append(moves, p)
		}
	}
	
	return moves
}

func (q Queen) GetPosition() grid.Cell {
	return q.position
}

func (q Queen) GetColor() piece.Color {
	return q.color
}
