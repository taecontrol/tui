package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
	components "github.com/taecontrol/tui/layout"
	"github.com/taecontrol/tui/navigation/item"
	"github.com/taecontrol/tui/navigation/sidebar"
	"os"
)

type Model struct {
	Width, Height int
	Sidebar       sidebar.Model
}

func (Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var navCmd tea.Cmd
	m.Sidebar, navCmd = m.Sidebar.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, navCmd
}

func (m Model) View() string {
	return components.NewLayout(
		m.Height,
		m.Width,
		"Sidebar example",
		components.WithLeftSidebar(m.Sidebar.View()),
	).Render()
}

func main() {
	width, height, _ := term.GetSize(os.Stdout.Fd())

	p := tea.NewProgram(Model{
		Width:  width,
		Height: height,
		Sidebar: sidebar.New(
			width/6,
			height,
			"My App",
			[]item.Model{
				{Label: "Home", ShortcutKey: "h", IsActive: true},
				{Label: "Test", ShortcutKey: "t"},
				{Label: "About Us", ShortcutKey: "a"},
			},
		),
	})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
