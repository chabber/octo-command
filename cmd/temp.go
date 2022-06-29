package cmd

/*var flagBedTemp, flagToolTemp, flagTempProfile string

// connectCmd represents the start command
var tempCmd = &cobra.Command{
	Use:   "temp {[set] | [monitor]}",
	Short: "Temperature interface for OctoPrint server",
	Long:  "Monitor and set bed and tool temperatures",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p := octoSvc.GetTempProfile(args[0])

		octoSvc.SetBedTemp(p.BedTemp)
		octoSvc.SetToolTemp(p.ToolTemp)
	},
}

var setSubCmd = &cobra.Command{
	Use:                   "set {[--bed=BED_TEMP] [--t=TOOL_TEMP] | [--profile=TEMP_PROFILE]}",
	Short:                 "Set temperature",
	Long:                  "Set temperature of bed/tool",
	DisableFlagsInUseLine: true,
	Run:                   setTemp,
}

func setTemp(cmd *cobra.Command, args []string) {
	// set temp based on profile
	if flagTempProfile != "" {
		p := octoSvc.GetTempProfile(flagTempProfile)
		octoSvc.SetBedTemp(p.BedTemp)
		octoSvc.SetToolTemp(p.ToolTemp)
	} else {
		if flagToolTemp != "" {

		}

		if flagBedTemp != "" {
			f, _ := strconv.ParseFloat(flagBedTemp, 64)
			octoSvc.SetBedTemp(f)
		}
	}
}



func init() {
	tempCmd.AddCommand(monitorSubCmd)
	tempCmd.AddCommand(setSubCmd)
	RootCmd.AddCommand(tempCmd)
}*/
