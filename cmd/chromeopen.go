package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(chromeOpenCmd)
}

var chromeOpenCmd = &cobra.Command{
	Use:   "chrome",
	Short: "open in Chrome",
	Run: func(cmd *cobra.Command, args []string) {
		htmlfile := args[0]
		execmd := exec.Command("open", "-a", "Google Chrome", string(htmlfile))
		execmd.Run()
	},
}
