package chess

import (
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewKnight(t *testing.T) {
	type args struct {
		s *Square
		c Color
	}
	tests := []struct {
		name string
		args args
		want *Knight
	}{
		{"White Knight at A1", args{NewSquare(A, 1), White}, &Knight{Name:"knight", Position:NewSquare(A, 1), Code:KnightPiece, Color:White}},
		{"Black Knight at H1", args{NewSquare(H, 1), Black}, &Knight{Name:"knight", Position:NewSquare(H, 1), Code:KnightPiece, Color:Black}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewKnight(tt.args.s, tt.args.c)
			assert.Equal(t, tt.want.Name, got.Name, "Name")
			assert.Equal(t, tt.want.Code, got.Code, "Code")
			assert.Equal(t, tt.want.Color, got.Color, "Color")
			assert.Equal(t, tt.want.Position.File, got.Position.File, "Position.File")
			assert.Equal(t, tt.want.Position.Rank, got.Position.Rank, "Position.Rank")

		})
	}
}

func TestKnight_ValidMoves(t *testing.T) {
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
			k := &Knight{
				Name:     tt.fields.Name,
				Code:     tt.fields.Code,
				Color:    tt.fields.Color,
				Position: tt.fields.Position,
			}
			if got := k.ValidMoves(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Knight.ValidMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKnight_GetCode(t *testing.T) {
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
			k := &Knight{
				Name:     tt.fields.Name,
				Code:     tt.fields.Code,
				Color:    tt.fields.Color,
				Position: tt.fields.Position,
			}
			if got := k.GetCode(); got != tt.want {
				t.Errorf("Knight.GetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKnight_GetColor(t *testing.T) {
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
			k := &Knight{
				Name:     tt.fields.Name,
				Code:     tt.fields.Code,
				Color:    tt.fields.Color,
				Position: tt.fields.Position,
			}
			if got := k.GetColor(); got != tt.want {
				t.Errorf("Knight.GetColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKnight_CurrentPosition(t *testing.T) {
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
			k := &Knight{
				Name:     tt.fields.Name,
				Code:     tt.fields.Code,
				Color:    tt.fields.Color,
				Position: tt.fields.Position,
			}
			if got := k.CurrentPosition(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Knight.CurrentPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
