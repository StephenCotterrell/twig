// Package app model provides necessary functionality for using bubble tea to render an interface
package app

import (
	"fmt"
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/StephenCotterrell/twig/cmd/internal/wg"
)

type Model struct {
	Profiles []wg.Profile
	Selected map[int]struct{}
	cursor   int
}

func InitialModel() Model {
	cfg := wg.DefaultConfig()
	profiles, err := wg.DiscoverProfiles(cfg.WireGuardDir)
	if err != nil {
		log.Fatal("failed to discover profiles")
	}
	return Model{
		Profiles: profiles,
		Selected: make(map[int]struct{}),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() tea.View {
	s := "WireGuard Options: \n"

	cfg := wg.DefaultConfig()
	profiles, err := wg.DiscoverProfiles(cfg.WireGuardDir)
	if err != nil {
		log.Fatal("failed to discover profiles")
	}

	for i, profile := range profiles {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.Selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, profile)
	}

	return tea.NewView(s)
}
