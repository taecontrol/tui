package main

import (
	tea "github.com/charmbracelet/bubbletea"
	components "github.com/taecontrol/tui/layout"
	"github.com/taecontrol/tui/navigation/item"
	"github.com/taecontrol/tui/navigation/navbar"
)

type About struct {
	Width  int
	Height int
	Navbar navbar.Model
}

func (About) Init() tea.Cmd {
	return nil
}

func (m About) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var navCmd tea.Cmd
	m.Navbar, navCmd = m.Navbar.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case item.UpdatedMsg:
		if msg.Item.ShortcutKey == "h" {
			home := Home{
				Width:  m.Width,
				Height: m.Height,
				Navbar: m.Navbar,
			}

			return home, navCmd
		}
	}

	return m, navCmd
}

func (m About) View() string {
	return components.NewLayout(
		m.Height,
		m.Width,
		"About",
		components.WithHeader(m.Navbar.View()),
	).Render()
}
