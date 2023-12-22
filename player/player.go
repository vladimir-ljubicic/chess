package player

import (
	"github.com/chess/board/grid"
	"github.com/chess/piece"
	"github.com/samber/lo"
	"time"
)

type Color int

const (
	White Color = iota
	Black
)

type Player struct {
	Color     Color
	Clock     time.Duration
	Increment time.Duration
}

func (pl Player) Move(p *piece.Piece, c grid.Cell) {
	if lo.Contains(p.LegalMoves, c) {
		p.Position = c
	}
}

func NewPlayer(color Color, clock time.Duration, increment time.Duration) *Player {
	return &Player{
		Color:     color,
		Clock:     clock,
		Increment: increment,
	}
}
