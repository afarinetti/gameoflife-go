package main

type CellState bool

const (
	Dead  CellState = false
	Alive           = true
)

func (state CellState) String() string {
	switch state {
	case Dead:
		return "DEAD"
	case Alive:
		return "ALIVE"
	default:
		return "UNKNOWN"
	}
}
