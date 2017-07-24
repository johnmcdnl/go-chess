package chess

import (
	"testing"
)

func BenchmarkKnight_ValidMoves(b *testing.B) {
	board, _ := BoardFEN("4k3/8/1q6/3N4/8/8/8/4K3 w - - 0 1")
	knight := board.GetSquare(D,5).CurrentPiece

	for i:=0; i<b.N; i++{
		knight.ValidMoves(board)
	}
}