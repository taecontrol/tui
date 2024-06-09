package navbar

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	title string
	items []Item
}

func New(title string, items []Item) Model {
	return Model{
		title: title,
		items: items,
	}
}

func (Model) Init() tea.Cmd {
	return nil
}

func (n Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	updatedItems := make([]Item, len(n.items))
	var itemCmd tea.Cmd

	for i, item := range n.items {
		updatedItem, cmd := item.Update(msg)
		if cmd != nil {
			itemCmd = cmd
		}
		updatedItems[i] = updatedItem
	}

	n.items = updatedItems

	if itemCmd != nil {
		return n, itemCmd
	}

	return n, nil
}

func (n Model) View() string {
	var menuString string

	if len(n.title) > 0 {
		menuString = lipgloss.NewStyle().
			PaddingLeft(2).
			PaddingRight(2).
			BorderStyle(lipgloss.NormalBorder()).
			BorderRight(true).
			Render(n.title)
	}

	for _, item := range n.items {
		if item.IsLast {
			menuString += lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				BorderRight(true).
				Render(item.View())
		} else {
			menuString += item.View()
		}
	}

	return lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#999")).
		Render(menuString)
}
