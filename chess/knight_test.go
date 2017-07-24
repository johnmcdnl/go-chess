package chess

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestKnight_ValidMoves(t *testing.T) {
	boardFEN := func(fen FEN) *Board {
		b, _ := BoardFEN(fen)
		return b
	}

	type fields struct {
		Coordinate Coordinate
		Color      Color
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
			fields{Coordinate{E, 4}, White},
			args{b: boardFEN("8/8/8/8/4N3/8/8/8 w - - 0 1")},
			[]string{"Ne4c5", "Ne4d6", "Ne4f6", "Ne4g5", "Ne4g3", "Ne4f2", "Ne4d2", "Ne4c3"},
		},
		{
			"Center Capture 8 directions",
			fields{Coordinate{E, 4}, White},
			args{b: boardFEN("8/8/3p1p2/2p3p1/4N3/2p3p1/3p1p2/8 w - - 0 1")},
			[]string{"Ne4xc5", "Ne4xd6", "Ne4xf6", "Ne4xg5", "Ne4xg3", "Ne4xf2", "Ne4xd2", "Ne4xc3"},
		},
		{
			"A knight can jump",
			fields{Coordinate{E, 4}, White},
			args{b: boardFEN("8/8/2BpQpB1/2pqbnp1/2PPNPP1/2prbrp1/2PpPpP1/8 w - - 0 1")},
			[]string{"Ne4xc5", "Ne4xd6", "Ne4xf6", "Ne4xg5", "Ne4xg3", "Ne4xf2", "Ne4xd2", "Ne4xc3"},
		},
		{
			"A knight pinned in the corner",
			fields{Coordinate{A, 1}, White},
			args{b: boardFEN("8/8/8/8/8/8/8/N7 w - - 0 1")},
			[]string{"Na1b3", "Na1c2"},
		},
		{
			"A knight almost trapped in the corner",
			fields{Coordinate{B, 2}, White},
			args{b: boardFEN("8/8/8/8/8/8/1N6/8 w - - 0 1")},
			[]string{"Nb2a4", "Nb2c4", "Nb2d3", "Nb2d1"},
		},
		{
			"From the starting position",
			fields{Coordinate{B, 1}, White},
			args{b: boardFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
			[]string{"Nb1a3", "Nb1c3"},
		},
		{
			"Can't capture a King",
			fields{Coordinate{D, 6}, White},
			args{b: boardFEN("2q1k3/8/3N4/8/8/8/8/8 w - - 0 1")},
			[]string{"Nd6xc8", "Nd6b7", "Nd6b5", "Nd6c4", "Nd6e4", "Nd6f5", "Nd6f7"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.args.b.GetSquare(tt.fields.Coordinate.File, tt.fields.Coordinate.Rank)
			k, _ := NewKnight(s, tt.fields.Color)
			aMoves := k.ValidMoves(tt.args.b)

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
