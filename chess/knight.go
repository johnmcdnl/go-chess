package chess

type Knight Piece

func NewKnight(s *Square, c Color) *Knight {
	var k Knight
	k.Name = "knight"
	k.Code = KnightPiece
	k.Color = c
	k.Position = s
	return &k
}

func (k *Knight) ValidMoves(b *Board) []*Move {
	var moves []*Move

	pos := k.Position
	currentFile := pos.File
	currentRank := pos.Rank

	if m, err := NewMove(pos, b.GetSquare(currentFile+2, currentRank+1)); err == nil {
		moves = append(moves, m)
	}

	if m, err := NewMove(pos, b.GetSquare(currentFile+2, currentRank-1)); err == nil {
		moves = append(moves, m)
	}

	if m, err := NewMove(pos, b.GetSquare(currentFile+1, currentRank+2)); err == nil {
		moves = append(moves, m)
	}
	if m, err := NewMove(pos, b.GetSquare(currentFile+1, currentRank-2)); err == nil {
		moves = append(moves, m)
	}
	if m, err := NewMove(pos, b.GetSquare(currentFile-2, currentRank+1)); err == nil {
		moves = append(moves, m)
	}
	if m, err := NewMove(pos, b.GetSquare(currentFile-2, currentRank-1)); err == nil {
		moves = append(moves, m)
	}
	if m, err := NewMove(pos, b.GetSquare(currentFile-1, currentRank+2)); err == nil {
		moves = append(moves, m)
	}
	if m, err := NewMove(pos, b.GetSquare(currentFile-1, currentRank-2)); err == nil {
		moves = append(moves, m)
	}
	return moves
}

func (k *Knight) GetCode() int {
	return k.Code
}

func (k *Knight) GetColor() Color {
	return k.Color
}

func (k *Knight) CurrentPosition() *Square {
	return k.Position
}
