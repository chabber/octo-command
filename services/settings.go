package services

import (
	"octo-command/data"
	"octo-command/models"
)

type settingsService struct {
	sdp data.StorageDataPort
}

func NewSettingsService(data data.StorageDataPort) SettingsService {
	return &settingsService{
		sdp: data,
	}
}

func (ss settingsService) GetConfig() (*models.Config, error) {
	c, _ := ss.sdp.GetConfig()

	return c, nil
}

func (ss settingsService) SavePrinterProfile(p models.PrinterProfile) error {
	return ss.sdp.SavePrinterProfile(p)
}

func (ss settingsService) GetPrinterProfile(n string) (*models.PrinterProfile, error) {
	return ss.sdp.GetPrinterProfile(n)
}

func (ss settingsService) GetPrinterProfiles() ([]models.PrinterProfile, error) {
	p, err := ss.sdp.GetPrinterProfiles()

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (ss settingsService) GetDefaultPrinterProfile() (*models.PrinterProfile, error) {
	profiles, err := ss.sdp.GetPrinterProfiles()

	if err != nil {
		return nil, err
	}

	for _, p := range profiles {
		if p.Default {
			return &p, nil
		}
	}

	return nil, nil
}

func (ss settingsService) DeletePrinterProfile(n string) error {
	return ss.sdp.DeletePrinterProfile(n)
}

func (ss settingsService) SaveTempProfile(t models.TempProfile) {
	ss.sdp.SaveTempProfile(t)
}

func (ss settingsService) GetTempProfile(n string) (models.TempProfile, error) {
	return ss.sdp.GetTempProfile(n)
}

func (ss settingsService) GetTempProfiles() ([]models.TempProfile, error) {
	return ss.sdp.GetTempProfiles()
}

func (ss settingsService) DeleteTempProfile(n string) error {
	return ss.sdp.DeleteTempProfile(n)
}
