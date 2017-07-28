package chess

import (
	"github.com/satori/go.uuid"
)

type Pawn struct {
	BasePiece
}

func NewPawn(s *Square, c Color) (*Pawn, error) {
	var p Pawn
	p.ID = uuid.NewV4().String()
	p.Name = "Pawn"
	p.Color = c
	p.Position = s
	p.Type = PawnType
	return &p, nil
}

func (p *Pawn) CurrentPosition() *Square {
	return p.Position
}

func (p *Pawn) PieceColor() Color {
	return p.Color
}

func (p *Pawn) PieceType() PieceType {
	return p.Type
}

func (p *Pawn) Move(b *Board, c Coordinate) {
	p.Position = b.GetSquare(c.File, c.Rank)
	b.GetSquare(c.File, c.Rank).CurrentPiece = p
}

func (p *Pawn) ValidMoves(b *Board) []*Move {
	var moves []*Move

	cp := p.CurrentPosition()
	var direction int
	var isFirstMove bool

	if p.Color == White {
		direction = 1
		if cp.Rank == 2 {
			isFirstMove = true
		}

	}
	if p.Color == Black {
		direction = -1
		if cp.Rank == 7 {
			isFirstMove = true
		}
	}

	if isFirstMove {
		if m := NewMove(cp, b.GetSquare(cp.File, cp.Rank + (2 * direction))); m != nil {
			moves = append(moves, m)
		}
	}

	if m := NewMove(cp, b.GetSquare(cp.File, cp.Rank + (1 * direction))); m != nil {
		moves = append(moves, m)
	}

	leftCapture := b.GetSquare(cp.File - 1, cp.Rank + (1 * direction))
	if leftCapture!=nil && leftCapture.CurrentPiece != nil {
		if m := NewMove(cp, leftCapture); m != nil {
			moves = append(moves, m)
		}
	}
	rightCapture := b.GetSquare(cp.File + 1, cp.Rank + (1 * direction))
	if rightCapture!=nil && rightCapture.CurrentPiece != nil {
		if m := NewMove(cp, rightCapture); m != nil {
			moves = append(moves, m)
		}
	}

	return moves
}
