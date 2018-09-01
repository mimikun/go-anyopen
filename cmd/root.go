package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(versionCmd)
}

// RootCmd is RootCmd
var RootCmd = &cobra.Command{
	Use:   "urlopen",
	Short: "A Simple CLI Tool to open HTML file",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of urlopen",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("urlopen v1.1")
	},
}
