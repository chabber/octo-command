package main

import (
	"octo-command/cmd"

	cobraprompt "github.com/stromland/cobra-prompt"
)

var simplePrompt = &cobraprompt.CobraPrompt{
	RootCmd:                  cmd.RootCmd,
	AddDefaultExitCommand:    true,
	DisableCompletionCommand: true,
	OnErrorFunc: func(err error) {
		cmd.RootCmd.PrintErrln(err)
	},
}

func main() {
	simplePrompt.Run()
}
