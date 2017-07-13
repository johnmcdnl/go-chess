package dev

type pieceForTestImplemented struct{}

func newPieceForTestImplemented() Piece {
	return pieceForTestImplemented{}
}
func (p pieceForTestImplemented)GetName() string {
	return "pieceImplemented"
}
