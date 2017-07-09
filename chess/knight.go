package chess

type Knight Piece

func NewKnight(c Color) *Knight {
	var k Knight
	k.Name = "knight"
	k.Color = c
	return &k
}

func (k *Knight)ValidMoves(b *Board)[]*Move {
	var arr []*Move
	return arr
}