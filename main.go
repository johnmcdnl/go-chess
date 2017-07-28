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
	fmt.Println()
	fmt.Println()
	fmt.Println()
	var turn = 1
	var playerColor = chess.White
	for ply := 1; ply <= 10; ply++ {
		move := game.Players[int(playerColor)-1].PickMove(game.Board)
		fmt.Println(turn, " ", move.PGNName())
		move.Apply(game.Board)

		if playerColor==chess.White{
			playerColor=chess.Black
		}else{
			playerColor=chess.White
		}

		if ply % 2 == 0 {
			turn++
			//board.Print()
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
	}

	board.Print()

}
