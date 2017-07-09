package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	var b = NewBoard()
	assert.NotNil(t, b.Squares)
	assert.Len(t, *b.Squares, 64)
}
