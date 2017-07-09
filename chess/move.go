package chess

import "fmt"

type Move struct {
	origin      *Square
	destination *Square
	name        string
}

func NewMove(origin, destination *Square) (*Move, error) {

	if origin == nil {
		return nil, fmt.Errorf("Invalid origin %s", fmt.Sprint(origin))
	}

	if destination == nil {
		return nil, fmt.Errorf("Invalid destination %s", fmt.Sprint(destination))
	}

	if origin.ChessPiece == nil {
		return nil, fmt.Errorf("Invalid ChessPiece %s", fmt.Sprint(origin.ChessPiece))
	}

	var m Move
	m.origin = origin
	m.destination = destination

	if err := m.isValid(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	m.generateName()

	return &m, nil
}

func (m *Move)isValid() (error) {

	if m.destination.ChessPiece == nil {
		return nil
	}

	if m.origin.ChessPiece != nil &&  m.origin.ChessPiece.GetColor() == m.destination.ChessPiece.GetColor() {
		return fmt.Errorf("Cannot capture your own color piece %s", fmt.Sprint(m.origin.ChessPiece, m.destination.ChessPiece.CurrentPosition().Name))
	}

	return nil
}

func (m *Move) generateName() {

	if m.origin == nil || m.origin.ChessPiece == nil {
		return
	}

	var pieceCode string
	switch m.origin.ChessPiece.GetCode() {
	default:
		return
	case PawnPiece:
		pieceCode = ""
	case KnightPiece:
		pieceCode = "N"
	case BishopPiece:
		pieceCode = "B"
	case RookPiece:
		pieceCode = "R"
	case QueenPiece:
		pieceCode = "Q"
	case KingPiece:
		pieceCode = "K"

	}
	m.name = fmt.Sprint(pieceCode, m.origin.Name, m.destination.Name)
}

func (m *Move) Name() string {
	return m.name
}
