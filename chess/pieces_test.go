package chess

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestValidPieces(t *testing.T) {
	pieces := ValidPieces()
	assert.Len(t, pieces, 12 )

	assert.Contains(t, pieces, WhiteRook)
	assert.Contains(t, pieces, WhiteBishop)
	assert.Contains(t, pieces, WhiteKnight)
	assert.Contains(t, pieces, WhiteQueen)
	assert.Contains(t, pieces, WhiteKing)
	assert.Contains(t, pieces, WhitePawn)

	assert.Contains(t, pieces, BlackRook)
	assert.Contains(t, pieces, BlackBishop)
	assert.Contains(t, pieces, BlackKnight)
	assert.Contains(t, pieces, BlackQueen)
	assert.Contains(t, pieces, BlackKing)
	assert.Contains(t, pieces, BlackPawn)
}
