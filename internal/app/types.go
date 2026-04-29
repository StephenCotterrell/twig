package app

import "github.com/StephenCotterrell/twig/internal/wg"

type Model struct {
	Config        wg.Config
	ProfileStates []wg.ProfileState
	Selected      int
	cursor        int
	width         int
	height        int
	DetailContent string
}

type (
	wgUpdateMsg            string
	profileStatesLoadedMsg []wg.ProfileState
	wgTickMsg              struct{}
	wgDownMsg              struct {
		Attempted []string
		Down      []string
		Failed    map[string]error
	}
	wgUpMsg string
)
