package main

import (
	//"flag"
	//"fmt"
	//"strconv"

	//"github.com/vladimirvivien/go-netbox/netbox"
	//"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/millbj92/synctl/pkg/management"
	"github.com/millbj92/synctl/pkg/models/tasks"
)

func main() {
	app := cli.NewApp()
	app.Name = "synctl"
	app.Usage = "A tool for managing housekeeping tasks"
	app.Description = "Synctl is a multi-purpose command line interface for housekeeping and syncronization on both local and remote machines."
	app.Version = "0.0.1"
	app.Authors = []*cli.Author{
		&cli.Author{
			Name: "Brandon Miller",
			Email: "brandon@brandonmiller.io",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:    "help",
			Aliases: []string{"h"},
			Usage:  "Show help",
			Action: Help,
		},
		{
			Name:    "delete",
			Flags:   []cli.Flag{
				&cli.PathFlag{
					Name:  "path",
					Category: "File Management",
					Usage: "Path to file or directory",
					Aliases: []string{"-p", "--path"},
					Required: true,
				},
				&cli.StringFlag{
					Name: "include",
					Category: "File Management",
					Usage: "Include files matching pattern",
					Aliases: []string{"-i", "--include"},
				},
				&cli.StringFlag{
					Name: "exclude",
					Category: "File Management",
					Usage: "Exclude files matching pattern",
					Aliases: []string{"-e", "--exclude"},
				},
				&cli.BoolFlag{
					Name:  "recursive",
					Aliases: []string{"-r", "--recursive"},
					Usage: "Recurse through directories",
				},
			},
			Aliases: []string{"rm"},
			Usage:  "Delete object(s)",
			Action:  func(c *cli.Context) error {
				return management.DeleteFiles(tasks)
			},
		},
		{
			Name:   "copy",
			Flags:  []cli.Flag{
				&cli.PathFlag{
					Name:  "source",
					Category: "File Management",
					Usage: "Path to file or directory",
					Aliases: []string{"-p", "--path"},
					Required: true,
				},
				&cli.PathFlag{
					Name:  "destination",
					Category: "File Management",
					Usage: "Destination path",
					Aliases: []string{"-d", "--destination"},
				},
				&cli.StringFlag{
					Name: "include",
					Category: "File Management",
					Usage: "Include files matching pattern",
					Aliases: []string{"-i", "--include"},
				},
				&cli.StringFlag{
					Name: "exclude",
					Category: "File Management",
					Usage: "Exclude files matching pattern",
					Aliases: []string{"-e", "--exclude"},
				},
				&cli.BoolFlag{
					Name:  "recursive",
					Aliases: []string{"-r", "--recursive"},
					Usage: "Recurse through directories",
				},
			},
			Aliases: []string{"cp"},
			Usage:  "Copy object(s)",
			Action: management.CopyFiles,
		},
		{
			Name:   "move",
			Flags:  []cli.Flag{
				&cli.PathFlag{
					Name:  "source",
					Category: "File Management",
					Usage: "Path to file or directory that wil be the source",
					Aliases: []string{"-p", "--path"},
					Required: true,
				},
				&cli.PathFlag{
					Name:  "destination",
					Category: "File Management",
					Usage: "Path to file or directory that wil be the destination",
					Aliases: []string{"-d", "--destination"},
					Required: true,
				},
				&cli.StringFlag{
					Name: "include",
					Category: "File Management",
					Usage: "Include files matching pattern",
					Aliases: []string{"-i", "--include"},
				},
				&cli.StringFlag{
					Name: "exclude",
					Category: "File Management",
					Usage: "Exclude files matching pattern",
					Aliases: []string{"-e", "--exclude"},
				},
				&cli.BoolFlag{
					Name:  "recursive",
					Aliases: []string{"-r", "--recursive"},
					Usage: "Recurse through directories",
				},
			},
			Aliases: []string{"mv"},
			Usage:  "Move object(s)",
			Action: management.MoveFiles,
		},
		{
			Name:   "sync",
			Flags:  []cli.Flag{
				&cli.PathFlag{
					Name:  "source",
					Category: "File Management",
					Usage: "Path to file or directory that wil be the source",
					Aliases: []string{"-p", "--path"},
					Required: true,
				},
				&cli.PathFlag{
					Name:  "destination",
					Category: "File Management",
					Usage: "Path to file or directory that wil be the destination",
					Aliases: []string{"-d", "--destination"},
					Required: true,
				},
				&cli.StringFlag{
					Name: "include",
					Category: "File Management",
					Usage: "Include files matching pattern",
					Aliases: []string{"-i", "--include"},
				},
				&cli.StringFlag{
					Name: "exclude",
					Category: "File Management",
					Usage: "Exclude files matching pattern",
					Aliases: []string{"-e", "--exclude"},
				},
				&cli.BoolFlag{
					Name:  "recursive",
					Aliases: []string{"-r", "--recursive"},
					Usage: "Recurse through directories",
				},
				&cli.BoolFlag{
					Name: "force",
					Category: "File Management",
					Usage: "Force overwrite",
					Aliases: []string{"-f", "--force"},
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
					Aliases: []string{"-dd", "--destructive"},
				},
			},
			Aliases: []string{"sync"},
			Usage:  "Sync object(s)",
			Action: management.SyncFiles,
		},
		{
			Name: "rename",
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:  "source",
					Category: "File Management",
					Usage: "Path to file or directory that wil be the source",
					Aliases: []string{"-p", "--path"},
					Required: true,
				},
				&cli.StringFlag{
					Name: "include",
					Category: "File Management",
					Usage: "Pattern to include files (e.g. *.txt)",
					Aliases: []string{"-i", "--include"},
				},
				&cli.StringFlag{
					Name: "exclude",
					Category: "File Management",
					Usage: "Pattern to exclude files (e.g. *.exe)",
					Aliases: []string{"-e", "--exclude"},
				},
				&cli.StringFlag{
					Name: "prefix",
					Category: "File Management",
					Usage: "Prefix to add to file names",
					Aliases: []string{"-p", "--prefix"},
				},
				&cli.StringFlag{
					Name: "extension",
					Category: "File Management",
					Usage: "Rename files with this filetype extension (e.g. .rtf)",
					Aliases: []string{"-e", "--extension"},
				},
				&cli.BoolFlag{
					Name:  "recursive",
					Aliases: []string{"-r", "--recursive"},
					Usage: "Recurse through directories",
				},
			},
			Aliases: []string{"rename"},
			Usage:  "Rename files",
			Action: management.RenameFiles,
		},
	  }
	  app.Action = Help
	}



