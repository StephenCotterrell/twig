package app

import (
	"fmt"
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
		case "d":
			return m, m.wgDownCmd()
		case "u":
			return m, m.wgUpCmd()
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

	case wgDownMsg:
		return m, tea.Batch(
			wgShowCmd(),
			m.wgRefreshProfileStatesCmd(),
		)

	case wgUpMsg:
		return m, tea.Batch(
			wgShowCmd(),
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

func (m Model) wgDownCmd() tea.Cmd {
	return func() tea.Msg {
		msg := wgDownMsg{
			Failed: make(map[string]error),
		}

		for _, state := range m.ProfileStates {
			if !state.IsActive {
				continue
			}
			name := state.Profile.Name
			msg.Attempted = append(msg.Attempted, name)

			if err := wg.Down(state); err != nil {
				msg.Failed[name] = err
				continue
			}

			msg.Down = append(msg.Down, name)
		}

		return msg
	}
}

func (m Model) wgUpCmd() tea.Cmd {
	return func() tea.Msg {
		if m.Selected < 0 || m.Selected >= len(m.ProfileStates) {
			return wgUpMsg("No profile selected")
		}
		err := wg.Up(m.ProfileStates[m.Selected])
		if err != nil {
			return wgUpMsg(fmt.Sprintf("%v", err))
		}
		return wgUpMsg("complete")
	}
}
