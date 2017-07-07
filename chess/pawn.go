package chess

type Pawn struct {
	*Piece
}

func NewPawn(currentPosition *Square, color int) *Pawn {
	return &Pawn{&Piece{
		Type:            PieceTypePawn,
		Color:           color,
		CurrentPosition: currentPosition,
	},
	}
}

func (p *Pawn) ValidMoves() *Moves {
	switch p.Color {
	default:
		return nil
	case White :
		return p.whitePawnMoves()
	case Black :
		return p.blackPawnMoves()
	}
}

func (p *Pawn)whitePawnMoves() *Moves {
	var moves = new(Moves)
	var m = Move{
		From: p.CurrentPosition,
	}

	cFile := p.CurrentPosition.FileLetter
	cRank := p.CurrentPosition.RankNumber

	if cRank == 2 {
		m.To = GlobalBoard.GetSquare(cFile, cRank + 2)
		moves.BestMoves = append(moves.BestMoves, m)
	}

	if cRank + 1 <= 8 {
		m.To = GlobalBoard.GetSquare(cFile, cRank + 1)
		moves.BestMoves = append(moves.BestMoves, m)
	}

	return moves
}

func (p *Pawn)blackPawnMoves() *Moves {
	var moves = new(Moves)
	var m = Move{
		From: p.CurrentPosition,
	}

	cFile := p.CurrentPosition.FileLetter
	cRank := p.CurrentPosition.RankNumber

	if cRank == 7 {
		m.To = GlobalBoard.GetSquare(cFile, cRank - 2)
		moves.BestMoves = append(moves.BestMoves, m)
	}

	if cRank - 1 >= 1 {
		m.To = GlobalBoard.GetSquare(cFile, cRank - 1)
		moves.BestMoves = append(moves.BestMoves, m)
	}

	return moves
}

func (p *Pawn) GetPiece() *Piece {
	return p.Piece
}