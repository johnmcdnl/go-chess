package chess

import (
	"github.com/satori/go.uuid"
)

type Bishop struct {
	BasePiece
}

func NewBishop(s *Square, c Color) (*Bishop, error) {
	var bishop Bishop
	bishop.ID = uuid.NewV4().String()
	bishop.Name = "Bishop"
	bishop.Color = c
	bishop.Position = s
	bishop.Type = BishopType
	return &bishop, nil
}

func (bishop *Bishop) CurrentPosition() *Square {
	return bishop.Position
}

func (bishop *Bishop) PieceColor() Color {
	return bishop.Color
}

func (bishop *Bishop) PieceType() PieceType {
	return bishop.Type
}

func (bishop *Bishop)Move(b *Board, c Coordinate) {
	bishop.Position = b.GetSquare(c.File, c.Rank)
	b.GetSquare(c.File, c.Rank).CurrentPiece = bishop

}

func (bishop *Bishop) ValidMoves(board *Board) []*Move {
	var moves []*Move

	cFile := bishop.CurrentPosition().File
	cRank := bishop.CurrentPosition().Rank

	//up right
	for i := 1; i <= 8; i++ {
		if cFile + i <= 8 && cRank + i <= 8 {
			d := board.GetSquare(cFile + i, cRank + i)
			if m := NewMove(bishop.CurrentPosition(), d); m != nil {
				moves = append(moves, m)

			}
			if d.CurrentPiece != nil {
				break
			}
		}
	}

	//up left
	for i := 1; i <= 8; i++ {
		if cFile - i >= 1 && cRank + i <= 8 {
			d := board.GetSquare(cFile - i, cRank + i)
			if m := NewMove(bishop.CurrentPosition(), d); m != nil {
				moves = append(moves, m)

			}
			if d.CurrentPiece != nil {
				break
			}

		}
	}

	//down right
	for i := 1; i <= 8; i++ {
		if cFile + i <= 8 && cRank - i >= 1 {
			d := board.GetSquare(cFile + i, cRank - i)
			if m := NewMove(bishop.CurrentPosition(), d); m != nil {
				moves = append(moves, m)
			}
			if d.CurrentPiece != nil {
				break
			}
		}
	}

	//down right
	for i := 1; i <= 8; i++ {
		if cFile - i >= 1 && cRank - i >= 1 {
			d := board.GetSquare(cFile - i, cRank - i)
			if m := NewMove(bishop.CurrentPosition(), d); m != nil {
				moves = append(moves, m)
			}
			if d.CurrentPiece != nil {
				break
			}
		}
	}

	return moves
}
