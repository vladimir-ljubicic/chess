package game

import (
	"github.com/chess/board"
	"github.com/chess/player"
	"time"
)

type PieceMove struct {
	Piece piece.Piece
	From  grid.Cell
	To    grid.Cell
}

type Game struct {
	Board    board.Board
	Players  []*player.Player
	MoveList []PieceMove
}

type GameOptions struct {
	time        time.Duration
	increment   time.Duration
	toMove      player.Color
	boardSchema string
}

type GameOption func(o *GameOptions)

var DefaultOptions = GameOptions{
	time:        10 * time.Minute,
	increment:   0 * time.Second,
	boardSchema: "default schema",
}

func WithTime(t int) func(o *GameOptions) {
	return func(o *GameOptions) {
		o.time = time.Duration(t) * time.Minute
	}
}

func WithIncrement(i int) func(o *GameOptions) {
	return func(o *GameOptions) {
		o.increment = time.Duration(i) * time.Second
	}
}

func WithBoardSchema(bs string) func(o *GameOptions) {
	return func(o *GameOptions) {
		o.boardSchema = bs
	}
}

func NewGame(options ...GameOption) Game {
	o := DefaultOptions
	for _, fn := range options {
		fn(&o)
	}
	return Game{
		Board: board.NewBoard(o.boardSchema),
		Players: []*player.Player{
			player.NewPlayer(player.White, o.time, o.increment),
			player.NewPlayer(player.Black, o.time, o.increment),
		},
	}
}
