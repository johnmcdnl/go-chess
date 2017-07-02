package chess

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestBoard_GetPositions(t *testing.T) {
	b := NewBoard()
	positions := b.SANPositions()
	assert.Len(t, positions, 64)

	for _, position := range positions {
		fmt.Sprint(position)

		switch position.Name {
		default:
			assert.Fail(t, "Unknown Position", fmt.Sprintln(position))

		case "A1":
			assert.Equal(t, position.X, 1)
			assert.Equal(t, position.Y, 1)
		case "A2":
			assert.Equal(t, position.X, 1)
			assert.Equal(t, position.Y, 2)
		case "A3":
			assert.Equal(t, position.X, 1)
			assert.Equal(t, position.Y, 3)
		case "A4":
			assert.Equal(t, position.X, 1)
			assert.Equal(t, position.Y, 4)
		case "A5":
			assert.Equal(t, position.X, 1)
			assert.Equal(t, position.Y, 5)
		case "A6":
			assert.Equal(t, position.X, 1)
			assert.Equal(t, position.Y, 6)
		case "A7":
			assert.Equal(t, position.X, 1)
			assert.Equal(t, position.Y, 7)
		case "A8":
			assert.Equal(t, position.X, 1)
			assert.Equal(t, position.Y, 8)

		case "B1":
			assert.Equal(t, position.X, 2)
			assert.Equal(t, position.Y, 1)
		case "B2":
			assert.Equal(t, position.X, 2)
			assert.Equal(t, position.Y, 2)
		case "B3":
			assert.Equal(t, position.X, 2)
			assert.Equal(t, position.Y, 3)
		case "B4":
			assert.Equal(t, position.X, 2)
			assert.Equal(t, position.Y, 4)
		case "B5":
			assert.Equal(t, position.X, 2)
			assert.Equal(t, position.Y, 5)
		case "B6":
			assert.Equal(t, position.X, 2)
			assert.Equal(t, position.Y, 6)
		case "B7":
			assert.Equal(t, position.X, 2)
			assert.Equal(t, position.Y, 7)
		case "B8":
			assert.Equal(t, position.X, 2)
			assert.Equal(t, position.Y, 8)

		case "C1":
			assert.Equal(t, position.X, 3)
			assert.Equal(t, position.Y, 1)
		case "C2":
			assert.Equal(t, position.X, 3)
			assert.Equal(t, position.Y, 2)
		case "C3":
			assert.Equal(t, position.X, 3)
			assert.Equal(t, position.Y, 3)
		case "C4":
			assert.Equal(t, position.X, 3)
			assert.Equal(t, position.Y, 4)
		case "C5":
			assert.Equal(t, position.X, 3)
			assert.Equal(t, position.Y, 5)
		case "C6":
			assert.Equal(t, position.X, 3)
			assert.Equal(t, position.Y, 6)
		case "C7":
			assert.Equal(t, position.X, 3)
			assert.Equal(t, position.Y, 7)
		case "C8":
			assert.Equal(t, position.X, 3)
			assert.Equal(t, position.Y, 8)

		case "D1":
			assert.Equal(t, position.X, 4)
			assert.Equal(t, position.Y, 1)
		case "D2":
			assert.Equal(t, position.X, 4)
			assert.Equal(t, position.Y, 2)
		case "D3":
			assert.Equal(t, position.X, 4)
			assert.Equal(t, position.Y, 3)
		case "D4":
			assert.Equal(t, position.X, 4)
			assert.Equal(t, position.Y, 4)
		case "D5":
			assert.Equal(t, position.X, 4)
			assert.Equal(t, position.Y, 5)
		case "D6":
			assert.Equal(t, position.X, 4)
			assert.Equal(t, position.Y, 6)
		case "D7":
			assert.Equal(t, position.X, 4)
			assert.Equal(t, position.Y, 7)
		case "D8":
			assert.Equal(t, position.X, 4)
			assert.Equal(t, position.Y, 8)

		case "E1":
			assert.Equal(t, position.X, 5)
			assert.Equal(t, position.Y, 1)
		case "E2":
			assert.Equal(t, position.X, 5)
			assert.Equal(t, position.Y, 2)
		case "E3":
			assert.Equal(t, position.X, 5)
			assert.Equal(t, position.Y, 3)
		case "E4":
			assert.Equal(t, position.X, 5)
			assert.Equal(t, position.Y, 4)
		case "E5":
			assert.Equal(t, position.X, 5)
			assert.Equal(t, position.Y, 5)
		case "E6":
			assert.Equal(t, position.X, 5)
			assert.Equal(t, position.Y, 6)
		case "E7":
			assert.Equal(t, position.X, 5)
			assert.Equal(t, position.Y, 7)
		case "E8":
			assert.Equal(t, position.X, 5)
			assert.Equal(t, position.Y, 8)

		case "F1":
			assert.Equal(t, position.X, 6)
			assert.Equal(t, position.Y, 1)
		case "F2":
			assert.Equal(t, position.X, 6)
			assert.Equal(t, position.Y, 2)
		case "F3":
			assert.Equal(t, position.X, 6)
			assert.Equal(t, position.Y, 3)
		case "F4":
			assert.Equal(t, position.X, 6)
			assert.Equal(t, position.Y, 4)
		case "F5":
			assert.Equal(t, position.X, 6)
			assert.Equal(t, position.Y, 5)
		case "F6":
			assert.Equal(t, position.X, 6)
			assert.Equal(t, position.Y, 6)
		case "F7":
			assert.Equal(t, position.X, 6)
			assert.Equal(t, position.Y, 7)
		case "F8":
			assert.Equal(t, position.X, 6)
			assert.Equal(t, position.Y, 8)

		case "G1":
			assert.Equal(t, position.X, 7)
			assert.Equal(t, position.Y, 1)
		case "G2":
			assert.Equal(t, position.X, 7)
			assert.Equal(t, position.Y, 2)
		case "G3":
			assert.Equal(t, position.X, 7)
			assert.Equal(t, position.Y, 3)
		case "G4":
			assert.Equal(t, position.X, 7)
			assert.Equal(t, position.Y, 4)
		case "G5":
			assert.Equal(t, position.X, 7)
			assert.Equal(t, position.Y, 5)
		case "G6":
			assert.Equal(t, position.X, 7)
			assert.Equal(t, position.Y, 6)
		case "G7":
			assert.Equal(t, position.X, 7)
			assert.Equal(t, position.Y, 7)
		case "G8":
			assert.Equal(t, position.X, 7)
			assert.Equal(t, position.Y, 8)

		case "H1":
			assert.Equal(t, position.X, 8)
			assert.Equal(t, position.Y, 1)
		case "H2":
			assert.Equal(t, position.X, 8)
			assert.Equal(t, position.Y, 2)
		case "H3":
			assert.Equal(t, position.X, 8)
			assert.Equal(t, position.Y, 3)
		case "H4":
			assert.Equal(t, position.X, 8)
			assert.Equal(t, position.Y, 4)
		case "H5":
			assert.Equal(t, position.X, 8)
			assert.Equal(t, position.Y, 5)
		case "H6":
			assert.Equal(t, position.X, 8)
			assert.Equal(t, position.Y, 6)
		case "H7":
			assert.Equal(t, position.X, 8)
			assert.Equal(t, position.Y, 7)
		case "H8":
			assert.Equal(t, position.X, 8)
			assert.Equal(t, position.Y, 8)
		}
	}

}
