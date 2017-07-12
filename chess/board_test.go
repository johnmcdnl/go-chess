package chess

import (
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	tests := []struct {
		name string
		want *Board
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoard(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_setStartingPosition(t *testing.T) {
	type fields struct {
		Squares *[]*Square
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				Squares: tt.fields.Squares,
			}
			b.setStartingPosition()
		})
	}
}

func TestBoard_SetFEN(t *testing.T) {
	type fields struct {
		Squares *[]*Square
	}
	type args struct {
		fen FEN
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				Squares: tt.fields.Squares,
			}
			b.SetFEN(tt.args.fen)
		})
	}
}

func TestBoard_GetFen(t *testing.T) {
	type fields struct {
		Squares *[]*Square
	}
	tests := []struct {
		name   string
		fields fields
		want   FEN
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				Squares: tt.fields.Squares,
			}
			if got := b.GetFen(); got != tt.want {
				t.Errorf("Board.GetFen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_GetSquare(t *testing.T) {
	type fields struct {
		Squares *[]*Square
	}
	type args struct {
		file int
		rank int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Square
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				Squares: tt.fields.Squares,
			}
			if got := b.GetSquare(tt.args.file, tt.args.rank); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.GetSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
