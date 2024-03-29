package tui

type item struct {
	title string
	desc  string
}

func (i item) Title() string { return i.title }

func (i item) Description() string {
	return i.desc
}

func (i item) FilterValue() string { return i.title }
