package main

import (
	"fmt"
	"github.com/johnmcdnl/go-chess/chess"
)

func main() {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	b, err := chess.NewBoard()
	fmt.Println(err)

	for i, s := range b.Squares {
		fmt.Println(i, s.File, s.Rank, s.CurrentPiece)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	for _, s := range b.Squares {
		if s.CurrentPiece != nil && s.CurrentPiece.PieceType() == chess.RookType {
			fmt.Println()
			for _, m := range s.CurrentPiece.ValidMoves(b) {
				fmt.Println(m.Printer())
			}
			fmt.Println()
		}
	}

}
