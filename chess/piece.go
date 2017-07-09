package chess

const (
	_ = iota
	PawnPiece
	KnightPiece
	BishopPiece
	RookPiece
	QueenPiece
	KingPiece
)

type ChessPiece interface {
	ValidMoves(b *Board) []*Move
	GetCode() int
	GetColor() Color
	CurrentPosition() *Square
}

type Piece struct {
	Name     string
	Code     int
	Color    Color
	Position *Square
}

func NewPiece(name string) *Piece {
	var p Piece
	p.Name = name
	return &p
}
