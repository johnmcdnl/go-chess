package chess

type Pawn Piece

func NewPawn(s *Square, c Color) *Pawn {
	var b Pawn
	b.Name = "Pawn"
	b.Code = PawnPiece
	b.Position = s
	b.Color = c

	return &b
}

func (p *Pawn) ValidMoves(board *Board) []*Move {
	var arr []*Move
	return arr
}
func (p *Pawn) GetCode() int {
	return p.Code
}

func (p *Pawn) GetColor() Color {
	return p.Color
}

func (p *Pawn) CurrentPosition() *Square {
	return p.Position
}
