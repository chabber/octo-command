package services

import (
	"octo-command/octo/data"
	"octo-command/octo/models"
)

type OctoService struct {
	octoData        OctoDataPort
	storageDataPort StorageDataPort
}

func (os *OctoService) AddTempProfile(n string, bed float64, tool float64) error {
	tp := models.TempProfile{
		Name:     n,
		BedTemp:  bed,
		ToolTemp: tool,
	}

	data.SaveTempProfile(tp)

	return nil
}

func (os *OctoService) SaveServerProfile(p models.ServerProfile) {
	data.SaveServerProfile(p)
}

func (os *OctoService) GetDefaultServerProfile() (*models.ServerProfile, error) {
	return data.GetDefaultServerProfile()
}

func (os *OctoService) GetServerProfile(n string) (*models.ServerProfile, error) {
	return data.GetServerProfile(n)
}

func (os *OctoService) DeleteServerProfile(n string) (err error) {
	return data.DeleteServerProfile(n)
}

func (os *OctoService) GetServerProfiles() []models.ServerProfile {
	return data.GetServerProfiles()
}

func (os *OctoService) GetTempProfile(n string) models.TempProfile {
	return data.GetTempProfile(n)
}
