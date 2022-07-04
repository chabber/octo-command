package cmd

import (
	"fmt"
	"octo-command/cmd/util"
	consts "octo-command/domain"
	"octo-command/models"
	"octo-command/services"
	"strconv"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

var flagServerDefault bool

func NewAddCmd(svc services.SettingsService) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Command for managing server and temp profiles",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	addServerSubCmd := &cobra.Command{
		Use:   "printer",
		Short: "Add printer profile",
		Run:   runAddPrinterSubCmd(svc),
	}
	addServerSubCmd.Flags().BoolVarP(&flagServerDefault, "default", "d", false, "Set profile as default")
	addCmd.AddCommand(addServerSubCmd)

	addTempSubCmd := &cobra.Command{
		Use:   "temp",
		Short: "Add temperature profile",
		Run:   runAddTempSubCmd(svc),
	}
	addCmd.AddCommand(addTempSubCmd)

	return addCmd
}

func runAddTempSubCmd(svc services.SettingsService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {
		var bedTemp, toolTemp float64
		fmt.Printf("args1: %s, args2: %s\n", args[1], args[2])

		// validate bed temperature
		bedTemp, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Println("Bed temperature must be a valid number.")
			return
		}
		if bedTemp > consts.MAX_BED_TEMPERATURE {
			fmt.Println("Temperature exceeds max bed temperature of ", consts.MAX_BED_TEMPERATURE)
		}

		// validate tool temperature
		toolTemp, err = strconv.ParseFloat(args[2], 64)
		if err != nil {
			fmt.Println("Tool temperature must be a valid number.L ", err)
			return
		}
		if toolTemp > consts.MAX_TOOL_TEMPERATURE {
			fmt.Println("Temperature exceeds max tool temperature of ", consts.MAX_TOOL_TEMPERATURE)
			return
		}

		var profile = models.TempProfile{
			Name:     args[0],
			BedTemp:  bedTemp,
			ToolTemp: toolTemp,
		}

		svc.SaveTempProfile(profile)
	}
}

func runAddPrinterSubCmd(svc services.SettingsService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {

		newProfile := getProfileInfo()

		// if setting this profile as default, must make sure no other profiles are marked as default already
		// if there is, remove as default
		if newProfile.Default {
			currentProfiles, err := svc.GetPrinterProfiles()

			if err != nil {
				fmt.Println("error getting server profiles to update to non-default: %v", err)
				return
			}
			for _, p := range currentProfiles {
				if p.Default {
					p.Default = false
					svc.SavePrinterProfile(p)
					break
				}
			}
		}

		svc.SavePrinterProfile(newProfile)
		fmt.Printf("Printer profile with name '%s' successfully saved.\n\n", newProfile.Name)
	}
}

func getProfileInfo() models.PrinterProfile {
	var profile models.PrinterProfile

	profile.Name = prompt.Input("Printer name: ", completer)
	var printerType string
	for printerType != "1" {
		printerType = prompt.Input("Printer type ('1' for OctoPrint): ", completer)
		if printerType != "1" {
			fmt.Println("Only OctoPrint is supported at this time")
		}
	}

	switch printerType {
	case "1":
		{
			profile.Url = prompt.Input("URL (do no include 'http://'): ", completer)
			profile.ApiKey = prompt.Input("API Key: ", completer)
			var def string
			for def != "y" && def != "n" {
				def = prompt.Input("Default (y/n): ", completer)
				if def != "y" && def != "n" {
					fmt.Println("Please enter y/n")
				}
			}
			profile.Default = def == "y"
			//profile.Type = consts.OCTO_PRINT
		}
	}

	return profile
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
