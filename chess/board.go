package chess

import (
	"fmt"
)

type Board struct {
	Squares []*Square
}

func NewEmptyBoard() (*Board, error) {
	var b Board

	if err := b.build(8, 8); err != nil {
		return nil, err
	}

	return &b, nil
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
		if (s.File == A && s.Rank == 1) || (s.File == H && s.Rank == 1) {
			r, _ := NewRook(s, White)
			s.SetPiece(r)
		}

		//White
		if (s.File == B && s.Rank == 1) || (s.File == G && s.Rank == 1) {
			k, _ := NewKnight(s, White)
			s.SetPiece(k)
		}

		//White
		if s.File == E && s.Rank == 1  {
			k, _ := NewKing(s, White)
			s.SetPiece(k)
		}

		//DUMMY
		if s.File == A && s.Rank == 3  {
			k, _ := NewRook(s, Black)
			s.SetPiece(k)
		}



		//Black
		if (s.File == A && s.Rank == 8) || (s.File == H && s.Rank == 8) {
			r, _ := NewRook(s, Black)
			s.SetPiece(r)
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
