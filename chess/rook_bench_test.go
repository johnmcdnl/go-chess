package chess

import "testing"

func BenchmarkRook_ValidMoves(b *testing.B) {
	e5 := GlobalBoard.GetSquare(5, 5)
	e5.CurrentPiece = NewRook(e5, White)

	for i := 0; i < b.N; i++ {
		e5.CurrentPiece.ValidMoves()
	}
}

