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

	baseItemStyle = lipgloss.NewStyle().PaddingLeft(1)

	// activeItemStyle = baseItemStyle.Foreground(lipgloss.Color("10")).Bold(true)

	// selectedItemStyle = baseItemStyle.Foreground(lipgloss.Color("12"))

	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true)
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

	if len(m.ProfileStates) == 0 {
		s.WriteString("No profiles found\n")
		return s.String()
	}

	fmt.Fprintf(&s, "WireGuard Options: \n\n")

	for i, p := range m.ProfileStates {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}

		checked := " "
		if i == m.Selected {
			checked = "X"
		}

		cursorPart := cursorStyle.Render(cursor)

		rowStyle := baseItemStyle

		if i == m.Selected {
			rowStyle = rowStyle.Foreground(lipgloss.Color("12"))
		}

		if p.IsActive {
			rowStyle = rowStyle.Bold(true).Foreground(lipgloss.Color("10"))
		}

		rowPart := rowStyle.Render(fmt.Sprintf(" [%s] %s", checked, p.Profile.Name))

		fmt.Fprintln(&s, cursorPart+rowPart)

	}

	fmt.Fprintf(&s, "\nPress q to quit.")
	return s.String()
}

func (m Model) rightPaneView() string {
	var s strings.Builder
	if len(m.ProfileStates) == 0 {
		s.WriteString("Details\n\nNo profile selected")
		return s.String()
	}

	return "wg show\n\n" + m.DetailContent
}
