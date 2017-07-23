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

	if !m.IsValid() {
		return nil
	}

	return &m
}

func (m *Move) Printer() string {
	return fmt.Sprint("Origin: ", m.Origin.PrettyPrint(), " Destination: ", m.Destination.PrettyPrint())
}

func (m *Move)PGNName() string {

	if m.Origin.CurrentPiece==nil{
		return "I AM NIL"
	}

	var pieceNameCode string
	switch m.Origin.CurrentPiece.PieceType() {
	default:
		fmt.Println(m.Origin.CurrentPiece.PieceType())
		fmt.Println(m.Origin.CurrentPiece.CurrentPosition())
		fmt.Println(m.Origin.CurrentPiece.PieceColor())
		return fmt.Sprintln("I AM INVALID")
	case KingType:
		pieceNameCode="K"
	case QueenType:
		pieceNameCode="Q"
	case RookType:
		pieceNameCode="R"
	case BishopType:
		pieceNameCode="B"
	case KnightType:
		pieceNameCode="N"
	case PawnType:
		pieceNameCode="I AM A PAWN"
	}
	return fmt.Sprint(pieceNameCode)
}

func (m *Move) IsValid() bool {

	if m.Destination.CurrentPiece == nil {
		return true
	}

	if m.Origin.CurrentPiece.PieceColor() == m.Destination.CurrentPiece.PieceColor() {
		return false
	}

	if m.Destination.CurrentPiece.PieceType() == KingType {
		return false
	}

	if m.Origin.CurrentPiece.PieceColor() != m.Destination.CurrentPiece.PieceColor() {
		return true
	}

	return true
}

func (m *Move) Apply() {

}
