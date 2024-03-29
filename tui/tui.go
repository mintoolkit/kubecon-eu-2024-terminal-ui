package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/mintoolkit/mint/pkg/report"
)

type state int

const (
	defaultState state = iota
	searchState
	MouseScrollSpeed = 3
	ListProportion   = 0.3
)

type model struct {
	width, height int

	list     list.Model
	layers   list.Model
	viewport viewport.Model

	statusMessage string
	ready         bool
	now           string

	limit int64 // scan size

	keyMap
	state
}

func New(data report.XrayCommand) (*model, error) {
	var items []list.Item
	for _, imageInfo := range data.ImageStack {
		image := item{
			title: imageInfo.FullName,
			desc:  imageInfo.NewSizeHuman,
		}
		items = append(items, image)
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "XRay Viewer"
	l.SetShowHelp(false)
	l.SetShowStatusBar(false)
	l.SetShowFilter(false)
	l.SetFilteringEnabled(false)

	var layersItems []list.Item
	for _, layerInfo := range data.ImageLayers {
		layer := item{
			title: fmt.Sprintf("%d", layerInfo.Index),
			desc:  layerInfo.ID,
		}
		layersItems = append(layersItems, layer)
	}

	layers := list.New(layersItems, list.NewDefaultDelegate(), 0, 0)

	return &model{
		list:   l,
		layers: layers,
		limit:  30,

		keyMap: defaultKeyMap(),
		state:  defaultState,
	}, nil
}
