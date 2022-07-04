package main

import (
	"bufio"
	"fmt"
	"octo-command/cmd"
	"octo-command/data"
	"octo-command/models"
	"octo-command/services"
	"os"

	"github.com/common-nighthawk/go-figure"
	cobraprompt "github.com/stromland/cobra-prompt"
)

var settingsSvc services.SettingsService
var printerSvc *services.PrinterService

var simplePrompt = &cobraprompt.CobraPrompt{
	RootCmd:                  cmd.RootCmd,
	AddDefaultExitCommand:    true,
	DisableCompletionCommand: true,

	OnErrorFunc: func(err error) {
		cmd.RootCmd.PrintErrln(err)
	},
}

func main() {
	settingsSvc = services.NewSettingsService(
		data.NewScribbleDataService(),
	)
	printerSvc = services.NewPrinterService(nil)
	cmd.Execute(printerSvc, settingsSvc)
	simplePrompt.Run()
}

func onboard() *models.PrinterProfile {
	myFigure := figure.NewColorFigure("OctoCommand", "", "blue", true)
	myFigure.Print()

	// attempt to load config
	// if not found, force user through onboarding
	c, err := settingsSvc.GetConfig()
	if err != nil {
		fmt.Println("error retrieving config: %v", err)
	}

	// if found, we are done and leave user at the prompt
	if c != nil {
		return nil
	}

	fmt.Println()
	fmt.Print("First off, there are a few pieces of information needed about your printer and server.  Let's get an OctoPrint server profile set up.\n")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Name of printer: ")
	printerName, _ := reader.ReadString('\n')
	fmt.Print("OctoPrint server address (don't iclude http://): ")
	url, _ := reader.ReadString('\n')
	fmt.Print("OctoPrint API key: ")
	apiKey, _ := reader.ReadString('\n')

	s := models.PrinterProfile{
		Name:    printerName,
		Url:     url,
		ApiKey:  apiKey,
		Default: true,
	}

	settingsSvc.SavePrinterProfile(s)

	return &s
}
