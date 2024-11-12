package main

import (
	"fmt"
	"strings"
)

type Sim struct {
	grid       *Grid
	Generation uint
}

func NewSim(numRows uint, numCols uint) *Sim {
	return &Sim{NewGrid(numRows, numCols), 0}
}

func NewSimWithGrid(numRows uint, numCols uint, initGrid []uint8) *Sim {
	return &Sim{NewGridWithInit(numRows, numCols, initGrid), 0}
}

func (s *Sim) IsCellAlive(row uint, col uint) bool {
	return s.grid.Get(row, col) == Alive
}

func (s *Sim) AnyCellAlive() bool {
	anyCellAlive := false
	for row := uint(0); row < s.grid.NumRows; row++ {
		for col := uint(0); col < s.grid.NumCols; col++ {
			if s.IsCellAlive(row, col) {
				anyCellAlive = true
				break
			}
		}
	}
	return anyCellAlive
}

func (s *Sim) NeighborCount(row uint, col uint) uint {
	count := uint(0)
	newRow := uint(0)
	newCol := uint(0)

	// 0 1 2
	// 3 X 4
	// 5 6 7

	// check the top left neighbor
	if (row > 0) && (col > 0) {
		newRow = row - 1
		newCol = col - 1

		if s.IsCellAlive(newRow, newCol) {
			count += 1
		}
	}

	// check the top center neighbor
	if row > 0 {
		newRow = row - 1
		newCol = col

		if s.IsCellAlive(newRow, newCol) {
			count += 1
		}
	}

	// check the top right neighbor
	if (row > 0) && ((col + 1) < s.grid.NumCols) {
		newRow = row - 1
		newCol = col + 1

		if s.IsCellAlive(newRow, newCol) {
			count += 1
		}
	}

	// check left neighbor
	if col > 0 {
		newRow = row
		newCol = col - 1

		if s.IsCellAlive(newRow, newCol) {
			count += 1
		}
	}

	// check right neighbor
	if (col + 1) < s.grid.NumCols {
		newRow = row
		newCol = col + 1

		if s.IsCellAlive(newRow, newCol) {
			count += 1
		}
	}

	// check bottom left neighbor
	if ((row + 1) < s.grid.NumRows) && (col > 0) {
		newRow = row + 1
		newCol = col - 1

		if s.IsCellAlive(newRow, newCol) {
			count += 1
		}
	}

	// check bottom center neighbor
	if (row + 1) < s.grid.NumRows {
		newRow = row + 1
		newCol = col

		if s.IsCellAlive(newRow, newCol) {
			count += 1
		}
	}

	// check bottom left neighbor
	if ((row + 1) < s.grid.NumRows) && ((col + 1) < s.grid.NumCols) {
		newRow = row + 1
		newCol = col + 1

		if s.IsCellAlive(newRow, newCol) {
			count += 1
		}
	}

	return count
}

func (s *Sim) ApplyRules(row uint, col uint) []Operation {
	var operations []Operation

	// determine the number of live neighbors to the current cell
	neighborCount := s.NeighborCount(row, col)

	// determine if the current cell is alive
	alive := s.IsCellAlive(row, col)

	if alive {
		// RULES FOR LIVE CELLS //////////////////////////////////////////

		switch {
		// rule 1: any live cell with fewer than two live neighbors dies,
		//          as if caused by under-population.
		case neighborCount < 2:
			operations = append(operations, NewOperation(row, col, Dead))

		// rule 2: any live cell with two or three live neighbors lives on
		//          to the next generation.
		case neighborCount <= 3:
			// do nothing, cell lives
			break

		// rule 3: any live cell with more than three neighbors dies, as if
		//          caused by overcrowding.
		default:
			operations = append(operations, NewOperation(row, col, Dead))
		}
	} else {
		// RULES FOR DEAD CELLS //////////////////////////////////////////

		// rule 4: any dead cell with exactly three live neighbors becomes
		//          a live cell, as if by reproduction.
		if neighborCount == 3 {
			operations = append(operations, NewOperation(row, col, Alive))
		}
	}

	return operations
}

func (s *Sim) Step() {
	var operations []Operation

	// increment the sim's generation
	s.Generation += 1

	for row := uint(0); row < s.grid.NumRows; row++ {
		for col := uint(0); col < s.grid.NumCols; col++ {
			// apply rules to the cell
			operationsCell := s.ApplyRules(row, col)

			// add any resultant operations to this step's operations list
			operations = append(operations, operationsCell...)
		}
	}

	// apply any operations for the step to the grid
	for _, operation := range operations {
		s.grid.Set(operation.Row, operation.Col, operation.State)
	}
}

func (s *Sim) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Generation: %v\n", s.Generation))
	b.WriteString(s.grid.String() + "\n")
	b.WriteString(fmt.Sprintf("Any cell alive?: %v\n", s.AnyCellAlive()))
	return b.String()
}
