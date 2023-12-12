package bishop

import (
	"fmt"
	"github.com/chess/board"
	"github.com/chess/grid"
	"github.com/chess/piece"
)

type Bishop struct {
	color      piece.Color
	position   grid.Cell
	LegalMoves []grid.Cell
}

func New(startPosition grid.Coordinates, color piece.Color) Bishop {
	return Bishop{
		color:      color,
		position:   grid.Cell{Coordinates: startPosition},
		LegalMoves: nil,
	}
}

func (bs Bishop) GetLegalMoves(b board.Board) (moves []grid.Cell) {
	for _, diagonalDirection := range grid.DiagonalDirections {
		move, found := grid.DiagonalMovements[diagonalDirection]
		if !found {
			panic(fmt.Sprintf("No movement method for specified diagonal direction: %s", diagonalDirection))
		}

		p := bs.position
		for {
			p = move(p)

			if !b.Grid.IsValidCell(p) {
				break
			}

			if encountered := b.GetPieceOn(p); encountered != nil {
				if (*encountered).GetColor() != bs.color {
					moves = append(moves, p)
				}
				break
			}

			moves = append(moves, p)
		}
	}

	return moves
}

func (bs Bishop) GetPosition() grid.Cell {
	return bs.position
}

func (bs Bishop) GetColor() piece.Color {
	return bs.color
}
