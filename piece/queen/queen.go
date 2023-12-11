package queen

import (
	"github.com/chess/grid"
	"github.com/chess/piece"
	"github.com/samber/lo"
)

type Queen struct {
	color      piece.Color
	position   grid.Cell
	LegalMoves []grid.Cell
}

func New(startPosition grid.Coordinates, color piece.Color) Queen {
	return Queen{
		color:      color,
		position:   grid.Cell{Coordinates: startPosition},
		LegalMoves: nil,
	}
}

func (q Queen) GetLegalMoves(g grid.Grid) (moves []grid.Cell) {
	row := g.GetRow(q.position.X)
	moves = append(moves, row...)

	column := g.GetColumn(q.position.Y)
	moves = append(moves, column...)

	diagonals := g.GetDiagonals(q.position)
	moves = append(moves, diagonals...)

	//	Filter out current grid position
	moves = lo.Filter(moves, func(c grid.Cell, _ int) bool {
		return c.Coordinates != q.position.Coordinates
	})

	return moves
}

func (q Queen) GetPosition() grid.Cell {
	return q.position
}

func (q Queen) GetColor() piece.Color {
	return q.color
}
