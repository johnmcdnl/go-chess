package dev

type pieceForTestImplemented struct{}

func newPieceForTestImplemented() Piece {
	return pieceForTestImplemented{}
}
func (p pieceForTestImplemented)GetName() string {
	return "pieceImplemented"
}

func (p pieceForTestImplemented)GetColor() Color {
	return White
}

func (p pieceForTestImplemented)PieceType() PieceType {
	return Queen
}