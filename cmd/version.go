package cmd

import (
	"fmt"

	"github.com/cuttle-ai/web-starter/version"
	"github.com/spf13/cobra"
)

/* This fine contains the version command of the application */

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of web-starter",
	Long:  `All software has versions. This is web-starter's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cuttle.ai web-starter", version.Default, "-- HEAD")
	},
}
