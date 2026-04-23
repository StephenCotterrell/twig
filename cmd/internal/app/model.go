// Package app model provides necessary functionality for using bubble tea to render an interface
package app

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	"github.com/StephenCotterrell/twig/cmd/internal/wg"
)

var (
	leftPaneStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			Padding(0, 1)

	rightPaneStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			Padding(0, 1)
)

type Model struct {
	Profiles      []wg.Profile
	Selected      int
	cursor        int
	width         int
	height        int
	detailContent string
}

func InitialModel(profiles []wg.Profile) Model {
	return Model{
		Profiles: profiles,
		Selected: -1,
		cursor:   0,
		width:    80,
		height:   20,
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
	}

	return m, nil
}

func (m Model) leftPaneView() string {
	var s strings.Builder
	s.WriteString("Profiles\n\n")

	if len(m.Profiles) == 0 {
		s.WriteString("No profiles found\n")
		return s.String()
	}

	fmt.Fprintf(&s, "WireGuard Options: \n")

	for i, p := range m.Profiles {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}

		checked := " "
		if i == m.Selected {
			checked = "x"
		}

		fmt.Fprintf(&s, "%s [%s] %s\n", cursor, checked, p.Name)
	}

	fmt.Fprintf(&s, "\nPress q to quit.")
	return s.String()
}

func (m Model) rightPaneView() string {
	var s strings.Builder
	if len(m.Profiles) == 0 {
		s.WriteString("Details\n\nNo profile selected")
		return s.String()
	}
	return "wg show\n\n" + m.detailContent
}

func (m Model) View() tea.View {
	leftWidth := m.width / 5 * 2
	rightWidth := m.width - leftWidth - 1

	left := leftPaneStyle.Width(leftWidth).Height(m.height - 2).Render(m.leftPaneView())
	right := rightPaneStyle.Width(rightWidth).Height(m.height - 2).Render(m.rightPaneView())

	return tea.NewView(lipgloss.JoinHorizontal(lipgloss.Top, left, right))
}
