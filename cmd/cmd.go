//Package cmd have the commands offered by web-starter application
package cmd

import (
	"fmt"
	"os"

	web_server "github.com/cuttle-ai/web-starter/cmd/web-server"

	"github.com/spf13/cobra"
)

/* This file contains the main app command specs */

//rootCmd is the root command of the application
var rootCmd = &cobra.Command{
	Use:   "web-starter",
	Short: "web-starter is a web-server code genertaion tool",
	Long:  `It is the generator used across cuttle.ai web server`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(web_server.WebServerCmd)
}

//Execute the main application command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
