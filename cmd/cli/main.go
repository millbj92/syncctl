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

	"github.com/gofiber/fiber/v2"
	"github.com/millbj92/synctl/internal/configs"
	"github.com/millbj92/synctl/pkg/commands/server"
	"github.com/millbj92/synctl/pkg/management"
	"github.com/millbj92/synctl/pkg/middleware"
	"github.com/millbj92/synctl/pkg/models/tasks"
	"github.com/millbj92/synctl/pkg/routes"
)

func main() {
	var fiberApp *fiber.App = nil
	app := cli.NewApp()
	app.Name = "synctl"
	app.HelpName = "synctl"
	app.Compiled = time.Now()
	app.Usage = "A tool for managing housekeeping tasks"
	app.UsageText = "synctl - [global options] command [command options]"
	app.ArgsUsage = "[command options]"
	app.Description = "Synctl is a multi-purpose command line interface for housekeeping and syncronization on both local and remote machines."
	app.Version = "0.0.1"
	app.Authors = []*cli.Author{
		{
			Name: "Brandon Miller",
			Email: "brandon@brandonmiller.io",
		},
	}
	// app.CustomAppHelpTemplate = `
	// NAME: {{.Name}} - {{.Usage}}
	// USAGE: {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
	// {{if .Version}}VERSION: {{.Version}}
	// {{end}}{{if len .Authors}}
	// AUTHOR(S):
	// {{range .Authors}}{{ . }}{{end}}
	// {{end}}{{if .Commands}}
	// COMMANDS:
	// {{range .Commands}}{{if not .HideHelp}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}{{ "\n" }}{{end}}{{end}}
	// {{end}}{{if .VisibleFlags}}
	// GLOBAL OPTIONS:
	// {{range .VisibleFlags}}{{.}}
	// {{end}}{{end}}{{if .Copyright }}`
	app.Copyright = `Copyright (c) 2022 Enterforge, Inc. All rights reserved.`
	app.Suggest = true
	app.EnableBashCompletion = true
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
		{
			Name:    "help",
			Aliases: []string{"h"},
			Usage:  "Show help",
			Action: Help,
		},
		server.ServerCommands(),
		// {
		// 	Name: "server",
		// 	Aliases: []string{"--sv"},
		// 	Usage:  "Start server",
		// 	UsageText: "server [command] [flags]",
		// 	Description: "Starts as a master node to relay all commands to remotes",
		// 	ArgsUsage: "[command] [flags]",
		// 	Subcommands: []*cli.Command{
		// 		&cli.Command{
		// 			Name: "start",
		// 			Action: func(c *cli.Context) error {
		// 				fiberApp, err := StartServer()
		// 				if err != nil {
		// 					return err
		// 				}
		// 				fiberApp.Listen(fmt.Sprintf("%s:%d", c.String("host"), c.Int("port")))
		// 				return nil
		// 			},
		// 			Flags: []cli.Flag{
		// 				&cli.IntFlag{
		// 					Name: "port",
		// 					Aliases: []string{"p"},
		// 					EnvVars: []string{"SERVER_PORT"},
		// 					Value: 8101,
		// 					DefaultText: "8101",
		// 				},
		// 				&cli.StringFlag{
		// 					Name: "host",
		// 					Category: "Server",
		// 					Usage: "Host to listen on",
		// 					EnvVars: []string{"SERVER_HOST"},
		// 					DefaultText: "Server will listen on all interfaces",
		// 					Value: "localhost",
		// 					Aliases: []string{"hs"},
		// 				},
		// 			},
		// 		},
		// 		&cli.Command{
		// 			Name: "stop",
		// 			Action: func(c *cli.Context) error {
		// 				if fiberApp == nil {
		// 					return nil
		// 				}
		// 				return StopServer(fiberApp)
		// 			},
		// 		},
		// 	},
		// 	SkipFlagParsing: false,
		// 	HideHelp: false,
		// 	HelpName: "server",
		// 	BashComplete: func(c *cli.Context) {
		// 		fmt.Fprintf(c.App.Writer, "--host, h, --port, p")
		// 	},
		// 	Before: func(c *cli.Context) error {
		// 		if c.Args().First() == "start" {
		// 		fmt.Fprintf(c.App.Writer, "\n\n\t\t🚀 OH LAWD WE COMIN\n")
		// 		}

		// 		if c.Args().First() == "stop" {
		// 			fmt.Fprintf(c.App.Writer, "\n\n\t\t👋 So long, space cowboy!\n")
		// 		}
		// 		return nil
		// 	},
		// 	After: func(c *cli.Context) error {
		// 		if c.Args().First() == "start" {
		// 		  fmt.Fprintf(c.App.Writer, "Server successfully started. Welcome aboard!\n")
		// 		}
		// 		if c.Args().First() == "stop" {
		// 			fmt.Fprintf(c.App.Writer, "Server successfully stopped. See you next time!\n")
		// 		}
		// 		return nil
		// 	},
		// 	Action: func(c *cli.Context) error {
		// 		c.Command.FullName()
		// 		c.Command.HasName("server")
		// 		c.Command.Names()
		// 		c.Command.VisibleFlags()
		// 		return nil
		// 	},
		// 	OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
		// 		fmt.Fprintf(c.App.Writer, "Error: %s\n", err.Error())
		// 		return err
		// 	},
		// },
		{
			Name: "files",
			Subcommands: []*cli.Command{
				//Copy
				{
					Name:   "copy",
					Flags:  []cli.Flag{
						&cli.PathFlag{
							Name:  "source",
							Category: "File Management",
							Usage: "Path to file or directory",
							Aliases: []string{"p"},
							Required: true,
						},
						&cli.PathFlag{
							Name:  "destination",
							Category: "File Management",
							Usage: "Destination path",
							Aliases: []string{"d"},
						},
						&cli.StringFlag{
							Name: "include",
							Category: "File Management",
							Usage: "Include files matching pattern",
							Aliases: []string{"i"},
						},
						&cli.StringFlag{
							Name: "exclude",
							Category: "File Management",
							Usage: "Exclude files matching pattern",
							Aliases: []string{"e"},
						},
						&cli.BoolFlag{
							Name:  "recursive",
							Aliases: []string{"r"},
							Usage: "Recurse through directories",
						},
						&cli.BoolFlag{
							Name:  "force",
							Aliases: []string{"f"},
							Usage: "Force copy",
						},
					},
					Aliases: []string{"cp"},
					Usage:  "Copy object(s)",
					Action: func (c *cli.Context) error {
						return management.CopyFiles(tasks.CopyArgs{
							TaskArgs: tasks.TaskArgs{
								Source: c.String("source"),
								Include: c.String("include"),
								Exclude: c.String("exclude"),
								Recursive: c.Bool("recursive"),
								Force: c.Bool("force"),
							},
							Destination: c.String("destination"),
						})
					},
				},
				//Move
				{
					Name:   "move",
					Flags:  []cli.Flag{
						&cli.PathFlag{
							Name:  "source",
							Category: "File Management",
							Usage: "Path to file or directory that wil be the source",
							Aliases: []string{"p"},
							Required: true,
						},
						&cli.PathFlag{
							Name:  "destination",
							Category: "File Management",
							Usage: "Path to file or directory that wil be the destination",
							Aliases: []string{"d" },
							Required: true,
						},
						&cli.StringFlag{
							Name: "include",
							Category: "File Management",
							Usage: "Include files matching pattern",
							Aliases: []string{"i"},
						},
						&cli.StringFlag{
							Name: "exclude",
							Category: "File Management",
							Usage: "Exclude files matching pattern",
							Aliases: []string{"e"},
						},
						&cli.BoolFlag{
							Name:  "recursive",
							Aliases: []string{"r"},
							Usage: "Recurse through directories",
						},
						&cli.BoolFlag{
							Name:  "force",
							Aliases: []string{"f"},
							Usage: "Force move",
						},
					},
					Aliases: []string{"mv"},
					Usage:  "Move object(s)",
					Action: func (c *cli.Context) error {
						return management.MoveFiles(tasks.MoveArgs{
							TaskArgs: tasks.TaskArgs{
								Source: c.String("source"),
								Include: c.String("include"),
								Exclude: c.String("exclude"),
								Recursive: c.Bool("recursive"),
								Force: c.Bool("force"),
							},
							Destination: c.String("destination"),
						})
					},
				},
				//Rename
				{
					Name: "rename",
					Flags: []cli.Flag{
						&cli.PathFlag{
							Name:  "source",
							Category: "File Management",
							Usage: "Path to file or directory that wil be the source",
							Aliases: []string{"p"},
							Required: true,
						},
						&cli.StringFlag{
							Name: "include",
							Category: "File Management",
							Usage: "Pattern to include files (e.g. *.txt)",
							Aliases: []string{"i"},
						},
						&cli.StringFlag{
							Name: "exclude",
							Category: "File Management",
							Usage: "Pattern to exclude files (e.g. *.exe)",
							Aliases: []string{"e"},
						},
						&cli.StringFlag{
							Name: "prefix",
							Category: "File Management",
							Usage: "Prefix to add to file names",
							Aliases: []string{"p"},
						},
						&cli.StringFlag{
							Name: "extension",
							Category: "File Management",
							Usage: "Rename files with this filetype extension (e.g. .rtf)",
							Aliases: []string{"e"},
						},
						&cli.BoolFlag{
							Name:  "recursive",
							Aliases: []string{"r"},
							Usage: "Recurse through directories",
						},
					},
					Aliases: []string{"rn"},
					Usage:  "Rename files",
					Action: func (c *cli.Context) error {
						return management.RenameFiles(tasks.RenameArgs{
							TaskArgs: tasks.TaskArgs{
								Source: c.String("source"),
								Include: c.String("include"),
								Exclude: c.String("exclude"),
								Recursive: c.Bool("recursive"),
							},
							Prefix: c.String("prefix"),
							Extension: c.String("extension"),
						})
					},
				},
				//TODO: Sync
				{
					Name:   "sync",
					Flags:  []cli.Flag{
						&cli.PathFlag{
							Name:  "source",
							Category: "File Management",
							Usage: "Path to file or directory that wil be the source",
							Aliases: []string{"p"},
							Required: true,
						},
						&cli.PathFlag{
							Name:  "destination",
							Category: "File Management",
							Usage: "Path to file or directory that wil be the destination",
							Aliases: []string{"d", },
							Required: true,
						},
						&cli.StringFlag{
							Name: "include",
							Category: "File Management",
							Usage: "Include files matching pattern",
							Aliases: []string{"i"},
						},
						&cli.StringFlag{
							Name: "exclude",
							Category: "File Management",
							Usage: "Exclude files matching pattern",
							Aliases: []string{"e"},
						},
						&cli.BoolFlag{
							Name:  "recursive",
							Aliases: []string{"r"},
							Usage: "Recurse through directories",
						},
						&cli.BoolFlag{
							Name: "force",
							Category: "File Management",
							Usage: "Force overwrite",
							Aliases: []string{"f"},
						},
						&cli.BoolFlag{
							Name: "delete",
							Category: "File Management",
							Usage: "Delete files in destination that are not in source",
							Aliases: []string{"-dd", "--delete-destination"},
						},
						&cli.BoolFlag{
							Name: "destructive",
							Category: "File Management",
							Usage: "Delete files in source after syncing",
							Aliases: []string{"-dd"},
						},
					},
					Aliases: []string{"sc"},
					Usage:  "Sync object(s)",
					Action: func (c *cli.Context) error {
						return management.SyncFiles(tasks.SyncArgs{
							TaskArgs: tasks.TaskArgs{
								Source: c.String("source"),
								Include: c.String("include"),
								Exclude: c.String("exclude"),
								Recursive: c.Bool("recursive"),
								Force: c.Bool("force"),
							},
							Delete: c.Bool("delete"),
							Destination: c.String("destination"),
							Destructive: c.Bool("destructive"),
						})
					},
				},
				//Delete
				{
					Name:    "delete",
					Flags:   []cli.Flag{
						&cli.PathFlag{
							Name:  "path",
							Category: "File Management",
							Usage: "Path to file or directory",
							Aliases: []string{"p"},
							Required: true,
						},
						&cli.StringFlag{
							Name: "include",
							Category: "File Management",
							Usage: "Include files matching pattern",
							Aliases: []string{"i"},
						},
						&cli.StringFlag{
							Name: "exclude",
							Category: "File Management",
							Usage: "Exclude files matching pattern",
							Aliases: []string{"e"},
						},
						&cli.BoolFlag{
							Name:  "recursive",
							Aliases: []string{"r"},
							Usage: "Recurse through directories",
						},
						&cli.BoolFlag{
							Name:  "force",
							Aliases: []string{"f"},
							Usage: "Force delete",
						},
					},
					Aliases: []string{"rm"},
					Usage:  "Delete object(s)",
					Action:  func(c *cli.Context) error {
						return management.DeleteFiles(tasks.TaskArgs{
							Source: c.String("path"),
							Include: c.String("include"),
							Exclude: c.String("exclude"),
							Recursive: c.Bool("recursive"),
							Force: c.Bool("force"),
						})
					},
				},
				{
					Name: "list",
					Flags: []cli.Flag{
						&cli.StringFlag{

							Name: "--path",
							Category: "File Management",
							Usage: "Path to file or directory",
							Aliases: []string{"p"},
							Required: true,
						},
						&cli.StringFlag{
							Name: "include",
							Category: "File Management",
							Usage: "Include files matching pattern",
							Aliases: []string{"i"},
							Required: false,
						},
						&cli.StringFlag{
							Name: "exclude",
							Category: "File Management",
							Usage: "Exclude files matching pattern",
							Aliases: []string{"e"},
							Required: false,
						},
						&cli.BoolFlag{
							Name: "recursive",
							Aliases: []string{"r"},
							Usage: "Recurse through directories",
							Required: false,
						},
					},
					Aliases: []string{"ls"},
					Usage: "List object(s)",
					Action: func(c *cli.Context) error {
						return management.ListFiles(tasks.TaskArgs{
							Source: c.String("path"),
							Include: c.String("include"),
							Exclude: c.String("exclude"),
							Recursive: c.Bool("recursive"),
						})

				},
			},
		},
	  },
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

func StartServer() (*fiber.App, error) {
	config := configs.ConfigureFiber()

	app := fiber.New(config)
	middleware.LoadMiddleware(app)
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	return app, nil
}

func StopServer(a *fiber.App) error {
	return a.Shutdown()
}
