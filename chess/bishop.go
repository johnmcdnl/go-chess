package chess

type Bishop Piece

func NewBishop(s *Square, c Color) *Bishop {
	var b Bishop
	b.Name = "bishop"
	b.Code = BishopPiece
	b.Position = s
	b.Color = c

	s.ChessPiece = &b
	return &b
}

func (bishop *Bishop) ValidMoves(board *Board) []*Move {
	var moves []*Move

	cRank := bishop.Position.Rank
	cFile := bishop.Position.File

	//up left
	for i := 1; i < 8; i++ {
		tFile := cFile - i
		tRank := cRank + i
		if tFile <= 1 && tRank <= 1 && tFile >= 8 && tRank >= 8 {
			break
		}
		m, err := NewMove(bishop.Position, board.GetSquare(tFile, tRank))
		if err != nil {
			break
		}
		moves = append(moves, m)
	}

	//up right
	for i := 1; i < 8; i++ {
		tFile := cFile + i
		tRank := cRank + i
		if tFile <= 1 && tRank <= 1 && tFile >= 8 && tRank >= 8 {
			break
		}
		m, err := NewMove(bishop.Position, board.GetSquare(tFile, tRank))
		if err != nil {
			break
		}
		moves = append(moves, m)
	}

	//down left
	for i := 1; i < 8; i++ {
		tFile := cFile - i
		tRank := cRank - i
		if tFile <= 1 && tRank <= 1 && tFile >= 8 && tRank >= 8 {
			break
		}
		m, err := NewMove(bishop.Position, board.GetSquare(tFile, tRank))
		if err != nil {
			break
		}
		moves = append(moves, m)
	}

	//down right
	for i := 1; i < 8; i++ {
		tFile := cFile + i
		tRank := cRank - i
		if tFile <= 1 && tRank <= 1 && tFile >= 8 && tRank >= 8 {
			break
		}
		m, err := NewMove(bishop.Position, board.GetSquare(tFile, tRank))
		if err != nil {
			break
		}
		moves = append(moves, m)
	}

	return moves
}
func (bishop *Bishop) GetCode() int {
	return bishop.Code
}

func (bishop *Bishop) GetColor() Color {
	return bishop.Color
}

func (bishop *Bishop) CurrentPosition() *Square {
	return bishop.Position
}
