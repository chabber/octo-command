package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var profileName, profileUrl, profileApiKey string
var longList bool

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server related commands",
	Long:  "Commands for connecting to and maintaining profiles of OctoPrint servers",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var addSubCmd = &cobra.Command{
	Use:   "add [-n name] [-u url] [-k api_key]",
	Short: "Add an OctoPrint server profile",
	Run: func(cmd *cobra.Command, args []string) {
		octoSvc.AddServerProfile(profileName, profileUrl, profileApiKey)
	},
}

var listSubCmd = &cobra.Command{
	Use:   "list",
	Short: "List all server profiles",
	Run: func(cmd *cobra.Command, args []string) {
		profiles := octoSvc.GetServerProfiles()

		// if long listing, print column headers
		if longList {
			fmt.Printf("%-15s | %-25s | %-30s\n", "Name", "URL", "API Key")
			fmt.Printf("%-15s + %-25s + %-30s\n", strings.Repeat("-", 15), strings.Repeat("-", 25), strings.Repeat("-", 30))
		}
		for _, p := range profiles {
			if longList {
				fmt.Printf("%-15s | %-25s | %-30s\n", p.Name, p.Url, p.ApiKey)
			} else {
				fmt.Printf("%s\n", p.Name)
			}
		}
	},
}

func init() {
	addSubCmd.Flags().StringVarP(&profileName, "name", "n", "", "Name for saved server")
	addSubCmd.Flags().StringVarP(&profileUrl, "url", "u", "", "URL for saved server")
	addSubCmd.Flags().StringVarP(&profileApiKey, "apikey", "k", "", "API key for saved server")

	listSubCmd.Flags().BoolVarP(&longList, "long", "l", false, "Show detailed listing")

	serverCmd.AddCommand(addSubCmd)
	serverCmd.AddCommand(listSubCmd)

	RootCmd.AddCommand(serverCmd)
}
