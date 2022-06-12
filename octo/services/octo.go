package services

import (
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

func (os *OctoService) AddServer(n string, u string, k string) {
	s := models.Server{
		Name:   n,
		Url:    u,
		ApiKey: k,
	}

	data.SaveServer(s)
}

func (osvc *OctoService) UploadFile(src string, dst string) {
	if osvc.client == nil {
		fmt.Println("Not connected to OctoPrint service")
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
func (os *OctoService) ListFiles(dir string) {
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
	s := data.GetServer(n)
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
