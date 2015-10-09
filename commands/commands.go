package commands

import (
	"github.com/codegangsta/cli"
)

var (
	commands []cli.Command
)

func Commands() []cli.Command {
	return commands
}

func AddCommand(command cli.Command){
	commands = append(commands, command)
}
