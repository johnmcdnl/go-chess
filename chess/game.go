package chess

type Game struct {
	Players [2]*Player
	Board   *Board
}

func NewGame() *Game {
	var g Game

	g.Board, _ = BoardFEN(StartingPositionFEN)
	g.Players[0] = NewPlayer("WhiteP1", White)
	g.Players[1] = NewPlayer("BlackP2", Black)

	return &g
}
