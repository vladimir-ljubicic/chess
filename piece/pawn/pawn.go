package pawn

import (
	"github.com/chess/board"
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

func (p Pawn) GetLegalMoves(b board.Board) (moves []grid.Cell) {
	moves = append(moves, p.position.Up())

	piece := b.GetPieceOn(p.position.Up().Left())
	if piece != nil && (*piece).GetColor() != p.color {
		moves = append(moves, p.position.Up().Left())
	}

	piece = b.GetPieceOn(p.position.Up().Right())
	if piece != nil && (*piece).GetColor() != p.color {
		moves = append(moves, p.position.Up().Right())
	}

	//	figure out how to distinguish white from black moves here
	if p.position.X == 1 || p.position.X == b.Grid.GetDimensions()-2 {
		moves = append(moves, p.position.Up().Up())
	}

	return moves
}

func (p Pawn) GetPosition() grid.Cell {
	return p.position
}

func (p Pawn) GetColor() piece.Color {
	return p.color
}
