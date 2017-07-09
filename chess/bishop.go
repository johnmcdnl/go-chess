package chess

type Bishop Piece

func NewBishop(s *Square, c Color) *Bishop {
	var b Bishop
	b.Name = "bishop"
	b.Code = BishopPiece
	b.Position = s
	b.Color = c

	return &b
}

func (bishop *Bishop) ValidMoves(board *Board) []*Move {
	var arr []*Move
	return arr
}
func (bishop *Bishop) GetCode() int {
	return bishop.Code
}

func (bishop *Bishop) GetColor() Color {
	return bishop.Color
}

func (bishop *Bishop) CurrentPosition() *Square {
	return bishop.Position
}
