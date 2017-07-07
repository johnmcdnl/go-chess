package chess

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestPawn_ValidMoves_WhiteOpening(t *testing.T) {
	e2 := GlobalBoard.GetSquare(5, 2)
	e2.CurrentPiece = NewPawn(e2, White)
	bestMoves := e2.CurrentPiece.ValidMoves().BestMoves

	validMoves := []string{"e2e4","e2e3"}
	assert.Len(t, bestMoves, len(validMoves))
	for _, foundMove := range bestMoves {
		assert.Contains(t, validMoves, foundMove.Name())
	}
	assert.NotEqual(t, bestMoves[0].Name(), bestMoves[1].Name(), fmt.Sprintln(bestMoves[0].Name(), bestMoves[1].Name()))
}

func TestPawn_ValidMoves_WhiteMidGame(t *testing.T) {
	e3 := GlobalBoard.GetSquare(5, 3)
	e3.CurrentPiece = NewPawn(e3, White)
	bestMoves := e3.CurrentPiece.ValidMoves().BestMoves

	validMoves := []string{"e3e4"}
	assert.Len(t, bestMoves, len(validMoves))
	for _, foundMove := range bestMoves {
		assert.Contains(t, validMoves, foundMove.Name())
	}
}

func TestPawn_ValidMoves_BlackOpening(t *testing.T) {
	e7 := GlobalBoard.GetSquare(5, 7)
	e7.CurrentPiece = NewPawn(e7, Black)
	bestMoves := e7.CurrentPiece.ValidMoves().BestMoves

	validMoves := []string{"e7e6","e7e5"}
	assert.Len(t, bestMoves, len(validMoves))
	for _, foundMove := range bestMoves {
		assert.Contains(t, validMoves, foundMove.Name())
	}
	assert.NotEqual(t, bestMoves[0].Name(), bestMoves[1].Name(), fmt.Sprintln(bestMoves[0].Name(), bestMoves[1].Name()))
}

func TestPawn_ValidMoves_BlackMidGame(t *testing.T) {
	e6 := GlobalBoard.GetSquare(5, 6)
	e6.CurrentPiece = NewPawn(e6, Black)
	bestMoves := e6.CurrentPiece.ValidMoves().BestMoves

	validMoves := []string{"e6e5"}
	assert.Len(t, bestMoves, len(validMoves))
	for _, foundMove := range bestMoves {
		assert.Contains(t, validMoves, foundMove.Name())
	}
}