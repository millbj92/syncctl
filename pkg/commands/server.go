package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/millbj92/synctl/internal/configs"
	"github.com/millbj92/synctl/pkg/management"
	"github.com/millbj92/synctl/pkg/middleware"
	"github.com/millbj92/synctl/pkg/routes"
	"github.com/urfave/cli/v2"
)

func ServerCommands() *cli.Command {
	var fiberApp *fiber.App = nil

    return &cli.Command{
		Name: "server",
		Aliases: []string{"--sv"},
		Usage:  "Start server",
		UsageText: "server [command] [flags]",
		Description: "Starts as a master node to relay all commands to remotes",
		ArgsUsage: "[command] [flags]",
		Subcommands: []*cli.Command{
			&cli.Command{
				Name: "start",
				Action: func(c *cli.Context) error {
					fiberApp = StartServer(c.String("host"), c.Int("port"))
					return nil
				},
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name: "port",
						Aliases: []string{"p"},
						EnvVars: []string{"SERVER_PORT"},
						Value: 8101,
						DefaultText: "8101",
					},
					&cli.StringFlag{
						Name: "host",
						Category: "Server",
						Usage: "Host to listen on",
						EnvVars: []string{"SERVER_HOST"},
						DefaultText: "Server will listen on all interfaces",
						Value: "localhost",
						Aliases: []string{"hs"},
					},
				},
			},
			&cli.Command{
				Name: "stop",
				Action: func(c *cli.Context) error {
					if fiberApp == nil {
						return nil
					}
					return StopServer(fiberApp)
				},
			},
		},
		SkipFlagParsing: false,
		HideHelp: false,
		HelpName: "server",
		BashComplete: func(c *cli.Context) {
			fmt.Fprintf(c.App.Writer, "--host, h, --port, p")
		},
		Before: func(c *cli.Context) error {
			if c.Args().First() == "start" {
			fmt.Fprintf(c.App.Writer, "\n\n\t\t🚀 OH LAWD WE COMIN\n")
			}

			if c.Args().First() == "stop" {
				fmt.Fprintf(c.App.Writer, "\n\n\t\t👋 So long, space cowboy!\n")
			}
			return nil
		},
		After: func(c *cli.Context) error {
			if c.Args().First() == "start" {
			  fmt.Fprintf(c.App.Writer, "Server successfully started. Welcome aboard!\n")
			}
			if c.Args().First() == "stop" {
				fmt.Fprintf(c.App.Writer, "Server successfully stopped. See you next time!\n")
			}
			return nil
		},
		Action: func(c *cli.Context) error {
			c.Command.FullName()
			c.Command.HasName("server")
			c.Command.Names()
			c.Command.VisibleFlags()
			return nil
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			fmt.Fprintf(c.App.Writer, "Error: %s\n", err.Error())
			return err
		},
	};
}


func StartServer(host string, port int) *fiber.App {
	config := configs.ConfigureFiber()

	app := fiber.New(config)
	middleware.LoadMiddleware(app)
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)
	management.StartServer(app, host, port)

	return app
}

func StopServer(a *fiber.App) error {
	return a.Shutdown()
}

