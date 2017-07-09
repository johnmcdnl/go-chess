package chess

type ChessPiece interface {
	ValidMoves(b *Board) []*Move
}

type Piece struct {
	Name  string
	Color Color
}

func NewPiece(name string) *Piece {
	var p Piece
	p.Name = name
	return &p
}
