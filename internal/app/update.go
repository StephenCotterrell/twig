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
			if m.cursor < len(m.ProfileStates)-1 {
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

	case profileStatesLoadedMsg:
		m.ProfileStates = msg
		return m, m.wgRefreshProfileStatesCmd()

	case wgTickMsg:
		return m, tea.Batch(
			wgShowCmd(),
			wgShowPoller(),
			m.wgRefreshProfileStatesCmd(),
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
	return tea.Tick(5*time.Second, func(time.Time) tea.Msg {
		return wgTickMsg{}
	})
}

func (m Model) wgRefreshProfileStatesCmd() tea.Cmd {
	return func() tea.Msg {
		client := wg.Client{Config: m.Config}
		profileStates, err := client.RefreshProfileStates()
		if err != nil {
			return nil
		}

		return profileStatesLoadedMsg(profileStates)
	}
}
