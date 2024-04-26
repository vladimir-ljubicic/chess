package test_utils

import (
	"github.com/chess/board"
	"github.com/chess/board/grid"
	"github.com/chess/piece"
)

func GetTestPosition() board.Board {
	b := board.NewBoard("")
	b.Pieces = append(b.Pieces,
		piece.Piece{
			Type:       piece.Pawn,
			Color:      piece.White,
			Position:   grid.Cell{X: 5, Y: 1},
			LegalMoves: nil,
		},
		piece.Piece{
			Type:       piece.Pawn,
			Color:      piece.White,
			Position:   grid.Cell{X: 6, Y: 1},
			LegalMoves: nil,
		},
		piece.Piece{
			Type:       piece.Pawn,
			Color:      piece.White,
			Position:   grid.Cell{X: 7, Y: 1},
			LegalMoves: nil,
		},
		piece.Piece{
			Type:       piece.Pawn,
			Color:      piece.White,
			Position:   grid.Cell{X: 4, Y: 3},
			LegalMoves: nil,
		},
		piece.Piece{
			Type:       piece.Pawn,
			Color:      piece.White,
			Position:   grid.Cell{X: 2, Y: 2},
			LegalMoves: nil,
		},
		piece.Piece{
			Type:       piece.Pawn,
			Color:      piece.White,
			Position:   grid.Cell{X: 2, Y: 4},
			LegalMoves: nil,
		},
		// Black =>
		piece.Piece{
			Type:       piece.Pawn,
			Color:      piece.Black,
			Position:   grid.Cell{X: 7, Y: 6},
			LegalMoves: nil,
		},
		piece.Piece{
			Type:       piece.Pawn,
			Color:      piece.Black,
			Position:   grid.Cell{X: 6, Y: 6},
			LegalMoves: nil,
		},
		piece.Piece{
			Type:       piece.Pawn,
			Color:      piece.Black,
			Position:   grid.Cell{X: 5, Y: 3},
			LegalMoves: nil,
		},
		piece.Piece{
			Type:       piece.Pawn,
			Color:      piece.Black,
			Position:   grid.Cell{X: 3, Y: 4},
			LegalMoves: nil,
		},
		piece.Piece{
			Type:       piece.Pawn,
			Color:      piece.Black,
			Position:   grid.Cell{X: 1, Y: 5},
			LegalMoves: nil,
		},
	)

	return b
}
