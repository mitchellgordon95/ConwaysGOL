package quadtree

import (
	"errors"
)

// A leaf node represents one cell of the board. It is either alive or dead.
type leafNode bool

func LeafNode(val bool) Node {
	return leafNode(val)
}

func (ln leafNode) Level() uint {
	return 0
}

func (ln leafNode) SetValue(x, y int64, value bool) (Node, error) {
	if x != 0 || y != 0 {
		return nil, errors.New("leafNode: grid location out of bound")
	}

	return LeafNode(value), nil
}

func (ln leafNode) GetValue(x, y int64) (bool, error) {
	if x != 0 || y != 0 {
		return false, errors.New("leafNode: grid location out of bound")
	}

	return bool(ln), nil
}

func (leafNode) NW() Node {
	return nil
}
func (leafNode) NE() Node {
	return nil
}
func (leafNode) SW() Node {
	return nil
}
func (leafNode) SE() Node {
	return nil
}
