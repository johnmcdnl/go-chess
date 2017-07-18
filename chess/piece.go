package chess

type Basic struct {
	ID       string
	Color    Color
	Name     string
	Position *Square
}

type Piece interface {
	ValidMoves(b *Board) []*Move
	CurrentPosition() *Square
}
