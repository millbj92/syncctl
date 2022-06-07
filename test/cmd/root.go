/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Synctl [command] [flags]",
	Short: "Sanity: Restored.",
	Long: `A cli application and server for housekeeping and syncronization on both local and remote machines..

	Commands:
	Note: You will not be able to perform any action on the server
	until you have ran the server at least once and logged in successfully.
	The system will generate you a one-time token that you will need to use to login.
	Once logged in, you will need to reset it.

		[Security and Authentication]
		auth [--username --password --t token (optional) -k key (ssh, optional) -g key (gpg, optional)]
		server
		--init [-u username] 			Used for first time setup. Will create a new default user and password.
		[-p port] [-h host] [-l log-level]
		gpg-keys [--list, --add, --remove, --import, --export]
		ssh-keys [--list, --add, --remove, --import, --export] (must be at least 4096 bits)

		[Users and Groups]
		user [--list, --add, --remove, --modify, --role [-rid role-id]] (default: 0 - DISABLED)]
		group [--list, --add, --remove, --modify --add-user [-uid], --remove-user [-uid], --add-role [-rid], --remove-role [-rid]]
		roles [--list, --add, --remove, --modify --add-user [-uid], --remove-user [-uid]]
		permissions [--list, --add, --remove, --modify --add-to-role [-rid], --remove-from-role [-rid]]

		[Remote Connections]
		ssh-keys [--list, --add, --remove, --import, --export]
		gpg-keys [--list, --add, --remove, --import, --export] (must be at least 4096 bits)
	    connection [--list, --add, --remove, --modify --add-user [-uid], --remove-user [-uid],
		--add-role [-rid], --remove-role [-rid]]

		[Tasks]
		task [--list, --remove, --update, --run, --stop, --restart, --pause, --resume, --logs [-tid task-id], --logs-tail [-tid task-id], --status [-tid task-id], --logs-download [-tid task-id]]
		task --add:
		 [-s source]	  		Source directory
		 [-o destination] 		Destination directory (if mode = copy, move, or sync)
		 [-r recursive]			If agent should recurse through subdirectories, or just perform actions on the top level directory
		 [-m mode]              Available actions: [scan, copy, move, delete, sync, rename]
		 [-c compression]       If agent should compress files before sending them to the destination.
		 [-e exclude]			If agent should exclude file_types from the transfer. (.mp3, .png, etc)
		 [-i include]  			If agent should explicitly include file_types in the transfer. (.mp3, .png, etc)
		 [-f file]		        If agent should perform actions on a single file.
		 [-l log-level]			Log level for the agent.
		 [-n name]				Name of the task.
		 [-i interval]			Interval at which task should be ran (in CRON format)
		 [-d description]		Description of the task.
		 [-D destructive]		If the task should be delete files after moving or copying.

		[Links]
		link --list, --add, --remove, --update

		[Remote Config]
		Used to view and edit configuration valuess on remote devices.
		remote-config --list, --add, --remove, --update, --get, --set, --get-all, --set-all
	For example:

		$ synctl server -p 8080 -h localhost -l verbose

	The line above will start a relay server on localhost, on port 8080, with verbose logging enabled.


   `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
