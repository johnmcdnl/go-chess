package chess

import (
	"testing"
)

func BenchmarkKnight_ValidMovesNearCenter(b *testing.B) {
	board, _ := BoardFEN("4k3/8/1q6/3N4/8/8/8/4K3 w - - 0 1")
	knight := board.GetSquare(D, 5).CurrentPiece

	for i := 0; i < b.N; i++ {
		knight.ValidMoves(board)
	}
}

func BenchmarkKnight_ValidMovesFirstColumn(b *testing.B) {
	board, _ := BoardFEN("4k3/8/8/8/N7/8/8/4K3 w - - 0 1")
	knight := board.GetSquare(A, 4).CurrentPiece

	for i := 0; i < b.N; i++ {
		knight.ValidMoves(board)
	}
}

func BenchmarkKnight_ValidMovesSecondColumn(b *testing.B) {
	board, _ := BoardFEN("4k3/8/8/8/1N6/8/8/4K3 w - - 0 1")
	knight := board.GetSquare(B, 4).CurrentPiece

	for i := 0; i < b.N; i++ {
		knight.ValidMoves(board)
	}
}