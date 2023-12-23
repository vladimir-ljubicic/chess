package knight

import (
	"github.com/chess/board"
	"github.com/chess/board/grid"
	"github.com/chess/piece"
	"github.com/samber/lo"
)

type Knight struct{}

func (k Knight) UpdateLegalMoves(p *piece.Piece, b board.Board) {
	moves := []grid.Cell{
		p.Position.Up().Up().Left(),
		p.Position.Up().Up().Right(),
		p.Position.Down().Down().Left(),
		p.Position.Down().Down().Right(),
		p.Position.Left().Left().Up(),
		p.Position.Left().Left().Down(),
		p.Position.Right().Right().Up(),
		p.Position.Right().Right().Down(),
	}

	//	Filter out invalid cells
	moves = lo.Filter(moves, func(c grid.Cell, _ int) bool {
		if !b.Grid.IsValidCell(c) {
			return false
		}

		if piece := b.GetPieceOn(c); piece != nil && (*piece).Color == p.Color {
			return false
		}

		return true
	})

	p.LegalMoves = moves
}
