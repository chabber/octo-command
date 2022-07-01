package cmd

import (
	"octo-command/cmd/util"
	"octo-command/services"

	"github.com/spf13/cobra"
)

func NewConnectCmd(ps *services.PrinterService, ss services.SettingsService) *cobra.Command {
	connectCmd := &cobra.Command{
		Use:   "connect [profile]",
		Args:  cobra.ExactArgs(1),
		Short: "Connect to printer or server",
		Run:   runConnectCmd(ps, ss),
	}

	return connectCmd
}

func runConnectCmd(ps *services.PrinterService, ss services.SettingsService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {
		profile, _ := ss.GetPrinterProfile(args[0])

		ps.Connect(profile)
	}
}
