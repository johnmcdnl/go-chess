package chess

import (
	"testing"
)

func BenchmarkRook_ValidMoves(b *testing.B) {
	board, _ := BoardFEN("3nk3/8/8/3R4/8/8/8/4K3 w - - 0 1")
	rook := board.GetSquare(D,5).CurrentPiece

	for i:=0; i<b.N; i++{
		rook.ValidMoves(board)
	}
}