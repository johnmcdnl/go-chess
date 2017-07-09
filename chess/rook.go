package chess

type Rook Piece

func NewRook(c Color) *Rook {
	var r Rook
	r.Name = "rook"
	r.Color = c
	return &r
}

func (r *Rook)ValidMoves(b *Board) []*Move {
	var arr []*Move
	return arr
}