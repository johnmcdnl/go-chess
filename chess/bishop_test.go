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
		{"White Bishop at A1", args{NewSquare(A, 1), White}, &Bishop{Name:"bishop", Position:NewSquare(A, 1), Code:BishopPiece, Color:White}},
		{"Black Bishop at H1", args{NewSquare(H, 1), Black}, &Bishop{Name:"bishop", Position:NewSquare(H, 1), Code:BishopPiece, Color:Black}},
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
		fields fields
		args   args
		want   []*Move
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
			if got := bishop.ValidMoves(tt.args.board); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bishop.ValidMoves() = %v, want %v", got, tt.want)
			}
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
