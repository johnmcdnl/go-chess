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

func (b *Board) setStartingPosition() {
	for _, s := range *b.Squares {
		if s.Rank == 7 {
			cp := NewPawn(s, White)
			s.ChessPiece = cp
		}
		if s.Rank == 2 {
			cp := NewPawn(s, Black)
			s.ChessPiece = cp
		}

		if s.Name == "b1" || s.Name == "g1" {
			cp := NewKnight(s, White)
			s.ChessPiece = cp
		}
		if s.Name == "b8" || s.Name == "g8" {
			cp := NewKnight(s, Black)
			s.ChessPiece = cp
		}

		if s.Name == "a1" || s.Name == "h1" {
			cp := NewRook(s, White)
			s.ChessPiece = cp
		}
		if s.Name == "a8" || s.Name == "h8" {
			cp := NewRook(s, Black)
			s.ChessPiece = cp
		}

		if s.Name == "c1" || s.Name == "f1" {
			cp := NewBishop(s, White)
			s.ChessPiece = cp
		}
		if s.Name == "c8" || s.Name == "f8" {
			cp := NewBishop(s, Black)
			s.ChessPiece = cp
		}

	}
}

func (b *Board) SetFEN(fen FEN) {
	if !fen.IsValid() {
		return
	}
}

func (b *Board) GetFen() FEN {
	return StartingPositionFEN
}

func (b *Board) GetSquare(file, rank int) *Square {
	if file < 1 || file > 8 || rank < 1 || rank > 8 {
		return nil
	}
	for _, s := range *b.Squares {
		if s.File == file && s.Rank == rank {
			return s
		}
	}
	return nil
}
