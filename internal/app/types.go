package app

import "github.com/StephenCotterrell/twig/internal/wg"

type Model struct {
	Status        string
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
)

type wgDownMsg struct {
	Result wg.DownResult
	Err    error
}

type wgUpMsg struct {
	Result wg.UpResult
	Err    error
}
