package board

import (
	"github.com/chess/grid"
	"github.com/chess/piece"
	"regexp"
	"strconv"
	"strings"
)

type Board struct {
	grid   grid.Grid
	Pieces []piece.Piece
}

func NewBoard(pieceSchema string) Board {

	return Board{
		grid:   grid.NewGrid(8),
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

type Player int

const (
	White Player = iota
	Black
)

func (b Board) IsInCheck(color Player) {
	//king := lo.Find()
	//opposingPieces := lo.Filter(b.Pieces, func(p piece.Piece, _ int)bool{
	//	return b.
	//})
	//for _, piece := range b.Pieces
}
