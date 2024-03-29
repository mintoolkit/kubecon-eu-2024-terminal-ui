package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mintoolkit/kubecon-eu-2024-terminal-ui/tui"
	"github.com/mintoolkit/mint/pkg/report"
)

func NewReport(rawData []byte) report.XrayCommand {
	var report report.XrayCommand
	err := json.Unmarshal(rawData, &report)
	if err != nil {
		fmt.Errorf("problem with unmarshal XrayCommand: %v", err)
	}
	return report
}

func readData(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func main() {
	rawData := readData("data/xray.slim.report.json")
	data := NewReport(rawData)

	model, err := tui.New(data)
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion())
	if err := p.Start(); err != nil {
		log.Fatal("start failed: ", err)
	}
}
