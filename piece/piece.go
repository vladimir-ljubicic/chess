package piece

import (
	"github.com/chess/grid"
)

type Color int

const (
	White Color = iota
	Black
)

type Piece interface {
	GetLegalMoves(g grid.Grid) []grid.Cell
}
