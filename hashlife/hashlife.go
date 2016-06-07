package hashlife

import (
	"fmt"
	"github.com/mitchellgordon95/ConwaysGOL/common"
	qt "github.com/mitchellgordon95/ConwaysGOL/quadtree"
)

// An implementation of a GOL board using the HashLife algorithm
type hashLife struct {
	qt.Node
}

// Get an instance of the hashlife board
func NewHashLifeBoard() common.GolBoard {
	// The width of the board is 2^64. This requires a quad
	// tree with 65 levels. However, we want to compute the whole board,
	// not just the subnode of width 2^63. So we need a tree with
	// 66 levels
	return hashLife{qt.EmptyTree(66)}
}

// Returns a copy of the board with cell in position (x,y) alive
func (hl hashLife) AddCell(x, y int64) common.GolBoard {
	node, err := hl.SetValue(x, y, true)

	if err != nil {
		fmt.Println("error: %s", err.Error())
		return nil
	}

	return hashLife{node}
}

// Returns a copy of the board with cell in position (x,y) dead
func (hl hashLife) KillCell(x, y int64) common.GolBoard {
	node, err := hl.SetValue(x, y, false)

	if err != nil {
		fmt.Println("error: %s", err.Error())
		return nil
	}

	return hashLife{node}
}

// returns whether or not a cell is alive
func (hl hashLife) IsAlive(x, y int64) bool {
	val, err := hl.GetValue(x, y)

	if err != nil {
		fmt.Println("error: %s", err.Error())
		// Just assume out of bounds is dead
		return false
	}

	return val
}

// Returns a copy of the board stepped to the next state of the simulation
var deadNode qt.Node = qt.EmptyTree(64)

func (hl hashLife) Step() common.GolBoard {
	next := NextGeneration(hl.Node)

	// We have to pad the result with dead cells, since
	// NextGeneration returns a node one level down
	return hashLife{
		qt.QuadNode(
			qt.QuadNode(deadNode, deadNode, deadNode, next.NW()),
			qt.QuadNode(deadNode, deadNode, next.NE(), deadNode),
			qt.QuadNode(deadNode, next.SW(), deadNode, deadNode),
			qt.QuadNode(next.SE(), deadNode, deadNode, deadNode),
		),
	}
}

func (hl hashLife) Clear() common.GolBoard {
	return NewHashLifeBoard()
}
