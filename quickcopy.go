package main

import (
	"github.com/codegangsta/cli"
	"os"
	"github.com/pkg/errors"
	"github.com/SeriyBg/quick-copy/copy"
)

var (
	NoSrcSpecified = errors.New("No source specified!")
	NoDstSpecified = errors.New("No destination specified")
)

func main() {
	app := cli.NewApp()
	app.Name = "quick-copy"
	app.Usage = "Copy a directory or file"
	app.ArgsUsage = "[src SOURCE_FILE_OR_DIRECTORY] [dst DESTINATION_DIRECTORY]"
	app.Action = func(c *cli.Context) {
		src := c.Args().Get(0)
		dst := c.Args().Get(1)
		copy, err := copy.Copier(src)
		if err != nil {
			app.ErrWriter.Write([]byte(err.Error()))
			return
		}
		copy(src, dst)
	}
	app.Run(os.Args)
}
