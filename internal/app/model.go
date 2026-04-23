// Package app model provides necessary functionality for using bubble tea to render an interface
package app

import (
	tea "charm.land/bubbletea/v2"
)

func InitialModel() Model {
	return Model{
		Selected: -1,
		cursor:   0,
		width:    100,
		height:   20,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		wgShowCmd(),
		wgShowPoller(),
		wgDiscoverProfilesCmd(),
	)
}
