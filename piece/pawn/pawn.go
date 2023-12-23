package pawn

import (
	"github.com/chess/board"
	"github.com/chess/board/grid"
	"github.com/chess/piece"
)

type Pawn struct{}

func (pw Pawn) UpdateLegalMoves(p *piece.Piece, b board.Board) {
	var moves []grid.Cell
	moves = append(moves, p.Position.Up())

	piece := b.GetPieceOn(p.Position.Up().Left())
	if piece != nil && (*piece).Color != p.Color {
		moves = append(moves, p.Position.Up().Left())
	}

	piece = b.GetPieceOn(p.Position.Up().Right())
	if piece != nil && (*piece).Color != p.Color {
		moves = append(moves, p.Position.Up().Right())
	}

	//	figure out how to distinguish white from black moves here
	if p.Position.X == 1 || p.Position.X == b.Grid.GetDimensions()-2 {
		moves = append(moves, p.Position.Up().Up())
	}

	p.LegalMoves = moves
}
