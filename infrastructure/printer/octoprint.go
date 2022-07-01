package printer

import (
	"errors"
	"fmt"
	"log"
	"octo-command/models"
	"os"

	"github.com/chabber/go-octoprint"
)

type octoServerService struct {
	client *octoprint.Client
}

func NewOctoServerService(c *octoprint.Client) *octoServerService {
	return &octoServerService{
		client: c,
	}
}

func (os *octoServerService) PrintFile(f string) error {
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

func (os *octoServerService) GetToolTemp() ([]*models.Temperature, error) {
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

func (os *octoServerService) GetBedTemp() (*models.Temperature, error) {
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

func (osvc *octoServerService) UploadFile(src string, dst string) {
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

func (os *octoServerService) ToolState() {
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

func (os *octoServerService) GetFiles(dir string) []models.FileInformation {
	r := octoprint.FilesRequest{
		Location:  octoprint.Local,
		Recursive: true,
	}
	resp, err := r.Do(os.client)
	if err != nil {
		log.Printf("Error connecting to OctoPrint: %s", err)
	}

	results := getFileInfo(resp.Files)

	return results
}

// Extract and convert octoprint file structure to my own
func getFileInfo(fResp []*octoprint.FileInformation) []models.FileInformation {
	var files []models.FileInformation

	for _, f := range fResp {
		var file models.FileInformation
		file.Name = f.Name
		file.IsFolder = f.IsFolder()
		files = append(files, file)
		if f.IsFolder() {
			file.Children = getFileInfo(f.Children)
		}
	}

	return files
}

func (os *octoServerService) Connect(s models.PrinterProfile) (state *string, err error) {
	os.client = octoprint.NewClient(s.Url, s.ApiKey)

	r := octoprint.ConnectionRequest{}
	resp, err := r.Do(os.client)
	if err != nil {
		return nil, err
	}

	return (*string)(&resp.Current.State), nil
}

func (os *octoServerService) Home() {
	if os.client == nil {
		fmt.Println("Not connected to OctoPrint service")
		return
	}
	r := octoprint.PrintHeadHomeRequest{
		Axes: []octoprint.Axis{octoprint.XAxis, octoprint.YAxis},
	}
	r.Do(os.client)
}

func (os *octoServerService) SetBedTemp(t float64) {
	if os.client == nil {
		fmt.Println("Not connected to OctoPrint service")
		return
	}
	r := octoprint.BedTargetRequest{
		Target: float64(t),
	}
	r.Do(os.client)
}

func (os *octoServerService) SetToolTemp(t float64) {
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
