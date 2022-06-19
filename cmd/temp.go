package cmd

import (
	"github.com/spf13/cobra"
)

// connectCmd represents the start command
var tempCmd = &cobra.Command{
	Use:   "temp [temp_profile]",
	Short: "Temperature interface for OctoPrint server",
	Long:  "Monitor and set bed and tool temperatures",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p := octoSvc.GetTempProfile(args[0])

		octoSvc.SetBedTemp(p.BedTemp)
		octoSvc.SetToolTemp(p.ToolTemp)
	},
}

func init() {
	RootCmd.AddCommand(tempCmd)
}
