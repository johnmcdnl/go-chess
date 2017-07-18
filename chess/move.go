package chess

import "fmt"

type Move struct {
	Origin      *Square
	Destination *Square
}

//New returns a Move or and error if the provided squares are invalid
func NewMove(origin, destination *Square) (*Move) {

	if origin==nil || destination ==nil{
		return nil
	}

	var m Move
	m.Origin = origin
	m.Destination = destination
	return &m
}

func (m *Move)Printer() {
	fmt.Println("Origin: ", m.Origin.File, m.Origin.Rank, "Destination: ", m.Destination.File, m.Destination.Rank)
}

func (m *Move)IsValid() bool {
	return true
}

func (m *Move)Apply() {

}