package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg struct {
	err error
}

type scanMsg struct {
	items []list.Item
}

func (m model) scanCmd() tea.Cmd {
	return func() tea.Msg {
		return scanMsg{items: []list.Item{}}
	}

}

type countMsg struct {
	count int
}

func (m model) countCmd() tea.Cmd {
	return func() tea.Msg {
		return countMsg{count: 1}
	}
}

type tickMsg struct {
	t string
}

func (m model) tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(_ time.Time) tea.Msg {
		return tickMsg{t: time.Now().Format("2006-01-02 15:04:05")}
	})
}
