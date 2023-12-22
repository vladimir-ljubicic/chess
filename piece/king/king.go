package king

import (
	"github.com/chess/board"
	grid2 "github.com/chess/board/grid"
	"github.com/chess/piece"
)

type King struct{}

func (k King) GetLegalMoves(p piece.Piece, b board.Board) (moves []grid2.Cell) {
	//	This will be by far the trickiest one. King cannot move into check, so we need
	//	to check if any other opposing piece has kings destination as a Legal move.

	return moves
}
