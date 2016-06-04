package main

import (
	"github.com/mitchellgordon95/ConwaysGOL/common"
	"github.com/mitchellgordon95/ConwaysGOL/display"
	gm "github.com/mitchellgordon95/ConwaysGOL/game_manager"
	"github.com/mitchellgordon95/ConwaysGOL/hashlife"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "ConwaysGOL"
	app.Usage = "Runs the Conway's Game of Life simulation"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file,f",
			Usage: "a json file to read the initial configuration from, defaults to an empty board",
		},
		cli.BoolFlag{
			Name:  "gui,g",
			Usage: "show the game board in a gui window",
		},
		cli.IntFlag{
			Name:  "size,s",
			Usage: "The size of the gameboard to show. Defaults to 16. Note that this is just the view, the actual size is 2^64",
		},
	}
	app.Action = func(c *cli.Context) error {
		file := c.String("file")
		var board common.GolBoard
		if file != "" {
			// TODO: read from file
			board = nil
		} else {
			board = hashlife.NewHashLifeBoard()
		}

		displayer := display.NewTextDisplayer(os.Stdout)

		size := c.Int("size")
		if size == 0 {
			size = 16
		}

		gm.NewTextManager(board, os.Stdin, displayer, int64(size)).Manage()

		return nil
	}

	app.Run(os.Args)
}
