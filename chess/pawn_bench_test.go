package chess

import (
	"testing"
)

func BenchmarkPawn_ValidMoves(b *testing.B) {
	board, _ := BoardFEN("4k3/8/8/3p4/4P3/8/8/4K3 w - - 0 1")
	pawn := board.GetSquare(E,4).CurrentPiece

	for i:=0; i<b.N; i++{
		pawn.ValidMoves(board)
	}
}