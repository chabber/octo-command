package cmd

import (
	"fmt"
	"octo-command/octo"
	"strconv"

	"github.com/spf13/cobra"
)

// connectCmd represents the start command
var addTempCmd = &cobra.Command{
	Use:   "addtemp [name] [bed_temp] [tool_temp]",
	Short: "Add temp profile",
	Long:  "Add a bed and tool temp as a temp profile",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		var bedTemp, toolTemp float64
		fmt.Printf("args1: %s, args2: %s\n", args[1], args[2])

		// validate bed temperature
		bedTemp, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Println("Bed temperature must be a valid number.")
			return
		}
		if bedTemp > octo.MAX_BED_TEMPERATURE {
			fmt.Println("Temperature exceeds max bed temperature of ", octo.MAX_BED_TEMPERATURE)
		}

		// validate tool temperature
		toolTemp, err = strconv.ParseFloat(args[2], 64)
		if err != nil {
			fmt.Println("Tool temperature must be a valid number.L ", err)
			return
		}
		if toolTemp > octo.MAX_TOOL_TEMPERATURE {
			fmt.Println("Temperature exceeds max tool temperature of ", octo.MAX_TOOL_TEMPERATURE)
			return
		}

		octoSvc.AddTempProfile(args[0], bedTemp, toolTemp)
	},
}

func init() {
	RootCmd.AddCommand(addTempCmd)
}
