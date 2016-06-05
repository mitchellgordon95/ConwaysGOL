package common

/**
A game of life board is a 2d plane centered at (0, 0)
that extends to the max size of a signed 64-bit integer in
all directions

The board does not wrap at the edges
**/

// GolBoard contains the state of the grid of cells in Conway's Game Of Life
type GolBoard interface {
	// Returns a copy of the board with cell in position (x,y) alive
	AddCell(int64, int64) GolBoard

	// Returns a copy of the board with cell in position (x,y) dead
	KillCell(int64, int64) GolBoard

	// returns whether or not a cell is alive
	IsAlive(int64, int64) bool

	// Returns a copy of the board stepped to the next state of the simulation
	Step() GolBoard

	// Returns an empty board
	Clear() GolBoard
}
