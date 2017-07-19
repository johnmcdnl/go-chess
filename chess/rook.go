package chess

import "github.com/satori/go.uuid"

type Rook struct {
	Basic
}

func NewRook(s *Square, c Color) (*Rook, error) {
	var r Rook
	r.ID = uuid.NewV4().String()
	r.Name = "Rook"
	r.Color = c
	r.Position = s
	r.Type = RookType
	return &r, nil
}

func (r*Rook) CurrentPosition() *Square {
	return r.Position
}

func (r *Rook)PieceColor() Color {
	return r.Color
}

func (r *Rook)PieceType() PieceType {
	return r.Type
}

func (r *Rook) ValidMoves(b *Board) []*Move {
	var moves []*Move

	return moves
}

