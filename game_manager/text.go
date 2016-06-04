package game_manager

import (
	"bufio"
	"errors"
	"github.com/mitchellgordon95/ConwaysGOL/common"
	"github.com/mitchellgordon95/ConwaysGOL/display"
	"io"
	"strconv"
	"strings"
)

// Manages the game state based on the input the user types
type textManager struct {
	board common.GolBoard
	*bufio.Reader
	display.Displayer
	size int64
}

/*
Creates a new text manager to manage the game state.
Takes a game board, a reader to get text from the user, a game displayer,
and the size of the game board to display, centered at 0
*/
func NewTextManager(board common.GolBoard, read io.Reader, displayer display.Displayer, size int64) GolManager {
	return &textManager{board, bufio.NewReader(read), displayer, size}
}

func (tm *textManager) Manage() {
	tm.greet()

	for {
		text, err := tm.ReadString('\n')
		if err != nil {
			tm.ShowMessage("Oops! Something went wrong. " + err.Error())
		}
		text = strings.TrimSpace(text)
		tokens := strings.Split(text, " ")

		if len(tokens) == 0 {
			continue
		}

		switch tokens[0] {
		case "show":
			tm.showBoard()
		case "quit":
			tm.ShowMessage("Bye!")
			return
		case "alive":
			tm.aliveCell(tokens[1:])
		case "kill":
			tm.deadCell(tokens[1:])
		case "next":
			tm.nextBoard(tokens[1:])
		default:
			tm.ShowMessage("Invalid command.")
		}
	}

}

func (tm *textManager) showBoard() {
	half_size := tm.size / 2
	tm.Display(tm.board, -half_size, -half_size, half_size, half_size)
}

func (tm *textManager) aliveCell(tokens []string) {
	if len(tokens) < 2 {
		tm.ShowMessage("Not enough arguments")
		return
	}
	x, y, err := parseCoordinates(tokens)
	if err != nil {
		tm.ShowMessage(err.Error())
		return
	}
	tm.board = tm.board.AddCell(x, y)
	tm.ShowMessage("Set cell to alive!")
}

func (tm *textManager) deadCell(tokens []string) {
	if len(tokens) < 2 {
		tm.ShowMessage("Not enough arguments")
	}
	x, y, err := parseCoordinates(tokens)
	if err != nil {
		tm.ShowMessage(err.Error())
		return
	}
	tm.board = tm.board.KillCell(x, y)
	tm.ShowMessage("Set cell to be dead!")
}

func (tm *textManager) nextBoard(tokens []string) {
	if len(tokens) == 0 {
		tm.board = tm.board.Step()
	} else {
		steps, err := strconv.ParseUint(tokens[0], 10, 64)
		if err != nil {
			tm.ShowMessage("Invalid number of steps")
			return
		}
		for i := uint64(0); i < steps; i++ {
			tm.board = tm.board.Step()
		}
	}
	tm.showBoard()
}

func parseCoordinates(tokens []string) (int64, int64, error) {
	x, err := strconv.ParseInt(tokens[0], 10, 64)
	if err != nil {
		return 0, 0, errors.New("Invalid x coordinate")
	}
	y, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return 0, 0, errors.New("Invalid y coordinate")
	}
	return x, y, nil
}

func (tm *textManager) greet() {
	tm.ShowMessage("Welcome to Conway's Game of Life!")
	tm.ShowMessage("Enter \"show\" to show the current game board")
	tm.ShowMessage("Enter \"next\" to go to the next step in the simulation")
	tm.ShowMessage("Enter \"next [steps]\" to do a certain number of steps in the simulation")
	tm.ShowMessage("Enter \"alive [x] [y]\" to set the cell at (x,y) as alive")
	tm.ShowMessage("Enter \"kill [x] [y]\" to kill the cell at (x,y)")
	tm.ShowMessage("Enter \"quit\" to quit")
}
