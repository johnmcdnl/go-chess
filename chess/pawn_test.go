package chess

import (
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewPawn(t *testing.T) {
	type args struct {
		s *Square
		c Color
	}
	tests := []struct {
		name string
		args args
		want *Pawn
	}{
		{"White Pawn at A1", args{NewSquare(A, 1), White}, &Pawn{Name:"pawn", Position:NewSquare(A, 1), Code:PawnPiece, Color:White}},
		{"Black Pawn at H1", args{NewSquare(H, 1), Black}, &Pawn{Name:"pawn", Position:NewSquare(H, 1), Code:PawnPiece, Color:Black}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPawn(tt.args.s, tt.args.c)
			assert.Equal(t, tt.want.Name, got.Name, "Name")
			assert.Equal(t, tt.want.Code, got.Code, "Code")
			assert.Equal(t, tt.want.Color, got.Color, "Color")
			assert.Equal(t, tt.want.Position.File, got.Position.File, "Position.File")
			assert.Equal(t, tt.want.Position.Rank, got.Position.Rank, "Position.Rank")
		})
	}
}

func TestPawn_ValidMoves(t *testing.T) {
	type fields struct {
		Name     string
		Code     int
		Color    Color
		Position *Square
	}
	type args struct {
		b *Board
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Move
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pawn{
				Name:     tt.fields.Name,
				Code:     tt.fields.Code,
				Color:    tt.fields.Color,
				Position: tt.fields.Position,
			}
			if got := p.ValidMoves(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pawn.ValidMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPawn_whiteMoves(t *testing.T) {
	type fields struct {
		Name     string
		Code     int
		Color    Color
		Position *Square
	}
	type args struct {
		b *Board
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Move
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pawn{
				Name:     tt.fields.Name,
				Code:     tt.fields.Code,
				Color:    tt.fields.Color,
				Position: tt.fields.Position,
			}
			if got := p.whiteMoves(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pawn.whiteMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPawn_blackMoves(t *testing.T) {
	type fields struct {
		Name     string
		Code     int
		Color    Color
		Position *Square
	}
	type args struct {
		b *Board
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Move
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pawn{
				Name:     tt.fields.Name,
				Code:     tt.fields.Code,
				Color:    tt.fields.Color,
				Position: tt.fields.Position,
			}
			if got := p.blackMoves(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pawn.blackMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPawn_GetCode(t *testing.T) {
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
			p := &Pawn{
				Name:     tt.fields.Name,
				Code:     tt.fields.Code,
				Color:    tt.fields.Color,
				Position: tt.fields.Position,
			}
			if got := p.GetCode(); got != tt.want {
				t.Errorf("Pawn.GetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPawn_GetColor(t *testing.T) {
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
			p := &Pawn{
				Name:     tt.fields.Name,
				Code:     tt.fields.Code,
				Color:    tt.fields.Color,
				Position: tt.fields.Position,
			}
			if got := p.GetColor(); got != tt.want {
				t.Errorf("Pawn.GetColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPawn_CurrentPosition(t *testing.T) {
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
			p := &Pawn{
				Name:     tt.fields.Name,
				Code:     tt.fields.Code,
				Color:    tt.fields.Color,
				Position: tt.fields.Position,
			}
			if got := p.CurrentPosition(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pawn.CurrentPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
