package chess

import "fmt"

type Moves struct {
	PossibleMoves []*Square
	BestMoves     []*Move
}

type Move struct {
	From *Square
	To   *Square
}

func (m *Move)Name() string {
	return fmt.Sprint(m.From.Name, m.To.Name)
}

func (m *Move)AddIfValid() (canMoveThisDirection bool) {
	m.To.CurrentPiece.GetPiece().Color = m.From.CurrentPiece.GetPiece().Color


	return false
}

