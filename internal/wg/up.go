package wg

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func Up(profileState ProfileState) error {
	if profileState.IsActive {
		return errors.New("can only call up on an inactive profile")
	}

	cmd := exec.Command("wg-quick", "up", profileState.Profile.Name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		if msg := strings.TrimSpace(string(out)); msg != "" {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return err
	}

	return nil
}

func UpProfile(state ProfileState) UpResult {
	result := UpResult{
		Failed: make(map[string]error),
	}

	name := state.Profile.Name
	result.Attempted = append(result.Attempted, name)

	if err := Up(state); err != nil {
		result.Failed[name] = err
	} else {
		result.Up = append(result.Up, name)
	}

	return result
}
