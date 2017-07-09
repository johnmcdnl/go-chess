package chess

type Bishop Piece

func NewBishop(c Color) *Bishop {
	var b Bishop
	b.Name = "bishop"
	b.Color = c
	return &b
}

func (bishop *Bishop)ValidMoves(board *Board) []*Move {
	var arr []*Move
	return arr
}