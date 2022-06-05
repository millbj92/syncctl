/*
Synctl - a command line interface for file system
management on both local and remote machines.

Usage:
	synctl [command] [arguments]

	commands:
		--init [path]	Initialize a new sync path
		--add [path]	Add a new sync path
		--rm [path]	    Remove a sync path
		--ls		    List all sync paths
*/
package cli

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func Run() error {
	app := &cli.App{
		Name:  "synctl",
		Usage: "A command line interface for file system management on both local and remote machines.",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello, World!")
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Errorf("Error: %v", err)
		return err
	}
	return nil
}
