package ports

import "octo-command/models"

// Local storage used by the command line tool
type StorageDataPort interface {
	GetConfig() (*models.Config, error)

	// server profiles
	SavePrinterProfile(models.PrinterProfile) error
	GetPrinterProfile(name string) (*models.PrinterProfile, error)
	GetPrinterProfiles() []models.PrinterProfile
	GetDefaultPrinterProfile() (*models.PrinterProfile, error)
	DeletePrinterProfile(string) error

	// temperature profiles
	SaveTempProfile(models.TempProfile)
	GetTempProfile(string) models.TempProfile
	GetTempProfiles() []models.TempProfile
	DeleteTempProfile(string) error
}
