package chess

import "testing"

func BenchmarkKnight_ValidMoves(b *testing.B) {
	var board = NewBoard()
	s := NewSquare(4, 4)
	k := NewKnight(s, White)

	for i := 0; i <= b.N; i++ {
		k.ValidMoves(board)
	}

}