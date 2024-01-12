package pawn

import (
	"github.com/chess/board"
	"github.com/chess/board/grid"
	"github.com/chess/piece"
	"github.com/samber/lo"
)

type Pawn struct {
	MoveHistory []piece.Move
}

func (pw Pawn) GetLegalMoves(p piece.Piece, b board.Board) []grid.Cell {
	var moves []grid.Cell

	//	move once if vacant
	if p.Color == piece.White {
		up := p.Position.Up()
		if b.Grid.IsValidCell(up) && b.GetPieceOn(up) == nil {
			moves = append(moves, up)
		}
	} else {
		down := p.Position.Down()
		if b.Grid.IsValidCell(down) && b.GetPieceOn(down) == nil {
			moves = append(moves, down)
		}
	}

	//	capture left
	if p.Color == piece.White {
		x := b.GetPieceOn(p.Position.Up().Left())
		if x != nil && (*x).Color != p.Color {
			moves = append(moves, p.Position.Up().Left())
		}
	} else {
		x := b.GetPieceOn(p.Position.Down().Left())
		if x != nil && (*x).Color != p.Color {
			moves = append(moves, p.Position.Down().Left())
		}
	}

	//	capture right
	if p.Color == piece.White {
		x := b.GetPieceOn(p.Position.Up().Right())
		if x != nil && (*x).Color != p.Color {
			moves = append(moves, p.Position.Up().Right())
		}
	} else {
		x := b.GetPieceOn(p.Position.Down().Right())
		if x != nil && (*x).Color != p.Color {
			moves = append(moves, p.Position.Down().Right())
		}
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

	if len(pw.MoveHistory) > 0 {
		lastMove, err := lo.Last(pw.MoveHistory)
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
	}

	return moves
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
