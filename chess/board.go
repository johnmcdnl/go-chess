package chess

import "fmt"

type Board struct {
	Squares []*Square
}

type Square struct {
	ID           int
	RankNumber   int
	FileLetter   int
	Name         string
	CurrentPiece int
}

func NewBoard() *Board {
	var b Board
	b.setupSquares()
	b.setInitialBoard()
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
			s.Name = fmt.Sprint(fileNumToLetter(s.RankNumber), s.FileLetter)
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
		return "A"
	case 2:
		return "B"
	case 3:
		return "C"
	case 4:
		return "D"
	case 5:
		return "E"
	case 6:
		return "F"
	case 7:
		return "G"
	case 8:
		return "H"
	}
}

func (b *Board) setInitialBoard() {
	for _, square := range b.Squares {
		if square.RankNumber == 2 {
			square.CurrentPiece = WhitePawn
		}
		if square.RankNumber == 7 {
			square.CurrentPiece = BlackPawn
		}

		if square.RankNumber == 1 {
			if square.FileLetter == 1 || square.FileLetter == 8 {
				square.CurrentPiece = WhiteRook
			}
			if square.FileLetter == 2 || square.FileLetter == 7 {
				square.CurrentPiece = WhiteKnight
			}
			if square.FileLetter == 3 || square.FileLetter == 6 {
				square.CurrentPiece = WhiteBishop
			}
			if square.FileLetter == 4 {
				square.CurrentPiece = WhiteQueen
			}
			if square.FileLetter == 5 {
				square.CurrentPiece = WhiteKing
			}
		}

		if square.RankNumber == 8 {
			if square.FileLetter == 1 || square.FileLetter == 8 {
				square.CurrentPiece = BlackRook
			}
			if square.FileLetter == 2 || square.FileLetter == 7 {
				square.CurrentPiece = BlackKnight
			}
			if square.FileLetter == 3 || square.FileLetter == 6 {
				square.CurrentPiece = BlackBishop
			}
			if square.FileLetter == 4 {
				square.CurrentPiece = BlackQueen
			}
			if square.FileLetter == 5 {
				square.CurrentPiece = BlackKing
			}
		}
	}
}
