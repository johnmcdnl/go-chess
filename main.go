package main

import (
	"fmt"

	"github.com/johnmcdnl/go-chess/chess"
)

func main() {
	fmt.Println("Welcome to go-chess")
	b := chess.NewBoard()
	chess.GlobalBoard = b

	for _, square := range b.Squares {
		if square.CurrentPiece != nil {
			square.CurrentPiece.ValidMoves()
		}
	}
}
