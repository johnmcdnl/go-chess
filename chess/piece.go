package chess

type BasePiece struct {
	ID       string
	Color    Color
	Name     string
	Position *Square
	Type     PieceType
}

type PieceType int

const (
	_ = PieceType(iota)
	KingType
	QueenType
	RookType
	BishopType
	KnightType
	PawnType
)

type Piece interface {
	ValidMoves(b *Board) []*Move
	CurrentPosition() *Square
	PieceColor() Color
	PieceType() PieceType
	Move(b *Board, newPosition Coordinate)
}
