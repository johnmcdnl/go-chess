package chess

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func init() {
	NewBoard()

}

func TestBishop_ValidMoves(t *testing.T) {
	e5 := GlobalBoard.GetSquare(5, 5)
	e5.CurrentPiece = NewBishop(e5, White)
	bestMoves := e5.CurrentPiece.ValidMoves().BestMoves

	validMoves := []string{"e5b8", "e5c7", "e5d6", "e5f4", "e5g3", "e5h2", "e5a1", "e5b2", "e5c3", "e5d4", "e5f6", "e5g7", "e5h8"}
	assert.Len(t, bestMoves, len(validMoves))
	for _, foundMove := range bestMoves {
		assert.Contains(t, validMoves, foundMove.Name())
	}
	assert.NotEqual(t, bestMoves[0].Name(), bestMoves[1].Name(), fmt.Sprintln(bestMoves[0].Name(), bestMoves[1].Name()))

}