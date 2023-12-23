package board

import (
	"github.com/chess/board/grid"
	"github.com/chess/piece"
	"github.com/samber/lo"
	"regexp"
	"strconv"
	"strings"
)

type MovementEngine interface {
	UpdateLegalMoves(p *piece.Piece, b Board)
}

type Board struct {
	Grid   grid.Grid
	Pieces []piece.Piece
}

func NewBoard(pieceSchema string) Board {
	return Board{
		Grid:   grid.NewGrid(8),
		Pieces: nil,
	}
}

const pattern = "[A-ha-h][1-8]"

var mapBoardRowToXCoordinate = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
}

func ToGridCell(space string) *grid.Cell {
	matched, stdErr := regexp.MatchString(pattern, space)
	if stdErr != nil {
		panic(stdErr)
	}
	if !matched {
		return nil
	}

	elems := strings.Split(space, "")
	row := mapBoardRowToXCoordinate[elems[0]]
	column, stdErr := strconv.Atoi(elems[1])
	if stdErr != nil {
		panic(stdErr)
	}

	return &grid.Cell{
		Coordinates: grid.Coordinates{
			X: row,
			Y: column + 1,
		}}
}

func (b Board) IsInCheck(player piece.Color) bool {
	//king, _ := lo.Find(b.Pieces, func(piece piece.Piece) bool {
	//	king, isKing := piece.(king.King)
	//
	//	return isKing && king.GetColor() == player
	//})
	//
	//opposingPieces := lo.Filter(b.Pieces, func(p piece.Piece, _ int) bool {
	//	return p.GetColor() != player
	//})
	//for _, piece := range opposingPieces {
	//	if lo.Contains(piece.UpdateLegalMoves(b), king.GetPosition()) {
	//		return true
	//	}
	//}

	return false
}

func (b Board) GetPieceOn(c grid.Cell) *piece.Piece {
	if p, found := lo.Find(b.Pieces, func(piece piece.Piece) bool {
		return piece.Position == c
	}); found {
		return &p
	}

	return nil
}
