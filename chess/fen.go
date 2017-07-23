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

	reverse := func(ss []string) {
		last := len(ss) - 1
		for i := 0; i < len(ss) / 2; i++ {
			ss[i], ss[last - i] = ss[last - i], ss[i]
		}
	}
	reverse(fenPiecePlacements)

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

	boardInfoArray := strings.Split(boardInfo, " ")

	switch boardInfoArray[0] {
	default:
		panic("unknown color")
	case "w":
		b.ActiveColor = White
	case "b":
		b.ActiveColor = Black
	}

	//Castling rights
	var castlingRights = CastlingRights{
		WhiteKingSideAvailable:false,
		WhiteQueenSideAvailable : false,
		BlackKingSideAvailable : false,
		BlackQueenSideAvailable : false,
	}
	for _, cr := range boardInfoArray[1] {
		switch string(cr) {
		default:
			panic("unknown castling rights")
		case "-":
			castlingRights.WhiteKingSideAvailable = false
			castlingRights.WhiteQueenSideAvailable = false
			castlingRights.BlackKingSideAvailable = false
			castlingRights.BlackQueenSideAvailable = false
		case "K":
			castlingRights.WhiteKingSideAvailable = true
		case "Q":
			castlingRights.WhiteQueenSideAvailable = true
		case "k":
			castlingRights.BlackKingSideAvailable = true
		case "q":
			castlingRights.BlackQueenSideAvailable = true
		}
	}
	b.CastlingRights = &castlingRights


	//en passant
	if boardInfoArray[2] == "-" {
		b.EnPassantSquare = nil
	} else {
		enPassantFileRankArr := strings.Split(boardInfoArray[2], "")
		fmt.Println(enPassantFileRankArr)

		enPassantRank, _ := strconv.Atoi(enPassantFileRankArr[1])

		switch strings.ToLower(enPassantFileRankArr[0]) {
		default:
			panic("unknown")
		case "a":
			b.EnPassantSquare = b.GetSquare(A, enPassantRank)
		case "b":
			b.EnPassantSquare = b.GetSquare(B, enPassantRank)
		case "c":
			b.EnPassantSquare = b.GetSquare(C, enPassantRank)
		case "d":
			b.EnPassantSquare = b.GetSquare(D, enPassantRank)
		case "e":
			b.EnPassantSquare = b.GetSquare(E, enPassantRank)
		case "f":
			b.EnPassantSquare = b.GetSquare(F, enPassantRank)
		case "g":
			b.EnPassantSquare = b.GetSquare(G, enPassantRank)
		case "h":
			b.EnPassantSquare = b.GetSquare(H, enPassantRank)
		}

	}

	b.HalfMoveClock, _ = strconv.Atoi(boardInfoArray[3])
	b.FullMoveNumber, _ = strconv.Atoi(boardInfoArray[4])
}