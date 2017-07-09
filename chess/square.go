package chess

import "fmt"

type Square struct {
	Name       string
	*Coordinate
	ChessPiece ChessPiece
}

func NewSquare(file int, rank int) *Square {
	var s Square
	s.Coordinate = NewCoordinate(file, rank)
	s.Name = fmt.Sprint(string('a' - 1 + file), rank)

	return &s
}
