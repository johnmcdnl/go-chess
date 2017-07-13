package dev

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

type Square struct {
	id         int
	File       int
	Rank       int
	name       string
	isOccupied bool
}

func NewSquare(file, rank int) (*Square, error) {

	if file <= 0 || file > 8 {
		return nil, fmt.Errorf("Invalid File %d", file)
	}
	if rank <= 0 || rank > 8 {
		return nil, fmt.Errorf("Invalid Rank %d", rank)
	}

	var s Square
	s.File = file
	s.Rank = rank
	s.name = fmt.Sprint(string('a' - 1 + s.File), s.Rank)
	s.id = file * 8 + rank - 8 - 1
	return &s, nil
}

func (s *Square)GetID() int {
	return s.id
}

func (s *Square)Name() string {
	return s.name
}

func (s *Square)IsOccupied() bool {
	return s.isOccupied
}

