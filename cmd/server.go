package cmd

import (
	"fmt"
	"octo-command/octo/data"
	"octo-command/octo/models"
	"strings"

	"github.com/spf13/cobra"
)

var flagLongList bool
var flagServerDefault bool

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server related commands",
	Long:  "Commands for connecting to and maintaining profiles of OctoPrint servers",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

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

var saveSubCmd = &cobra.Command{
	Use:   "save [name] [url] [api_key]",
	Short: "Save/update an OctoPrint server profile",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		newProfile := models.ServerProfile{
			Name:    args[0],
			Url:     args[1],
			ApiKey:  args[2],
			Default: flagServerDefault,
		}

		// if setting this profile as default, must make sure no other profiles are marked as default already
		// if there is, remove as default
		if newProfile.Default {
			currentProfiles := octoSvc.GetServerProfiles()
			for _, p := range currentProfiles {
				if p.Default {
					p.Default = false
					data.SaveServerProfile(p)
					break
				}
			}
		}

		octoSvc.SaveServerProfile(newProfile)
	},
}

var listSubCmd = &cobra.Command{
	Use:   "list",
	Short: "List all server profiles",
	Run: func(cmd *cobra.Command, args []string) {
		profiles := octoSvc.GetServerProfiles()

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
	},
}

var deleteSubCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete server profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := octoSvc.DeleteServerProfile(args[0])
		if err != nil {
			fmt.Printf("Error deleting profile: %s", err)
		}
	},
}

func init() {
	// list sub-command
	listSubCmd.Flags().BoolVarP(&flagLongList, "long", "l", false, "Show detailed listing")
	serverCmd.AddCommand(listSubCmd)

	// add sub-command
	saveSubCmd.Flags().BoolVarP(&flagServerDefault, "default", "d", false, "Set profile as default")
	serverCmd.AddCommand(saveSubCmd)

	// delete sub-command
	serverCmd.AddCommand(deleteSubCmd)

	// connect sub-command
	serverCmd.AddCommand(connectSubCmd)

	RootCmd.AddCommand(serverCmd)
}
