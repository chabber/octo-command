package cmd

import (
	"github.com/spf13/cobra"
)

var homeSubCmd = &cobra.Command{
	Use:   "home",
	Short: "Home position",
	Long:  "Return the print head to home position",
	Run: func(cmd *cobra.Command, args []string) {
		octoSvc.Home()
	},
}
