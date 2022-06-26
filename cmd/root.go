package cmd

import (
	"fmt"
	"octo-command/octo/data"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:           "OctoCommand",
	Short:         "CLI for OctoPrint",
	Long:          `A command line interface for interacting with OctoPrint`,
	SilenceErrors: true,
	SilenceUsage:  false,
	Run:           onboard,
}

func onboard(cmd *cobra.Command, args []string) {
	fmt.Println("Welcomd to OctoCommand!")
	// attempt to load config
	// if not found, force user through onboarding
	c, _ := data.GetConfig()

	// if found, we are done and leave user at the prompt
	if c != nil {
		return
	}

	fmt.Println()
	fmt.Println("First off, there are a few pieces of information needed about your printer and server.")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	fmt.Println("in execute()")
	err := RootCmd.Execute()
	if err != nil {
		fmt.Print("exiting!!!!")
		os.Exit(1)
	}
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
