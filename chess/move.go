package chess

import "fmt"

type Move struct {
	Origin      *Square
	Destination *Square
}

//New returns a Move or and error if the provided squares are invalid
func NewMove(origin, destination *Square) *Move {

	if origin == nil || destination == nil {
		return nil
	}
	var m Move
	m.Origin = origin
	m.Destination = destination

	if !m.IsValid(){
		return nil
	}

	return &m
}

func (m *Move) Printer() string {
	return fmt.Sprint("Origin: ", m.Origin.PrettyPrint(), " Destination: ", m.Destination.PrettyPrint())
}

func (m *Move) IsValid() bool {

	if m.Destination.CurrentPiece==nil{
		return true
	}

	if m.Origin.CurrentPiece.PieceColor() == m.Destination.CurrentPiece.PieceColor() {
		fmt.Println("func (m *Move) IsValid() bool  this is your own piece", m.Printer())
		return false
	}

	if m.Destination.CurrentPiece.PieceType() == KingType {
		fmt.Println("func (m *Move) IsValid() bool  Cannot Capture the King", m.Printer())
		return false
	}

	if m.Origin.CurrentPiece.PieceColor() != m.Destination.CurrentPiece.PieceColor() {
		fmt.Println("func (m *Move) IsValid() bool  Capturing MOVES!!!", m.Printer())
		return true
	}

	return true
}

func (m *Move) Apply() {

}
