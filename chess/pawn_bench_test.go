package chess

import "testing"

func BenchmarkPawn_ValidMoves_FirstMove(b *testing.B) {
	e2 := GlobalBoard.GetSquare(5, 2)
	e2.CurrentPiece = NewPawn(e2, White)

	for i:=0; i<b.N;i++{
		e2.CurrentPiece.ValidMoves()
	}
}

func BenchmarkPawn_ValidMoves_MidGame(b *testing.B) {
	e3 := GlobalBoard.GetSquare(5, 3)
	e3.CurrentPiece = NewPawn(e3, White)

	for i:=0; i<b.N;i++{
		e3.CurrentPiece.ValidMoves()
	}
}