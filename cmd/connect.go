package cmd

import "github.com/spf13/cobra"

// connectCmd represents the start command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to OctoPrint server",
	Long:  "Connect to OctoPrint server by server name.  Use `addserver` first.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		octoSvc.Connect(args[0])
	},
}

func init() {
	RootCmd.AddCommand(connectCmd)
}
