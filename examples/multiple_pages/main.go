package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
	"github.com/taecontrol/tui/navigation/item"
	"github.com/taecontrol/tui/navigation/navbar"
	"os"
)

func main() {
	width, height, _ := term.GetSize(os.Stdout.Fd())

	p := tea.NewProgram(Home{
		Width:  width,
		Height: height,
		Navbar: navbar.New(
			"My App",
			[]item.Model{
				{Label: "Home", ShortcutKey: "h", IsActive: true},
				{Label: "About Us", ShortcutKey: "a"},
			},
		),
	})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
