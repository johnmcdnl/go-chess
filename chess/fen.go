package chess

type FEN string

const StartingPositionFEN = FEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

func (f *FEN) IsValid() bool {
	return true
}
