package chess

import (
	"testing"
)

func BenchmarkKing_ValidMoves(b *testing.B) {
	board, _ := BoardFEN("4k3/8/8/4p3/4K3/8/8/8 w - - 0 1")
	king := board.GetSquare(E, 4).CurrentPiece

	for i := 0; i < b.N; i++ {
		king.ValidMoves(board)
	}
}