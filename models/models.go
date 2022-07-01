package models

import "octo-command/domain"

// PrinterProfile holds settings for all different types of printers. So, some values won't be
// present in all profiles.  Front end will have to hanlde getting the correct values for the
// type of printer.
type PrinterProfile struct {
	// common values
	Name    string
	Type    domain.PrinterCommType
	Default bool

	// Octo Print values
	Url    string
	ApiKey string

	// Direct values
}

type TempProfile struct {
	Name     string  `json: "Name"`
	BedTemp  float64 `json: "BedTemp"`
	ToolTemp float64 `json: "ToolTemp"`
}

type Temperature struct {
	Label  string
	Actual float64
	Target float64
	Offset float64
}

type Config struct {
	NumHotEnds int
	Onboarded  bool
}

type FileInformation struct {
	Name     string
	Children []FileInformation
	IsFolder bool
}
