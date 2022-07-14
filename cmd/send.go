package cmd

import (
	"fmt"
	"octo-command/cmd/util"
	"octo-command/services"

	"github.com/spf13/cobra"
)

func NewSendCmd(ps *services.PrinterService) *cobra.Command {
	sendCmd := &cobra.Command{
		Use:   "send [command]",
		Short: "Send commands directly to printer",
		Args:  cobra.ExactArgs(1),
		Run:   send(ps),
	}

	return sendCmd
}

func send(ps *services.PrinterService) util.RunFunc {
	return func(cmd *cobra.Command, args []string) {
		err := ps.SendCommand(args[0])

		if err != nil {
			fmt.Println("error while sending command: ", err)
		}
	}
}
