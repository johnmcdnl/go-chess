package main

import (
	"github.com/johnmcdnl/go-chess/chess"
	"fmt"
)

func main() {
	game := chess.NewGame()
	fmt.Println(game)
	board := game.Board

	fmt.Println("board.ActiveColor", board.ActiveColor)
	fmt.Println("board.CastlingRights", board.CastlingRights)
	fmt.Println(board.EnPassantSquare)
	fmt.Println(board.HalfMoveClock)
	fmt.Println(board.FullMoveNumber)

	var turn = 1
	for i := 0; i <= 50; i++ {
		move := game.Players[i % 2].PickMove(game.Board)
		fmt.Println(turn, " ", move.PGNName())
		move.Apply(game.Board)

		if i % 2 == 0 {
			turn++
			board.Print()
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
	}

	board.Print()

}
