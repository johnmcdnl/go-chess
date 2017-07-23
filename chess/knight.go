package chess

import (
	"github.com/satori/go.uuid"
)

type Knight struct {
	Basic
}

func NewKnight(s *Square, c Color) (*Knight, error) {
	var k Knight
	k.ID = uuid.NewV4().String()
	k.Name = "knight"
	k.Color = c
	k.Position = s
	k.Type = KnightType
	return &k, nil
}

func (k *Knight) CurrentPosition() *Square {
	return k.Position
}

func (k *Knight)PieceColor() Color {
	return k.Color
}

func (k *Knight)PieceType() PieceType {
	return k.Type
}


func (k *Knight) ValidMoves(b *Board) []*Move {

	var moves []*Move

	if m := NewMove(k.CurrentPosition(), b.GetSquare(k.CurrentPosition().File + 2, k.CurrentPosition().Rank + 1)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(k.CurrentPosition().File + 2, k.CurrentPosition().Rank - 1)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(k.CurrentPosition().File + 1, k.CurrentPosition().Rank + 2)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(k.CurrentPosition().File + 1, k.CurrentPosition().Rank - 2)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(k.CurrentPosition().File - 1, k.CurrentPosition().Rank + 2)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(k.CurrentPosition().File - 1, k.CurrentPosition().Rank - 2)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(k.CurrentPosition().File - 2, k.CurrentPosition().Rank + 1)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(k.CurrentPosition().File - 2, k.CurrentPosition().Rank - 1)); m != nil {
		moves = append(moves, m)
	}

	return moves
}
