package king

import (
	"github.com/chess/board"
	"github.com/chess/board/grid"
	"github.com/chess/piece"
	"github.com/samber/lo"
)

type King struct{}

func (k King) UpdateLegalMoves(p *piece.Piece, b board.Board) {
	var moves []grid.Cell

	moves = append(moves,
		p.Position.Up(),
		p.Position.Down(),
		p.Position.Left(),
		p.Position.Right(),
		p.Position.Up().Right(),
		p.Position.Up().Left(),
		p.Position.Down().Right(),
		p.Position.Down().Left(),
	)

	opposingPieces := lo.Filter(b.Pieces, func(pc piece.Piece, _ int) bool {
		return pc.Color != p.Color
	})

	for _, op := range opposingPieces {
		diff, _ := lo.Difference(moves, op.LegalMoves)
		moves = diff
	}

	// if len(moves) == 0 signify checkmate , otherwise =>
	p.LegalMoves = moves
}
