package printer

import "octo-command/octo/models"

// Interface to OctoPrint server
type PrinterDataPort interface {
	PrintFile(string) error
	GetToolTemp() []*models.Temperature
	GetBedTemp() (*models.Temperature, error)
	UploadFile(src string, dst string)
	ToolState()
	GetFiles(dir string) []models.FileInformation
	Connect(models.ServerProfile) (state *string, err error)
	Home()
	SetBedTemp(temp float64)
	SetToolTemp(temp float64)
}
