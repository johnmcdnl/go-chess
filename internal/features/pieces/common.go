package pieces

import (
	. "github.com/gucumber/gucumber"
	"github.com/johnmcdnl/go-chess/chess"
	"github.com/stretchr/testify/assert"
	"strings"
	"strconv"
	"fmt"
)

var b *chess.Board
var currentPiece chess.Piece

func Board() *chess.Board {
	return b
}

func init() {
	Given(`^an empty board$`, func() {
		board, err := chess.NewEmptyBoard()
		assert.NoError(T, err)
		b = board
	})

	When(`^a "(.+?)" "(.+?)" is placed on the square "(.+?)"$`, func(color, piece, square string) {

		var c chess.Color
		var s *chess.Square

		codes := strings.Split(square, "")

		var fileNum int
		switch codes[0]{
		case "A":
			fileNum = 1
		case "B":
			fileNum = 2
		case "C":
			fileNum = 3
		case "D":
			fileNum = 4
		case "E":
			fileNum = 5
		case "F":
			fileNum = 6
		case "G":
			fileNum = 7
		case "H":
			fileNum = 8
		}

		rank, _ := strconv.Atoi(codes[1])

		s = Board().GetSquare(fileNum, rank)

		switch color{
		case "white":
			c = chess.White
		case "black":
			c = chess.Black
		}

		var p chess.Piece
		switch piece{
		case "rook":
			p, _ = chess.NewRook(s, c)
		case "knight":
			p, _ = chess.NewKnight(s, c)
		case "king":
			p, _ = chess.NewKing(s, c)

		}

		s.SetPiece(p)
		currentPiece = p
	})

	And(`^we print the board$`, func() {
		for i, s := range Board().Squares {
			fmt.Println(i, s.File, s.Rank, s.CurrentPiece)
		}
	})
}