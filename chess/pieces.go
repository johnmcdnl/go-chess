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

type Pawn struct {
	*Piece
}

func NewPawn(currentPosition *Square, color int) *Pawn {
	return &Pawn{&Piece{
		Type:  PieceTypePawn,
		Color: color,
		CurrentPosition:currentPosition,
	},
	}
}
func (p *Pawn) ValidMoves() *Moves {
	var moves *Moves
	return moves
}
func (p *Pawn) GetPiece() *Piece {
	return p.Piece
}

