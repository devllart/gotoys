package main

import (
	"os"

	"github.com/devllart/gotoys/move"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gotoys"
	app.Commands = []cli.Command{
		{
			Name:   "mv",
			Usage:  "move file and fix imports",
			Action: move.CliRun,
		},
	}
	app.Run(os.Args)
	app.Usage = "Golang tolls, but we called it go toys"
}
