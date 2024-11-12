package main

import "fmt"

type Operation struct {
	Row   uint
	Col   uint
	State CellState
}

func NewOperation(row uint, col uint, state CellState) Operation {
	return Operation{Row: row, Col: col, State: state}
}

func (op *Operation) String() string {
	return fmt.Sprintf("Operation[row: %d, col: %d, state: %s]", op.Row, op.Col, op.State)
}
