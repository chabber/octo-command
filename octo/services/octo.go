package services

import (
	"fmt"
	"log"
	"octocommand/octo/data"
	"octocommand/octo/models"

	"github.com/ugurgudelek/go-octoprint"
)

type OctoService struct {
	client *octoprint.Client
}

func (os *OctoService) AddServer(n string, u string, k string) {
	s := models.Server{
		Name:   n,
		Url:    u,
		ApiKey: k,
	}

	data.SaveServer(s)
}

func (os *OctoService) ToolState() {
	if os.client == nil {
		fmt.Println("Not connected to OctoPrint service")
		return
	}
	r := octoprint.ToolStateRequest{
		History: false,
		Limit:   100,
	}

	ts, err := r.Do(os.client)

	if err != nil {
		log.Printf("Error getting tool state: %s", err)
	}

	fmt.Print(string(ts.UnmarshalJSON())

}

func (os *OctoService) Connect(u string, k string, n string) {
	if n != "" {
		s := data.GetServer(n)
		os.client = octoprint.NewClient(s.Url, s.ApiKey)
	} else {
		os.client = octoprint.NewClient(u, k)
	}

	r := octoprint.ConnectionRequest{}
	s, err := r.Do(os.client)
	if err != nil {
		log.Printf("Error connecting to OctoPrint: %s", err)
	}

	fmt.Printf("Connection State: %q\n", s.Current.State)
}

func (os *OctoService) Home() {
	if os.client == nil {
		fmt.Println("Not connected to OctoPrint service")
		return
	}
	r := octoprint.PrintHeadHomeRequest{
		Axes: []octoprint.Axis{octoprint.XAxis, octoprint.YAxis},
	}
	r.Do(os.client)
}

func (os *OctoService) BedTemp(t float64) {
	if os.client == nil {
		fmt.Println("Not connected to OctoPrint service")
		return
	}
	r := octoprint.BedTargetRequest{
		Target: t,
	}
	r.Do(os.client)
}
