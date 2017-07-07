package chess

import "testing"

func BenchmarkBishop_ValidMoves(b *testing.B) {
	e5 := GlobalBoard.GetSquare(5, 5)
	e5.CurrentPiece = NewBishop(e5, White)

	for i:=0; i<b.N;i++{
		e5.CurrentPiece.ValidMoves()
	}
}