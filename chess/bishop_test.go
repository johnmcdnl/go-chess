package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBishop_ValidMoves(t *testing.T) {
	boardFEN := func(fen FEN) *Board {
		if !fen.IsValid() {
			assert.Fail(t, "invalid fen")
		}
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
			fields{Coordinate{E, 4}},
			args{b: boardFEN("7k/8/8/8/4B3/8/8/K7 w - - 0 1")},
			[]string{"Be4b1", "Be4c2", "Be4d3", "Be4f5", "Be4g6", "Be4h7", "Be4a8", "Be4b7", "Be4c6", "Be4d5", "Be4f3", "Be4g2", "Be4h1"},
		},
		{
			"Center Capture 4 directions",
			fields{Coordinate{E, 4}},
			args{b: boardFEN("7k/8/8/3r1n2/4B3/3q1p2/8/K7 w - - 0 1")},
			[]string{"Be4xd5", "Be4xd3", "Be4xf3", "Be4xf5"},
		},
		{
			"A bishop blocked in by it's own pieces",
			fields{Coordinate{E, 4}},
			args{b: boardFEN("7k/8/8/3R1B2/4B3/3Q1P2/8/K7 w - - 0 1")},
			[]string{},
		},
		{
			"A bishop in the corner",
			fields{Coordinate{A, 1}},
			args{b: boardFEN("k7/8/8/8/8/8/8/b6K w - - 0 1")},
			[]string{"Ba1b2", "Ba1c3", "Ba1d4", "Ba1e5", "Ba1f6", "Ba1g7", "Ba1h8"},
		},
		{
			"A bishop on the left side of the board",
			fields{Coordinate{A, 5}},
			args{b: boardFEN("k7/8/8/B7/8/8/8/7K w - - 0 1")},
			[]string{"Ba5b6", "Ba5c7", "Ba5d8", "Ba5b4", "Ba5c3", "Ba5d2", "Ba5e1"},
		},
		{
			"A bishop on the right side of the board",
			fields{Coordinate{H, 6}},
			args{b: boardFEN("k7/8/7b/8/8/8/8/7K w - - 0 1")},
			[]string{"Bh6g7", "Bh6f8", "Bh6g5", "Bh6f4", "Bh6e3", "Bh6d2", "Bh6c1"},
		},
		{
			"A bishop almost in the corner",
			fields{Coordinate{B, 2}},
			args{b: boardFEN("k7/8/8/8/8/8/1b6/7K w - - 0 1")},
			[]string{"Bb2a1", "Bb2c3", "Bb2d4", "Bb2e5", "Bb2f6", "Bb2g7", "Bb2h8", "Bb2a3", "Bb2c1"},
		},
		{
			"A bishop at the start of a game",
			fields{Coordinate{C, 1}},
			args{b: boardFEN(StartingPositionFEN)},
			[]string{},
		},
		{
			"A bishop cannot capture king",
			fields{Coordinate{C, 6}},
			args{b: boardFEN("q3k3/8/2B5/8/8/8/8/4K3 w - - 0 1")},
			[]string{"Bc6xa8", "Bc6b7", "Bc6d5", "Bc6e4", "Bc6f3", "Bc6g2", "Bc6h1", "Bc6a4", "Bc6b5", "Bc6d7"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bishop := tt.args.b.GetSquare(tt.fields.Coordinate.File, tt.fields.Coordinate.Rank).CurrentPiece
			aMoves := bishop.ValidMoves(tt.args.b)
			assert.Equal(t, BishopType, bishop.PieceType(), "Wrong piece type")

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
