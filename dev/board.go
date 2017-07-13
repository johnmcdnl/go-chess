package dev

type Board struct {
	Squares []*Square
}

func NewBoard() *Board {
	var b Board
	for file := A; file <= H; file++ {
		for rank := 1; rank <= 8; rank++ {
			s, err := NewSquare(file, rank)
			if err != nil {
				panic(err)
			}
			b.Squares = append(b.Squares, s)
		}
	}
	return &b
}
