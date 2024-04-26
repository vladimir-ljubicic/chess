package pawn

import (
	"testing"

	"github.com/chess/board/grid"
	"github.com/chess/piece"
	"github.com/chess/piece/pawn/test_utils"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func Test_PawnMove(t *testing.T) {
	b := test_utils.GetTestPosition()

	//	White pawn
	pawnH2, _ := lo.Find(b.Pieces, func(p piece.Piece) bool {
		return p.Position == grid.Cell{X: 7, Y: 1}
	})

	movable := Pawn{}
	moves := movable.GetLegalMoves(pawnH2, b)

	require.Contains(t, moves, grid.Cell{X: 7, Y: 2})

	//	Black pawn
	pawnH7, _ := lo.Find(b.Pieces, func(p piece.Piece) bool {
		return p.Position == grid.Cell{X: 7, Y: 6}
	})

	moves = movable.GetLegalMoves(pawnH7, b)
	require.Contains(t, moves, grid.Cell{X: 7, Y: 5})
}

func Test_PawnDoubleMove(t *testing.T) {
	b := test_utils.GetTestPosition()

	//	White pawn
	pawnH2, _ := lo.Find(b.Pieces, func(p piece.Piece) bool {
		return p.Position == grid.Cell{X: 7, Y: 1}
	})

	movable := Pawn{}
	moves := movable.GetLegalMoves(pawnH2, b)

	require.Contains(t, moves, grid.Cell{X: 7, Y: 3})

	//	Black pawn
	pawnH7, _ := lo.Find(b.Pieces, func(p piece.Piece) bool {
		return p.Position == grid.Cell{X: 7, Y: 6}
	})

	moves = movable.GetLegalMoves(pawnH7, b)
	require.Contains(t, moves, grid.Cell{X: 7, Y: 4})
}

func Test_PawnCaptureLeft(t *testing.T) {
	b := test_utils.GetTestPosition()

	pawnE4, _ := lo.Find(b.Pieces, func(p piece.Piece) bool {
		return p.Position == grid.Cell{X: 4, Y: 3}
	})

	movable := Pawn{}
	moves := movable.GetLegalMoves(pawnE4, b)

	require.Contains(t, moves, grid.Cell{X: 3, Y: 4})
}

func Test_PawnCaptureRight(t *testing.T) {
	b := test_utils.GetTestPosition()

	pawnE4, _ := lo.Find(b.Pieces, func(p piece.Piece) bool {
		return p.Position == grid.Cell{X: 3, Y: 4}
	})

	movable := Pawn{}
	moves := movable.GetLegalMoves(pawnE4, b)

	require.Contains(t, moves, grid.Cell{X: 4, Y: 3})
}

func Test_PawnEnPassantWhite(t *testing.T) {
	b := test_utils.GetTestPosition()

	pawnC3, _ := lo.Find(b.Pieces, func(p piece.Piece) bool {
		return p.Position == grid.Cell{X: 2, Y: 4}
	})
	movable := Pawn{
		MoveHistory: []piece.Move{
			{
				Piece: piece.Piece{
					Type:     piece.Pawn,
					Color:    piece.Black,
					Position: grid.Cell{X: 3, Y: 4},
				},
				From: grid.Cell{X: 3, Y: 6},
				To:   grid.Cell{X: 3, Y: 4},
			},
		},
	}
	moves := movable.GetLegalMoves(pawnC3, b)

	require.Contains(t, moves, grid.Cell{X: 3, Y: 5})
}

func Test_PawnEnPassantBlack(t *testing.T) {
	b := test_utils.GetTestPosition()

	pawnF4, _ := lo.Find(b.Pieces, func(p piece.Piece) bool {
		return p.Position == grid.Cell{X: 5, Y: 3}
	})
	movable := Pawn{
		MoveHistory: []piece.Move{
			{
				Piece: piece.Piece{
					Type:     piece.Pawn,
					Color:    piece.White,
					Position: grid.Cell{X: 4, Y: 3},
				},
				From: grid.Cell{X: 4, Y: 1},
				To:   grid.Cell{X: 4, Y: 3},
			},
		},
	}
	moves := movable.GetLegalMoves(pawnF4, b)

	require.Contains(t, moves, grid.Cell{X: 4, Y: 2})
}
