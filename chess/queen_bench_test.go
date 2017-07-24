package chess

import (
	"testing"
)

func BenchmarkQueen_ValidMoves(b *testing.B) {
	board, _ := BoardFEN("4k3/7n/8/8/8/3Q4/8/4K3 w - - 0 1")
	queen := board.GetSquare(D,3).CurrentPiece

	for i:=0; i<b.N; i++{
		queen.ValidMoves(board)
	}
}