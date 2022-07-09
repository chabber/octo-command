package printer

import (
	"fmt"
	"octo-command/cmd/util"
	"octo-command/services"
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

	monitorSubCmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor temperature",
		Long:  "Monitor tool and bed temperatures",
		Args:  cobra.NoArgs,
		Run:   runMonitorSubCmd(ps),
	}
	printerCmd.AddCommand(monitorSubCmd)

	return printerCmd
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
					toolTemps, _ := ps.GetToolTemp()
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
