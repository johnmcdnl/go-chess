package dev

import (
	"fmt"
	"encoding/json"
	"errors"
)

type Move struct {
	origin      *Square
	destination *Square
}

func NewMove(origin, destination *Square) (*Move, error) {
	var m Move
	m.origin = origin
	m.destination = destination
	return &m, nil

}

func (m *Move)Name() string {
	return fmt.Sprintln(m.origin.piece.GetName(), m.origin.Name(), m.destination.Name())
}

func (m *Move)Apply() error {
	if isValid, err := m.IsValid(); !isValid {
		return err
	}
	if err := m.destination.NewPiece(m.origin.piece); err != nil {
		return err
	}
	m.origin.piece = nil
	m.origin.isOccupied = false
	return nil
}

func (m *Move)IsValid() (bool, error) {

	d := m.destination
	o := m.origin

	if d == nil {
		return false, errors.New("nil destination")
	}

	if o == nil {
		return false, errors.New("nil origin")
	}

	if !d.IsOccupied() {
		return true, nil
	}

	if d.IsOccupied() && o.GetPiece().GetColor() == d.GetPiece().GetColor() {
		js, err := json.Marshal(o.GetPiece())
		fmt.Println(err, string(js))
		jp, err := json.Marshal(d.GetPiece())
		fmt.Println(err, string(jp))
		return false, errors.New((fmt.Sprintln("Can't capture your own piece: ", string(js), string(jp))))
	}

	if d.IsOccupied() && o.GetPiece().GetColor() != d.GetPiece().GetColor() && d.GetPiece().PieceType() == KingPiece {
		js, err := json.Marshal(o.GetPiece())
		fmt.Println(err, string(js))
		jp, err := json.Marshal(d.GetPiece())
		fmt.Println(err, string(jp))
		return false, errors.New((fmt.Sprintln("Can't capture the king: ", string(js), string(jp))))
	}

	return true, nil
}