package wg

import (
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	WireGuardDir string
}

func DefaultConfig() Config {
	return Config{
		WireGuardDir: "/etc/wireguard",
	}
}

func DiscoverProfiles(dir string) ([]Profile, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var profiles []Profile
	for _, entry := range entries {
		if filepath.Ext(entry.Name()) != ".conf" {
			continue
		}

		name := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
		profiles = append(profiles, Profile{
			Name: name,
			Path: filepath.Join(dir, entry.Name()),
		})
	}

	return profiles, nil
}
