package services

import "octo-command/models"

// Local storage used by the command line tool
type SettingsService interface {
	GetConfig() (*models.Config, error)

	// server profiles
	SavePrinterProfile(models.PrinterProfile) error
	GetPrinterProfile(name string) (*models.PrinterProfile, error)
	GetPrinterProfiles() ([]models.PrinterProfile, error)
	GetDefaultPrinterProfile() (*models.PrinterProfile, error)
	DeletePrinterProfile(string) error

	// temperature profiles
	SaveTempProfile(models.TempProfile)
	GetTempProfile(string) (models.TempProfile, error)
	GetTempProfiles() ([]models.TempProfile, error)
	DeleteTempProfile(string) error
}
