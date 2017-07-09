package chess

type Coordinate struct {
	File int
	Rank int
}

func NewCoordinate(file int, rank int) *Coordinate {
	var c Coordinate
	c.File = file
	c.Rank = rank
	return &c
}
