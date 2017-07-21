package chess

import "github.com/satori/go.uuid"

type Bishop struct {
	Basic
}

func NewBishop(s *Square, c Color) (*Bishop, error) {
	var bishop Bishop
	bishop.ID = uuid.NewV4().String()
	bishop.Name = "Bishop"
	bishop.Color = c
	bishop.Position = s
	bishop.Type = BishopType
	return &bishop, nil
}

func (bishop *Bishop) CurrentPosition() *Square {
	return bishop.Position
}

func (bishop *Bishop)PieceColor() Color {
	return bishop.Color
}

func (bishop *Bishop)PieceType() PieceType {
	return bishop.Type
}

func (bishop *Bishop) ValidMoves(board *Board) []*Move {
	var moves []*Move

	return moves
}

