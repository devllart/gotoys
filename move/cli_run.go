package move

import (
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func CliRun(c *cli.Context) {
	oldPath := c.Args().Get(0)
	newPath := c.Args().Get(1)
	moveFile(oldPath, newPath)

	if strings.HasPrefix(oldPath, "./") {
		oldPath = oldPath[2:]
	}
	if strings.HasPrefix(newPath, "./") {
		newPath = newPath[2:]
	}
	files := getGoFilesList(".")

	replaceInFile(files, oldPath, newPath)
}

func moveFile(oldLocation, newLocation string) {
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}
