package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestPawn_ValidMoves(t *testing.T) {
	boardFEN := func(fen FEN) *Board {
		b, _ := BoardFEN(fen)
		return b
	}

	type fields struct {
		board      *Board
		Coordinate Coordinate
	}

	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			"A White Pawn may advance 2 square on it's first move",
			fields{boardFEN(StartingPositionFEN), Coordinate{A, 2}},
			[]string{"a2a3", "a2a4"},
		},
		{
			"A Black Pawn may advance 2 square on it's first move",
			fields{boardFEN(StartingPositionFEN), Coordinate{A, 7}},
			[]string{"a7a6", "a7a5"},
		},

		{
			"A White Pawn may only advance 1 square at a time",
			fields{boardFEN("k7/8/3p4/8/8/3P4/8/K7 w - - 0 1"), Coordinate{D, 3}},
			[]string{"d3d4"},
		},
		{
			"A Black Pawn may only advance 1 square at a time",
			fields{boardFEN("k7/8/3p4/8/8/3P4/8/K7 w - - 0 1"), Coordinate{D, 6}},
			[]string{"d6d5"},
		},
		{
			"A white pawn may cannot capture a piece directly in front of itself",
			fields{boardFEN("k7/8/8/3p4/3P4/8/8/K7 w - - 0 1"), Coordinate{D, 4}},
			[]string{},
		},
		{
			"A black pawn may cannot capture a piece directly in front of itself",
			fields{boardFEN("k7/8/8/3p4/3P4/8/8/K7 w - - 0 1"), Coordinate{D, 5}},
			[]string{},
		},
		{
			"A white pawn may capture a piece diagonally one square in front itself",
			fields{boardFEN("k7/8/8/2npn3/2NPN3/8/8/K7 w - - 0 1"), Coordinate{D, 4}},
			[]string{"d4xc5", "d4xe5"},
		},
		{
			"A black pawn may capture a piece diagonally one square in front itself",
			fields{boardFEN("k7/8/8/2npn3/2NPN3/8/8/K7 w - - 0 1"), Coordinate{D, 5}},
			[]string{"d5xc4", "d5xe4"},
		},
		{
			"A white pawn may not capture a piece that is more than diagonal square away",
			fields{boardFEN("k7/8/1n3n2/3p4/3P4/1N3N2/8/K7 w - - 0 1"), Coordinate{D, 4}},
			[]string{},
		},
		{
			"A black pawn may not capture a piece that is more than diagonal square away",
			fields{boardFEN("k7/8/1n3n2/3p4/3P4/1N3N2/8/K7 w - - 0 1"), Coordinate{D, 5}},
			[]string{},
		},
		{
			"A white pawn may not move backwards",
			fields{boardFEN("k7/8/8/3p4/3P4/8/8/K7 w - - 0 1"), Coordinate{D, 4}},
			[]string{},
		},
		{
			"A black pawn may not move backwards",
			fields{boardFEN("k7/8/8/3p4/3P4/8/8/K7 w - - 0 1"), Coordinate{D, 5}},
			[]string{},
		},
		{
			"A white pawn may not jump over it's own piece",
			fields{boardFEN("k7/8/8/2RRR3/2RPR3/2RRR3/8/K7 w - - 0 1"), Coordinate{D, 4}},
			[]string{},
		},
		{
			"A black pawn may not jump over it's own piece",
			fields{boardFEN("k7/8/2rrr3/2rpr3/2rrr3/8/8/K7 w - - 0 1"), Coordinate{D, 5}},
			[]string{},
		},

		{
			"A white pawn may not capture a oponents king",
			fields{boardFEN("8/8/6p1/5KRR/1rrk4/2P5/8/8 w - - 0 1"), Coordinate{C, 3}},
			[]string{"c3xb4"},
		},
		{
			"A black pawn may not capture a oponents king",
			fields{boardFEN("8/8/6p1/5KRR/1rrk4/2P5/8/8 w - - 0 1"), Coordinate{G, 6}},
			[]string{"g6xh5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pawn := tt.fields.board.GetSquare(tt.fields.Coordinate.File, tt.fields.Coordinate.Rank).CurrentPiece
			aMoves := pawn.ValidMoves(tt.fields.board)
			assert.Equal(t, PawnType, pawn.PieceType(), "Wrong piece type")


			var aMovePGNS []string
			for _, m := range aMoves {
				aMovePGNS = append(aMovePGNS, m.PGNName())
			}

			assert.Equal(t, len(tt.want), len(aMoves), "Incorrect Number of Moves", fmt.Sprintln(tt.want), fmt.Sprintln(aMovePGNS))



			for _, aMove := range aMoves {
				assert.Contains(t, tt.want, aMove.PGNName(), "Expected Slice doesn't contain actual")
			}

			for _, eMove := range tt.want {
				assert.Contains(t, aMovePGNS, eMove, "Actual Slice doesn't contain expected")
			}
		})
	}
}
