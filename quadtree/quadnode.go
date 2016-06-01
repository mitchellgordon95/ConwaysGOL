package quadtree

import (
	"errors"
	"fmt"
)

func init() {
	// Initialize the node cache to be empty
	nodeCache = map[quadNode]*quadNode{}
}

// nodeCache stores canonical copies of all quadnodes to reduce redundant memory consumption
var nodeCache map[quadNode]*quadNode

// A quadNode points to the four sub-sections of the game board that it contains
type quadNode struct {
	// Pointers to the sub-sections of the node
	nw, ne, sw, se Node
	// Level of the node in the tree. Determines the size of the board the node is responsible for.
	level uint
}

// TODO: Write garbage collection for cache
func PrintCache() {
	fmt.Printf("%v", nodeCache)
	fmt.Println()
}

// Returns a new tree node. Caches the resulting node so that only one canonical copy of each node exists at any time.
func QuadNode(nw, ne, sw, se Node) Node {
	node := quadNode{nw, ne, sw, se, nw.Level() + 1}
	cached, ok := nodeCache[node]

	if !ok {
		nodeCache[node] = &node
		return &node
	} else {
		return cached
	}
}

func (qn *quadNode) Level() uint {
	return qn.level
}

func outOfBound(x, y, subsectionSize int64) bool {
	return x >= subsectionSize || x < -subsectionSize || y >= subsectionSize || y < -subsectionSize
}

func (qn *quadNode) SetValue(x, y int64, val bool) (Node, error) {
	// If the level is 65 or above, construct a smaller subnode centered at the current one
	// and call that. Then reconstruct it into its original size
	if qn.level > 64 {
		// Assume anything outside of the addressable bounds is empty
		empty := qn.NW().NW()
		res, err := QuadNode(qn.NW().SE(), qn.NE().SW(), qn.SW().NE(), qn.SE().NW()).SetValue(x, y, val)
		return QuadNode(
			QuadNode(empty, empty, empty, res.NW()),
			QuadNode(empty, empty, res.NE(), empty),
			QuadNode(empty, res.SW(), empty, empty),
			QuadNode(res.SE(), empty, empty, empty),
		), err
	}

	// if the level is 64 or above, don't check bounds because we can't possibly be out of them
	if qn.level < 64 {
		// The width of a subsection. Note this fits in an int64 since the level is < 64
		subsectionSize := int64(1) << (qn.level - 1)
		if outOfBound(x, y, subsectionSize) {
			return nil, errors.New("quadNode: grid location out of bound")
		}
	}

	var posOffset, negOffset int64
	if qn.level == 1 {
		// Note: level 1 is a special case because we're using integer division (i.e. 1 / 2 == 0)
		posOffset = 1
		negOffset = 0
	} else {
		offset := int64(1) << (qn.level - 2)
		posOffset = offset
		negOffset = -offset
	}

	var out, subNode Node
	var err error
	switch {
	case x < 0 && y >= 0:
		subNode, err = qn.nw.SetValue(x+posOffset, y+negOffset, val)
		out = QuadNode(subNode, qn.ne, qn.sw, qn.se)
	case x >= 0 && y >= 0:
		subNode, err = qn.ne.SetValue(x+negOffset, y+negOffset, val)
		out = QuadNode(qn.nw, subNode, qn.sw, qn.se)
	case x < 0 && y < 0:
		subNode, err = qn.sw.SetValue(x+posOffset, y+posOffset, val)
		out = QuadNode(qn.nw, qn.ne, subNode, qn.se)
	case x >= 0 && y < 0:
		subNode, err = qn.se.SetValue(x+negOffset, y+posOffset, val)
		out = QuadNode(qn.nw, qn.ne, qn.sw, subNode)
	}

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (qn *quadNode) GetValue(x, y int64) (bool, error) {
	// If the level is 65 or above, construct a smaller subnode centered at the current one
	// and call that.
	if qn.level > 64 {
		return QuadNode(qn.NW().SE(), qn.NE().SW(), qn.SW().NE(), qn.SE().NW()).GetValue(x, y)
	}

	// if the level is 64 or above, don't check bounds because we can't possibly be out of them
	if qn.level < 64 {
		// The width of a subsection. Note this fits in an int64 since the level is < 64
		subsectionSize := int64(1) << (qn.level - 1)
		if outOfBound(x, y, subsectionSize) {
			return false, errors.New("quadNode: grid location out of bound")
		}
	}

	var posOffset, negOffset int64
	if qn.level == 1 {
		// Note: level 1 is a special case because we're using integer division (i.e. 1 / 2 == 0)
		posOffset = 1
		negOffset = 0
	} else {
		offset := int64(1) << (qn.level - 2)
		posOffset = offset
		negOffset = -offset
	}

	var val bool
	var err error
	switch {
	case x < 0 && y >= 0:
		val, err = qn.nw.GetValue(x+posOffset, y+negOffset)
	case x >= 0 && y >= 0:
		val, err = qn.ne.GetValue(x+negOffset, y+negOffset)
	case x < 0 && y < 0:
		val, err = qn.sw.GetValue(x+posOffset, y+posOffset)
	case x >= 0 && y < 0:
		val, err = qn.se.GetValue(x+negOffset, y+posOffset)
	}

	if err != nil {
		return false, err
	}

	return val, nil
}

func (qn *quadNode) NW() Node {
	return qn.nw
}
func (qn *quadNode) NE() Node {
	return qn.ne
}
func (qn *quadNode) SW() Node {
	return qn.sw
}
func (qn *quadNode) SE() Node {
	return qn.se
}
