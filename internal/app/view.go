package app

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

var (
	leftPaneStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			Padding(0, 1)

	rightPaneStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			Padding(0, 1)
)

func (m Model) View() tea.View {
	leftWidth := m.width / 5 * 2
	rightWidth := m.width - leftWidth - 1

	left := leftPaneStyle.Width(leftWidth).Height(m.height - 2).Render(m.leftPaneView())
	right := rightPaneStyle.Width(rightWidth).Height(m.height - 2).Render(m.rightPaneView())

	return tea.NewView(lipgloss.JoinHorizontal(lipgloss.Top, left, right))
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

	return "wg show\n\n" + m.DetailContent
}
