package chess

import "github.com/satori/go.uuid"

type King struct {
	Basic
}

func NewKing(s *Square, c Color) (*King, error) {
	var k King
	k.ID = uuid.NewV4().String()
	k.Name = "king"
	k.Color = c
	k.Position = s
	k.Type = KingType
	return &k, nil
}

func (k *King) CurrentPosition() *Square {
	return k.Position
}

func (k *King)PieceColor() Color {
	return k.Color
}

func (k *King)PieceType() PieceType {
	return k.Type
}

func (k *King) ValidMoves(b *Board) []*Move {
	var moves []*Move

	if m := NewMove(k.CurrentPosition(), b.GetSquare(k.CurrentPosition().File - 1, k.CurrentPosition().Rank + 1)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(k.CurrentPosition().File + 0, k.CurrentPosition().Rank + 1)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(k.CurrentPosition().File + 1, k.CurrentPosition().Rank + 1)); m != nil {
		moves = append(moves, m)
	}

	if m := NewMove(k.CurrentPosition(), b.GetSquare(k.CurrentPosition().File + 1, k.CurrentPosition().Rank)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(k.CurrentPosition().File - 1, k.CurrentPosition().Rank)); m != nil {
		moves = append(moves, m)
	}

	if m := NewMove(k.CurrentPosition(), b.GetSquare(k.CurrentPosition().File - 1, k.CurrentPosition().Rank - 1)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(k.CurrentPosition().File + 0, k.CurrentPosition().Rank - 1)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(k.CurrentPosition().File + 1, k.CurrentPosition().Rank - 1)); m != nil {
		moves = append(moves, m)
	}
	return moves
}

