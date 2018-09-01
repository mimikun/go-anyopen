package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(vivaldiOpenCmd)
}

var vivaldiOpenCmd = &cobra.Command{
	Use:   "vivaldi",
	Short: "open in Vivaldi",
	Run: func(cmd *cobra.Command, args []string) {
		htmlfile := args[0]
		execmd := exec.Command("open", "-a", "Vivaldi", string(htmlfile))
		execmd.Run()
	},
}
