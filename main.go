package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mintoolkit/mint/pkg/report"
)

type slimreport struct {
	ImageStacks []list.Item
}

func NewReport(rawData []byte) slimreport {
	var report report.XrayCommand
	var model slimreport
	err := json.Unmarshal(rawData, &report)
	if err != nil {
		fmt.Errorf("problem with unmarshal XrayCommand: %s", err)
	}
	for _, imageInfo := range report.ImageStack {
		image := item{
			title: imageInfo.FullName,
			desc:  imageInfo.NewSizeHuman,
		}
		model.ImageStacks = append(model.ImageStacks, image)
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
	if err != nil {
		log.Fatal(err)
	}
	return b
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {
	rawData := readData("data/xray.slim.report.json")
	data := NewReport(rawData)

	m := model{list: list.New(data.ImageStacks, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "X-ray Slim"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
