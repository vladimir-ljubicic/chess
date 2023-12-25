package pawn

import (
	"github.com/chess/board"
	"github.com/chess/board/grid"
	"github.com/chess/piece"
	"github.com/samber/lo"
)

type Pawn struct {
	moveList []piece.Move
}

func (pw Pawn) UpdateLegalMoves(p *piece.Piece, b board.Board) {
	var moves []grid.Cell

	//	move once if vacant
	if p.Color == piece.White {
		if b.GetPieceOn(p.Position.Up()) == nil {
			moves = append(moves, p.Position.Up())
		}
	} else {
		if b.GetPieceOn(p.Position.Down()) == nil {
			moves = append(moves, p.Position.Down())
		}
	}

	//	capture left
	x := b.GetPieceOn(p.Position.Up().Left())
	if x != nil && (*x).Color != p.Color {
		moves = append(moves, p.Position.Up().Left())
	}

	//	capture right
	x = b.GetPieceOn(p.Position.Up().Right())
	if x != nil && (*x).Color != p.Color {
		moves = append(moves, p.Position.Up().Right())
	}

	//	double move from starting position if both subsequent squares are vacant
	if isPawnStartingPosition(p.Color, p.Position) {
		if p.Color == piece.White {
			if b.GetPieceOn(p.Position.Up()) == nil && b.GetPieceOn(p.Position.Up().Up()) == nil {
				moves = append(moves,
					p.Position.Up().Up(),
				)
			}
		} else {
			if b.GetPieceOn(p.Position.Down()) == nil && b.GetPieceOn(p.Position.Down().Down()) == nil {
				moves = append(moves,
					p.Position.Down().Down(),
				)
			}
		}
	}

	lastMove, err := lo.Last(pw.moveList)
	if err != nil {
		panic(err)
	}

	//	capture "en passant"
	if isDoubleMove(lastMove) && areAdjacentPawns(p.Position, lastMove.Piece.Position) {
		if lastMove.Piece.Position.X < p.Position.X {
			moves = append(moves, p.Position.Up().Left())
		} else {
			moves = append(moves, p.Position.Up().Right())
		}
	}
	// we need to clear captured pieces off the board, this applies to all other scenarios as well

	p.LegalMoves = moves
}

func isPawnStartingPosition(c piece.Color, cell grid.Cell) bool {
	if c == piece.White {
		return cell.Y == 1
	} else {
		return cell.Y == 6
	}
}

func isDoubleMove(move piece.Move) bool {
	if isPawnStartingPosition(move.Piece.Color, move.From) {
		return false
	}

	return (move.Piece.Color == piece.White && move.To.Y == 3) || (move.Piece.Color == piece.Black && move.To.Y == 4)
}

func areAdjacentPawns(i grid.Cell, j grid.Cell) bool {
	return i.X == j.X+1 || i.X == j.X-1
}
