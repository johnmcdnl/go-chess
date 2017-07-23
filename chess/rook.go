package chess

import "github.com/satori/go.uuid"

type Rook struct {
	BasePiece
}

func NewRook(s *Square, c Color) (*Rook, error) {
	var r Rook
	r.ID = uuid.NewV4().String()
	r.Name = "Rook"
	r.Color = c
	r.Position = s
	r.Type = RookType

	s.SetPiece(&r)
	return &r, nil
}

func (r*Rook) CurrentPosition() *Square {
	return r.Position
}

func (r *Rook)PieceColor() Color {
	return r.Color
}

func (r *Rook)PieceType() PieceType {
	return r.Type
}

func (r *Rook) ValidMoves(b *Board) []*Move {
	var moves []*Move

	//up
	for rank := r.Position.Rank + 1; rank <= 8; rank++ {
		d := b.GetSquare(r.CurrentPosition().File, rank)
		if m := NewMove(r.CurrentPosition(), d); m != nil {
			moves = append(moves, m)
		}
		if d.CurrentPiece != nil {
			break
		}
	}

	//up
	for rank := r.Position.Rank - 1; rank >= 1; rank-- {
		d := b.GetSquare(r.CurrentPosition().File, rank)
		if m := NewMove(r.CurrentPosition(), d); m != nil {
			moves = append(moves, m)
		}
		if d.CurrentPiece != nil {
			break
		}
	}


	//right
	for file := r.Position.File + 1; file <= 8; file++ {
		d := b.GetSquare(file, r.CurrentPosition().Rank)
		if m := NewMove(r.CurrentPosition(), d); m != nil {
			moves = append(moves, m)
		}
		if d.CurrentPiece != nil {
			break
		}
	}

	//right
	for file := r.Position.File - 1; file >= 1; file-- {
		d := b.GetSquare(file, r.CurrentPosition().Rank)
		if m := NewMove(r.CurrentPosition(), d); m != nil {
			moves = append(moves, m)
		}
		if d.CurrentPiece != nil {
			break
		}
	}

	return moves
}

