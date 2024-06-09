package navbar

import (
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ItemUpdatedMsg struct {
	Item Item
}

type Item struct {
	Label       string
	ShortcutKey string
	IsActive    bool
	IsLast      bool
}

func (i Item) Init() tea.Cmd {
	return nil
}

func (i Item) Update(msg tea.Msg) (Item, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == i.ShortcutKey {
			i.IsActive = true
			return i, i.ItemUpdated
		}
	case ItemUpdatedMsg:
		if i.ShortcutKey != msg.Item.ShortcutKey {
			i.IsActive = false
		}
		return i, nil
	}

	return i, nil
}

func (i Item) View() string {
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

func (i Item) ItemUpdated() tea.Msg {
	return ItemUpdatedMsg{
		Item: i,
	}
}
