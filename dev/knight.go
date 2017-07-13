package dev

type Knight struct {
	BasePiece
}

func (k *Knight)GetName() string {
	return k.name
}

func (k *Knight)GetColor() Color {
	return k.color
}

func (k *Knight)PieceType() PieceType {
	return k.pieceType
}

func (k *Knight)ValidMoves() []*Move {
	return k.getValidMoves()
}

func (k *Knight)getValidMoves() []*Move {
	var moves []*Move
	return moves
}