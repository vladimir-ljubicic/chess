package bishop

import (
	"github.com/chess/grid"
	"github.com/chess/piece"
)

type Bishop struct {
	Color      piece.Color
	Position   grid.Cell
	LegalMoves []grid.Cell
}

func New(startPosition grid.Coordinates, color piece.Color) Bishop {
	return Bishop{
		Color:      color,
		Position:   grid.Cell{Coordinates: startPosition},
		LegalMoves: nil,
	}
}

func (b Bishop) GetLegalMoves(g grid.Grid) (moves []grid.Cell) {
	//	it's better if grid returns all diagonal cells and after that we filter out initial cell here
	//	maybe use ascending + descending diagonal
	return g.GetDiagonals(b.Position)
}
