package services

import (
	"octo-command/infrastructure/printer"
	"octo-command/models"
)

type PrinterService struct {
	pdp printer.PrinterDataPort
}

func NewPrinterService(data printer.PrinterDataPort) *PrinterService {
	return &PrinterService{
		pdp: data,
	}
}

func (ps *PrinterService) PrintFile(f string) error {
	return ps.pdp.PrintFile(f)
}

func (ps *PrinterService) GetToolTemp() ([]*models.Temperature, error) {
	return ps.pdp.GetToolTemp()
}

func (ps *PrinterService) GetBedTemp() (*models.Temperature, error) {
	return ps.pdp.GetBedTemp()
}

func (ps *PrinterService) UploadFile(src string, dst string) {
	ps.pdp.UploadFile(src, dst)
}

func (ps *PrinterService) ToolState() {
	panic("not implemented") // TODO: Implement
}

func (ps *PrinterService) GetFiles(dir string) []models.FileInformation {
	panic("not implemented") // TODO: Implement
}

func (ps *PrinterService) Connect(p *models.PrinterProfile) (state *string, err error) {
	return ps.pdp.Connect(*p)
}

func (ps *PrinterService) Home() {
	ps.pdp.Home()
}

func (ps *PrinterService) SetBedTemp(temp float64) {
	// TODO: check max temp
	ps.pdp.SetBedTemp(temp)
}

func (ps *PrinterService) SetToolTemp(temp float64) {
	// TODO: check max temp
	ps.pdp.SetToolTemp(temp)
}
