package dev

type Color int

const White = Color(0)
const Black = Color(1)

type PieceType int

const KingPiece = PieceType(1)
const QueenPiece = PieceType(2)
const RookPiece = PieceType(3)
const BishopPiece = PieceType(4)
const KnightPiece = PieceType(5)
const PawnPiece = PieceType(6)

type Move struct {

}

type BasePiece struct {
	name      string
	color     Color
	pieceType PieceType
}

type Piece interface {
	GetName() string
	GetColor() Color
	PieceType() PieceType
	ValidMoves() []*Move
}