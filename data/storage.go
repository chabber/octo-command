package data

import (
	"octo-command/domain"
	ds "octo-command/infrastructure/datastore"
	"octo-command/models"
	"octo-command/ports"
)

type storageDataService struct {
	db ds.DatabasePort
}

func NewStorageDataPort(d ds.DatabasePort) ports.StorageDataPort {
	return &storageDataService{
		db: d,
	}
}

func (sds *storageDataService) GetConfig() (*models.Config, error) {
	var config = models.Config{}
	sds.db.Get(domain.CONFIG_RESOURCE, domain.CONFIG_COLLECTION, config)

	return &config, nil
}

func (sds *storageDataService) SaveServerProfile(s models.ServerProfile) error {
	return sds.db.Save(domain.SERVER_PROFILE_COLLECTION, s.Name, s)
}

func (sds *storageDataService) GetServerProfile(n string) (*models.ServerProfile, error) {
	var p models.ServerProfile

	err := sds.db.Get(n, domain.SERVER_PROFILE_COLLECTION, p)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (sds *storageDataService) GetServerProfiles() []models.ServerProfile {
	var ps []interface{}
	sds.db.GetAll(domain.SERVER_PROFILE_COLLECTION, ps)

	var profiles []models.ServerProfile
	for _, p := range ps {
		profiles = append(profiles, p.(models.ServerProfile))
	}

	return profiles
}

func (sds *storageDataService) GetDefaultServerProfile() (*models.ServerProfile, error) {
	var defaultProfile *models.ServerProfile

	// get all profiles
	profiles := sds.GetServerProfiles()

	for _, p := range profiles {
		if p.Default {
			defaultProfile = &p
			break
		}
	}

	return defaultProfile, nil
}

func (sds *storageDataService) DeleteServerProfile(n string) error {
	sds.db.Delete(domain.SERVER_PROFILE_COLLECTION, n)

	return nil
}

func (sds *storageDataService) SaveTempProfile(s models.TempProfile) {
	sds.db.Save(s.Name, domain.TEMP_PROFILE_COLLECTION, s)
}

func (sds *storageDataService) GetTempProfile(n string) models.TempProfile {
	rec := models.TempProfile{}
	sds.db.Get(n, domain.TEMP_PROFILE_COLLECTION, rec)

	return rec
}

func (sds *storageDataService) DeleteTempProfile(n string) error {
	return sds.db.Delete(domain.TEMP_PROFILE_COLLECTION, n)
}

func (sds *storageDataService) GetTempProfiles() []models.TempProfile {
	var ps []interface{}
	sds.db.GetAll(domain.TEMP_PROFILE_COLLECTION, ps)

	var profiles []models.TempProfile
	for _, p := range ps {
		profiles = append(profiles, p.(models.TempProfile))
	}

	return profiles
}
