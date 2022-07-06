package cmd

import (
	"octo-command/services"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:           "OctoCommand",
	Short:         "CLI for OctoPrint",
	Long:          `A command line interface for interacting with OctoPrint`,
	SilenceErrors: true,
	SilenceUsage:  false,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(pSvc *services.PrinterService, sSvc services.SettingsService) error {
	registerCommands(pSvc, sSvc)

	return RootCmd.Execute()
}

func registerCommands(pSvc *services.PrinterService, sSvc services.SettingsService) {
	RootCmd.AddCommand(NewAddCmd(sSvc))
	RootCmd.AddCommand(NewDeleteCmd(sSvc))
	RootCmd.AddCommand(NewListCmd(sSvc))
	RootCmd.AddCommand(NewConnectCmd(pSvc, sSvc))
	RootCmd.AddCommand(NewSetCmd(pSvc, sSvc))
	RootCmd.AddCommand(NewMonitorCmd(pSvc))
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.3dprinter.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
