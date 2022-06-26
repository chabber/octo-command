package models

type ServerProfile struct {
	Name    string
	Url     string
	ApiKey  string
	Default bool
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
}
