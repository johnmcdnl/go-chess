package chess

import "github.com/satori/go.uuid"

type Queen struct {
	Basic
}

func NewQueen(s *Square, c Color) (*Queen, error) {
	var q Queen
	q.ID = uuid.NewV4().String()
	q.Name = "Queen"
	q.Color = c
	q.Position = s
	q.Type = QueenType
	return &q, nil
}

func (q *Queen) CurrentPosition() *Square {
	return q.Position
}

func (q *Queen)PieceColor() Color {
	return q.Color
}

func (q *Queen)PieceType() PieceType {
	return q.Type
}

func (q *Queen) ValidMoves(b *Board) []*Move {
	var moves []*Move

	return moves
}

