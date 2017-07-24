package chess

import (
	"github.com/satori/go.uuid"
)

type Queen struct {
	BasePiece
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

	cFile := q.CurrentPosition().File
	cRank := q.CurrentPosition().Rank

	//up right
	for i := 1; i <= 8; i++ {
		if cFile + i <= 8 && cRank + i <= 8 {
			d := b.GetSquare(cFile + i, cRank + i)
			if m := NewMove(q.CurrentPosition(), d); m != nil {
				moves = append(moves, m)

			}
			if d.CurrentPiece != nil {
				break
			}
		}
	}

	//up left
	for i := 1; i <= 8; i++ {
		if cFile - i >= 1 && cRank + i <= 8 {
			d := b.GetSquare(cFile - i, cRank + i)
			if m := NewMove(q.CurrentPosition(), d); m != nil {
				moves = append(moves, m)

			}
			if d.CurrentPiece != nil {
				break
			}

		}
	}

	//down right
	for i := 1; i <= 8; i++ {
		if cFile + i <= 8 && cRank - i >= 1 {
			d := b.GetSquare(cFile + i, cRank - i)
			if m := NewMove(q.CurrentPosition(), d); m != nil {
				moves = append(moves, m)
			}
			if d.CurrentPiece != nil {
				break
			}
		}
	}

	//down right
	for i := 1; i <= 8; i++ {
		if cFile - i >= 1 && cRank - i >= 1 {
			d := b.GetSquare(cFile - i, cRank - i)
			if m := NewMove(q.CurrentPosition(), d); m != nil {
				moves = append(moves, m)
			}
			if d.CurrentPiece != nil {
				break
			}
		}
	}
	return moves
}

