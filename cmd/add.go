package cmd

import (
	"fmt"
	"octo-command/cmd/util"
	consts "octo-command/domain"
	"octo-command/models"
	"octo-command/services"
	"strconv"

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
		Use:   "add printer",
		Short: "Add printer profile",
		Run:   runAddServerSubCmd(svc),
	}
	addServerSubCmd.Flags().BoolVarP(&flagServerDefault, "default", "d", false, "Set profile as default")
	addCmd.AddCommand(addServerSubCmd)

	addTempSubCmd := &cobra.Command{
		Use:   "add temp",
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

func runAddServerSubCmd(svc services.SettingsService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {
		newProfile := models.PrinterProfile{
			Name:    args[0],
			Url:     args[1],
			ApiKey:  args[2],
			Default: flagServerDefault,
		}

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
	}
}
