package app

import (
	"errors"
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
			return m, m.wgDownSelectedCmd()
		case "ctrl+d":
			return m, m.wgDownActiveCmd()
		case "u":
			return m, m.wgUpCmd()
		}

	case wgUpdateMsg:
		m.DetailContent = string(msg)
		return m, nil

	case profileStatesLoadedMsg:
		m.ProfileStates = msg
		return m, nil

	case wgTickMsg:
		return m, tea.Batch(
			wgShowCmd(),
			m.wgRefreshProfileStatesCmd(),
			wgShowPoller(),
		)

	case wgDownMsg:
		if msg.Err != nil {
			m.Status = msg.Err.Error()
		} else {
			m.Status = downStatus(msg.Result)
		}
		return m, tea.Batch(
			wgShowCmd(),
			m.wgRefreshProfileStatesCmd(),
		)

	case wgUpMsg:
		if msg.Err != nil {
			m.Status = msg.Err.Error()
		} else {
			m.Status = upStatus(msg.Result)
		}
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
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
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

func (m Model) wgDownSelectedCmd() tea.Cmd {
	return func() tea.Msg {
		if m.Selected < 0 || m.Selected >= len(m.ProfileStates) {
			return wgDownMsg{
				Err: errors.New("no profile selected"),
			}
		}
		return wgDownMsg{
			Result: wg.DownProfiles([]wg.ProfileState{m.ProfileStates[m.Selected]}),
		}
	}
}

func (m Model) wgDownActiveCmd() tea.Cmd {
	return func() tea.Msg {
		return wgDownMsg{
			Result: wg.DownActive(m.ProfileStates),
		}
	}
}

func (m Model) wgUpCmd() tea.Cmd {
	return func() tea.Msg {
		if m.Selected < 0 || m.Selected >= len(m.ProfileStates) {
			return wgUpMsg{
				Err: errors.New("no profile selected"),
			}
		}
		return wgUpMsg{
			Result: wg.UpProfile(m.ProfileStates[m.Selected]),
		}
	}
}

func upStatus(result wg.UpResult) string {
	switch {
	case len(result.Attempted) == 0:
		return "No profile to connect"
	case len(result.Failed) == 0 && len(result.Up) == 1:
		return fmt.Sprintf("Connected %s", result.Up[0])
	case len(result.Up) == 0 && len(result.Failed) == 1:
		name, err := singleFailure(result.Failed)
		return fmt.Sprintf("Failed to connect to %s: %v", name, err)
	case len(result.Failed) == 0:
		return fmt.Sprintf("Connected %d tunnels", len(result.Up))
	case len(result.Up) == 0:
		return fmt.Sprintf("Failed to connect %d tunnel(s)", len(result.Failed))
	default:
		return fmt.Sprintf("Connected %d tunnel(s), failed %d", len(result.Up), len(result.Failed))
	}
}

func downStatus(result wg.DownResult) string {
	switch {
	case len(result.Attempted) == 0:
		return "No active tunnels"
	case len(result.Failed) == 0 && len(result.Down) == 1:
		return fmt.Sprintf("Disconnected %s", result.Down[0])
	case len(result.Down) == 0 && len(result.Failed) == 1:
		name, err := singleFailure(result.Failed)
		return fmt.Sprintf("Failed to disconnect from %s: %v", name, err)
	case len(result.Failed) == 0:
		return fmt.Sprintf("Disconnected %d tunnels", len(result.Down))
	case len(result.Down) == 0:
		return fmt.Sprintf("Failed to disconnect %d tunnel(s)", len(result.Failed))
	default:
		return fmt.Sprintf("Disconnected %d tunnel(s), failed %d", len(result.Down), len(result.Failed))
	}
}

func singleFailure(failed map[string]error) (string, error) {
	for name, err := range failed {
		return name, err
	}
	return "", nil
}
