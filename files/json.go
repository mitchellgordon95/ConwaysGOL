package files

import (
	"encoding/json"
	"github.com/mitchellgordon95/ConwaysGOL/common"
	"io/ioutil"
)

type JsonBoard struct {
	AliveCells [][]int64
}

// Load a file onto the board, with respect to the given starting position
func LoadJson(board common.GolBoard, filename string, centerX, centerY int64) (common.GolBoard, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var jb JsonBoard
	json.Unmarshal(file, &jb)

	for _, cell := range jb.AliveCells {
		board = board.AddCell(centerX+cell[0], centerY+cell[1])
	}

	return board, nil
}
