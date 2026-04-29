package wg

import (
	"errors"
	"os/exec"
)

func Up(profileState ProfileState) error {
	if profileState.IsActive {
		return errors.New("can only call up on an inactive profile")
	}

	cmd := exec.Command("wg-quick", "up", profileState.Profile.Name)
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}
