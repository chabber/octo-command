package cmd

import (
	"octo-command/cmd/util"
	"octo-command/services"

	"github.com/spf13/cobra"
)

func NewDeleteCmd(svc services.SettingsService) *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Server related commands",
		Long:  "Commands for connecting to and maintaining profiles of OctoPrint servers",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	deleteServerSubCmd := &cobra.Command{
		Use:   "delete server",
		Short: "Delete server profile",
		Args:  cobra.ExactArgs(1),
		Run:   runDeleteServerSubCmd(svc),
	}
	deleteCmd.AddCommand(deleteServerSubCmd)

	deleteTempSubCmd := &cobra.Command{
		Use:   "delete temp",
		Short: "Delete server profile",
		Args:  cobra.ExactArgs(1),
		Run:   runDeleteTempSubCmd(svc),
	}
	deleteCmd.AddCommand(deleteTempSubCmd)

	return deleteCmd
}

func runDeleteServerSubCmd(svc services.SettingsService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {
		svc.DeleteServerProfile(args[0])
	}
}

func runDeleteTempSubCmd(svc services.SettingsService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {

	}
}
