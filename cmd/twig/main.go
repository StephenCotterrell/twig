package main

import (
	"fmt"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/StephenCotterrell/twig/cmd/internal/app"
	"github.com/StephenCotterrell/twig/cmd/internal/wg"
)

func main() {
	cfg := wg.DefaultConfig()
	profiles, err := wg.DiscoverProfiles(cfg.WireGuardDir)
	if err != nil {
		log.Fatal("failed to discover profiles")
	}
	p := tea.NewProgram(app.InitialModel(profiles))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
