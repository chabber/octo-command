package cmd

import (
	"fmt"
	"octo-command/cmd/util"
	"octo-command/services"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const (
	LINE_UP     string = "\033[1A"
	LINE_CLEAR  string = "\x1b[2K"
	HIDE_CURSOR string = "\033[?25l"
	SHOW_CURSOR string = "\033[?25h"
)

func NewMonitorCmd(ps *services.PrinterService) *cobra.Command {
	monitorCmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor activity of printer",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	monitorTempCmd := &cobra.Command{
		Use:   "temp",
		Short: "Montior temp of printer",
		Run:   runMonitorTempCmd(ps),
	}
	monitorCmd.AddCommand(monitorTempCmd)

	return monitorCmd
}

func runMonitorTempCmd(ps *services.PrinterService) util.RunFunc {
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
		fmt.Print("\n\n", SHOW_CURSOR)
	}
}
