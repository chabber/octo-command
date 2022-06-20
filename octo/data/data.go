package data

import (
	"encoding/json"
	"fmt"
	"octo-command/octo/models"

	scribble "github.com/nanobox-io/golang-scribble"
)

const (
	SERVER_PROFILE_STORE string = ".store/server_profiles"
	TEMP_PROFILE_STORE   string = ".store/temp_profiles"
)

func SaveServerProfile(s models.ServerProfile) {
	db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	db.Write(SERVER_PROFILE_STORE, s.Name, s)
}

func GetDefaultServerProfile() (*models.ServerProfile, error) {
	var defaultProfile *models.ServerProfile

	// get all profiles
	profiles := GetServerProfiles()

	for _, p := range profiles {
		if p.Default {
			defaultProfile = &p
			break
		}
	}

	return defaultProfile, nil
}

func GetServerProfile(n string) (*models.ServerProfile, error) {
	db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}

	rec := &models.ServerProfile{}
	err = db.Read(SERVER_PROFILE_STORE, n, rec)
	if err != nil {
		return nil, fmt.Errorf("error loading profile")
	}

	return rec, nil
}

func GetServerProfiles() []models.ServerProfile {
	db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	var profiles []models.ServerProfile

	ps, _ := db.ReadAll(SERVER_PROFILE_STORE)
	for _, p := range ps {
		var profile models.ServerProfile
		json.Unmarshal([]byte(p), &profile)

		profiles = append(profiles, profile)
	}

	return profiles
}

func DeleteServerProfile(n string) (err error) {
	db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	err = db.Delete(SERVER_PROFILE_STORE, n)
	if err != nil {
		return err
	}

	return nil
}

func SaveTempProfile(s models.TempProfile) {
	db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	db.Write(TEMP_PROFILE_STORE, s.Name, s)
}

func GetTempProfile(n string) models.TempProfile {
	db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	rec := models.TempProfile{}
	db.Read("TEMP_PROFILE_STORE", n, &rec)

	return rec
}
