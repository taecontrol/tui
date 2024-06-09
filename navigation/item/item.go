package item

import (
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type UpdatedMsg struct {
	Item Model
}

type Model struct {
	Label       string
	ShortcutKey string
	IsActive    bool
}

func (i Model) Init() tea.Cmd {
	return nil
}

func (i Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == i.ShortcutKey {
			i.IsActive = true
			return i, i.ItemUpdated
		}
	case UpdatedMsg:
		if i.ShortcutKey != msg.Item.ShortcutKey {
			i.IsActive = false
		}
		return i, nil
	}

	return i, nil
}

func (i Model) View() string {
	style := lipgloss.NewStyle()
	if i.IsActive {
		style = style.Bold(true)
	} else {
		style = style.Foreground(lipgloss.Color("#999"))
	}

	labelString := style.Render(i.Label)
	shortcutKeyString := lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#999")).
		Render(fmt.Sprintf("[%s]", i.ShortcutKey))

	return lipgloss.NewStyle().
		PaddingLeft(2).
		PaddingRight(2).
		Render(fmt.Sprintf("%s %s", shortcutKeyString, labelString))
}

func (i Model) ItemUpdated() tea.Msg {
	return UpdatedMsg{
		Item: i,
	}
}
