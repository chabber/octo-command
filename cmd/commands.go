package cmd

import (
	"fmt"
	"log"
	svc "octo-command/octo/services"

	"github.com/spf13/cobra"
)

var (
	octoSvc *svc.OctoService

	// flags
	name   string
	url    string
	apiKey string
)

var printFileCmd = &cobra.Command{
	Use:   "printfile",
	Short: "Print a file",
	Long:  "Send request to OctoPrint server to print selected file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		octoSvc.PrintFile(args[0])
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

/*var setBedTempCmd = &cobra.Command{
	Use:   "bedtemp",
	Short: "Bed temperature",
	Long:  "Set the bed temperature",
	Run: func(cmd *cobra.Command, args []string) {
		if f, err := strconv.ParseInt(args[0], 10, 64); err == nil {
			if f <= octo.MAX_BED_TEMPERATURE {
				octoSvc.SetBedTemp(f)
			} else {
				fmt.Println("Temperature exceeds max bed temperature of", octo.MAX_BED_TEMPERATURE)
			}
		} else {
			fmt.Println("Temperature must be a valid number.")
		}
	},
}*/

var addServerCmd = &cobra.Command{
	Use:   "add",
	Short: "Add server",
	Long:  "Add an OctoPrint server as saved server",
	Run: func(cmd *cobra.Command, args []string) {
		octoSvc.AddServer(name, url, apiKey)
	},
}

var getToolTempCmd = &cobra.Command{
	Use:   "gettooltemp",
	Short: "Get tool temp",
	Long:  "Retrieve tool temperatures from OctoPrint server",
	Run: func(cmd *cobra.Command, args []string) {
		temps, err := octoSvc.GetToolTemp()

		if err != nil {
			log.Println(err)
			return
		}

		for _, t := range temps {
			fmt.Printf("[%s]\n", t.Label)
			fmt.Printf("Actual: %.2f, Target: %.2f, Offset: %.2f\n", t.Actual, t.Target, t.Offset)
		}
	},
}

var listFilesCmd = &cobra.Command{
	Use:   "list",
	Short: "List files",
	Long:  "List file from OctoPrint server",
	Run: func(cmd *cobra.Command, args []string) {
		octoSvc.GetFiles(args[0])
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

	RootCmd.AddCommand(homeCmd)
	//RootCmd.AddCommand(setBedTempCmd)
	RootCmd.AddCommand(addServerCmd)
	RootCmd.AddCommand(toolStateCmd)
	RootCmd.AddCommand(uploadFileCmd)
	RootCmd.AddCommand(listFilesCmd)
	RootCmd.AddCommand(getToolTempCmd)
	RootCmd.AddCommand(printFileCmd)

	// Add server command
	addServerCmd.Flags().StringVarP(&name, "name", "n", "", "Name for saved server")
	addServerCmd.Flags().StringVarP(&url, "url", "u", "", "URL for saved server")
	addServerCmd.Flags().StringVarP(&apiKey, "apikey", "k", "", "API key for saved server")
}
