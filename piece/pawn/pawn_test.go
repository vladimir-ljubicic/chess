package pawn

import (
	"github.com/chess/board/grid"
	"github.com/chess/piece"
	"github.com/chess/piece/test_utils"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_PawnMove(t *testing.T) {
	b := test_utils.GetTestPosition()

	//	White pawn
	pawnH2, _ := lo.Find(b.Pieces, func(p piece.Piece) bool {
		return p.Position == grid.Cell{
			grid.Coordinates{
				X: 7,
				Y: 1,
			}}
	})

	movable := Pawn{}
	moves := movable.GetLegalMoves(pawnH2, b)

	require.Contains(t, moves, grid.Cell{
		grid.Coordinates{
			X: 7,
			Y: 2,
		},
	})

	//	Black pawn
	pawnH7, _ := lo.Find(b.Pieces, func(p piece.Piece) bool {
		return p.Position == grid.Cell{
			grid.Coordinates{
				X: 7,
				Y: 6,
			}}
	})

	moves = movable.GetLegalMoves(pawnH7, b)
	require.Contains(t, moves, grid.Cell{
		grid.Coordinates{
			X: 7,
			Y: 5,
		},
	})
}

func Test_PawnDoubleMove(t *testing.T) {
	b := test_utils.GetTestPosition()

	//	White pawn
	pawnH2, _ := lo.Find(b.Pieces, func(p piece.Piece) bool {
		return p.Position == grid.Cell{
			grid.Coordinates{
				X: 7,
				Y: 1,
			}}
	})

	movable := Pawn{}
	moves := movable.GetLegalMoves(pawnH2, b)

	require.Contains(t, moves, grid.Cell{
		grid.Coordinates{
			X: 7,
			Y: 3,
		},
	})

	//	Black pawn
	pawnH7, _ := lo.Find(b.Pieces, func(p piece.Piece) bool {
		return p.Position == grid.Cell{
			grid.Coordinates{
				X: 7,
				Y: 6,
			}}
	})

	moves = movable.GetLegalMoves(pawnH7, b)
	require.Contains(t, moves, grid.Cell{
		grid.Coordinates{
			X: 7,
			Y: 4,
		},
	})
}

func Test_PawnCaptureLeft(t *testing.T) {
	b := test_utils.GetTestPosition()

	pawnE4, _ := lo.Find(b.Pieces, func(p piece.Piece) bool {
		return p.Position == grid.Cell{
			grid.Coordinates{
				X: 4,
				Y: 3,
			}}
	})

	movable := Pawn{}
	moves := movable.GetLegalMoves(pawnE4, b)

	require.Contains(t, moves, grid.Cell{
		grid.Coordinates{
			X: 3,
			Y: 4,
		},
	})
}

func Test_PawnCaptureRight(t *testing.T) {
	b := test_utils.GetTestPosition()

	pawnE4, _ := lo.Find(b.Pieces, func(p piece.Piece) bool {
		return p.Position == grid.Cell{
			grid.Coordinates{
				X: 3,
				Y: 4,
			}}
	})

	movable := Pawn{}
	moves := movable.GetLegalMoves(pawnE4, b)

	require.Contains(t, moves, grid.Cell{
		grid.Coordinates{
			X: 4,
			Y: 3,
		},
	})
}
