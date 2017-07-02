package chess

import (
	"fmt"
	"strconv"
)

type Board struct {
	coordinates []Coordinate
}

type SANPosition struct {
	Name string
	Coordinate
}

type Coordinate struct {
	X int
	Y int
}

func NewBoard() *Board {
	var b Board
	coordinates := []Coordinate{}
	for yRank := 1; yRank <= 8; yRank++ {
		for xFile := 1; xFile <= 8; xFile++ {
			coordinates = append(coordinates, Coordinate{xFile, yRank})
		}
	}
	b.coordinates = coordinates
	return &b
}

func (b *Board) SANPositions() []SANPosition {
	coordinates := b.coordinates
	var sanPositions []SANPosition
	for _, coordinate := range coordinates {
		var s SANPosition
		s.Coordinate = coordinate
		s.Name = fmt.Sprintf("%s%s", string('A'-1+coordinate.X), strconv.Itoa(coordinate.Y))
		sanPositions = append(sanPositions, s)
	}
	return sanPositions
}
