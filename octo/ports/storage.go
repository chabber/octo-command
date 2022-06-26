package ports

import "octo-command/octo/models"

// Local storage used by the command line tool
type StorageDataPort interface {
	GetConfig() (*models.Config, error)
	SaveServerProfile(models.ServerProfile)
	GetServerProfile(name string) (*models.ServerProfile, error)
	GetServerProfiles() []models.ServerProfile
	GetDefaultServerProfile() (*models.ServerProfile, error)
	DeleteServerProfile(string) error
	SaveTempProfile(models.TempProfile)
	GetTempProfile(string) models.TempProfile
}
