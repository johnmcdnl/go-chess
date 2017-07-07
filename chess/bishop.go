package chess

type Bishop struct {
	*Piece
}

func NewBishop(currentPosition *Square, color int) *Bishop {
	return &Bishop{&Piece{
		Type:            PieceTypeBishop,
		Color:           color,
		CurrentPosition: currentPosition,
	},
	}
}
func (b *Bishop) ValidMoves() *Moves {
	var moves = new(Moves)
	moves.CurrentSquare = b.CurrentPosition

	//left up      file down, rank up
	for x := 1; x <= 8; x++ {
		cFile := b.CurrentPosition.FileLetter
		cRank := b.CurrentPosition.RankNumber

		tFile := cFile - x
		tRank := cRank + x

		if tFile > 0 && tFile <= 8 && tRank > 0 && tRank <= 8 {
			s := GlobalBoard.GetSquare(tFile, tRank)
			var m = Move{
				From: b.CurrentPosition,
				To:   s,
			}
			moves.BestMoves = append(moves.BestMoves, &m)
		}
	}
	//right up      file up, rank up
	for x := 1; x <= 8; x++ {
		cFile := b.CurrentPosition.FileLetter
		cRank := b.CurrentPosition.RankNumber

		tFile := cFile + x
		tRank := cRank + x

		if tFile > 0 && tFile <= 8 && tRank > 0 && tRank <= 8 {
			s := GlobalBoard.GetSquare(tFile, tRank)
			var m = Move{
				From: b.CurrentPosition,
				To:   s,
			}
			moves.BestMoves = append(moves.BestMoves, &m)
		}
	}

	//left down      file down, rank down
	for x := 1; x <= 8; x++ {
		cFile := b.CurrentPosition.FileLetter
		cRank := b.CurrentPosition.RankNumber

		tFile := cFile - x
		tRank := cRank - x

		if tFile > 0 && tFile <= 8 && tRank > 0 && tRank <= 8 {
			s := GlobalBoard.GetSquare(tFile, tRank)
			var m = Move{
				From: b.CurrentPosition,
				To:   s,
			}
			moves.BestMoves = append(moves.BestMoves, &m)
		}
	}

	//right down      file up, rank down
	for x := 1; x <= 8; x++ {
		cFile := b.CurrentPosition.FileLetter
		cRank := b.CurrentPosition.RankNumber

		tFile := cFile + x
		tRank := cRank - x

		if tFile > 0 && tFile <= 8 && tRank > 0 && tRank <= 8 {
			s := GlobalBoard.GetSquare(tFile, tRank)
			var m = Move{
				From: b.CurrentPosition,
				To:   s,
			}
			moves.BestMoves = append(moves.BestMoves, &m)
		}
	}

	return moves
}
func (r *Bishop) GetPiece() *Piece {
	return r.Piece
}
