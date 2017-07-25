package chess

import "fmt"

const (
	_ = iota
	A
	B
	C
	D
	E
	F
	G
	H
)

type Coordinate struct {
	File int
	Rank int
}

type Square struct {
	ID           int
	Coordinate
	CurrentPiece Piece
}

func NewSquare(file, rank int) (*Square, error) {
	var s Square
	s.File = file
	s.Rank = rank
	s.ID = GetSquareID(file, rank)
	return &s, nil
}

func GetSquareID(file, rank int) int {
	if file < 1 || file > 8 || rank < 1 || rank > 8 {
		return -1
	}

	id := ((rank - 1) * 8) + file - 1
	if id < 0 || id > 63 {
		return -1
	}
	return id
}

func (s *Square) SetPiece(p Piece) {
	s.CurrentPiece = p
}

func (s *Square)PrettyPrint() string {
	var str string

	switch s.File {
	case 1: str += "a"
	case 2:str += "b"
	case 3:str += "c"
	case 4:str += "d"
	case 5:str += "e"
	case 6:str += "f"
	case 7:str += "g"
	case 8:str += "h"
	}

	str += fmt.Sprint(s.Rank)

	return str
}

func (s *Square)Name() string {
	var str string

	switch s.File {
	case 1: str += "a"
	case 2:str += "b"
	case 3:str += "c"
	case 4:str += "d"
	case 5:str += "e"
	case 6:str += "f"
	case 7:str += "g"
	case 8:str += "h"
	}

	str += fmt.Sprint(s.Rank)

	return str
}