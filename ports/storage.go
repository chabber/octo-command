package ports

import "octo-command/models"

// Local storage used by the command line tool
type StorageDataPort interface {
	GetConfig() (*models.Config, error)

	// server profiles
	SaveServerProfile(models.ServerProfile) error
	GetServerProfile(name string) (*models.ServerProfile, error)
	GetServerProfiles() []models.ServerProfile
	GetDefaultServerProfile() (*models.ServerProfile, error)
	DeleteServerProfile(string) error

	// temperature profiles
	SaveTempProfile(models.TempProfile)
	GetTempProfile(string) models.TempProfile
	GetTempProfiles() []models.TempProfile
	DeleteTempProfile(string) error
}
