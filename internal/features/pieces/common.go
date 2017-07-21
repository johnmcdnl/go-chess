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
		b = nil
		board, err := chess.NewEmptyBoard()
		assert.NoError(T, err)
		b = board
	})

	When(`^a "(.+?)" "(.+?)" is placed on the square "(.+?)"$`, func(color, piece, square string) {

		var c chess.Color
		var s *chess.Square

		file, rank := SquareNameToCoordinates(square)
		s = Board().GetSquare(file, rank)

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
		case "queen":
			p, _ = chess.NewQueen(s, c)
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

	Then(`^the piece has moves available to "(.+?)"$`, func(valid string) {
		expectedMoves := strings.Split(valid, " ")
		fmt.Println(expectedMoves)
		actualMoves := currentPiece.ValidMoves(Board())

		assert.Equal(T, len(expectedMoves), len(actualMoves), "Invalid number of moves")

		for _, actual := range actualMoves {
			af := actual.Destination.File
			ar := actual.Destination.Rank
			var foundActualMove bool
			for _, expected := range expectedMoves {
				ef, er := SquareNameToCoordinates(expected)
				if ef == af && er == ar {
					foundActualMove = true
					break
				}
			}

			assert.True(T, foundActualMove, "Unexpected Move: ", actual.Printer())
		}

	})
}

func SquareNameToCoordinates(s string) (file, rank int) {
	codes := strings.Split(s, "")

	switch strings.ToUpper(codes[0]){
	case "A":
		file = 1
	case "B":
		file = 2
	case "C":
		file = 3
	case "D":
		file = 4
	case "E":
		file = 5
	case "F":
		file = 6
	case "G":
		file = 7
	case "H":
		file = 8
	}

	rank, _ = strconv.Atoi(codes[1])
	return file, rank
}