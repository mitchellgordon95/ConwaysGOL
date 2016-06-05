package hashlife

import (
	qt "github.com/mitchellgordon95/ConwaysGOL/quadtree"
)

// A cache containing the results of computing the next generation for a node
var generationCache map[qt.Node]qt.Node

func init() {
	generationCache = map[qt.Node]qt.Node{}
}

// Returns the next generation of life for a node one level down the tree, centered at the given node
func NextGeneration(node qt.Node) qt.Node {
	// If we have a cached result, use that
	cached, ok := generationCache[node]
	if ok {
		return cached
	}

	if node.Level() == 2 {
		return baseCase(node)
	}

	// First, we construct 9 nodes two levels down that encompass the area we're trying to generate
	n00 := centeredSubnode(node.NW())
	n01 := centeredHorizontal(node.NW(), node.NE())
	n02 := centeredSubnode(node.NE())
	n10 := centeredVertical(node.NW(), node.SW())
	n11 := centeredSubSubnode(node)
	n12 := centeredVertical(node.NE(), node.SE())
	n20 := centeredSubnode(node.SW())
	n21 := centeredHorizontal(node.SW(), node.SE())
	n22 := centeredSubnode(node.SE())

	// Then we construct four nodes one level down out of those 9 nodes.
	// Each of these sub nodes will have a centered sub node that makes up one quadrant
	// of the node we're trying to compute
	out := qt.QuadNode(
		NextGeneration(qt.QuadNode(n00, n01, n10, n11)),
		NextGeneration(qt.QuadNode(n01, n02, n11, n12)),
		NextGeneration(qt.QuadNode(n10, n11, n20, n21)),
		NextGeneration(qt.QuadNode(n11, n12, n21, n22)),
	)

	// Store the result for future calls
	generationCache[node] = out
	return out
}

// For the base case of the hashlife algorithm, when the node is level 2, just run the simulation normally to find the next generation of the centered subnode
func baseCase(node qt.Node) qt.Node {
	// For each of the four leaf nodes in the center, do the simulation
	return qt.QuadNode(
		nextState(node, -1, 0),
		nextState(node, 0, 0),
		nextState(node, -1, -1),
		nextState(node, 0, -1),
	)
}

// Return the next state for a leaf node at the given coordinates
func nextState(node qt.Node, x, y int64) qt.Node {
	aliveNeighbors := 0

	var err error
	var alive bool

	for cur_x := x - 1; cur_x < x+2; cur_x++ {
		for cur_y := y - 1; cur_y < y+2; cur_y++ {
			// skip the node we're trying to calculate
			if cur_x == x && cur_y == y {
				continue
			}

			alive, err = node.GetValue(cur_x, cur_y)

			if alive {
				aliveNeighbors += 1
			}
		}
	}

	alive, err = node.GetValue(x, y)

	// TODO: Log this; it shouldn't happen
	if err != nil {
		return nil
	}

	if alive {
		// If we're alive and we have 2 or three alive neighbors, we're alive
		if aliveNeighbors == 2 || aliveNeighbors == 3 {
			return qt.LeafNode(true)
		}
		// otherwise we're dead
		return qt.LeafNode(false)
	} else {
		// If we're dead and we have 3 alive neighbors, we're alive
		if aliveNeighbors == 3 {
			return qt.LeafNode(true)
		}
		// otherwise we're still dead
		return qt.LeafNode(false)
	}
}

// Given two nodes side by side on the grid, returns a node one level down centered vertically and on the boundary of the two nodes
func centeredHorizontal(w, e qt.Node) qt.Node {
	if w.Level() < 2 || e.Level() < 2 || e.Level() != w.Level() {
		return nil
	}

	return qt.QuadNode(w.NE().SE(), e.NW().SW(), w.SE().NE(), e.SW().NW())
}

// Given two nodes side by side on the grid, returns a node one level down centered vertically and on the boundary of the two nodes
func centeredVertical(n, s qt.Node) qt.Node {
	if n.Level() < 2 || s.Level() < 2 || s.Level() != n.Level() {
		return nil
	}

	return qt.QuadNode(n.SW().SE(), n.SE().SW(), s.NW().NE(), s.NE().NW())
}

// Returns a node one level down the tree centered at the middle of the current node, or nil if impossible
func centeredSubnode(node qt.Node) qt.Node {
	if node.Level() < 2 {
		return nil
	}

	return qt.QuadNode(node.NW().SE(), node.NE().SW(), node.SW().NE(), node.SE().NW())
}

// Returns a node two levels down the tree centered at the middle of the current node, or nil if impossible
func centeredSubSubnode(node qt.Node) qt.Node {
	if node.Level() < 3 {
		return nil
	}

	return qt.QuadNode(
		node.NW().SE().SE(),
		node.NE().SW().SW(),
		node.SW().NE().NE(),
		node.SE().NW().NW(),
	)
}
