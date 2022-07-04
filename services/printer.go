package services

import (
	"octo-command/infrastructure/printer"
	"octo-command/models"
)

type PrinterService struct {
	Pdp printer.PrinterDataPort
}

func NewPrinterService(data printer.PrinterDataPort) *PrinterService {
	return &PrinterService{
		Pdp: data,
	}
}

func (ps *PrinterService) PrintFile(f string) error {
	return ps.Pdp.PrintFile(f)
}

func (ps *PrinterService) GetToolTemp() ([]*models.Temperature, error) {
	return ps.Pdp.GetToolTemp()
}

func (ps *PrinterService) GetBedTemp() (*models.Temperature, error) {
	return ps.Pdp.GetBedTemp()
}

func (ps *PrinterService) UploadFile(src string, dst string) {
	ps.Pdp.UploadFile(src, dst)
}

func (ps *PrinterService) ToolState() {
	panic("not implemented") // TODO: Implement
}

func (ps *PrinterService) GetFiles(dir string) []models.FileInformation {
	panic("not implemented") // TODO: Implement
}

func (ps *PrinterService) Connect(p *models.PrinterProfile) (state *string, err error) {
	return ps.Pdp.Connect(*p)
}

func (ps *PrinterService) Home() {
	ps.Pdp.Home()
}

func (ps *PrinterService) SetBedTemp(temp float64) {
	// TODO: check max temp
	ps.Pdp.SetBedTemp(temp)
}

func (ps *PrinterService) SetToolTemp(temp float64) {
	// TODO: check max temp
	ps.Pdp.SetToolTemp(temp)
}
