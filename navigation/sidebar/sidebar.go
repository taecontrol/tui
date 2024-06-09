package sidebar

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/taecontrol/tui/navigation/item"
)

type Model struct {
	width, height int
	title         string
	items         []item.Model
}

func New(width, height int, title string, items []item.Model) Model {
	return Model{
		width:  width,
		height: height,
		title:  title,
		items:  items,
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
	var menuItems []string

	if len(m.title) > 0 {
		menuItems = append(
			menuItems,
			lipgloss.NewStyle().
				PaddingLeft(2).
				PaddingRight(2).
				BorderStyle(lipgloss.NormalBorder()).
				BorderBottom(true).
				Width(m.width).
				Render(m.title),
		)
	}

	for _, i := range m.items {
		renderedItem := lipgloss.NewStyle().
			PaddingTop(2).
			Render(i.View())

		menuItems = append(menuItems, renderedItem)
	}

	return lipgloss.NewStyle().
		Height(m.height - len(m.items)).
		Width(m.width).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#999")).
		Render(lipgloss.JoinVertical(lipgloss.Top, menuItems...))
}
