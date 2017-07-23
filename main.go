package main

import (
	"github.com/johnmcdnl/go-chess/chess"
	"fmt"
)

func main() {
	board, _ := chess.NewEmptyBoard()
	chess.NewFEN("k1r4P/1r1r2P1/r3rP2/4Pr2/3P2r1/2P4r/1P2K1r1/P4r2 w - - 0 1").Apply(board)

	fmt.Println("board.ActiveColor", board.ActiveColor)
	fmt.Println("board.CastlingRights", board.CastlingRights)
	fmt.Println(board.EnPassantSquare)
	fmt.Println(board.HalfMoveClock)
	fmt.Println(board.FullMoveNumber)

	board.Print()
}
