package chess

type Rook Piece

func NewRook(s *Square, c Color) *Rook {
	var r Rook
	r.Name = "rook"
	r.Color = c
	r.Position = s
	r.Code = RookPiece
	return &r
}

func (r *Rook) ValidMoves(b *Board) []*Move {
	var arr []*Move
	return arr
}

func (r *Rook) GetCode() int {
	return r.Code
}

func (r *Rook) GetColor() Color {
	return r.Color
}
func (r *Rook) CurrentPosition() *Square {
	return r.Position
}
