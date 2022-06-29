package list

import (
	"fmt"
	"octo-command/cmd/util"
	"octo-command/services"
	"strings"

	"github.com/spf13/cobra"
)

var flagLongList bool

func NewCmd(svc services.SettingsService) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list {server | temp}",
		Short: "List saved servers or temperature profiles",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	listServersSubCmd := &cobra.Command{
		Use:   "list server",
		Short: "List server profiles",
		Run:   runListServerSubCmd(svc),
	}
	listCmd.AddCommand(listServersSubCmd)

	listTempsSubCmd := &cobra.Command{
		Use:   "list temp",
		Short: "List temperature profiles",
		Run:   runListTempsSubCmd(svc),
	}

	return listCmd
}

func runListTempsSubCmd(svc services.SettingsService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {
		profiles := svc.GetTempProfiles()

		if flagLongList {

		}
	}
}

func runListServerSubCmd(svc services.SettingsService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {
		profiles := svc.GetServerProfiles()

		// if long listing, print column headers
		if flagLongList {
			fmt.Printf("%-15s | %-25s | %-30s | %-7s\n", "Name", "URL", "API Key", "Default")
			fmt.Printf("%-15s+%-25s+%-30s+%-7s\n", strings.Repeat("-", 16), strings.Repeat("-", 27), strings.Repeat("-", 32), strings.Repeat("-", 8))
		}

		// print formatted list of profiles
		for _, p := range profiles {
			if flagLongList {
				fmt.Printf("%-15s | %-25s | %-30s | %-7v\n", p.Name, p.Url, p.ApiKey, p.Default)
			} else {
				fmt.Printf("%s\n", p.Name)
			}
		}
	}
}
