package commands
import (
	"github.com/codegangsta/cli"
	"fmt"
)

func init(){
	AddCommand(cli.Command{
		Name: "command",
		Usage: "Show all command",
		Action: func (c *cli.Context){
			for _, command := range Commands() {
				fmt.Println(command.Name, " - ", command.Usage)
			}
		},
	})
}
