package chess

import (
	"testing"
)

func TestGetSquareID(t *testing.T) {
	type args struct {
		file int
		rank int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"A1", args{A, 1}, 0},
		{"H8", args{H, 8}, 63},
		{"C5", args{C, 5}, 34},
		{"Invalid Low", args{0, 0}, -1},
		{"Invalid High", args{9, 9}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSquareID(tt.args.file, tt.args.rank); got != tt.want {
				t.Errorf("GetSquareID() = %v, want %v", got, tt.want)
			}
		})
	}
}
