package printer

import (
	"fmt"
	"octo-command/cmd/util"
	"octo-command/services"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var flagTempProfile string
var flagBedTemp string
var flagToolTemp string

const (
	LINE_UP     string = "\033[1A"
	LINE_CLEAR  string = "\x1b[2K"
	HIDE_CURSOR string = "\033[?25l"
	SHOW_CURSOR string = "\033[?25h"
)

func NewCmd(ps services.PrinterService, ss services.SettingsService) *cobra.Command {
	printerCmd := &cobra.Command{
		Use:   "printer",
		Short: "Various commands for the connected printer",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	printFileSubCmd := &cobra.Command{
		Use:   "printfile [filename]",
		Short: "Print a file",
		Long:  "Send request to OctoPrint server to print file",
		Args:  cobra.ExactArgs(1),
		Run:   runPrintFileSubCmd(ps),
	}
	printerCmd.AddCommand(printFileSubCmd)

	homeSubCmd := &cobra.Command{
		Use:   "home",
		Short: "Home position",
		Long:  "Return the print head to home position",
		Run:   runHomeSubCmd(ps),
	}
	printerCmd.AddCommand(homeSubCmd)

	tempCmd := &cobra.Command{
		Use:   "temp {[set] | [monitor]}",
		Short: "Temperature interface for OctoPrint server",
		Long:  "Monitor and set bed and tool temperatures",
		Args:  cobra.ExactArgs(1),
		Run:   runTempCmd(ps, ss),
	}
	printerCmd.AddCommand(tempCmd)

	monitorSubCmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor temperature",
		Long:  "Monitor tool and bed temperatures",
		Args:  cobra.NoArgs,
		Run:   runMonitorSubCmd(ps),
	}
	printerCmd.AddCommand(monitorSubCmd)

	setSubCmd := &cobra.Command{
		Use:                   "set {[--bed=BED_TEMP] [--t=TOOL_TEMP] | [--profile=TEMP_PROFILE]}",
		Short:                 "Set temperature",
		Long:                  "Set temperature of bed/tool",
		DisableFlagsInUseLine: true,
		Run:                   runSetSubCmd(ps, ss),
	}
	printerCmd.AddCommand(setSubCmd)

	return printerCmd
}

func runSetSubCmd(ps services.PrinterService, ss services.SettingsService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {
		// set temp based on profile
		if flagTempProfile != "" {
			p := ss.GetTempProfile(flagTempProfile)
			ps.SetBedTemp(p.BedTemp)
			ps.SetToolTemp(p.ToolTemp)
		} else {
			if flagToolTemp != "" {
				f, _ := strconv.ParseFloat(flagBedTemp, 64)
				ps.SetToolTemp(f)
			}

			if flagBedTemp != "" {
				f, _ := strconv.ParseFloat(flagBedTemp, 64)
				ps.SetBedTemp(f)
			}
		}
	}
}

func runMonitorSubCmd(ps services.PrinterService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {
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
					bedTemp, _ := ps.GetBedTemp()
					toolTemps := ps.GetToolTemp()
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
	}
}

func runTempCmd(ps services.PrinterService, ss services.SettingsService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {
		p := ss.GetTempProfile(args[0])

		ps.SetBedTemp(p.BedTemp)
		ps.SetToolTemp(p.ToolTemp)
	}
}

func runPrintFileSubCmd(svc services.PrinterService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {
		svc.PrintFile(args[0])
	}
}

func runHomeSubCmd(svc services.PrinterService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {
		svc.Home()
	}
}
