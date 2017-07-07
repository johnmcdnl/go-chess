package chess

import (
	"testing"
)

func BenchmarkKnight_ValidMoves(b *testing.B) {
	e5 := GlobalBoard.GetSquare(5, 5)
	e5.CurrentPiece = NewKnight(e5, White)

	for i:=0; i<b.N;i++{
		e5.CurrentPiece.ValidMoves()
	}
}