package chess

import (
	"testing"
)

func BenchmarkBishop_ValidMoves(b *testing.B) {
	board, _ := BoardFEN("q3k3/8/8/3B4/8/8/8/4K3 w - - 0 1")
	bishop := board.GetSquare(D,5).CurrentPiece

	for i:=0; i<b.N; i++{
		bishop.ValidMoves(board)
	}
}