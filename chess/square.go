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
	Coordinate
	CurrentPiece Piece
}

func NewSquare(file, rank int) (*Square, error) {
	var s Square
	s.File = file
	s.Rank = rank
	return &s, nil
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