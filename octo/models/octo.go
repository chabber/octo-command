package models

type ServerProfile struct {
	Name   string
	Url    string
	ApiKey string
}

type TempProfile struct {
	Name     string `json: "Name"`
	BedTemp  int    `json: "BedTemp"`
	ToolTemp int    `json: "ToolTemp"`
}

type Temperature struct {
	Label  string
	Actual float64
	Target float64
	Offset float64
}
