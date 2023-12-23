package king

import (
	"github.com/chess/board"
	"github.com/chess/piece"
)

type King struct{}

func (k King) UpdateLegalMoves(p *piece.Piece, b board.Board) {
	//	This will be by far the trickiest one. King cannot move into check, so we need
	//	to check if any other opposing piece has kings destination as a Legal move.

	return
}
