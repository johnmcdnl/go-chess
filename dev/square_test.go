package dev

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSquare(t *testing.T) {
	type args struct {
		file int
		rank int
	}
	tests := []struct {
		name     string
		args     args
		want     *Square
		hasError bool
	}{
		{"a1", args{A, 1}, &Square{File: A, Rank: 1, name: "a1", id: 0, isOccupied: false}, false},
		{"a2", args{A, 2}, &Square{File: A, Rank: 2, name: "a2", id: 1, isOccupied: false}, false},
		{"a3", args{A, 3}, &Square{File: A, Rank: 3, name: "a3", id: 2, isOccupied: false}, false},
		{"a4", args{A, 4}, &Square{File: A, Rank: 4, name: "a4", id: 3, isOccupied: false}, false},
		{"a5", args{A, 5}, &Square{File: A, Rank: 5, name: "a5", id: 4, isOccupied: false}, false},
		{"a6", args{A, 6}, &Square{File: A, Rank: 6, name: "a6", id: 5, isOccupied: false}, false},
		{"a7", args{A, 7}, &Square{File: A, Rank: 7, name: "a7", id: 6, isOccupied: false}, false},
		{"a8", args{A, 8}, &Square{File: A, Rank: 8, name: "a8", id: 7, isOccupied: false}, false},
		{"b1", args{B, 1}, &Square{File: B, Rank: 1, name: "b1", id: 8, isOccupied: false}, false},
		{"b2", args{B, 2}, &Square{File: B, Rank: 2, name: "b2", id: 9, isOccupied: false}, false},
		{"b3", args{B, 3}, &Square{File: B, Rank: 3, name: "b3", id: 10, isOccupied: false}, false},
		{"b4", args{B, 4}, &Square{File: B, Rank: 4, name: "b4", id: 11, isOccupied: false}, false},
		{"b5", args{B, 5}, &Square{File: B, Rank: 5, name: "b5", id: 12, isOccupied: false}, false},
		{"b6", args{B, 6}, &Square{File: B, Rank: 6, name: "b6", id: 13, isOccupied: false}, false},
		{"b7", args{B, 7}, &Square{File: B, Rank: 7, name: "b7", id: 14, isOccupied: false}, false},
		{"b8", args{B, 8}, &Square{File: B, Rank: 8, name: "b8", id: 15, isOccupied: false}, false},
		{"c1", args{C, 1}, &Square{File: C, Rank: 1, name: "c1", id: 16, isOccupied: false}, false},
		{"c2", args{C, 2}, &Square{File: C, Rank: 2, name: "c2", id: 17, isOccupied: false}, false},
		{"c3", args{C, 3}, &Square{File: C, Rank: 3, name: "c3", id: 18, isOccupied: false}, false},
		{"c4", args{C, 4}, &Square{File: C, Rank: 4, name: "c4", id: 19, isOccupied: false}, false},
		{"c5", args{C, 5}, &Square{File: C, Rank: 5, name: "c5", id: 20, isOccupied: false}, false},
		{"c6", args{C, 6}, &Square{File: C, Rank: 6, name: "c6", id: 21, isOccupied: false}, false},
		{"c7", args{C, 7}, &Square{File: C, Rank: 7, name: "c7", id: 22, isOccupied: false}, false},
		{"c8", args{C, 8}, &Square{File: C, Rank: 8, name: "c8", id: 23, isOccupied: false}, false},
		{"d1", args{D, 1}, &Square{File: D, Rank: 1, name: "d1", id: 24, isOccupied: false}, false},
		{"d2", args{D, 2}, &Square{File: D, Rank: 2, name: "d2", id: 25, isOccupied: false}, false},
		{"d3", args{D, 3}, &Square{File: D, Rank: 3, name: "d3", id: 26, isOccupied: false}, false},
		{"d4", args{D, 4}, &Square{File: D, Rank: 4, name: "d4", id: 27, isOccupied: false}, false},
		{"d5", args{D, 5}, &Square{File: D, Rank: 5, name: "d5", id: 28, isOccupied: false}, false},
		{"d6", args{D, 6}, &Square{File: D, Rank: 6, name: "d6", id: 29, isOccupied: false}, false},
		{"d7", args{D, 7}, &Square{File: D, Rank: 7, name: "d7", id: 30, isOccupied: false}, false},
		{"d8", args{D, 8}, &Square{File: D, Rank: 8, name: "d8", id: 31, isOccupied: false}, false},
		{"e1", args{E, 1}, &Square{File: E, Rank: 1, name: "e1", id: 32, isOccupied: false}, false},
		{"e2", args{E, 2}, &Square{File: E, Rank: 2, name: "e2", id: 33, isOccupied: false}, false},
		{"e3", args{E, 3}, &Square{File: E, Rank: 3, name: "e3", id: 34, isOccupied: false}, false},
		{"e4", args{E, 4}, &Square{File: E, Rank: 4, name: "e4", id: 35, isOccupied: false}, false},
		{"e5", args{E, 5}, &Square{File: E, Rank: 5, name: "e5", id: 36, isOccupied: false}, false},
		{"e6", args{E, 6}, &Square{File: E, Rank: 6, name: "e6", id: 37, isOccupied: false}, false},
		{"e7", args{E, 7}, &Square{File: E, Rank: 7, name: "e7", id: 38, isOccupied: false}, false},
		{"e8", args{E, 8}, &Square{File: E, Rank: 8, name: "e8", id: 39, isOccupied: false}, false},
		{"f1", args{F, 1}, &Square{File: F, Rank: 1, name: "f1", id: 40, isOccupied: false}, false},
		{"f2", args{F, 2}, &Square{File: F, Rank: 2, name: "f2", id: 41, isOccupied: false}, false},
		{"f3", args{F, 3}, &Square{File: F, Rank: 3, name: "f3", id: 42, isOccupied: false}, false},
		{"f4", args{F, 4}, &Square{File: F, Rank: 4, name: "f4", id: 43, isOccupied: false}, false},
		{"f5", args{F, 5}, &Square{File: F, Rank: 5, name: "f5", id: 44, isOccupied: false}, false},
		{"f6", args{F, 6}, &Square{File: F, Rank: 6, name: "f6", id: 45, isOccupied: false}, false},
		{"f7", args{F, 7}, &Square{File: F, Rank: 7, name: "f7", id: 46, isOccupied: false}, false},
		{"f8", args{F, 8}, &Square{File: F, Rank: 8, name: "f8", id: 47, isOccupied: false}, false},
		{"g1", args{G, 1}, &Square{File: G, Rank: 1, name: "g1", id: 48, isOccupied: false}, false},
		{"g2", args{G, 2}, &Square{File: G, Rank: 2, name: "g2", id: 49, isOccupied: false}, false},
		{"g3", args{G, 3}, &Square{File: G, Rank: 3, name: "g3", id: 50, isOccupied: false}, false},
		{"g4", args{G, 4}, &Square{File: G, Rank: 4, name: "g4", id: 51, isOccupied: false}, false},
		{"g5", args{G, 5}, &Square{File: G, Rank: 5, name: "g5", id: 52, isOccupied: false}, false},
		{"g6", args{G, 6}, &Square{File: G, Rank: 6, name: "g6", id: 53, isOccupied: false}, false},
		{"g7", args{G, 7}, &Square{File: G, Rank: 7, name: "g7", id: 54, isOccupied: false}, false},
		{"g8", args{G, 8}, &Square{File: G, Rank: 8, name: "g8", id: 55, isOccupied: false}, false},
		{"h1", args{H, 1}, &Square{File: H, Rank: 1, name: "h1", id: 56, isOccupied: false}, false},
		{"h2", args{H, 2}, &Square{File: H, Rank: 2, name: "h2", id: 57, isOccupied: false}, false},
		{"h3", args{H, 3}, &Square{File: H, Rank: 3, name: "h3", id: 58, isOccupied: false}, false},
		{"h4", args{H, 4}, &Square{File: H, Rank: 4, name: "h4", id: 59, isOccupied: false}, false},
		{"h5", args{H, 5}, &Square{File: H, Rank: 5, name: "h5", id: 60, isOccupied: false}, false},
		{"h6", args{H, 6}, &Square{File: H, Rank: 6, name: "h6", id: 61, isOccupied: false}, false},
		{"h7", args{H, 7}, &Square{File: H, Rank: 7, name: "h7", id: 62, isOccupied: false}, false},
		{"h8", args{H, 8}, &Square{File: H, Rank: 8, name: "h8", id: 63, isOccupied: false}, false},

		{"file too low", args{0, 8}, nil, true },
		{"file too high", args{9, 8}, nil, true},
		{"rank too low", args{H, 0}, nil, true},
		{"rank too high", args{H, 9}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSquare(tt.args.file, tt.args.rank)

			if tt.hasError {
				assert.Nil(t, got)
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tt.want.File, got.File, tt.want.name)
			assert.Equal(t, tt.want.Rank, got.Rank, tt.want.name)
			assert.Equal(t, tt.want.name, got.name, tt.want.name)
			assert.Equal(t, tt.want.id, got.id, tt.want.name)
			assert.Equal(t, tt.want.isOccupied, got.isOccupied, tt.want.name)
			assert.NoError(t, err)

		})
	}
}

