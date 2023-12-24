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


func (g Game) PinnedPiecesUpdateMoves(c player.Color) {
	k, found := lo.Find(g.Board.Pieces, func(p piece.Piece) bool {
		return p.Color == piece.Color(c) && p.Type == piece.King
	})
	if !found {
		panic("King not found")
	}

	for _, direction := range grid.Directions {
		var encounteredPieces []*piece.Piece
		move, found := grid.Movements[direction]
		if !found {
			panic(fmt.Sprintf("No movement method for specified direction: %s", direction))
		}

		x := k.Position
		for {
			x = move(x)

			if !g.Board.Grid.IsValidCell(x) {
				break
			}

			if encountered := g.Board.GetPieceOn(x); encountered != nil {
				encounteredPieces = append(encounteredPieces, encountered)
				if len(encounteredPieces) == 2 {
					break
				}
			}
		}

		if len(encounteredPieces) < 2 || encounteredPieces[0].Color != k.Color || encounteredPieces[1].Color == k.Color {
			continue
		}

		if encounteredPieces[1].Type == piece.Queen || encounteredPieces[1].Type == piece.Rook {
			encounteredPieces[0].LegalMoves = lo.Filter(encounteredPieces[0].LegalMoves,
				func(c grid.Cell, _ int) bool {
					if direction == grid.Right || direction == grid.Left {
						return c.Y == k.Position.Y
					} else {
						return c.X == k.Position.X
					}
				})
		}

	}

	for _, direction := range grid.DiagonalDirections {
		var encounteredPieces []*piece.Piece
		move, found := grid.DiagonalMovements[direction]
		if !found {
			panic(fmt.Sprintf("No movement method for specified direction: %s", direction))
		}

		x := k.Position
		for {
			x = move(x)

			if !g.Board.Grid.IsValidCell(x) {
				break
			}

			if encountered := g.Board.GetPieceOn(x); encountered != nil {
				encounteredPieces = append(encounteredPieces, encountered)
				if len(encounteredPieces) == 2 {
					break
				}
			}
		}

		if len(encounteredPieces) < 2 || encounteredPieces[0].Color != k.Color || encounteredPieces[1].Color == k.Color {
			continue
		}

		if encounteredPieces[1].Type == piece.Queen || encounteredPieces[1].Type == piece.Bishop {
			encounteredPieces[0].LegalMoves = lo.Filter(encounteredPieces[0].LegalMoves,
				func(c grid.Cell, _ int) bool {
					switch direction {
					case grid.RightAscending:
						return c.X > k.Position.X && c.Y > k.Position.Y && c.X-k.Position.X == c.Y-k.Position.Y
					case grid.RightDescending:
						return c.X > k.Position.X && c.Y < k.Position.Y && c.X-k.Position.X*-1 == c.Y-k.Position.Y
					case grid.LeftAscending:
						return c.X < k.Position.X && c.Y > k.Position.Y && c.X-k.Position.X*-1 == c.Y-k.Position.Y
					case grid.LeftDescending:
						return c.X < k.Position.X && c.Y < k.Position.Y && c.X-k.Position.X == c.Y-k.Position.Y
					default:
						return false
					}
				})
		}

	}
}

func (g Game) IsInCheck(c player.Color) bool {
	k, found := lo.Find(g.Board.Pieces, func(p piece.Piece) bool {
		return p.Color == piece.Color(c) && p.Type == piece.King
	})
	if !found {
		panic("King not found")
	}

	opposingPieces := lo.Filter(g.Board.Pieces, func(p piece.Piece, _ int) bool {
		return p.Color != piece.Color(c)
	})

	for _, op := range opposingPieces {
		if lo.Contains(op.LegalMoves, k.Position) {
			return true
		}
	}

	return false
}
