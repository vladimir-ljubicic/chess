package game

import (
	"fmt"
	"github.com/chess/board"
	"github.com/chess/board/grid"
	"github.com/chess/piece"
	"github.com/chess/piece/bishop"
	"github.com/chess/piece/king"
	"github.com/chess/piece/knight"
	"github.com/chess/piece/pawn"
	"github.com/chess/piece/queen"
	"github.com/chess/piece/rook"
	"github.com/chess/player"
	"github.com/samber/lo"
	"sync"
	"time"
)

type Game struct {
	Board       board.Board
	Players     []*player.Player
	MoveHistory []piece.Move
}

type GameOptions struct {
	time        time.Duration
	increment   time.Duration
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

func GetMovableForPieceType(t piece.Type, moveHistory []piece.Move) board.Movable {
	switch t {
	case piece.King:
		return king.King{MoveHistory: moveHistory}
	case piece.Queen:
		return queen.Queen{}
	case piece.Rook:
		return rook.Rook{}
	case piece.Bishop:
		return bishop.Bishop{}
	case piece.Knight:
		return knight.Knight{}
	case piece.Pawn:
		return pawn.Pawn{MoveHistory: moveHistory}
	default:
		panic(fmt.Sprintf("Movement engine not found for piece type %d", t))
	}
}

func (g Game) MovePiece(p *piece.Piece, moveTo grid.Cell) {
	if !lo.Contains(p.LegalMoves, moveTo) {
		return
	}
	// Capture
	g.Board.Pieces = lo.Reject(g.Board.Pieces, func(i piece.Piece, _ int) bool {
		return i.Position == moveTo
	})

	moveFrom := p.Position
	p.Move(moveTo)

	g.PostMoveActions(piece.Move{
		Piece: *p,
		From:  moveFrom,
		To:    moveTo,
	})
}

func (g Game) PostMoveActions(move piece.Move) {
	g.AddToMoveHistory(move)

	nextPlayer := player.Color(move.Piece.Color*-1 + 1)
	g.UpdateLegalMovesInParallel(nextPlayer)
	g.handleChecks(nextPlayer)
	g.handlePinnedPieces(nextPlayer)
}

func (g Game) AddToMoveHistory(move piece.Move) {
	g.MoveHistory = append(g.MoveHistory, move)
}

func (g Game) handlePinnedPieces(c player.Color) {
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

func (g Game) handleChecks(c player.Color) {
	k, found := lo.Find(g.Board.Pieces, func(p piece.Piece) bool {
		return p.Color == piece.Color(c) && p.Type == piece.King
	})
	if !found {
		panic("King not found")
	}

	opposingPieces := lo.Filter(g.Board.Pieces, func(p piece.Piece, _ int) bool {
		return p.Color != piece.Color(c)
	})

	var checkingPieces []piece.Piece
	for _, op := range opposingPieces {
		if lo.Contains(op.LegalMoves, k.Position) {
			checkingPieces = append(checkingPieces, op)
		}
	}

	if len(checkingPieces) == 0 {
		return
	}

	//	If double-check
	if len(checkingPieces) > 1 {
		//	If king has no legal moves end the game
		if len(k.LegalMoves) == 0 {
			//	Game ends
		} else {
			//	Otherwise clear legal moves of other pieces (king has to move)
			lo.ForEach(g.Board.Pieces, func(p piece.Piece, _ int) {
				if p.Color == k.Color {
					p.LegalMoves = nil
				}
			})
		}
	} else {
		//	Allow blocking moves only
		lo.ForEach(g.Board.Pieces, func(p piece.Piece, _ int) {
			if p.Color == k.Color {
				medialSquares := grid.GetMedialCells(checkingPieces[0].Position, k.Position)
				p.LegalMoves = lo.Intersect(medialSquares, p.LegalMoves)
			}
		})
	}
}

func (g Game) UpdateLegalMovesInParallel(c player.Color) {
	var wg sync.WaitGroup
	playerPieces := lo.Filter(g.Board.Pieces, func(p piece.Piece, _ int) bool {
		return p.Color == piece.Color(c)
	})
	for _, p := range playerPieces {
		wg.Add(1)
		go func(p piece.Piece) {
			defer wg.Done()
			movable := GetMovableForPieceType(p.Type, g.MoveHistory)
			moves := movable.GetLegalMoves(p, g.Board)
			_, index, _ := lo.FindIndexOf(g.Board.Pieces, func(item piece.Piece) bool {
				return item.Position == p.Position
			})

			p.LegalMoves = moves
			g.Board.Pieces[index] = p
		}(p)
	}
	wg.Wait()
}
