package chess

import (
	"fmt"
)

var GlobalBoard *Board

type Board struct {
	Squares []*Square
}

type Square struct {
	ID           int
	RankNumber   int
	FileLetter   int
	Name         string
	CurrentPiece ChessPiece
}

func NewBoard() *Board {
	var b Board
	b.setupSquares()
	b.setInitialBoard()
	GlobalBoard = &b
	return &b
}

func (b *Board) setupSquares() {
	var squares []*Square

	for rank := 1; rank <= 8; rank++ {
		for file := 1; file <= 8; file++ {
			var s Square
			s.FileLetter = file
			s.RankNumber = rank
			s.ID = (((rank * 8) + file ) - 8)
			s.Name = fmt.Sprint(fileNumToLetter(s.FileLetter), s.RankNumber)
			squares = append(squares, &s)
		}
	}

	b.Squares = squares
}

func fileNumToLetter(i int) string {

	switch i {
	default:
		return ""
	case 1:
		return "a"
	case 2:
		return "b"
	case 3:
		return "c"
	case 4:
		return "d"
	case 5:
		return "e"
	case 6:
		return "f"
	case 7:
		return "g"
	case 8:
		return "h"
	}
}

func (b *Board) setInitialBoard() {
	for _, square := range b.Squares {
		if square.RankNumber == 2 {
			square.CurrentPiece = NewPawn(square, White)
		}
		if square.RankNumber == 7 {
			square.CurrentPiece = NewPawn(square, Black)
		}

		if square.RankNumber == 1 {
			if square.FileLetter == 1 || square.FileLetter == 8 {
				square.CurrentPiece = NewRook(square, White)
			}
			//if square.FileLetter == 1 || square.FileLetter == 7 {
			//	square.CurrentPiece = GetPiece(WhiteKnight)
			//}
			if square.FileLetter == 3    ||square.FileLetter == 6 {
				square.CurrentPiece = NewBishop(square, White)
			}
			//if square.FileLetter == 1 {
			//	square.CurrentPiece = GetPiece(WhiteQueen)
			//}
			//if square.FileLetter == 1 {
			//	square.CurrentPiece = GetPiece(WhiteKing)
			//}
		}

		if square.RankNumber == 8 {
			if square.FileLetter == 1 || square.FileLetter == 8 {
				square.CurrentPiece = NewRook(square, Black)
			}
			//if square.FileLetter == 2 || square.FileLetter == 7 {
			//	square.CurrentPiece = GetPiece(BlackKnight)
			//}
			if square.FileLetter == 3    ||square.FileLetter == 6 {
				square.CurrentPiece = NewBishop(square, Black)
			}
			//if square.FileLetter == 4 {
			//	square.CurrentPiece = GetPiece(BlackQueen)
			//}
			//if square.FileLetter == 5 {
			//	square.CurrentPiece = GetPiece(BlackKing)
			//}
		}
	}
}

func (b *Board)GetSquare(file int, rank int) *Square {
	for _, s := range b.Squares {
		if s.FileLetter == file && s.RankNumber == rank {
			return s
		}
	}
	return nil
}