func TestSquare_GetID(t *testing.T) {
	type fields struct {
		id         int
		File       int
		Rank       int
		name       string
		isOccupied bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"valid-id", fields{10, 0, 0, "some-square", true}, 10 },
		{"large-id", fields{100, 0, 0, "some-square", false}, 100 },
		{"zero-id", fields{0, 0, 0, "some-square", false}, 0 },
		{"negative-id", fields{-1, 0, 0, "some-square", false}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Square{
				id:         tt.fields.id,
				File:       tt.fields.File,
				Rank:       tt.fields.Rank,
				name:       tt.fields.name,
				isOccupied: tt.fields.isOccupied,
			}
			if got := s.GetID(); got != tt.want {
				t.Errorf("Square.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare_Name(t *testing.T) {
	type fields struct {
		id         int
		File       int
		Rank       int
		name       string
		isOccupied bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"valid-name", fields{0, 0, 0, "my-valid-name", true}, "my-valid-name" },
		{"empty-name", fields{0, 0, 0, "", false}, "" },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Square{
				id:         tt.fields.id,
				File:       tt.fields.File,
				Rank:       tt.fields.Rank,
				name:       tt.fields.name,
				isOccupied: tt.fields.isOccupied,
			}
			if got := s.Name(); got != tt.want {
				t.Errorf("Square.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare_IsOccupied(t *testing.T) {
	type fields struct {
		id         int
		File       int
		Rank       int
		name       string
		isOccupied bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"is-occupied", fields{0, 0, 0, "some-square", true}, true },
		{"is-not-occupied", fields{0, 0, 0, "some-square", false}, false },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Square{
				id:         tt.fields.id,
				File:       tt.fields.File,
				Rank:       tt.fields.Rank,
				name:       tt.fields.name,
				isOccupied: tt.fields.isOccupied,
			}
			if got := s.IsOccupied(); got != tt.want {
				t.Errorf("Square.IsOccupied() = %v, want %v", got, tt.want)
			}
		})
	}
}
