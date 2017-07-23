package main

import (
	"github.com/johnmcdnl/go-chess/chess"
	"fmt"
)

func main() {
	board, _ := chess.NewEmptyBoard()
	chess.NewFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b - - 0 1").Apply(board)

	fmt.Println("board.ActiveColor", board.ActiveColor)
	fmt.Println("board.CastlingRights", board.CastlingRights)
	//fmt.Println(board.EnPassantSquare)
	//fmt.Println(board.HalfMoveClock)
	//fmt.Println(board.FullMoveNumber)

}
