package main

import (
	"bufio"
	"fmt"
	"octo-command/cmd"
	"octo-command/infrastructure/datastore"
	"octo-command/infrastructure/printer"
	"octo-command/octo/data"
	"octo-command/octo/models"
	"octo-command/octo/services"
	"os"

	"github.com/chabber/go-octoprint"
	"github.com/common-nighthawk/go-figure"
	scribble "github.com/nanobox-io/golang-scribble"
	cobraprompt "github.com/stromland/cobra-prompt"
)

var SettingsSvc services.SettingsService
var PrinterSvc services.PrinterService

var simplePrompt = &cobraprompt.CobraPrompt{
	RootCmd:                  cmd.RootCmd,
	AddDefaultExitCommand:    true,
	DisableCompletionCommand: true,

	OnErrorFunc: func(err error) {
		cmd.RootCmd.PrintErrln(err)
	},
}

func main() {
	// storage service
	db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	SettingsSvc = services.NewSettingsService(
		data.NewStorageDataPort(datastore.NewScribbleDatabaseService(db)),
	)

	// onboard first time users
	profile := onboard()

	// printer service
	c := octoprint.NewClient(profile.Url, profile.ApiKey)
	r := octoprint.ConnectionRequest{}
	resp, err := r.Do(c)
	if err != nil {
		return nil, err
	}

	PrinterSvc = *services.NewPrinterService(printer.NewOctoServerPort(c))

	simplePrompt.Run()
}

func onboard() *models.ServerProfile {
	myFigure := figure.NewColorFigure("OctoCommand", "", "blue", true)
	myFigure.Print()

	// attempt to load config
	// if not found, force user through onboarding
	c := settingsSvc.GetConfig()

	// if found, we are done and leave user at the prompt
	if c != nil {
		return nil
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

	settingsSvc.SaveServerProfile(s)

	return s
}
