package main

import "strings"

type Grid struct {
	NumCols uint
	NumRows uint
	data    []CellState
}

func NewGrid(numCols uint, numRows uint) *Grid {
	data := make([]CellState, numRows*numCols)
	return &Grid{NumCols: numCols, NumRows: numRows, data: data}
}

func NewGridWithInit(numCols uint, numRows uint, initGrid []uint8) *Grid {
	data := make([]CellState, numRows*numCols)
	for i, v := range initGrid {
		if v == 1 {
			data[i] = Alive
		} else {
			data[i] = Dead
		}
	}

	return &Grid{NumCols: numCols, NumRows: numRows, data: data}
}

func (g *Grid) cellToIndex(row uint, col uint) uint {
	return (row * g.NumCols) + col
}

func (g *Grid) Get(row uint, col uint) CellState {
	return g.data[g.cellToIndex(row, col)]
}

func (g *Grid) Set(row uint, col uint, state CellState) {
	g.data[g.cellToIndex(row, col)] = state
}

func (g *Grid) String() string {
	var b strings.Builder
	var divider = strings.Repeat("-", int(g.NumCols*2)+2)
	b.WriteString(divider + "\n")
	for row := uint(0); row < g.NumRows; row++ {
		b.WriteString("|")
		for col := uint(0); col < g.NumCols; col++ {
			if g.Get(row, col) == Alive {
				b.WriteString("◼ ")
			} else {
				b.WriteString("  ")
			}
		}
		b.WriteString("|\n")
	}
	b.WriteString(divider)
	return b.String()
}
