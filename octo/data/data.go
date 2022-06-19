package data

import (
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

func GetServerProfile(n string) models.ServerProfile {
	db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	rec := models.ServerProfile{}
	db.Read(SERVER_PROFILE_STORE, n, &rec)

	return rec
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
