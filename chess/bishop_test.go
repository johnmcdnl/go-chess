package chess

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBishop(t *testing.T) {
	type args struct {
		s *Square
		c Color
	}

	tests := []struct {
		name string
		args args
		want *Bishop
	}{
		{"White Bishop at A1", args{NewSquare(A, 1), White}, &Bishop{Name: "bishop", Position: NewSquare(A, 1), Code: BishopPiece, Color: White}},
		{"Black Bishop at H1", args{NewSquare(H, 1), Black}, &Bishop{Name: "bishop", Position: NewSquare(H, 1), Code: BishopPiece, Color: Black}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBishop(tt.args.s, tt.args.c)
			assert.Equal(t, tt.want.Name, got.Name, "Name")
			assert.Equal(t, tt.want.Code, got.Code, "Code")
			assert.Equal(t, tt.want.Color, got.Color, "Color")
			assert.Equal(t, tt.want.Position.File, got.Position.File, "Position.File")
			assert.Equal(t, tt.want.Position.Rank, got.Position.Rank, "Position.Rank")
		})
	}
}

func TestBishop_ValidMoves(t *testing.T) {
	//boardStartingPositionFEN := NewBoardWithFen(StartingPositionFEN)
	boardBlockedE5FEN := NewBoardWithFen(BlockedE5FEN)
	//boardOnlyKingsFEN := NewBoardWithFen(OnlyKingsFEN)
	type fields struct {
		Name     string
		Code     int
		Color    Color
		Position *Square
	}
	type args struct {
		board *Board
	}

	tests := []struct {
		name   string
		bishop *Bishop
		args   args
		want   []*Move
	}{
		//{"Bishop A1 Empty Board",
		//	NewBishop(boardOnlyKingsFEN.GetSquare(A, 1), White),
		//	args{NewBoardWithFen(FEN(OnlyKingsFEN))},
		//	[]*Move{
		//		{boardOnlyKingsFEN.GetSquare(A, 1), boardOnlyKingsFEN.GetSquare(B, 2), false, "Ba1b2"},
		//		{boardOnlyKingsFEN.GetSquare(A, 1), boardOnlyKingsFEN.GetSquare(C, 3), false, "Ba1c3"},
		//		{boardOnlyKingsFEN.GetSquare(A, 1), boardOnlyKingsFEN.GetSquare(D, 4), false, "Ba1d4"},
		//		{boardOnlyKingsFEN.GetSquare(A, 1), boardOnlyKingsFEN.GetSquare(E, 5), false, "Ba1e5"},
		//		{boardOnlyKingsFEN.GetSquare(A, 1), boardOnlyKingsFEN.GetSquare(F, 6), false, "Ba1f6"},
		//		{boardOnlyKingsFEN.GetSquare(A, 1), boardOnlyKingsFEN.GetSquare(G, 7), false, "Ba1g7"},
		//		{boardOnlyKingsFEN.GetSquare(A, 1), boardOnlyKingsFEN.GetSquare(H, 8), false, "Ba1h8"},
		//	},
		//},
		//{"Black Bishop C8 With Starting Position",
		//	NewBishop(boardStartingPositionFEN.GetSquare(C, 8), Black),
		//	args{NewBoardWithFen(FEN(StartingPositionFEN))},
		//	[]*Move{},
		//},

		{"Bishop D4 With Blocked E5 Square, Open Other Sides",
			NewBishop(boardBlockedE5FEN.GetSquare(D, 4), White),
			args{NewBoardWithFen(FEN(BlockedE5FEN))},
			[]*Move{
				{boardBlockedE5FEN.GetSquare(D, 4), boardBlockedE5FEN.GetSquare(A, 7), false, "Bd4a7"},
				{boardBlockedE5FEN.GetSquare(D, 4), boardBlockedE5FEN.GetSquare(B, 6), false, "Bd4b6"},
				{boardBlockedE5FEN.GetSquare(D, 4), boardBlockedE5FEN.GetSquare(C, 5), false, "Bd4c5"},

				{boardBlockedE5FEN.GetSquare(D, 4), boardBlockedE5FEN.GetSquare(E, 3), false, "Bd4a7"},
				{boardBlockedE5FEN.GetSquare(D, 4), boardBlockedE5FEN.GetSquare(F, 2), false, "Bd4f2"},
				{boardBlockedE5FEN.GetSquare(D, 4), boardBlockedE5FEN.GetSquare(G, 1), false, "Bd4g1"},

				{boardBlockedE5FEN.GetSquare(D, 4), boardBlockedE5FEN.GetSquare(C, 3), false, "Bd4c3"},
				{boardBlockedE5FEN.GetSquare(D, 4), boardBlockedE5FEN.GetSquare(B, 2), false, "Bd4b2"},
				{boardBlockedE5FEN.GetSquare(D, 4), boardBlockedE5FEN.GetSquare(A, 1), false, "Bd4a1"},

				{boardBlockedE5FEN.GetSquare(D, 4), boardBlockedE5FEN.GetSquare(E, 5), true, "Bd4e5"},

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bishop.ValidMoves(tt.args.board)

			assert.Len(t, got, len(tt.want))


			//for _, w := range tt.want {
			//	var namePassed bool
			//	for _, g := range got {
			//		if w.name == g.name {
			//			namePassed = true
			//			assert.Equal(t, w.IsCaptureMove, g.IsCaptureMove)
			//			assert.Equal(t, w.origin.File, g.origin.File)
			//			assert.Equal(t, w.origin.Rank, g.origin.Rank)
			//			assert.Equal(t, w.destination.File, g.destination.File)
			//			assert.Equal(t, w.destination.Rank, g.destination.Rank)
			//			break
			//		}
			//	}
			//
			//	assert.True(t, namePassed, fmt.Sprintln(w.Name()))
			//}

			//tt.args.board.Print()
		})
	}
}

func TestBishop_GetCode(t *testing.T) {
	type fields struct {
		Name     string
		Code     int
		Color    Color
		Position *Square
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bishop := &Bishop{
				Name:     tt.fields.Name,
				Code:     tt.fields.Code,
				Color:    tt.fields.Color,
				Position: tt.fields.Position,
			}
			if got := bishop.GetCode(); got != tt.want {
				t.Errorf("Bishop.GetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBishop_GetColor(t *testing.T) {
	type fields struct {
		Name     string
		Code     int
		Color    Color
		Position *Square
	}
	tests := []struct {
		name   string
		fields fields
		want   Color
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bishop := &Bishop{
				Name:     tt.fields.Name,
				Code:     tt.fields.Code,
				Color:    tt.fields.Color,
				Position: tt.fields.Position,
			}
			if got := bishop.GetColor(); got != tt.want {
				t.Errorf("Bishop.GetColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBishop_CurrentPosition(t *testing.T) {
	type fields struct {
		Name     string
		Code     int
		Color    Color
		Position *Square
	}
	tests := []struct {
		name   string
		fields fields
		want   *Square
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bishop := &Bishop{
				Name:     tt.fields.Name,
				Code:     tt.fields.Code,
				Color:    tt.fields.Color,
				Position: tt.fields.Position,
			}
			if got := bishop.CurrentPosition(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bishop.CurrentPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
