package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
	components "github.com/taecontrol/tui/layout"
	"github.com/taecontrol/tui/navigation/item"
	"github.com/taecontrol/tui/navigation/navbar"
	"os"
)

type Model struct {
	Width, Height int
	Navbar        navbar.Model
}

func (Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var navCmd tea.Cmd
	m.Navbar, navCmd = m.Navbar.Update(msg)

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
		"Navbar example",
		components.WithHeader(m.Navbar.View()),
	).Render()
}

func main() {
	width, height, _ := term.GetSize(os.Stdout.Fd())

	p := tea.NewProgram(Model{
		Width:  width,
		Height: height,
		Navbar: navbar.New(
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
