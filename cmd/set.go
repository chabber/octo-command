package cmd

import (
	"octo-command/cmd/util"
	"octo-command/services"
	"strconv"

	"github.com/spf13/cobra"
)

var flagBedTemp string
var flagToolTemp string
var flagTempProfile string
var flagTempOff bool

func NewSetCmd(ps *services.PrinterService, ss services.SettingsService) *cobra.Command {
	setCmd := &cobra.Command{
		Use:   "set",
		Short: "Commands for setting various features of printer",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	setTempCmd := &cobra.Command{
		Use:   "temp",
		Short: "Set bed and tool temps",
		Run:   runSetTempCmd(ps, ss),
	}
	setTempCmd.Flags().StringVarP(&flagBedTemp, "bed", "b", "", "Set bed temp")
	setTempCmd.Flags().StringVarP(&flagToolTemp, "tool", "t", "", "Set tool temp")
	setTempCmd.Flags().StringVarP(&flagTempProfile, "profile", "p", "", "Set temp by profile")
	setTempCmd.Flags().BoolVarP(&flagTempOff, "off", "o", false, "Turn off heaters")
	setCmd.AddCommand(setTempCmd)

	return setCmd
}

func runSetTempCmd(ps *services.PrinterService, ss services.SettingsService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {

		// set temp based on profile
		if flagTempProfile != "" {
			p, _ := ss.GetTempProfile(flagTempProfile)
			ps.SetBedTemp(p.BedTemp)
			ps.SetToolTemp(p.ToolTemp)
		} else {
			if flagTempOff {
				ps.SetToolTemp(0)
				ps.SetBedTemp(0)
			} else if flagToolTemp != "" {
				temp, _ := strconv.ParseFloat(flagToolTemp, 64)
				ps.SetToolTemp(temp)
			} else if flagBedTemp != "" {
				temp, _ := strconv.ParseFloat(flagBedTemp, 64)
				ps.SetBedTemp(temp)
			}
		}

	}
}
