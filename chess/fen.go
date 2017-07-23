package chess

import (
	"strings"
	"fmt"
	"strconv"
)

type FEN string

func NewFEN(f string) *FEN {
	fen := FEN(f)
	return &fen
}

func (f *FEN) IsValid() bool {
	return true
}

func (f *FEN)Apply(b *Board) {
	str := string(*f)
	arr := strings.Split(str, "/")

	var fenPiecePlacements []string
	var boardInfo string
	for i, p := range arr {
		if i == 7 {
			var rowInfo string
			splitStr := strings.Split(p, " ")
			rowInfo = splitStr[0]
			boardInfo = strings.Join(splitStr[1:], " ")
			fenPiecePlacements = append(fenPiecePlacements, rowInfo)
			continue
		}
		fenPiecePlacements = append(fenPiecePlacements, p)
	}


	//TODO board info stuff
	fmt.Println(boardInfo)

	reverse := func(ss []string) {
		last := len(ss) - 1
		for i := 0; i < len(ss) / 2; i++ {
			ss[i], ss[last - i] = ss[last - i], ss[i]
		}
	}
	reverse(fenPiecePlacements)

	for i, s := range b.Squares {
		fmt.Println(i, s)
	}

	boardIndex := 1
	for fenRank, fenPiecePlacement := range fenPiecePlacements {

		fenFile := strings.Split(fenPiecePlacement, "")

		for fenFileIndex, fenFileData := range fenFile {
			if empty, err := strconv.Atoi(fenFileData); err == nil {
				boardIndex += empty
				continue
			}
			var p Piece
			s := b.GetSquare(fenFileIndex + 1, fenRank + 1)
			switch fenFileData {
			case "K":
				p, _ = NewKing(s, White)
			case "Q":
				p, _ = NewQueen(s, White)
			case "R":
				p, _ = NewRook(s, White)
			case "B":
				p, _ = NewBishop(s, White)
			case "N":
				p, _ = NewKnight(s, White)
			case "P":
				p, _ = NewPawn(s, White)

			case "k":
				p, _ = NewKing(s, Black)
			case "q":
				p, _ = NewQueen(s, Black)
			case "r":
				p, _ = NewRook(s, Black)
			case "b":
				p, _ = NewBishop(s, Black)
			case "n":
				p, _ = NewKnight(s, Black)
			case "p":
				p, _ = NewPawn(s, Black)
			}
			s.SetPiece(p)
			boardIndex++
		}

	}

	for i, s := range b.Squares {
		fmt.Println(i, s)
	}

}