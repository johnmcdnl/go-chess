package main

import (
	"fmt"

	"github.com/johnmcdnl/go-chess/chess"
	"encoding/json"
)

func main() {
	fmt.Println("Welcome to go-chess")

	g, err := chess.NewGame()
	//pos, _ := g.Board.GetSANPosition("A1")
	j, _ := json.Marshal(g)
	fmt.Println(string(j), err)
	fmt.Println(len(g.Board.SANPositions))
}
