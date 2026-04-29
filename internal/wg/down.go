package wg

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func Down(profileState ProfileState) error {
	if !(profileState.IsActive) {
		return errors.New("can only call down on an active profile")
	}

	cmd := exec.Command("wg-quick", "down", profileState.Profile.Name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %w", strings.TrimSpace(string(out)), err)
	}

	return nil
}

func DownProfiles(profileStates []ProfileState) DownResult {
	result := DownResult{
		Failed: make(map[string]error),
	}
	for _, state := range profileStates {
		name := state.Profile.Name
		result.Attempted = append(result.Attempted, name)

		if err := Down(state); err != nil {
			result.Failed[name] = err
			continue
		}

		result.Down = append(result.Down, name)

	}
	return result
}

func DownActive(profileStates []ProfileState) DownResult {
	active := []ProfileState{}

	for _, state := range profileStates {
		if state.IsActive {
			active = append(active, state)
		}
	}
	return DownProfiles(active)
}
