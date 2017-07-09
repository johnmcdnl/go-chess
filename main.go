package main

import (
	"github.com/johnmcdnl/go-chess/chess"
	"fmt"
)

func main() {

	b := chess.NewBoard()
	//for i, s := range *b.Squares {
	//	if s != nil && s.Piece != nil {
	//		//fmt.Println(i, s.Name, s.Piece.Name)
	//	}
	//}
	fmt.Println()
	fmt.Println()
	for i, s := range *b.Squares {
		if s.ChessPiece != nil {
			fmt.Println(i, s.Name, fmt.Sprint(s.ChessPiece))
			fmt.Println(s.ChessPiece.ValidMoves(b))
		}
	}


}
