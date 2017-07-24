package chess

import (
	"fmt"
)

type CastlingRights struct {
	WhiteKingSideAvailable  bool
	WhiteQueenSideAvailable bool
	BlackKingSideAvailable  bool
	BlackQueenSideAvailable bool
}

type Board struct {
	Squares         []*Square
	ActiveColor     Color
	CastlingRights  *CastlingRights
	EnPassantSquare *Square
	HalfMoveClock   int
	FullMoveNumber  int
}

func NewEmptyBoard() (*Board, error) {
	var b Board

	if err := b.build(8, 8); err != nil {
		return nil, err
	}
	b.ActiveColor = White
	b.CastlingRights = &CastlingRights{
		WhiteKingSideAvailable:true,
		WhiteQueenSideAvailable:true,
		BlackKingSideAvailable:true,
		BlackQueenSideAvailable:true,
	}
	b.EnPassantSquare = nil
	b.HalfMoveClock = 0
	b.FullMoveNumber = 1

	return &b, nil
}

func NewBoard() (*Board, error) {
	var b Board

	if err := b.build(8, 8); err != nil {
		return nil, err
	}
	if err := b.LoadFromFEN(StartingPositionFEN); err != nil {
		return nil, err
	}

	return &b, nil
}

func BoardFEN(fen FEN) (*Board, error) {
	var b Board

	if err := b.build(8, 8); err != nil {
		return nil, err
	}
	if err := b.LoadFromFEN(fen); err != nil {
		return nil, err
	}

	return &b, nil
}

func (b *Board) build(files, ranks int) error {
	for rank := 1; rank <= ranks; rank++ {
		for file := 1; file <= files; file++ {
			s, err := NewSquare(file, rank)
			if err != nil {
				return err
			}
			b.Squares = append(b.Squares, s)
		}
	}
	return nil
}

func (b *Board) LoadFromFEN(f FEN) error {
	if !f.IsValid() {
		return fmt.Errorf("Invalid FEN %s", f)
	}
	f.Apply(b)
	return nil
}

func (b *Board) GetSquare(file, rank int) *Square {

	if file < 1 || file > 8 || rank < 1 || rank > 8 {
		return nil
	}

	id := ((rank - 1) * 8) + file - 1
	if id < 0 || id > 63 {
		return nil
	}

	return b.Squares[id]
}

func (b *Board)Print() {
	rank8 := fmt.Sprintf("8  |%s|%s|%s|%s|%s|%s|%s|%s|", sp(b.Squares[56]), sp(b.Squares[57]), sp(b.Squares[58]), sp(b.Squares[59]), sp(b.Squares[60]), sp(b.Squares[61]), sp(b.Squares[62]), sp(b.Squares[63]))
	rank7 := fmt.Sprintf("7  |%s|%s|%s|%s|%s|%s|%s|%s|", sp(b.Squares[48]), sp(b.Squares[49]), sp(b.Squares[50]), sp(b.Squares[51]), sp(b.Squares[52]), sp(b.Squares[53]), sp(b.Squares[54]), sp(b.Squares[55]))
	rank6 := fmt.Sprintf("6  |%s|%s|%s|%s|%s|%s|%s|%s|", sp(b.Squares[40]), sp(b.Squares[41]), sp(b.Squares[42]), sp(b.Squares[43]), sp(b.Squares[44]), sp(b.Squares[45]), sp(b.Squares[46]), sp(b.Squares[47]))
	rank5 := fmt.Sprintf("5  |%s|%s|%s|%s|%s|%s|%s|%s|", sp(b.Squares[32]), sp(b.Squares[33]), sp(b.Squares[34]), sp(b.Squares[35]), sp(b.Squares[36]), sp(b.Squares[37]), sp(b.Squares[38]), sp(b.Squares[39]))
	rank4 := fmt.Sprintf("4  |%s|%s|%s|%s|%s|%s|%s|%s|", sp(b.Squares[24]), sp(b.Squares[25]), sp(b.Squares[26]), sp(b.Squares[27]), sp(b.Squares[28]), sp(b.Squares[29]), sp(b.Squares[30]), sp(b.Squares[31]))
	rank3 := fmt.Sprintf("3  |%s|%s|%s|%s|%s|%s|%s|%s|", sp(b.Squares[16]), sp(b.Squares[17]), sp(b.Squares[18]), sp(b.Squares[19]), sp(b.Squares[20]), sp(b.Squares[21]), sp(b.Squares[22]), sp(b.Squares[23]))
	rank2 := fmt.Sprintf("2  |%s|%s|%s|%s|%s|%s|%s|%s|", sp(b.Squares[8]), sp(b.Squares[9]), sp(b.Squares[10]), sp(b.Squares[11]), sp(b.Squares[12]), sp(b.Squares[13]), sp(b.Squares[14]), sp(b.Squares[15]))
	rank1 := fmt.Sprintf("1  |%s|%s|%s|%s|%s|%s|%s|%s|", sp(b.Squares[0]), sp(b.Squares[1]), sp(b.Squares[2]), sp(b.Squares[3]), sp(b.Squares[4]), sp(b.Squares[5]), sp(b.Squares[6]), sp(b.Squares[7]))
	//empty := fmt.Sprintf(" %s %s %s %s %s %s %s %s|", "   ", "   ", "   ", "   ", "   ", "   ", "   ", "   ")
	files := fmt.Sprintf("    %s %s %s %s %s %s %s %s ", " A ", " B ", " C ", " D ", " E ", " F ", " G ", " H ")
	fmt.Println(rank8)
	fmt.Println(rank7)
	fmt.Println(rank6)
	fmt.Println(rank5)
	fmt.Println(rank4)
	fmt.Println(rank3)
	fmt.Println(rank2)
	fmt.Println(rank1)
	//fmt.Println(empty)
	fmt.Println(files)
}

func sp(s *Square) string {
	if s.CurrentPiece == nil {
		return "   "
	}
	switch s.CurrentPiece.PieceType() {
	case KingType:
		if s.CurrentPiece.PieceColor() == White {
			return " K "
		} else {
			return " k "
		}
	case QueenType:
		if s.CurrentPiece.PieceColor() == White {
			return " Q "
		} else {
			return " q "
		}
	case RookType:
		if s.CurrentPiece.PieceColor() == White {
			return " R "
		} else {
			return " r "
		}
	case BishopType:
		if s.CurrentPiece.PieceColor() == White {
			return " B "
		} else {
			return " b "
		}
	case KnightType:
		if s.CurrentPiece.PieceColor() == White {
			return " N "
		} else {
			return " n "
		}
	case PawnType:
		if s.CurrentPiece.PieceColor() == White {
			return " P "
		} else {
			return " p "
		}

	}
	return "   "
}