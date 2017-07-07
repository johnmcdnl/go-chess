package chess

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestKnight_ValidMoves_MiddleBoard(t *testing.T) {
	e5 := GlobalBoard.GetSquare(5, 5)
	e5.CurrentPiece = NewKnight(e5, White)
	bestMoves := e5.CurrentPiece.ValidMoves().BestMoves

	validMoves := []string{"e5c6","e5d7","e5f7","e5g6","e5g4","e5f3","e5d3","e5c4"}
	assert.Len(t, bestMoves, len(validMoves))
	for _, foundMove := range bestMoves {
		assert.Contains(t, validMoves, foundMove.Name())
	}
}

func TestKnight_ValidMoves_NearCorner(t *testing.T) {
	b2 := GlobalBoard.GetSquare(2, 2)
	b2.CurrentPiece = NewKnight(b2, White)
	bestMoves := b2.CurrentPiece.ValidMoves().BestMoves

	validMoves := []string{"b2a4","b2c4","b2d3","b2d1"}
	assert.Len(t, bestMoves, len(validMoves))
	for _, foundMove := range bestMoves {
		assert.Contains(t, validMoves, foundMove.Name())
	}
}
