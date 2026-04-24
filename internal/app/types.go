package app

import "github.com/StephenCotterrell/twig/internal/wg"

type Model struct {
	Config        wg.Config
	Profiles      []wg.Profile
	Selected      int
	cursor        int
	width         int
	height        int
	DetailContent string
}

type (
	wgUpdateMsg       string
	profilesLoadedMsg []wg.Profile
	wgTickMsg         struct{}
)
