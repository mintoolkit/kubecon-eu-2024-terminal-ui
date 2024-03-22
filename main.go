package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mintoolkit/mint/pkg/report"
)

type slimreport struct {
	XrayCommand report.Command
}

func NewReport(rawData []byte) slimreport {
	var model slimreport
	err := json.Unmarshal(rawData, &model)
	if err != nil {
		fmt.Errorf("problem with unmarshal slimreport")
	}
	return model
}

func readData(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := io.ReadAll(file)
	return b
}

func main() {
	rawData := readData("data/xray.slim.report.json")
	data := NewReport(rawData)

	m := model{list: list.New(data.XrayCommand.ImageStack, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "X-ray Slim"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
