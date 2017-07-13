package chess

import (
	"fmt"
)

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
	return &b
}

func NewBoardWithFen(fen FEN) *Board {
	b := NewBoard()
	b.SetFEN(fen)
	return b
}

func (b *Board) setStartingPosition() {
	for _, s := range *b.Squares {
		if s.Rank == 2 {
			cp := NewPawn(s, White)
			s.ChessPiece = cp
		}
		if s.Rank == 7 {
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
	if fen == StartingPositionFEN {
		b.setStartingPosition()
		return
	}
	if fen == OnlyKingsFEN {
		return
	}
	if fen == BlockedE5FEN {
		for _, s := range *b.Squares {
			if s.Name == "e5" {
				cp := NewPawn(s, Black)
				s.ChessPiece = cp
			} else {
				s.ChessPiece = nil
			}
		}
		return
	}

	panic("TODO NOT IMPLEMENTED SET FEN")

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

func (b *Board)Print() {

	PrintX(b.GetSquare(A, 8))
	PrintX(b.GetSquare(B, 8))
	PrintX(b.GetSquare(C, 8))
	PrintX(b.GetSquare(D, 8))
	PrintX(b.GetSquare(E, 8))
	PrintX(b.GetSquare(F, 8))
	PrintX(b.GetSquare(G, 8))
	PrintX(b.GetSquare(H, 8))
	fmt.Println()
	PrintX(b.GetSquare(A, 7))
	PrintX(b.GetSquare(B, 7))
	PrintX(b.GetSquare(C, 7))
	PrintX(b.GetSquare(D, 7))
	PrintX(b.GetSquare(E, 7))
	PrintX(b.GetSquare(F, 7))
	PrintX(b.GetSquare(G, 7))
	PrintX(b.GetSquare(H, 7))
	fmt.Println()
	PrintX(b.GetSquare(A, 6))
	PrintX(b.GetSquare(B, 6))
	PrintX(b.GetSquare(C, 6))
	PrintX(b.GetSquare(D, 6))
	PrintX(b.GetSquare(E, 6))
	PrintX(b.GetSquare(F, 6))
	PrintX(b.GetSquare(G, 6))
	PrintX(b.GetSquare(H, 6))
	fmt.Println()
	PrintX(b.GetSquare(A, 5))
	PrintX(b.GetSquare(B, 5))
	PrintX(b.GetSquare(C, 5))
	PrintX(b.GetSquare(D, 5))
	PrintX(b.GetSquare(E, 5))
	PrintX(b.GetSquare(F, 5))
	PrintX(b.GetSquare(G, 5))
	PrintX(b.GetSquare(H, 5))
	fmt.Println()
	PrintX(b.GetSquare(A, 4))
	PrintX(b.GetSquare(B, 4))
	PrintX(b.GetSquare(C, 4))
	PrintX(b.GetSquare(D, 4))
	PrintX(b.GetSquare(E, 4))
	PrintX(b.GetSquare(F, 4))
	PrintX(b.GetSquare(G, 4))
	PrintX(b.GetSquare(H, 4))
	fmt.Println()
	PrintX(b.GetSquare(A, 3))
	PrintX(b.GetSquare(B, 3))
	PrintX(b.GetSquare(C, 3))
	PrintX(b.GetSquare(D, 3))
	PrintX(b.GetSquare(E, 3))
	PrintX(b.GetSquare(F, 3))
	PrintX(b.GetSquare(G, 3))
	PrintX(b.GetSquare(H, 3))
	fmt.Println()
	PrintX(b.GetSquare(A, 2))
	PrintX(b.GetSquare(B, 2))
	PrintX(b.GetSquare(C, 2))
	PrintX(b.GetSquare(D, 2))
	PrintX(b.GetSquare(E, 2))
	PrintX(b.GetSquare(F, 2))
	PrintX(b.GetSquare(G, 2))
	PrintX(b.GetSquare(H, 2))
	fmt.Println()
	PrintX(b.GetSquare(A, 1))
	PrintX(b.GetSquare(B, 1))
	PrintX(b.GetSquare(C, 1))
	PrintX(b.GetSquare(D, 1))
	PrintX(b.GetSquare(E, 1))
	PrintX(b.GetSquare(F, 1))
	PrintX(b.GetSquare(G, 1))
	PrintX(b.GetSquare(H, 1))
	fmt.Println()
	fmt.Println()
	fmt.Println()

}

func PrintX(s *Square) {
	if (s.ChessPiece == nil) {
		fmt.Print(s.Name, "(", "", ") ", "\t")
	} else {
		fmt.Print(s.Name, "(", s.ChessPiece.GetCode(), ") ", "\t")
	}
}
