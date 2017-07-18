package chess


type Square struct {
	File         int
	Rank         int
	CurrentPiece Piece
}

func NewSquare(file, rank int) (*Square, error) {
	var s Square
	s.File = file
	s.Rank = rank
	return &s, nil
}

func (s *Square)SetPiece(p Piece) {
	s.CurrentPiece=p
}