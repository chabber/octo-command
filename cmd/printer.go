package cmd

import "github.com/spf13/cobra"

var printerCmd = &cobra.Command{
	Use:   "printer",
	Short: "Various commands for the connected printer",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
var printFileSubCmd = &cobra.Command{
	Use:   "printfile [filename]",
	Short: "Print a file",
	Long:  "Send request to OctoPrint server to print selected file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		octoSvc.PrintFile(args[0])
	},
}

var homeSubCmd = &cobra.Command{
	Use:   "home",
	Short: "Home position",
	Long:  "Return the print head to home position",
	Run: func(cmd *cobra.Command, args []string) {
		octoSvc.Home()
	},
}

func init() {
	printerCmd.AddCommand(homeSubCmd)
	printerCmd.AddCommand(printFileSubCmd)

	RootCmd.AddCommand(printerCmd)
}
