package main

import (
	"github.com/codegangsta/cli"
	"github.com/emaniacs/gron/commands"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gron"
	app.Usage = "$ gron <command> <options>"
	
	// app.Flags = ""
	app.Commands = commands.Commands()
	app.Run(os.Args)
}
