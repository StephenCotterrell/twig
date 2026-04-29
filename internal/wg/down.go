package wg

import (
	"errors"
	"os/exec"
)

func Down(profileState ProfileState) error {
	if !(profileState.IsActive) {
		return errors.New("can only call down on an active profile")
	}

	cmd := exec.Command("wg-quick", "down", profileState.Profile.Name)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}
