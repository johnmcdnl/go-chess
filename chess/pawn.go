package chess

import "github.com/satori/go.uuid"

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

func (p *Pawn)PieceColor() Color {
	return p.Color
}

func (p *Pawn)PieceType() PieceType {
	return p.Type
}

func (p *Pawn) ValidMoves(b *Board) []*Move {
	var moves []*Move

	return moves
}

