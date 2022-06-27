package cmd

import (
	"octo-command/octo/services"

	"github.com/spf13/cobra"
)

func NewCmd(svc services.PrinterService) *cobra.Command {
	printerCmd := &cobra.Command{
		Use:   "printer",
		Short: "Various commands for the connected printer",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	printFileSubCmd := &cobra.Command{
		Use:   "printfile [filename]",
		Short: "Print a file",
		Long:  "Send request to OctoPrint server to print file",
		Args:  cobra.ExactArgs(1),
		Run:   runPrintFileSubCmd(svc),
	}
	printerCmd.AddCommand(printFileSubCmd)

	homeSubCmd = &cobra.Command{
		Use:   "home",
		Short: "Home position",
		Long:  "Return the print head to home position",
		Run:   runHomeSubCmd(svc),
	}
	printerCmd.AddCommand(homeSubCmd)

	return printerCmd
}

func runPrintFileSubCmd(svc services.PrinterService) RunFunc {
	return func(cmd *cobra.Command, args []string) {
		svc.PrintFile(args[0])
	}
}

func runHomeSubCmd(svc services.PrinterService) RunFunc {
	return func(cmd *cobra.Command, args []string) {
		svc.Home()
	}
}
