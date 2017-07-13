package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "quick-copy"
	app.Usage = "Copy a directory or file"
	app.Action = func(c *cli.Context) {
	}
	app.Run(os.Args)
}
