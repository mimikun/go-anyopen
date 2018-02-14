package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(safariOpenCmd)
}

var safariOpenCmd = &cobra.Command{
	Use:   "safari",
	Short: "open in Safari",
	Run: func(cmd *cobra.Command, args []string) {
		htmlfile := args[0]
		execmd := exec.Command("open", "-a", "Safari", string(htmlfile))
		execmd.Run()
	},
}
