package chess

type Pawn Piece

func NewPawn(s *Square, c Color) *Pawn {
	var b Pawn
	b.Name = "Pawn"
	b.Code = PawnPiece
	b.Position = s
	b.Color = c

	return &b
}

func (p *Pawn) ValidMoves(b *Board) []*Move {
	switch p.Color {
	default:
		return []*Move{}
	case White:
		return p.whiteMoves(b)
	case Black:
		return p.blackMoves(b)
	}
}

func (p *Pawn) whiteMoves(b *Board) []*Move {
	var moves []*Move

	pos := p.Position
	currentFile := pos.File
	currentRank := pos.Rank

	if currentRank+1 < 8 {
		if m, err := NewMove(pos, b.GetSquare(currentFile, currentRank+1)); err == nil {
			moves = append(moves, m)
		}
	}

	if currentRank == 2 {
		if m, err := NewMove(pos, b.GetSquare(currentFile, currentRank+2)); err == nil {
			moves = append(moves, m)
		}
	}

	return moves
}

func (p *Pawn) blackMoves(b *Board) []*Move {
	var moves []*Move

	pos := p.Position
	currentFile := pos.File
	currentRank := pos.Rank

	if currentRank-1 > 0 {
		if m, err := NewMove(pos, b.GetSquare(currentFile, currentRank-1)); err == nil {
			moves = append(moves, m)
		}
	}

	if currentRank == 7 {
		if m, err := NewMove(pos, b.GetSquare(currentFile, currentRank-2)); err == nil {
			moves = append(moves, m)
		}
	}

	return moves
}

func (p *Pawn) GetCode() int {
	return p.Code
}

func (p *Pawn) GetColor() Color {
	return p.Color
}

func (p *Pawn) CurrentPosition() *Square {
	return p.Position
}
