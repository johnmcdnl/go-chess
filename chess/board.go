package chess

type Board struct {
	Squares *[]*Square
}

func NewBoard() *Board {
	var b Board
	var squares []*Square
	for rank := 1; rank <= 8; rank++ {
		for file := 1; file <= 8; file++ {
			s := NewSquare(file, rank)
			squares = append(squares, s)
		}
	}
	b.Squares = &squares
	b.setStartingPosition()
	return &b
}

func (b *Board)setStartingPosition() {
	for _, s := range *b.Squares {
		if s.Name == "b1" || s.Name == "g1" {
			k := NewKnight(White)
			s.ChessPiece = k
		}
		if s.Name == "b8" || s.Name == "g8" {
			k := NewKnight(Black)
			s.ChessPiece = k
		}

		if s.Name == "a1" || s.Name == "h1" {
			k := NewRook(White)
			s.ChessPiece = k
		}
		if s.Name == "a8" || s.Name == "h8" {
			k := NewRook(Black)
			s.ChessPiece = k
		}

		if s.Name == "c1" || s.Name == "f1" {
			k := NewBishop(White)
			s.ChessPiece = k
		}
		if s.Name == "c8" || s.Name == "f8" {
			k := NewBishop(Black)
			s.ChessPiece = k
		}
	}
}

func (b *Board)SetFEN(fen FEN) {
	if !fen.IsValid() {
		return
	}
}

func (b *Board)GetFen() FEN {
	return StartingPositionFEN
}
