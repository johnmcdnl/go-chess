package chess

const (
	NoPiece = 0
)

const (
	White = iota
	Black
)

const (
	_ = iota
	PieceTypePawn
	PieceTypeKnight
	PieceTypeBishop
	PieceTypeRook
	PieceTypeQueen
	PieceTypeKing
)
const (
	_ = iota
	WhitePawn
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing
	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
)

func ValidPieces() []int {
	return []int{
		WhitePawn, WhiteKing, WhiteQueen, WhiteBishop, WhiteKnight, WhiteRook,
		BlackPawn, BlackKing, BlackQueen, BlackBishop, BlackKnight, BlackRook,
	}
}

type ChessPiece interface {
	ValidMoves() *Moves
	GetPiece() *Piece
}

type Piece struct {
	Type            int
	Color           int
	Name            string
	CurrentPosition *Square
}

