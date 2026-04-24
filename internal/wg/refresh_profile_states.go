package wg

func (c Client) RefreshProfileStates() ([]ProfileState, error) {
	profiles, err := DiscoverProfiles(c.Config.WireGuardDir)
	if err != nil {
		return []ProfileState{}, err
	}

	activeInterfaces, err := ActiveInterfaces()
	if err != nil {
		return []ProfileState{}, err
	}

	activeByName := make(map[string]InterfaceStatus, len(activeInterfaces))
	for _, iface := range activeInterfaces {
		activeByName[iface.Name] = iface
	}

	profileStates := make([]ProfileState, 0, len(profiles))

	for _, profile := range profiles {
		state := ProfileState{
			Profile: profile,
		}
		if iface, ok := activeByName[profile.Name]; ok {
			state.IsActive = true
			state.Interface = &iface
		}

		profileStates = append(profileStates, state)
	}
	return profileStates, nil
}
