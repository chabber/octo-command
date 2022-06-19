package services

import (
	"errors"
	"fmt"
	"log"
	"octo-command/octo/data"
	"octo-command/octo/models"
	"os"

	"github.com/chabber/go-octoprint"
)

type OctoService struct {
	client *octoprint.Client
}

func (os *OctoService) AddTempProfile(n string, bed float64, tool float64) error {
	tp := models.TempProfile{
		Name:     n,
		BedTemp:  bed,
		ToolTemp: tool,
	}

	data.SaveTempProfile(tp)

	return nil
}

func (os *OctoService) AddServerProfile(n string, u string, k string) {
	s := models.ServerProfile{
		Name:   n,
		Url:    u,
		ApiKey: k,
	}

	data.SaveServerProfile(s)
}

func (os *OctoService) GetServerProfile(n string) models.ServerProfile {
	return data.GetServerProfile(n)
}

func (os *OctoService) GetTempProfile(n string) models.TempProfile {
	return data.GetTempProfile(n)
}

func (os *OctoService) PrintFile(f string) error {
	if os.client == nil {
		return errors.New("error: not connected to OctoPrint server")
	}

	sel := octoprint.SelectFileRequest{
		Location: octoprint.Local,
		Path:     f,
		Print:    true,
	}
	err := sel.Do(os.client)
	if err != nil {
		return err
	}

	return nil
}

func (os *OctoService) GetToolTemp() ([]*models.Temperature, error) {
	if os.client == nil {
		return nil, errors.New("error: not connected to OctoPrint server")
	}

	r := octoprint.ToolStateRequest{}

	resp, err := r.Do(os.client)

	if err != nil {
		return nil, err
	}

	var temps []*models.Temperature

	for k, v := range resp.Current {
		t := models.Temperature{
			Actual: v.Actual,
			Target: v.Target,
			Offset: v.Offset,
			Label:  k,
		}

		temps = append(temps, &t)
	}

	return temps, nil
}

func (os *OctoService) GetBedTemp() (*models.Temperature, error) {
	if os.client == nil {
		return nil, errors.New("error: not connected to OctoPrint server")
	}

	r := octoprint.BedStateRequest{}

	resp, err := r.Do(os.client)

	if err != nil {
		return nil, err
	}

	t := new(models.Temperature)

	if d, found := resp.Current["bed"]; found {
		t.Actual = d.Actual
		t.Target = d.Target
		t.Offset = d.Offset
	} else {
		return nil, errors.New("error: bed temperature not available")
	}

	return t, nil
}

func (osvc *OctoService) UploadFile(src string, dst string) {
	if osvc.client == nil {
		fmt.Println("Not connected to OctoPrint server")
		return
	}

	r := octoprint.UploadFileRequest{
		Select:   false,
		Location: octoprint.Local,
		Print:    false,
	}

	reader, err := os.Open(src)
	if err != nil {
		log.Printf("Error opening file for upload: %s", err)
		return
	}

	err = r.AddFile(dst, reader)
	if err != nil {
		log.Printf("Error adding file for upload: %s", err)
		return
	}

	_, err = r.Do(osvc.client)
	if err != nil {
		log.Printf("Error uploading file to server: %s", err)
		return
	}
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

	_, err := r.Do(os.client)

	if err != nil {
		log.Printf("Error getting tool state: %s", err)
	}
}
func (os *OctoService) GetFiles(dir string) {
	if os.client == nil {
		fmt.Println("Not connected to OctoPrint service")
		return
	}

	r := octoprint.FilesRequest{
		Location:  octoprint.Local,
		Recursive: true,
	}
	resp, err := r.Do(os.client)
	if err != nil {
		log.Printf("Error connecting to OctoPrint: %s", err)
	}

	printFileList(resp.Files, 0)
}

func printFileList(f []*octoprint.FileInformation, level int) {
	var indent string = ""
	for x := 0; x < level; x++ {
		indent = indent + "   "
	}

	for _, f := range f {
		if f.IsFolder() {
			fmt.Printf("%s[%s]\n", indent, f.Name)
			printFileList(f.Children, level+1)
		} else {
			fmt.Printf("%s%s\n", indent, f.Name)
		}
	}
}

func (os *OctoService) Connect(n string) {
	s := data.GetServerProfile(n)
	os.client = octoprint.NewClient(s.Url, s.ApiKey)

	r := octoprint.ConnectionRequest{}
	resp, err := r.Do(os.client)
	if err != nil {
		log.Printf("Error connecting to OctoPrint: %s", err)
	}

	fmt.Printf("Connection State: %q\n", resp.Current.State)
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

func (os *OctoService) SetBedTemp(t float64) {
	if os.client == nil {
		fmt.Println("Not connected to OctoPrint service")
		return
	}
	r := octoprint.BedTargetRequest{
		Target: float64(t),
	}
	r.Do(os.client)
}

func (os *OctoService) SetToolTemp(t float64) {
	if os.client == nil {
		fmt.Println("Not connected to OctoPrint service")
		return
	}

	targets := make(map[string]float64)
	targets["tool0"] = t

	r := octoprint.ToolTargetRequest{
		Targets: targets,
	}

	r.Do(os.client)
}
