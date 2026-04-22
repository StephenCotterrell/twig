package main

import (
	"fmt"
	"log"

	"github.com/StephenCotterrell/twig/cmd/internal/wg"
)

func main() {
	cfg := wg.DefaultConfig()

	profiles, err := wg.DiscoverProfiles(cfg.WireGuardDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range profiles {
		fmt.Println(p.Name)
		fmt.Println(p.Path)
	}
}
