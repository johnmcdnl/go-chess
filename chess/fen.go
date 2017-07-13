package chess

type FEN string

const StartingPositionFEN = FEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
const OnlyKingsFEN = FEN("4k3/8/8/8/8/8/8/4K3 w - - 0 1")
const BlockedE5FEN = FEN("4k3/8/8/4p3/3B4/8/8/4K3 w - - 0 1")

func (f *FEN) IsValid() bool {
	return true
}
