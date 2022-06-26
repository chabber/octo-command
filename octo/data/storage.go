package data

import (
	"encoding/json"
	"fmt"
	"octo-command/octo/db"
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
	db db.DatabasePort
}

func NewStorageDataPort(d db.DatabasePort) ports.StorageDataPort {
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

func (sds *storageDataService) SaveServerProfile(s models.ServerProfile) {
	sds.db.Save(SERVER_PROFILE_COLLECTION, s.Name, s)
}

func (sds *storageDataService) GetServerProfile(n string) (*models.ServerProfile, error) {
	db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}

	rec := &models.ServerProfile{}
	err = db.Read(SERVER_PROFILE_COLLECTION, n, rec)
	if err != nil {
		return nil, fmt.Errorf("error loading profile")
	}

	return rec, nil
}

func (sds *storageDataService) GetServerProfiles() []models.ServerProfile {
	db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	var profiles []models.ServerProfile

	ps, _ := db.ReadAll(SERVER_PROFILE_COLLECTION)
	for _, p := range ps {
		var profile models.ServerProfile
		json.Unmarshal([]byte(p), &profile)

		profiles = append(profiles, profile)
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

func (sds *storageDataService) DeleteServerProfile(n string) (err error) {
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
