package cmd

const (
	LINE_UP     string = "\033[1A"
	LINE_CLEAR  string = "\x1b[2K"
	HIDE_CURSOR string = "\033[?25l"
	SHOW_CURSOR string = "\033[?25h"
)

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

var monitorSubCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor temperature",
	Long:  "Monitor tool and bed temperatures",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		fmt.Printf("%-10s | %-10s | %-10s | %-10s\n", "Device", "Actual", "Target", "Offfset")
		fmt.Printf("%-10s+%-10s+%-10s+%-10s\n", strings.Repeat("-", 11), strings.Repeat("-", 12), strings.Repeat("-", 12), strings.Repeat("-", 9))
		fmt.Print(HIDE_CURSOR)

		quit := make(chan bool)

		go func() {
			for {
				select {
				case <-quit:
					fmt.Println()
					return
				default:
					bedTemp, _ := octoSvc.GetBedTemp()
					toolTemps, _ := octoSvc.GetToolTemp()
					fmt.Printf("%-10s | %-10v | %-10v | %-10v\n", "Bed", bedTemp.Actual, bedTemp.Target, bedTemp.Offset)
					for _, t := range toolTemps {
						fmt.Printf("%-10s | %-10v | %-10v | %-10v\n", t.Label, t.Actual, t.Target, t.Offset)
					}
					time.Sleep(3 * time.Second)
					fmt.Printf("%s%s", LINE_UP, LINE_UP)
				}
			}
		}()

		// wait for keyboard input to kill thread
		fmt.Scanf("%s")
		quit <- true
		fmt.Print(SHOW_CURSOR)
	},
}

func init() {
	tempCmd.AddCommand(monitorSubCmd)
	tempCmd.AddCommand(setSubCmd)
	RootCmd.AddCommand(tempCmd)
}*/
