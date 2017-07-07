package chess

type Rook struct {
	*Piece
}

func NewRook(currentPosition *Square, color int) *Rook {
	return &Rook{&Piece{
		Type:            PieceTypeRook,
		Color:           color,
		CurrentPosition: currentPosition,
	},
	}
}

func (r *Rook) ValidMoves() *Moves {
	var moves = new(Moves)
	var m = Move{
		From: r.CurrentPosition,
	}
	//vertical
	currentFile := r.GetPiece().CurrentPosition.FileLetter
	for i := 1; i <= 8; i++ {
		if i != currentFile {
			m.To = GlobalBoard.GetSquare(i, r.CurrentPosition.RankNumber)
			moves.BestMoves = append(moves.BestMoves, &m)
		}
	}

	//vertical
	currentRank := r.GetPiece().CurrentPosition.RankNumber
	for i := 1; i <= 8; i++ {
		if i != currentRank {
			m.To = GlobalBoard.GetSquare(r.CurrentPosition.FileLetter, i)
			moves.BestMoves = append(moves.BestMoves, &m)
		}
	}

	return moves
}

func (r *Rook) GetPiece() *Piece {
	return r.Piece
}
