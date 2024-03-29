package tui

import (
	"log"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cast"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	// global msg handling
	switch msg := msg.(type) {
	case errMsg:
		m.statusMessage = msg.err.Error()
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		log.Printf("Width: %d Height: %d", m.width, m.height)
		statusBarHeight := lipgloss.Height(m.statusView())
		height := m.height - statusBarHeight
		listViewWidth := cast.ToInt(ListProportion * float64(m.width))
		listWidth := listViewWidth - listViewStyle.GetHorizontalFrameSize()
		log.Printf("list width height: %d %d", listWidth, height)
		m.list.SetSize(listWidth, height)

		detailViewWidth := m.width - listWidth
		log.Printf("viewport: %d %d", detailViewWidth, height)
		m.layers.SetSize(detailViewWidth, height)
		m.viewport = viewport.New(detailViewWidth, height)
		m.viewport.MouseWheelEnabled = true
		m.viewport.SetContent(m.viewportContent(m.viewport.Width))
	case countMsg:
		m.ready = true
	}

	switch m.state {
	case defaultState:
		cmds = append(cmds, m.handleDefaultState(msg))
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *model) handleDefaultState(msg tea.Msg) tea.Cmd {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEscape:
			cmd = tea.Quit
			cmds = append(cmds, cmd)
		case tea.KeyCtrlC:
			cmd = tea.Quit
			cmds = append(cmds, cmd)
		case tea.KeyUp, tea.KeyDown, tea.KeyLeft, tea.KeyRight:
			m.list, cmd = m.list.Update(msg)
			cmds = append(cmds, cmd)
			m.layers, cmd = m.layers.Update(msg)
			cmds = append(cmds, cmd)
			// m.viewport.GotoTop()
			// m.viewport.SetContent(m.viewportContent(m.viewport.Width))
		}
	default:
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)

		m.layers, cmd = m.layers.Update(msg)
		cmds = append(cmds, cmd)

		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return tea.Batch(cmds...)
}
