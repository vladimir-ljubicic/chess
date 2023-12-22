package piece

import (
	"github.com/chess/board/grid"
)

type Color int

const (
	White Color = iota
	Black
)

type Type int

const (
	King Type = iota
	Queen
	Rook
	Bishop
	Knight
	Pawn
)

type Piece struct {
	Type       Type
	Color      Color
	Position   grid.Cell
	LegalMoves []grid.Cell
}
