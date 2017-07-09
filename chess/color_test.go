package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColors(t *testing.T) {
	assert.Equal(t, White, Color(0))
	assert.Equal(t, Black, Color(1))
	assert.NotEqual(t, White, Black)
}
