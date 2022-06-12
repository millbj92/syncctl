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

	_ "github.com/millbj92/synctl/docs"
	"github.com/millbj92/synctl/pkg/commands"
)

// @title Synctl Server
// @version 1.0
// @description The Synctl Server is meant to act as a bridge between the
// @description nodes in a synctl cluster. It will be used to sync files
// @description between nodes, schedule remote nodes for self cleanup, and
// @description perform other general housekeeping tasks.
// @contact.name Brandon Miller
// @contact.email brandon@brandonmiller.io
// @license.name MIT
// @license.url https://github.com/millbj92/synctl/blob/master/LICENSE
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
			"Brandon Miller",
			"brandon@brandonmiller.io",
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
			Usage:   "Show help",
			Action:  Help,
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

	err := cli.ShowAppHelp(c)
	if err != nil {
		return err
	}
	return nil
}
