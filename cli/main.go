package main

import (
	//"flag"
	//"fmt"
	//"strconv"

	//"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "synctl"
	app.Usage = "Sanity: restored."
	app.Description = "Synctl is a multi-purpose command line interface for housekeeping and syncronization on both local and remote machines."
	app.Version = "0.0.1"
	app.Authors = []*cli.Author{
		&cli.Author{
			Name: "Brandon Miller",
			Email: "brandon@brandonmiller.io",
		},
	}
	app.Commands = []*cli.Command{
			startWalk(),
	}
	app.Action = startWalk()
}


func startWalk() *cli.Command {

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




// func Help(ctx *cli.Context) error {
// 	ctx.App.Command("help").Run(ctx)

// 	return nil
// }


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
