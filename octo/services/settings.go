package services

import (
	"octo-command/octo/models"
	"octo-command/octo/ports"
)

type SettingsService struct {
	sdp ports.StorageDataPort
}

func NewSettingsService(data ports.StorageDataPort) SettingsService {
	return SettingsService{
		sdp: data,
	}
}

func (os *SettingsService) GetConfig() error {
	os.sdp.GetConfig()

	return nil
}

func (os *SettingsService) SaveTempProfile(t models.TempProfile) error {
	os.sdp.SaveTempProfile(t)

	return nil
}

func (ss *SettingsService) SaveServerProfile(p models.ServerProfile) {
	ss.sdp.SaveServerProfile(p)
}

func (ss *SettingsService) GetDefaultServerProfile() (*models.ServerProfile, error) {
	return ss.sdp.GetDefaultServerProfile()
}

func (ss *SettingsService) GetServerProfile(n string) (*models.ServerProfile, error) {
	return ss.sdp.GetServerProfile(n)
}

func (ss *SettingsService) DeleteServerProfile(n string) error {
	return ss.sdp.DeleteServerProfile(n)
}

func (ss *SettingsService) GetServerProfiles() []models.ServerProfile {
	return ss.sdp.GetServerProfiles()
}

func (ss *SettingsService) GetTempProfile(n string) models.TempProfile {
	return ss.sdp.GetTempProfile(n)
}
