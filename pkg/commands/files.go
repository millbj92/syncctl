package commands

import (
	"github.com/millbj92/synctl/pkg/management"
	"github.com/millbj92/synctl/pkg/models/tasks"
	"github.com/urfave/cli/v2"
)

func FileCommands() *cli.Command {
	return &cli.Command{
		Name: "files",
		Subcommands: []*cli.Command{
			//Copy
			{
				Name: "copy",
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:     "source",
						Category: "File Management",
						Usage:    "Path to file or directory",
						Aliases:  []string{"p"},
						Required: true,
					},
					&cli.PathFlag{
						Name:     "destination",
						Category: "File Management",
						Usage:    "Destination path",
						Aliases:  []string{"d"},
					},
					&cli.StringFlag{
						Name:     "include",
						Category: "File Management",
						Usage:    "Include files matching pattern",
						Aliases:  []string{"i"},
					},
					&cli.StringFlag{
						Name:     "exclude",
						Category: "File Management",
						Usage:    "Exclude files matching pattern",
						Aliases:  []string{"e"},
					},
					&cli.BoolFlag{
						Name:    "recursive",
						Aliases: []string{"r"},
						Usage:   "Recurse through directories",
					},
					&cli.BoolFlag{
						Name:    "force",
						Aliases: []string{"f"},
						Usage:   "Force copy",
					},
				},
				Aliases: []string{"cp"},
				Usage:   "Copy object(s)",
				Action: func(c *cli.Context) error {
					return management.CopyFiles(tasks.CopyArgs{
						TaskArgs: tasks.TaskArgs{
							Source:    c.String("source"),
							Include:   c.String("include"),
							Exclude:   c.String("exclude"),
							Recursive: c.Bool("recursive"),
							Force:     c.Bool("force"),
						},
						Destination: c.String("destination"),
					})
				},
			},
			//Move
			{
				Name: "move",
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:     "source",
						Category: "File Management",
						Usage:    "Path to file or directory that wil be the source",
						Aliases:  []string{"p"},
						Required: true,
					},
					&cli.PathFlag{
						Name:     "destination",
						Category: "File Management",
						Usage:    "Path to file or directory that wil be the destination",
						Aliases:  []string{"d"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "include",
						Category: "File Management",
						Usage:    "Include files matching pattern",
						Aliases:  []string{"i"},
					},
					&cli.StringFlag{
						Name:     "exclude",
						Category: "File Management",
						Usage:    "Exclude files matching pattern",
						Aliases:  []string{"e"},
					},
					&cli.BoolFlag{
						Name:    "recursive",
						Aliases: []string{"r"},
						Usage:   "Recurse through directories",
					},
					&cli.BoolFlag{
						Name:    "force",
						Aliases: []string{"f"},
						Usage:   "Force move",
					},
				},
				Aliases: []string{"mv"},
				Usage:   "Move object(s)",
				Action: func(c *cli.Context) error {
					return management.MoveFiles(tasks.MoveArgs{
						TaskArgs: tasks.TaskArgs{
							Source:    c.String("source"),
							Include:   c.String("include"),
							Exclude:   c.String("exclude"),
							Recursive: c.Bool("recursive"),
							Force:     c.Bool("force"),
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
						Name:     "source",
						Category: "File Management",
						Usage:    "Path to file or directory that wil be the source",
						Aliases:  []string{"p"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "include",
						Category: "File Management",
						Usage:    "Pattern to include files (e.g. *.txt)",
						Aliases:  []string{"i"},
					},
					&cli.StringFlag{
						Name:     "exclude",
						Category: "File Management",
						Usage:    "Pattern to exclude files (e.g. *.exe)",
						Aliases:  []string{"e"},
					},
					&cli.StringFlag{
						Name:     "prefix",
						Category: "File Management",
						Usage:    "Prefix to add to file names",
						Aliases:  []string{"p"},
					},
					&cli.StringFlag{
						Name:     "extension",
						Category: "File Management",
						Usage:    "Rename files with this filetype extension (e.g. .rtf)",
						Aliases:  []string{"e"},
					},
					&cli.BoolFlag{
						Name:    "recursive",
						Aliases: []string{"r"},
						Usage:   "Recurse through directories",
					},
				},
				Aliases: []string{"rn"},
				Usage:   "Rename files",
				Action: func(c *cli.Context) error {
					return management.RenameFiles(tasks.RenameArgs{
						TaskArgs: tasks.TaskArgs{
							Source:    c.String("source"),
							Include:   c.String("include"),
							Exclude:   c.String("exclude"),
							Recursive: c.Bool("recursive"),
						},
						Prefix:    c.String("prefix"),
						Extension: c.String("extension"),
					})
				},
			},
			//TODO: Sync
			{
				Name: "sync",
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:     "source",
						Category: "File Management",
						Usage:    "Path to file or directory that wil be the source",
						Aliases:  []string{"p"},
						Required: true,
					},
					&cli.PathFlag{
						Name:     "destination",
						Category: "File Management",
						Usage:    "Path to file or directory that wil be the destination",
						Aliases:  []string{"d"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "include",
						Category: "File Management",
						Usage:    "Include files matching pattern",
						Aliases:  []string{"i"},
					},
					&cli.StringFlag{
						Name:     "exclude",
						Category: "File Management",
						Usage:    "Exclude files matching pattern",
						Aliases:  []string{"e"},
					},
					&cli.BoolFlag{
						Name:    "recursive",
						Aliases: []string{"r"},
						Usage:   "Recurse through directories",
					},
					&cli.BoolFlag{
						Name:     "force",
						Category: "File Management",
						Usage:    "Force overwrite",
						Aliases:  []string{"f"},
					},
					&cli.BoolFlag{
						Name:     "delete",
						Category: "File Management",
						Usage:    "Delete files in destination that are not in source",
						Aliases:  []string{"-dd", "--delete-destination"},
					},
					&cli.BoolFlag{
						Name:     "destructive",
						Category: "File Management",
						Usage:    "Delete files in source after syncing",
						Aliases:  []string{"-dd"},
					},
				},
				Aliases: []string{"sc"},
				Usage:   "Sync object(s)",
				Action: func(c *cli.Context) error {
					return management.SyncFiles(tasks.SyncArgs{
						TaskArgs: tasks.TaskArgs{
							Source:    c.String("source"),
							Include:   c.String("include"),
							Exclude:   c.String("exclude"),
							Recursive: c.Bool("recursive"),
							Force:     c.Bool("force"),
						},
						Delete:      c.Bool("delete"),
						Destination: c.String("destination"),
						Destructive: c.Bool("destructive"),
					})
				},
			},
			//Delete
			{
				Name: "delete",
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:     "path",
						Category: "File Management",
						Usage:    "Path to file or directory",
						Aliases:  []string{"p"},
						Required: false,
						Hidden:   true,
					},
					&cli.StringFlag{
						Name:     "include",
						Category: "File Management",
						Usage:    "Include files matching pattern",
						Aliases:  []string{"i"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "exclude",
						Category: "File Management",
						Usage:    "Exclude files matching pattern",
						Aliases:  []string{"e"},
					},
					&cli.BoolFlag{
						Name:    "recursive",
						Aliases: []string{"r"},
						Usage:   "Recurse through directories",
					},
					&cli.BoolFlag{
						Name:    "force",
						Aliases: []string{"f"},
						Usage:   "Force delete",
					},
				},
				Aliases: []string{"rm"},
				Usage:   "Delete object(s)",
				Action: func(c *cli.Context) error {
					return management.DeleteFiles(tasks.TaskArgs{
						Source:    c.String("path"),
						Include:   c.String("include"),
						Exclude:   c.String("exclude"),
						Recursive: c.Bool("recursive"),
						Force:     c.Bool("force"),
					})
				},
			},
			{
				Name: "list",
				Flags: []cli.Flag{
					&cli.StringFlag{

						Name:     "--path",
						Category: "File Management",
						Usage:    "Path to file or directory",
						Aliases:  []string{"p"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "include",
						Category: "File Management",
						Usage:    "Include files matching pattern",
						Aliases:  []string{"i"},
						Required: false,
					},
					&cli.StringFlag{
						Name:     "exclude",
						Category: "File Management",
						Usage:    "Exclude files matching pattern",
						Aliases:  []string{"e"},
						Required: false,
					},
					&cli.BoolFlag{
						Name:     "recursive",
						Aliases:  []string{"r"},
						Usage:    "Recurse through directories",
						Required: false,
					},
				},
				Aliases: []string{"ls"},
				Usage:   "List object(s)",
				Action: func(c *cli.Context) error {
					return management.ListFiles(tasks.TaskArgs{
						Source:    c.String("path"),
						Include:   c.String("include"),
						Exclude:   c.String("exclude"),
						Recursive: c.Bool("recursive"),
					})

				},
			},
		},
	}
}
