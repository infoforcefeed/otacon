package main

import (
	"log"

	"github.com/infoforcefeed/otacon/config"
	"github.com/infoforcefeed/otacon/tracker"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	t := tracker.New(cfg)

	t.Run()
}