// func Server () *cli.Command{
//     return &cli.Command{
// 		Name: "server",
// 		Description: "Starts the control server",
// 		Usage: "Start a server to act as a relay between all systems in the cluster",
// 		Aliases: []string{"S"},
// 		Flags: []cli.Flag{
// 			&cli.StringFlag{
// 				Name: "port",
// 				Usage: "The port to listen on",
// 				Value: "8080",
// 				Aliases: []string{"p"},
// 			},
// 			&cli.StringFlag{
// 				Name: "host",
// 				Usage: "The host to listen on",
// 				Value: "localhost",
// 				Aliases: []string{"h"},
// 			},
// 			&cli.StringFlag{
// 				Name: "log-level",
// 				Usage: "The log level to use",
// 				Value: "verbose",
// 				Aliases: []string{"l"},
// 			},
// 		},
// 		Category: "server",
// 		Action: func(ctx *cli.Context) error {
// 			fmt.Println("Synctl Server is starting...")
// 			return nil
// 		},

// 	}
// }

// func Paths() error {
// 	return nil
// }




func Help(ctx *cli.Context) error {
	ctx.App.Command("help").Run(ctx)

	return nil
}


// func ParseArgs(oper string) error {
// 	// Parse the command line arguments
// 	flag.StringVar(&oper, "oper", "", "The operation to perform")
// 	flag.IntVar(&op1, "op1", 0, "The first operand")
// 	flag.IntVar(&op2, "op2", 0, "The second operand")
// 	flag.Parse()

// 	return nil
// }

// func Init(path int) error {

// }
