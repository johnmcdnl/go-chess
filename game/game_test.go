package game

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGame_SetInitialState(t *testing.T) {
	var g Game
	g.SetInitialState()
	fen := g.FEN()
	assert.Equal(t, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",fen )
}