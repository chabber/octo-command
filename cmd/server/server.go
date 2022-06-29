package server

import (
	"fmt"
	"octo-command/cmd/util"
	"octo-command/services"

	"github.com/spf13/cobra"
)

var flagLongList bool
var flagServerDefault bool

// NewCmd wraps the cobra command so the service can be passed in
// Create new server command and sub-commands
func NewCmd(svc services.SettingsService) *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Server related commands",
		Long:  "Commands for connecting to and maintaining profiles of OctoPrint servers",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	listSubCmd := &cobra.Command{
		Use:   "list",
		Short: "List all server profiles",
		Run:   runListSubCmd(svc),
	}
	listSubCmd.Flags().BoolVarP(&flagLongList, "long", "l", false, "Show detailed listing")
	serverCmd.AddCommand(listSubCmd)

	deleteSubCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete server profile",
		Args:  cobra.ExactArgs(1),
		Run:   runDeleteSubCmd(svc),
	}
	serverCmd.AddCommand(deleteSubCmd)

	return serverCmd
}

func runDeleteSubCmd(svc services.SettingsService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {
		err := svc.DeleteServerProfile(args[0])
		if err != nil {
			fmt.Printf("Error deleting profile: %s", err)
		}
	}
}

/*
var connectSubCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to OctoPrint server",
	Long:  "Connect to OctoPrint server by server name.  Use `server add` first to create profile.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var status *string
		var err error
		var profile *models.ServerProfile

		useDefault := len(args) == 0

		// load profile based on if user specified to use default or not
		if useDefault {
			profile, err = octoSvc.GetDefaultServerProfile()
		} else {
			profile, err = octoSvc.GetServerProfile(args[0])
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
				fmt.Println("No default server profile found, create one using `server add`")
			} else {
				fmt.Printf("No server profile found for name: '%s'\n", args[0])
			}
			return
		}

		// connect to server using profile
		status, err = octoSvc.Connect(profile)
		if err != nil {
			fmt.Printf("Error connecting to server: %s\n", err)
			return
		}

		fmt.Printf("Connection status: %s\n", *status)
	},
}
*/
