package data

import (
	"encoding/json"
	"fmt"
	"octo-command/domain"
	"octo-command/models"

	scribble "github.com/nanobox-io/golang-scribble"
)

type scribbleDataService struct {
	data *scribble.Driver
}

func NewScribbleDataService() StorageDataPort {
	// storage service
	db, err := scribble.New(".", nil)

	if err != nil {
		fmt.Println("error getting scribble db driver: ", err)
	}

	// TODO: error handling
	return &scribbleDataService{
		data: db,
	}
}

func (sds *scribbleDataService) GetConfig() (*models.Config, error) {
	var config = models.Config{}
	sds.data.Read(domain.CONFIG_COLLECTION, domain.CONFIG_RESOURCE, config)

	return &config, nil
}

func (sds *scribbleDataService) SavePrinterProfile(s models.PrinterProfile) error {
	return sds.data.Write(domain.SERVER_PROFILE_COLLECTION, s.Name, s)
}

func (sds *scribbleDataService) GetPrinterProfile(n string) (*models.PrinterProfile, error) {
	p := models.PrinterProfile{}

	err := sds.data.Read(domain.SERVER_PROFILE_COLLECTION, n, &p)
	if err != nil {
		fmt.Println("error reading printer profile from db: ", err)
		return nil, err
	}
	fmt.Println("getting profile for: ", n)
	fmt.Println("profile: ", p.Url)
	return &p, err
}

func (sds *scribbleDataService) GetPrinterProfiles() ([]models.PrinterProfile, error) {
	records, err := sds.data.ReadAll(domain.SERVER_PROFILE_COLLECTION)

	if err != nil {
		return nil, err
	}

	var profiles []models.PrinterProfile

	for _, record := range records {
		var profile models.PrinterProfile
		json.Unmarshal([]byte(record), &profile)

		profiles = append(profiles, profile)
	}

	return profiles, nil

}

func (sds *scribbleDataService) DeletePrinterProfile(n string) error {
	return sds.data.Delete(domain.SERVER_PROFILE_COLLECTION, n)
}

func (sds *scribbleDataService) SaveTempProfile(s models.TempProfile) error {
	return sds.data.Write(domain.SERVER_PROFILE_COLLECTION, s.Name, s)
}

func (sds *scribbleDataService) GetTempProfile(n string) (models.TempProfile, error) {
	rec := models.TempProfile{}
	sds.data.Read(domain.TEMP_PROFILE_COLLECTION, n, &rec)

	return rec, nil
}

func (sds *scribbleDataService) DeleteTempProfile(n string) error {
	return sds.data.Delete(domain.TEMP_PROFILE_COLLECTION, n)
}

func (sds *scribbleDataService) GetTempProfiles() ([]models.TempProfile, error) {
	records, err := sds.data.ReadAll(domain.TEMP_PROFILE_COLLECTION)

	if err != nil {
		return nil, err
	}

	var profiles []models.TempProfile

	for _, record := range records {
		var profile models.TempProfile
		json.Unmarshal([]byte(record), &profile)

		profiles = append(profiles, profile)
	}

	return profiles, nil
}
