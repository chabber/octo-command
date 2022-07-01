package datastore

import (
	"encoding/json"

	scribble "github.com/nanobox-io/golang-scribble"
)

type scribbleDatabaseService struct {
	db *scribble.Driver
}

func NewScribbleDatabaseService(drv *scribble.Driver) *scribbleDatabaseService {
	return &scribbleDatabaseService{
		db: drv,
	}
}

func (s *scribbleDatabaseService) Get(resource string, collection string, obj interface{}) error {
	err := s.db.Read(collection, resource, obj)

	return err
}

func (s *scribbleDatabaseService) GetAll(c string, r []interface{}) error {
	records, err := s.db.ReadAll(c)

	if err != nil {
		return err
	}

	for _, record := range records {
		var result interface{}
		json.Unmarshal([]byte(record), &result)

		r = append(r, result)
	}

	return nil
}

func (s *scribbleDatabaseService) Save(resource string, collection string, obj interface{}) error {
	err := s.db.Write(collection, resource, obj)

	return err
}

func (s *scribbleDatabaseService) Delete(collection string, resource string) error {
	err := s.db.Delete(collection, resource)

	return err
}
