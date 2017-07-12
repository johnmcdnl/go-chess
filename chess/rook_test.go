package chess

import (
	"reflect"
	"testing"
)

func TestNewRook(t *testing.T) {
	type args struct {
		s *Square
		c Color
	}
	tests := []struct {
		name string
		args args
		want *Rook
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRook(tt.args.s, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRook_ValidMoves(t *testing.T) {
	type args struct {
		b *Board
	}
	tests := []struct {
		name string
		r    *Rook
		args args
		want []*Move
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.ValidMoves(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rook.ValidMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRook_GetCode(t *testing.T) {
	tests := []struct {
		name string
		r    *Rook
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.GetCode(); got != tt.want {
				t.Errorf("Rook.GetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRook_GetColor(t *testing.T) {
	tests := []struct {
		name string
		r    *Rook
		want Color
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.GetColor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rook.GetColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRook_CurrentPosition(t *testing.T) {
	tests := []struct {
		name string
		r    *Rook
		want *Square
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.CurrentPosition(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rook.CurrentPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
