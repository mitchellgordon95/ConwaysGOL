/*
quadtree implements a quadtree for representing a 2d grid of cells
*/
package quadtree

// Node is a piece of the game board, which is stored in a quadtree
type Node interface {
	// Returns the level of the node
	Level() uint

	// Returns a copy of the node with the given cell in the node set to that value. (0,0) is the center of the node. A cell is identified by the coordinate of its lower left corner
	// Returns an error if the coordinate is out of bounds.
	SetValue(x, y int64, val bool) (Node, error)

	// Returns the value held by a cell contained by the node. (0,0) is the center of the node. A cell is identified by the coordinate of its lower left corner
	// Returns an error if the coordinate is out of bounds.
	GetValue(x, y int64) (bool, error)

	// Returns the subnode representing a quadrant of the node, or nil for leaves
	NW() Node
	NE() Node
	SW() Node
	SE() Node
}

// Returns an empty tree with the number of levels specified.
func EmptyTree(levels int) Node {
	if levels < 1 {
		return nil
	}

	// An empty cell
	var node Node = LeafNode(false)
	for i := 1; i < levels; i++ {
		// Build each level, pointing all quadrants to the empty node a level below the current one
		node = QuadNode(node, node, node, node)
	}

	return node
}
