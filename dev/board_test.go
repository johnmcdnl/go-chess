package dev

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	tests := []struct {
		name string
		want *Board
	}{
		{"valid-board",
			&Board{
				Squares:[]*Square{
					{File: A, Rank: 1, name: "a1", id: 0, isOccupied: false},
					{File: A, Rank: 2, name: "a2", id: 1, isOccupied: false},
					{File: A, Rank: 3, name: "a3", id: 2, isOccupied: false},
					{File: A, Rank: 4, name: "a4", id: 3, isOccupied: false},
					{File: A, Rank: 5, name: "a5", id: 4, isOccupied: false},
					{File: A, Rank: 6, name: "a6", id: 5, isOccupied: false},
					{File: A, Rank: 7, name: "a7", id: 6, isOccupied: false},
					{File: A, Rank: 8, name: "a8", id: 7, isOccupied: false},
					{File: B, Rank: 1, name: "b1", id: 8, isOccupied: false},
					{File: B, Rank: 2, name: "b2", id: 9, isOccupied: false},
					{File: B, Rank: 3, name: "b3", id: 10, isOccupied: false},
					{File: B, Rank: 4, name: "b4", id: 11, isOccupied: false},
					{File: B, Rank: 5, name: "b5", id: 12, isOccupied: false},
					{File: B, Rank: 6, name: "b6", id: 13, isOccupied: false},
					{File: B, Rank: 7, name: "b7", id: 14, isOccupied: false},
					{File: B, Rank: 8, name: "b8", id: 15, isOccupied: false},
					{File: C, Rank: 1, name: "c1", id: 16, isOccupied: false},
					{File: C, Rank: 2, name: "c2", id: 17, isOccupied: false},
					{File: C, Rank: 3, name: "c3", id: 18, isOccupied: false},
					{File: C, Rank: 4, name: "c4", id: 19, isOccupied: false},
					{File: C, Rank: 5, name: "c5", id: 20, isOccupied: false},
					{File: C, Rank: 6, name: "c6", id: 21, isOccupied: false},
					{File: C, Rank: 7, name: "c7", id: 22, isOccupied: false},
					{File: C, Rank: 8, name: "c8", id: 23, isOccupied: false},
					{File: D, Rank: 1, name: "d1", id: 24, isOccupied: false},
					{File: D, Rank: 2, name: "d2", id: 25, isOccupied: false},
					{File: D, Rank: 3, name: "d3", id: 26, isOccupied: false},
					{File: D, Rank: 4, name: "d4", id: 27, isOccupied: false},
					{File: D, Rank: 5, name: "d5", id: 28, isOccupied: false},
					{File: D, Rank: 6, name: "d6", id: 29, isOccupied: false},
					{File: D, Rank: 7, name: "d7", id: 30, isOccupied: false},
					{File: D, Rank: 8, name: "d8", id: 31, isOccupied: false},
					{File: E, Rank: 1, name: "e1", id: 32, isOccupied: false},
					{File: E, Rank: 2, name: "e2", id: 33, isOccupied: false},
					{File: E, Rank: 3, name: "e3", id: 34, isOccupied: false},
					{File: E, Rank: 4, name: "e4", id: 35, isOccupied: false},
					{File: E, Rank: 5, name: "e5", id: 36, isOccupied: false},
					{File: E, Rank: 6, name: "e6", id: 37, isOccupied: false},
					{File: E, Rank: 7, name: "e7", id: 38, isOccupied: false},
					{File: E, Rank: 8, name: "e8", id: 39, isOccupied: false},
					{File: F, Rank: 1, name: "f1", id: 40, isOccupied: false},
					{File: F, Rank: 2, name: "f2", id: 41, isOccupied: false},
					{File: F, Rank: 3, name: "f3", id: 42, isOccupied: false},
					{File: F, Rank: 4, name: "f4", id: 43, isOccupied: false},
					{File: F, Rank: 5, name: "f5", id: 44, isOccupied: false},
					{File: F, Rank: 6, name: "f6", id: 45, isOccupied: false},
					{File: F, Rank: 7, name: "f7", id: 46, isOccupied: false},
					{File: F, Rank: 8, name: "f8", id: 47, isOccupied: false},
					{File: G, Rank: 1, name: "g1", id: 48, isOccupied: false},
					{File: G, Rank: 2, name: "g2", id: 49, isOccupied: false},
					{File: G, Rank: 3, name: "g3", id: 50, isOccupied: false},
					{File: G, Rank: 4, name: "g4", id: 51, isOccupied: false},
					{File: G, Rank: 5, name: "g5", id: 52, isOccupied: false},
					{File: G, Rank: 6, name: "g6", id: 53, isOccupied: false},
					{File: G, Rank: 7, name: "g7", id: 54, isOccupied: false},
					{File: G, Rank: 8, name: "g8", id: 55, isOccupied: false},
					{File: H, Rank: 1, name: "h1", id: 56, isOccupied: false},
					{File: H, Rank: 2, name: "h2", id: 57, isOccupied: false},
					{File: H, Rank: 3, name: "h3", id: 58, isOccupied: false},
					{File: H, Rank: 4, name: "h4", id: 59, isOccupied: false},
					{File: H, Rank: 5, name: "h5", id: 60, isOccupied: false},
					{File: H, Rank: 6, name: "h6", id: 61, isOccupied: false},
					{File: H, Rank: 7, name: "h7", id: 62, isOccupied: false},
					{File: H, Rank: 8, name: "h8", id: 63, isOccupied: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewBoard()
			if !assert.Len(t, got.Squares, 64){
				return
			}

			for i := 0; i < 64; i++ {
				assert.Equal(t, got.Squares[i].GetID(), tt.want.Squares[i].GetID())
				assert.Equal(t, got.Squares[i].Name(), tt.want.Squares[i].Name())
				assert.Equal(t, got.Squares[i].IsOccupied(), tt.want.Squares[i].IsOccupied())
				assert.Equal(t, got.Squares[i].File, tt.want.Squares[i].File)
				assert.Equal(t, got.Squares[i].Rank, tt.want.Squares[i].Rank)
			}

		})
	}
}
