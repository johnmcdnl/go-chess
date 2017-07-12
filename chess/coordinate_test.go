package chess

import (
	"reflect"
	"testing"
)

func TestNewCoordinate(t *testing.T) {
	type args struct {
		file int
		rank int
	}
	tests := []struct {
		name string
		args args
		want *Coordinate
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCoordinate(tt.args.file, tt.args.rank); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCoordinate() = %v, want %v", got, tt.want)
			}
		})
	}
}
