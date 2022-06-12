package cmd

import (
	"fmt"
	"octo-command/octo"
	svc "octo-command/octo/services"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	octoSvc *svc.OctoService

	// flags
	name   string
	url    string
	apiKey string
)

// connectCmd represents the start command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to OctoPrint server",
	Long:  "Connect to OctoPrint server by supplying URL and API key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		octoSvc.Connect(args[0])
	},
}

var homeCmd = &cobra.Command{
	Use:   "home",
	Short: "Home position",
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
		if f, err := strconv.ParseFloat(args[0], 64); err == nil {
			if f <= octo.MAX_BED_TEMPERATURE {
				octoSvc.BedTemp(f)
			} else {
				fmt.Println("Temperature exceeds max bed temperature of", octo.MAX_BED_TEMPERATURE)
			}
		} else {
			fmt.Println("Temperature must be a valid number.")
		}
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

var listFilesCmd = &cobra.Command{
	Use:   "list",
	Short: "List files",
	Long:  "List file from OctoPrint server",
	Run: func(cmd *cobra.Command, args []string) {
		octoSvc.ListFiles(args[0])
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

var uploadFileCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload file",
	Long:  "Upload a file for printing to the OctoPrint server",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		src := args[0]
		dst := src
		if len(args) == 2 {
			dst = args[1]
		}
		octoSvc.UploadFile(src, dst)
	},
}

func init() {
	octoSvc = new(svc.OctoService)

	RootCmd.AddCommand(connectCmd)
	RootCmd.AddCommand(homeCmd)
	RootCmd.AddCommand(bedTempCmd)
	RootCmd.AddCommand(addServerCmd)
	RootCmd.AddCommand(toolStateCmd)
	RootCmd.AddCommand(uploadFileCmd)
	RootCmd.AddCommand(listFilesCmd)

	// Add server command
	addServerCmd.Flags().StringVarP(&name, "name", "n", "", "Name for saved server")
	addServerCmd.Flags().StringVarP(&url, "url", "u", "", "URL for saved server")
	addServerCmd.Flags().StringVarP(&apiKey, "apikey", "k", "", "API key for saved server")
}
