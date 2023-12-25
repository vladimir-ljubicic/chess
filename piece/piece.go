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

func (p *Piece) Move(c grid.Cell) {
	p.Position = c
}

type Move struct {
	Piece Piece
	From  grid.Cell
	To    grid.Cell
}
