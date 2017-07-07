package chess

type Knight struct {
	*Piece
}

func NewKnight(currentPosition *Square, color int) *Knight {
	return &Knight{&Piece{
		Type:            PieceTypeKnight,
		Color:           color,
		CurrentPosition: currentPosition,
	},
	}
}
func (k *Knight) ValidMoves() *Moves {
	var moves = new(Moves)
	var m = Move{
		From: k.CurrentPosition,
	}
	cFile := k.CurrentPosition.FileLetter
	cRank := k.CurrentPosition.RankNumber

	//RankUp2FileUp1
	tFile := cFile + 1
	tRank := cRank + 2
	if tFile > 0 && tFile <= 8 && tRank > 0 && tRank <= 8 {
		m.To = GlobalBoard.GetSquare(tFile, tRank)
		moves.BestMoves = append(moves.BestMoves, m)
	}

	//RankUp2FileDown1
	tFile = cFile - 1
	tRank = cRank + 2
	if tFile > 0 && tFile <= 8 && tRank > 0 && tRank <= 8 {
		m.To = GlobalBoard.GetSquare(tFile, tRank)
		moves.BestMoves = append(moves.BestMoves, m)
	}


	//RankDown2FileUp1
	tFile = cFile + 1
	tRank = cRank -2
	if tFile > 0 && tFile <= 8 && tRank > 0 && tRank <= 8 {
		m.To = GlobalBoard.GetSquare(tFile, tRank)
		moves.BestMoves = append(moves.BestMoves, m)
	}
	//RankDown2FileUp1
	tFile = cFile + 1
	tRank = cRank - 2
	if tFile > 0 && tFile <= 8 && tRank > 0 && tRank <= 8 {
		m.To = GlobalBoard.GetSquare(tFile, tRank)
		moves.BestMoves = append(moves.BestMoves, m)
	}


	//FileDown2RankUp1
	tFile = cFile - 2
	tRank = cRank + 1
	if tFile > 0 && tFile <= 8 && tRank > 0 && tRank <= 8 {
		m.To = GlobalBoard.GetSquare(tFile, tRank)
		moves.BestMoves = append(moves.BestMoves, m)
	}
	//FileDown2RankDown1
	tFile = cFile - 2
	tRank = cRank - 1
	if tFile > 0 && tFile <= 8 && tRank > 0 && tRank <= 8 {
		m.To = GlobalBoard.GetSquare(tFile, tRank)
		moves.BestMoves = append(moves.BestMoves, m)
	}

	//FileUp2RankUp1
	tFile = cFile +2
	tRank = cRank +1
	if tFile > 0 && tFile <= 8 && tRank > 0 && tRank <= 8 {
		m.To = GlobalBoard.GetSquare(tFile, tRank)
		moves.BestMoves = append(moves.BestMoves, m)
	}
	//FileUp2RankDown1
	tFile = cFile +2
	tRank = cRank -1
	if tFile > 0 && tFile <= 8 && tRank > 0 && tRank <= 8 {
		m.To = GlobalBoard.GetSquare(tFile, tRank)
		moves.BestMoves = append(moves.BestMoves, m)
	}


	return moves
}
func (r *Knight) GetPiece() *Piece {
	return r.Piece
}
