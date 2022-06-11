package data

import (
	"fmt"
	"octocommand/octo/models"

	scribble "github.com/nanobox-io/golang-scribble"
)

func SaveServer(s models.Server) {
	db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	db.Write("servers", s.Name, s)
}

func GetServer(n string) models.Server {
	db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	s := models.Server{}
	db.Read("servers", n, &s)

	return s
}
