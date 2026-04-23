package app

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/StephenCotterrell/twig/internal/wg"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.Profiles)-1 {
				m.cursor++
			}
		case "enter", "space":
			m.Selected = m.cursor
		case "esc":
			m.Selected = -1
		}

	case wgUpdateMsg:
		m.DetailContent = string(msg)
		return m, wgShowPoller()

	case profilesLoadedMsg:
		m.Profiles = msg
		return m, wgDiscoverProfilesCmd()

	case wgTickMsg:
		return m, tea.Batch(
			wgShowCmd(),
			wgShowPoller(),
		)
	}

	return m, nil
}

func wgShowCmd() tea.Cmd {
	return func() tea.Msg {
		out, err := wg.Show()
		if err != nil {
			return wgUpdateMsg("error")
		}
		return wgUpdateMsg(out)
	}
}

func wgShowPoller() tea.Cmd {
	return tea.Tick(3*time.Second, func(time.Time) tea.Msg {
		return wgTickMsg{}
	})
}

func wgDiscoverProfilesCmd() tea.Cmd {
	return func() tea.Msg {
		cfg := wg.DefaultConfig()
		profiles, err := wg.DiscoverProfiles(cfg.WireGuardDir)
		if err != nil {
			return nil
		}
		return profilesLoadedMsg(profiles)
	}
}
