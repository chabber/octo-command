package cmd

import (
	octo "octocommand/octo/services"

	"github.com/spf13/cobra"
)

var (
	octoSvc *octo.OctoService

	// flags
	bedTemp float64

	name   string
	url    string
	apiKey string

	serverName string
)

// connectCmd represents the start command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to OctoPrint server",
	Long:  "Connect to OctoPrint server by supplying URL and API key",
	Run: func(cmd *cobra.Command, args []string) {
		octoSvc.Connect(url, apiKey, name)
	},
}

var homeCmd = &cobra.Command{
	Use:   "home",
	Short: "Return the print head to home position",
	Long:  "Return the print head to home position",
	Run: func(cmd *cobra.Command, args []string) {
		octoSvc.Home()
	},
}

var bedTempCmd = &cobra.Command{
	Use:   "bedtemp",
	Short: "Bed temperature",
	Long:  "Set the bed temperature",
	Run: func(cmd *cobra.Command, args []string) {
		octoSvc.BedTemp(bedTemp)
	},
}

var addServerCmd = &cobra.Command{
	Use:   "add",
	Short: "Add server",
	Long:  "Add an OctoPrint server as saved server",
	Run: func(cmd *cobra.Command, args []string) {
		octoSvc.AddServer(name, url, apiKey)
	},
}

var toolStateCmd = &cobra.Command{
	Use:   "tool",
	Short: "Tool state",
	Long:  "Add an OctoPrint server as saved server",
	Run: func(cmd *cobra.Command, args []string) {
		octoSvc.ToolState()
	},
}

func init() {
	octoSvc = new(octo.OctoService)

	RootCmd.AddCommand(connectCmd)
	RootCmd.AddCommand(homeCmd)
	RootCmd.AddCommand(bedTempCmd)
	RootCmd.AddCommand(addServerCmd)
	RootCmd.AddCommand(toolStateCmd)

	// Connect command
	connectCmd.Flags().StringVarP(&apiKey, "apikey", "k", "", "API key for connecting to OctoPrint service")
	connectCmd.Flags().StringVarP(&url, "url", "u", "", "URL for OctoPrint service")
	connectCmd.Flags().StringVarP(&name, "name", "n", "", "Name of saved OctoPrint server")

	// Bed temp command
	bedTempCmd.Flags().Float64VarP(&bedTemp, "temp", "t", 215, "Bed temperature")

	// Add server command
	addServerCmd.Flags().StringVarP(&name, "name", "n", "", "Name for saved server")
	addServerCmd.Flags().StringVarP(&url, "url", "u", "", "URL for saved server")
	addServerCmd.Flags().StringVarP(&apiKey, "apikey", "k", "", "API key for saved server")
}
