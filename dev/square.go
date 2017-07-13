package dev

import (
	"fmt"
	"encoding/json"
)

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
	//TODO just use a nil check on piece vs isOccupied bool - retain method for convenience
	isOccupied bool
	piece      Piece
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

func (s *Square)GetPiece() Piece {
	return s.piece
}

//TODO test me - why did I do this before making pieces :-|
func (s *Square)NewPiece(p Piece) error {

	//TODO consider deleting this and let 'move.go' handle all moving logic
	if s.isOccupied && s.piece.GetColor() == p.GetColor() {
		js, err := json.Marshal(s)
		fmt.Println(err, string(js))
		jp, err := json.Marshal(p)
		fmt.Println(err, string(jp))
		return fmt.Errorf("Cannot capture your own piece: %s  %s:", string(js), string(jp))
	}

	if s.isOccupied && s.piece.GetColor() != p.GetColor() && s.piece.PieceType() == KingPiece {
		js, err := json.Marshal(s)
		fmt.Println(err, string(js))
		jp, err := json.Marshal(p)
		fmt.Println(err, string(jp))
		return fmt.Errorf("Cannot capture the King: %s  %s:", string(js), string(jp))
	}

	s.piece = p
	s.isOccupied = true
	return nil
}
