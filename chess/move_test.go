package chess

import (
	"testing"
)

func TestMove_PGNName(t *testing.T) {

	boardFEN := func(fen FEN) *Board {
		b, _ := BoardFEN(fen)
		return b
	}

	type fields struct {
		board       *Board
		Origin      Coordinate
		Destination Coordinate
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Qd8xd5", fields{board: boardFEN(FEN("3q4/8/8/3P4/8/8/8/8 w KQkq - 1 1")), Origin:Coordinate{D, 8}, Destination:Coordinate{D, 5}}, "Qd8xd5"},
		{"Qd8d5", fields{board: boardFEN(FEN("3q4/8/8/8/8/8/8/8 w KQkq - 1 1")), Origin:Coordinate{D, 8}, Destination:Coordinate{D, 5}}, "Qd8d5"},
		{"Qd8xb6", fields{board: boardFEN(FEN("3q4/8/1P6/8/8/8/8/8 w KQkq - 1 1")), Origin:Coordinate{D, 8}, Destination:Coordinate{B, 6}}, "Qd8xb6"},
		{"Ne3xd5", fields{board: boardFEN(FEN("8/8/8/3p4/8/4N3/8/8 w KQkq - 1 1")), Origin:Coordinate{E, 3}, Destination:Coordinate{D, 5}}, "Ne3xd5"},
		{"Ne3d5", fields{board: boardFEN(FEN("8/8/8/8/8/4N3/8/8 w KQkq - 1 1")), Origin:Coordinate{E, 3}, Destination:Coordinate{D, 5}}, "Ne3d5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Move{
				Origin:      tt.fields.board.GetSquare(tt.fields.Origin.File, tt.fields.Origin.Rank),
				Destination: tt.fields.board.GetSquare(tt.fields.Destination.File, tt.fields.Destination.Rank),
			}

			if got := m.PGNName(); got != tt.want {
				t.Errorf("Move.PGNName() = %v, want %v", got, tt.want)
			}
		})
	}
}

