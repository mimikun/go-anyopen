package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(firefoxOpenCmd)
}

var firefoxOpenCmd = &cobra.Command{
	Use:   "firefox",
	Short: "open in Firefox",
	Run: func(cmd *cobra.Command, args []string) {
		htmlfile := args[0]
		execmd := exec.Command("open", "-a", "Firefox", string(htmlfile))
		execmd.Run()
	},
}
