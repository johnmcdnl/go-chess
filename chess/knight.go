package chess

import (
	"github.com/satori/go.uuid"
)

type Knight struct {
	BasePiece
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

	cFile := k.CurrentPosition().File
	cRank := k.CurrentPosition().Rank

	if m := NewMove(k.CurrentPosition(), b.GetSquare(cFile + 2, cRank + 1)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(cFile + 2, cRank - 1)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(cFile + 1, cRank + 2)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(cFile + 1, cRank - 2)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(cFile - 1, cRank + 2)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(cFile - 1, cRank - 2)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(cFile - 2, cRank + 1)); m != nil {
		moves = append(moves, m)
	}
	if m := NewMove(k.CurrentPosition(), b.GetSquare(cFile - 2, cRank - 1)); m != nil {
		moves = append(moves, m)
	}

	return moves
}
