package chess

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
	WhitePawn = PieceTypePawn
	WhiteKnight = PieceTypeKnight
	WhiteBishop = PieceTypeBishop
	WhiteRook = PieceTypeRook
	WhiteQueen = PieceTypeQueen
	WhiteKing = PieceTypeKing
	BlackPawn = PieceTypePawn
	BlackKnight = PieceTypeKnight
	BlackBishop = PieceTypeBishop
	BlackRook = PieceTypeRook
	BlackQueen = PieceTypeQueen
	BlackKing = PieceTypeKing
)

func ValidPieces() []int {
	return []int{
		WhitePawn, WhiteKing, WhiteQueen, WhiteBishop, WhiteKnight, WhiteRook,
		BlackPawn, BlackKing, BlackQueen, BlackBishop, BlackKnight, BlackRook,
	}
}