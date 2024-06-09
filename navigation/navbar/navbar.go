package navbar

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/taecontrol/tui/navigation/item"
)

type Model struct {
	title string
	items []item.Model
}

func New(title string, items []item.Model) Model {
	return Model{
		title: title,
		items: items,
	}
}

func (Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	updatedItems := make([]item.Model, len(m.items))
	var itemCmd tea.Cmd

	for idx, i := range m.items {
		updatedItem, cmd := i.Update(msg)
		if cmd != nil {
			itemCmd = cmd
		}
		updatedItems[idx] = updatedItem
	}

	m.items = updatedItems

	if itemCmd != nil {
		return m, itemCmd
	}

	return m, nil
}

func (m Model) View() string {
	var menuString string

	if len(m.title) > 0 {
		menuString = lipgloss.NewStyle().
			PaddingLeft(2).
			PaddingRight(2).
			BorderStyle(lipgloss.NormalBorder()).
			BorderRight(true).
			Render(m.title)
	}

	for _, i := range m.items {
		menuString += i.View()
	}

	return lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#999")).
		Render(menuString)
}
