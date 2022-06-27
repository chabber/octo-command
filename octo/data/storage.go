package data

import (
	"fmt"
	ds "octo-command/infrastructure/datastore"
	"octo-command/octo/models"
	"octo-command/octo/ports"

	scribble "github.com/nanobox-io/golang-scribble"
)

const (
	SERVER_PROFILE_COLLECTION string = ".store/server_profiles"
	TEMP_PROFILE_COLLECTION   string = ".store/temp_profiles"
	CONFIG_COLLECTION         string = ".store"
	CONFIG_RESOURCE           string = "config"
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
	db, err := scribble.New(".", nil)
	if err != nil {
		return nil, err
	}

	rec := models.Config{}
	err = db.Read(CONFIG_COLLECTION, CONFIG_RESOURCE, rec)
	if err != nil {
		return nil, err
	}

	return &rec, nil
}

func (sds *storageDataService) SaveServerProfile(s models.ServerProfile) error {
	return sds.db.Save(SERVER_PROFILE_COLLECTION, s.Name, s)
}

func (sds *storageDataService) GetServerProfile(n string) (*models.ServerProfile, error) {
	var p models.ServerProfile

	err := sds.db.Get(n, SERVER_PROFILE_COLLECTION, p)

	if err != nil {
		return nil, err
	}

	return &p, nil

	/*db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}

	rec := &models.ServerProfile{}
	err = db.Read(SERVER_PROFILE_COLLECTION, n, rec)
	if err != nil {
		return nil, fmt.Errorf("error loading profile")
	}

	return rec, nil*/
}

func (sds *storageDataService) GetServerProfiles() []models.ServerProfile {
	var ps []interface{}
	sds.db.GetAll(SERVER_PROFILE_COLLECTION, ps)

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
	db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	err = db.Delete(SERVER_PROFILE_COLLECTION, n)
	if err != nil {
		return err
	}

	return nil
}

func (sds *storageDataService) SaveTempProfile(s models.TempProfile) {
	db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	db.Write(TEMP_PROFILE_COLLECTION, s.Name, s)
}

func (sds *storageDataService) GetTempProfile(n string) models.TempProfile {
	db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	rec := models.TempProfile{}
	db.Read("TEMP_PROFILE_STORE", n, &rec)

	return rec
}
