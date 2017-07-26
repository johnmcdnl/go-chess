package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueen_ValidMoves(t *testing.T) {
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
			args{b: boardFEN("7k/8/8/3Q4/8/8/8/K7 w - - 0 1")},
			[]string{
				"Qd5d1", "Qd5d2", "Qd5d3", "Qd5d4", "Qd5d6", "Qd5d7", "Qd5d8",
				"Qd5a5", "Qd5b5", "Qd5c5", "Qd5e5", "Qd5f5", "Qd5g5", "Qd5h5",
				"Qd5a2", "Qd5b3", "Qd5c4", "Qd5e6", "Qd5f7", "Qd5g8",
				"Qd5a8", "Qd5b7", "Qd5c6", "Qd5e4", "Qd5f3", "Qd5g2", "Qd5h1",
			},
		},
		{
			"In a Corner",
			fields{Coordinate{H, 1}},
			args{b: boardFEN("3Kk3/8/8/8/8/8/8/7q w - - 0 1")},
			[]string{
				"Qh1a1", "Qh1b1", "Qh1c1", "Qh1d1", "Qh1e1", "Qh1f1", "Qh1g1",
				"Qh1h2", "Qh1h3", "Qh1h4", "Qh1h5", "Qh1h6", "Qh1h7", "Qh1h8",
				"Qh1g2", "Qh1f3", "Qh1e4", "Qh1d5", "Qh1c6", "Qh1b7", "Qh1a8",
			},
		},
		{
			"Almost In a Corner",
			fields{Coordinate{G, 2}},
			args{b: boardFEN("3Kk3/8/8/8/8/8/6q1/8 w - - 0 1")},
			[]string{
				"Qg2a8", "Qg2b7", "Qg2c6", "Qg2d5", "Qg2e4", "Qg2f3", "Qg2h1",
				"Qg2a2", "Qg2b2", "Qg2c2", "Qg2d2", "Qg2e2", "Qg2f2", "Qg2h2",
				"Qg2g1", "Qg2g3", "Qg2g4", "Qg2g5", "Qg2g6", "Qg2g7", "Qg2g8",
				"Qg2f1", "Qg2h3",
			},
		},
		{
			"Queen Cannot Jump",
			fields{Coordinate{D, 5}},
			args{b: boardFEN("3Kk3/8/2pnn3/2nqn3/2nnn3/8/8/8 w - - 0 1")},
			[]string{},
		},
		{
			"Queen Capturing in multiple directions",
			fields{Coordinate{D, 5}},
			args{b: boardFEN("3Kk3/1P6/3pP3/2pqp3/2P5/2PP1P2/8/8 w - - 0 1")},
			[]string{"Qd5c6", "Qd5xb7", "Qd5xe6", "Qd5e4", "Qd5xf3", "Qd5d4", "Qd5xd3", "Qd5xc4"},
		},
		{
			"Queen cannot capture a king",
			fields{Coordinate{D, 5}},
			args{b: boardFEN("8/1P6/2KpP3/2pqk3/2P5/2PP1P2/8/8 w - - 0 1")},
			[]string{"Qd5xe6", "Qd5e4", "Qd5xf3", "Qd5d4", "Qd5xd3", "Qd5xc4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queen := tt.args.b.GetSquare(tt.fields.Coordinate.File, tt.fields.Coordinate.Rank).CurrentPiece
			aMoves := queen.ValidMoves(tt.args.b)
			assert.Equal(t, QueenType, queen.PieceType(), "Wrong piece type")

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
