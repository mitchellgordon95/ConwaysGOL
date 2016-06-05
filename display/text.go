package display

import (
	"fmt"
	"github.com/mitchellgordon95/ConwaysGOL/common"
	"io"
)

type textDisplayer struct {
	io.Writer
}

func NewTextDisplayer(writer io.Writer) Displayer {
	return &textDisplayer{writer}
}

// Displays the game board in text.
func (td *textDisplayer) Display(board common.GolBoard, min_x, min_y, max_x, max_y int64) {
	// Clear the screen
	fmt.Fprintf(td, "\033c")
	for y := max_y - 1; y >= min_y; y-- {
		for x := min_x; x < max_x; x++ {
			val := board.IsAlive(int64(x), int64(y))

			if val {
				fmt.Fprint(td, "O")
			} else {
				fmt.Fprint(td, " ")
			}

			if x < max_x-1 {
				if x == (max_x-min_x)/2+min_x-1 {
					// Print a line down the middle of the grid
					fmt.Fprint(td, "|")
				} else if y == (max_y-min_y)/2+min_y {
					// Print a line across the middle of the grid
					fmt.Fprint(td, "_")
				} else {
					fmt.Fprint(td, " ")
				}
			}
		}
		fmt.Fprintln(td)
	}
}

func (td *textDisplayer) ShowMessage(msg string) {
	fmt.Fprintln(td, msg)
}
