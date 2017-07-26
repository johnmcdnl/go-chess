package chess

import (
	"github.com/satori/go.uuid"
	"time"
	"math/rand"
)

type Player struct {
	ID    string
	Name  string
	Color Color
}

func NewPlayer(name string, color Color) *Player {
	var p Player
	p.ID = uuid.NewV4().String()
	p.Name = name
	p.Color = color

	return &p
}

func (p *Player)PickMove(b *Board) *Move {
	var myPieces  []Piece
	for _, square := range b.Squares {
		if square.CurrentPiece != nil && square.CurrentPiece.PieceColor() == p.Color {
			myPieces = append(myPieces, square.CurrentPiece)
		}
	}

	var moves []*Move

	var capturingMoves []*Move

	for _, myPiece := range myPieces {

		newMoves := myPiece.ValidMoves(b)

		for _, m := range newMoves {
			if m.IsCapturingMove() {
				capturingMoves = append(capturingMoves, m)
			}
		}

		moves = append(moves, newMoves...)

	}



	//TODO returning a random move isn't wise :)
	if len(capturingMoves) > 0 {
		return capturingMoves[rand.Intn(len(capturingMoves))]
	}

	return moves[rand.Intn(len(moves))]
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
