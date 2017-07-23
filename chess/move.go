package chess

import (
	"fmt"
	"strings"
)

type Move struct {
	Origin          *Square
	Destination     *Square
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

	if m.Origin.CurrentPiece == nil {
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
		pieceNameCode = "K"
	case QueenType:
		pieceNameCode = "Q"
	case RookType:
		pieceNameCode = "R"
	case BishopType:
		pieceNameCode = "B"
	case KnightType:
		pieceNameCode = "N"
	case PawnType:
		pieceNameCode = "I AM A PAWN"
	}
	//Qd8xd5

	var isCapture = ""
	if m.IsCapturingMove() {
		isCapture = "x"
	}
	return fmt.Sprintf("%s%s%s%s", pieceNameCode, strings.ToLower(m.Origin.Name()), isCapture, strings.ToLower(m.Destination.Name()))
}

func (m *Move)IsCapturingMove()bool{
	fmt.Println(m.Origin.Name(), m.Destination.Name())
	if m.Destination.CurrentPiece==nil{
		fmt.Println("Nothing to capture")
		return false
	}
	if m.Destination.CurrentPiece.PieceColor() == m.Origin.CurrentPiece.PieceColor(){
		fmt.Println("can't capture own color")
		return false
	}

	if m.Destination.CurrentPiece.PieceType()==KingType{
		fmt.Println("can't capture king")
		return false
	}

	if m.Destination.CurrentPiece!=nil{
		fmt.Println("this is a capture move")
		return true
	}
	panic("I missed something obvious")
	return false
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
