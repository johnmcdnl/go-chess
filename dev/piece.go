package dev

type Color int

const White = Color(0)
const Black = Color(1)

type PieceType int

const King = PieceType(1)
const Queen = PieceType(2)
const Rook = PieceType(3)
const Bishop = PieceType(4)
const Knight = PieceType(5)
const Pawn = PieceType(6)

type Move struct {

}

type BasePiece struct {
	name string
	color Color
	pieceType PieceType
}

type Piece interface {
	GetName() string
	GetColor() Color
	PieceType() PieceType
	ValidMoves() []*Move
}