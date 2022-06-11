package main

import (
	//"flag"
	//"fmt"
	//"strconv"

	//"github.com/vladimirvivien/go-netbox/netbox"
	//"github.com/sirupsen/logrus"

	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli/v2"

	"github.com/millbj92/synctl/pkg/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "synctl"
	app.HelpName = "synctl"
	app.Compiled = time.Now()
	app.Usage = "A tool for managing housekeeping tasks"
	app.UsageText = "synctl [global options] command [command options]"
	app.ArgsUsage = "[command options]"
	app.Description = "Synctl is a multi-purpose command line interface for housekeeping and syncronization on both local and remote machines."
	app.Version = "0.0.1"
	app.Authors = []*cli.Author{
		{
			Name: "Brandon Miller",
			Email: "brandon@brandonmiller.io",
		},
	}
	app.Copyright = `Copyright (c) 2022 Enterforge, Inc. All rights reserved.`
	app.Suggest = true
	app.EnableBashCompletion = true
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
		{
			Name:    "help",
			Aliases: []string{"-h"},
			Usage:  "Show help",
			Action: Help,
		},
		commands.ServerCommands(),
		commands.FileCommands(),
	}
	  app.Action = func(c *cli.Context) error {
		  return Help(c)
	  }
	  err := app.Run(os.Args)
	  if err != nil {
		  spew.Dump(err)
	  }
}




func Help(c *cli.Context) error {
	cli.ShowAppHelp(c)
	return nil
}
