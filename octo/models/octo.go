package models

type Server struct {
	Name   string
	Url    string
	ApiKey string
}

type Temperature struct {
	Label  string
	Actual float64
	Target float64
	Offset float64
}
