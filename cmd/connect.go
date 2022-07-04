package cmd

import (
	"fmt"
	"octo-command/cmd/util"
	"octo-command/infrastructure/printer"
	"octo-command/models"
	"octo-command/services"

	"github.com/chabber/go-octoprint"
	"github.com/spf13/cobra"
)

func NewConnectCmd(ps *services.PrinterService, ss services.SettingsService) *cobra.Command {
	connectCmd := &cobra.Command{
		Use:   "connect [profile]",
		Short: "Connect to printer or server",
		Run:   runConnectCmd(ps, ss),
	}

	return connectCmd
}

func runConnectCmd(ps *services.PrinterService, ss services.SettingsService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {
		var status *string
		var err error
		var profile *models.PrinterProfile

		// if no profile arguemnt was passed, we will try to use a default profile
		useDefault := len(args) == 0
		// load profile based on if user specified to use default or not
		if useDefault {
			profile, err = ss.GetDefaultPrinterProfile()
		} else {
			profile, err = ss.GetPrinterProfile(args[0])
		}

		// if any errors loading profile (does not include not finding the profile)
		if err != nil {
			if useDefault {
				fmt.Printf("Error getting default server profile: %s\n", err)
			} else {
				fmt.Printf("Error getting profile for name: '%s'\n", err)
			}
		}

		// if no profile was found
		if profile == nil {
			if useDefault {
				fmt.Println("No default server profile found, create one using `add server`")
			} else {
				fmt.Printf("No server profile found for name: '%s'\n", args[0])
			}
			return
		}

		// connect to server using profile
		ps.Pdp = printer.NewOctoServerService(octoprint.NewClient(profile.Url, profile.ApiKey))
		status, err = ps.Connect(profile)
		if err != nil {
			fmt.Printf("Error connecting to server: %s\n", err)
			return
		}

		ps.Connected = true

		fmt.Printf("Connection status: %s\n", *status)
	}
}
