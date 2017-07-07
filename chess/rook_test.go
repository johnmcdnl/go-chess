package chess

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestRook_ValidMoves_MiddleBoard(t *testing.T) {
	e5 := GlobalBoard.GetSquare(5, 5)
	e5.CurrentPiece = NewRook(e5, White)
	bestMoves := e5.CurrentPiece.ValidMoves().BestMoves

	validMoves := []string{"e5e1","e5e2","e5e3","e5e4","e5e6","e5e7","e5e8", "e5a5","e5b5","e5c5","e5d5","e5f5","e5g5","e5h5"}
	assert.Len(t, bestMoves, len(validMoves))
	for _, foundMove := range bestMoves {
		assert.Contains(t, validMoves, foundMove.Name())
	}
	assert.NotEqual(t, bestMoves[0].Name(), bestMoves[1].Name(), fmt.Sprintln(bestMoves[0].Name(), bestMoves[1].Name()))
}

func TestRook_ValidMoves_CornerBoard(t *testing.T) {
	a1 := GlobalBoard.GetSquare(1, 1)
	a1.CurrentPiece = NewRook(a1, White)
	bestMoves := a1.CurrentPiece.ValidMoves().BestMoves

	validMoves := []string{"a1a2","a1a3","a1a4","a1a5","a1a6","a1a7","a1a8", "a1b1","a1c1","a1e1","a1d1","a1f1","a1g1","a1h1"}
	assert.Len(t, bestMoves, len(validMoves))
	for _, foundMove := range bestMoves {
		assert.Contains(t, validMoves, foundMove.Name())
	}
	assert.NotEqual(t, bestMoves[0].Name(), bestMoves[1].Name(), fmt.Sprintln(bestMoves[0].Name(), bestMoves[1].Name()))
}