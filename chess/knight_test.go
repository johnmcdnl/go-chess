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
