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
	Use:   "go-anyopen",
	Short: "A Simple CLI Tool to open HTML file",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go-anyopen",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-anyopen v1.0")
	},
}
