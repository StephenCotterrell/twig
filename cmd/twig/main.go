package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/StephenCotterrell/twig/internal/app"
	"github.com/StephenCotterrell/twig/internal/wg"
)

func main() {
	cfg := wg.DefaultConfig()

	m := app.InitialModel(cfg)
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
