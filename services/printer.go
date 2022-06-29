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

func (ps *PrinterService) GetToolTemp() []*models.Temperature {
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

func (ps *PrinterService) Connect(_ *models.ServerProfile) (state string, err error) {
	panic("not implemented") // TODO: Implement
}

func (ps *PrinterService) Home() {
	panic("not implemented") // TODO: Implement
}

func (ps *PrinterService) SetBedTemp(temp float64) {
	panic("not implemented") // TODO: Implement
}

func (ps *PrinterService) SetToolTemp(temp float64) {
	panic("not implemented") // TODO: Implement
}
