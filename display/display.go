package display

import "github.com/mitchellgordon95/ConwaysGOL/common"

// Displays a GOL board
type Displayer interface {
	// Displays a chunk of the board, from min coordinates (inclusive) to max coordinates (exclusive)
	Display(board common.GolBoard, min_x, min_y, max_x, max_y int64)
	// Shows a message to the user
	ShowMessage(msg string)
}
