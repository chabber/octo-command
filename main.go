package main

import (
	"octocommand/cmd"
	"os"

	cobraprompt "github.com/stromland/cobra-prompt"
)

var simplePrompt = &cobraprompt.CobraPrompt{
	RootCmd:                  cmd.RootCmd,
	AddDefaultExitCommand:    true,
	DisableCompletionCommand: true,
}

func main() {
	if os.Args[1] == "i" {
		simplePrompt.Run()
	} else {
		cmd.RootCmd.Execute()
	}
}
