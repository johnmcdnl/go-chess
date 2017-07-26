package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRook_ValidMoves(t *testing.T) {
	boardFEN := func(fen FEN) *Board {
		b, _ := BoardFEN(fen)
		return b
	}

	type fields struct {
		Coordinate Coordinate
	}
	type args struct {
		b *Board
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			"Center Empty Board",
			fields{Coordinate{D, 5}},
			args{b: boardFEN("7k/8/8/3R4/8/8/8/K7 w - - 0 1")},
			[]string{
				"Rd5d1", "Rd5d2", "Rd5d3", "Rd5d4", "Rd5d6", "Rd5d7", "Rd5d8",
				"Rd5a5", "Rd5b5", "Rd5c5", "Rd5e5", "Rd5f5", "Rd5g5", "Rd5h5",
			},
		},
		{
			"In a Corner",
			fields{Coordinate{H, 1}},
			args{b: boardFEN("3Kk3/8/8/8/8/8/8/7r w - - 0 1")},
			[]string{
				"Rh1a1", "Rh1b1", "Rh1c1", "Rh1d1", "Rh1e1", "Rh1f1", "Rh1g1",
				"Rh1h2", "Rh1h3", "Rh1h4", "Rh1h5", "Rh1h6", "Rh1h7", "Rh1h8",
			},
		},
		{
			"In a different Corner",
			fields{Coordinate{A, 8}},
			args{b: boardFEN("r7/3Kk3/8/8/8/8/8/8 w - - 0 1")},
			[]string{
				"Ra8a1", "Ra8a2", "Ra8a3", "Ra8a4", "Ra8a5", "Ra8a6", "Ra8a7",
				"Ra8b8", "Ra8c8", "Ra8d8", "Ra8e8", "Ra8f8", "Ra8g8", "Ra8h8",
			},
		},
		{
			"Almost In a Corner",
			fields{Coordinate{G, 2}},
			args{b: boardFEN("3Kk3/8/8/8/8/8/6r1/8 w - - 0 1")},
			[]string{
				"Rg2g1", "Rg2g3", "Rg2g4", "Rg2g5", "Rg2g6", "Rg2g7", "Rg2g8",
				"Rg2a2", "Rg2b2", "Rg2c2", "Rg2d2", "Rg2e2", "Rg2f2", "Rg2h2",
			},
		},
		{
			"Rook Cannot Jump",
			fields{Coordinate{D, 5}},
			args{b: boardFEN("3Kk3/8/2pnn3/2nrn3/2nnn3/8/8/8 w - - 0 1")},
			[]string{},
		},
		{
			"Rook Capturing when pinned in",
			fields{Coordinate{D, 5}},
			args{b: boardFEN("3Kk3/1P6/3pP3/2prp3/2P5/2PP1P2/8/8 w - - 0 1")},
			[]string{"Rd5d4", "Rd5xd3"},
		},
		{
			"Rook cannot capture a king",
			fields{Coordinate{D, 5}},
			args{b: boardFEN("8/1P6/2KpP3/2prk3/2P5/2PP1P2/8/8 w - - 0 1")},
			[]string{"Rd5d4", "Rd5xd3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rook := tt.args.b.GetSquare(tt.fields.Coordinate.File, tt.fields.Coordinate.Rank).CurrentPiece
			aMoves := rook.ValidMoves(tt.args.b)

			assert.Equal(t, RookType, rook.PieceType(), "Wrong piece type")

			assert.Equal(t, len(tt.want), len(aMoves), "Incorrect Number of Moves")

			var aMovePGNS []string
			for _, m := range aMoves {
				aMovePGNS = append(aMovePGNS, m.PGNName())
			}

			for _, aMove := range aMoves {
				assert.Contains(t, tt.want, aMove.PGNName(), "Expected Slice doesn't contain actual")
			}

			for _, eMove := range tt.want {
				assert.Contains(t, aMovePGNS, eMove, "Actual Slice doesn't contain expected")
			}
		})
	}
}
