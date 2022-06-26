package main

import (
	"bufio"
	"fmt"
	"octo-command/cmd"
	"octo-command/octo/data"
	"octo-command/octo/models"
	"os"

	cobraprompt "github.com/stromland/cobra-prompt"
)

type OctoDataPort interface {
}

var simplePrompt = &cobraprompt.CobraPrompt{
	RootCmd:                  cmd.RootCmd,
	AddDefaultExitCommand:    true,
	DisableCompletionCommand: true,

	OnErrorFunc: func(err error) {
		cmd.RootCmd.PrintErrln(err)
	},
}

func main() {
	fmt.Println("Welcomd to OctoCommand!")
	// attempt to load config
	// if not found, force user through onboarding
	c, _ := data.GetConfig()

	// if found, we are done and leave user at the prompt
	if c != nil {
		return
	}

	fmt.Println()
	fmt.Println("First off, there are a few pieces of information needed about your printer and server.  Let's get an OctoPrint server profile set up.\n")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Name of printer: ")
	printerName, _ := reader.ReadString('\n')
	fmt.Print("OctoPrint server address (don't iclude http://): ")
	url, _ := reader.ReadString('\n')
	fmt.Print("OctoPrint API key: ")
	apiKey, _ := reader.ReadString('\n')

	s := models.ServerProfile{
		Name:    printerName,
		Url:     url,
		ApiKey:  apiKey,
		Default: true,
	}

	simplePrompt.Run()
}
