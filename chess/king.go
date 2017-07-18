package chess

import "github.com/satori/go.uuid"

type King struct {
	Basic
}

func NewKing(s *Square, c Color) (*Knight, error) {
	var k Knight
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
	return moves
}

