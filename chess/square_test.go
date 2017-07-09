package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSquare(t *testing.T) {
	a1 := NewSquare(1, 1)
	assert.Equal(t, a1.Name, "a1")
	d4 := NewSquare(4, 4)
	assert.Equal(t, d4.Name, "d4")
	h8 := NewSquare(8, 8)
	assert.Equal(t, h8.Name, "h8")
}
