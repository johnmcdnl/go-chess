package chess

import (
	"fmt"
)

type Board struct {
	Squares []*Square
}

func NewBoard() (*Board, error) {
	var b Board

	if err := b.build(8, 8); err != nil {
		return nil, err
	}
	if err := b.LoadFromFEN(FEN("TODO")); err != nil {
		return nil, err
	}

	return &b, nil
}

func (b *Board) build(files, ranks int) error {
	for rank := 1; rank <= ranks; rank++ {
		for file := 1; file <= files; file++ {
			s, err := NewSquare(file, rank)
			if err != nil {
				return err
			}
			b.Squares = append(b.Squares, s)
		}
	}
	return nil
}

func (b *Board) LoadFromFEN(f FEN) error {
	if !f.IsValid() {
		return fmt.Errorf("Invalid FEN %s", f)
	}

	for _, s := range b.Squares {
		//White
		if (s.File == B && s.Rank == 1) || (s.File == G && s.Rank == 1) {
			k, _ := NewKnight(s, White)
			s.SetPiece(k)
		}

		//White
		if s.File == D && s.Rank == 4  {
			k, _ := NewKing(s, White)
			s.SetPiece(k)
		}


		//Black
		if (s.File == B && s.Rank == 8) || (s.File == G && s.Rank == 8) {
			k, _ := NewKnight(s, Black)
			s.SetPiece(k)
		}

		//White
		if s.File == E && s.Rank == 8  {
			k, _ := NewKing(s, Black)
			s.SetPiece(k)
		}

	}

	return nil
}

func (b *Board) GetSquare(file, rank int) *Square {
	for _, s := range b.Squares {
		if s.File == file && s.Rank == rank {
			return s
		}
	}

	return nil
}
