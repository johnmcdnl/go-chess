package chess

import (
	"reflect"
	"testing"
)

func TestNewPiece(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Piece
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPiece(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPiece() = %v, want %v", got, tt.want)
			}
		})
	}
}
