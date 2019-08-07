//Package web-server has the commands for web server
package web_server

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

/* This fine contains the web-server command of the application */

//WebServerCmd for the web sevrer related web starter boilerplate code
var WebServerCmd = &cobra.Command{
	Use:   "web-server",
	Short: "Web-server code management",
	Long:  `Web server boilerplate code creation is done from this command`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	},
}
