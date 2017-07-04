package game

type Game struct {

}

func NewGame() *Game {
	return &Game{}
}

func (g *Game)SetInitialState() {

}

func (g *Game)FEN() string {
	return ""
}