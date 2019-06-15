package main

import (
	"fmt"
	"os"

	"github.com/tomocy/kibidango/config"
	"github.com/tomocy/kibidango/engine/booter"
	"github.com/tomocy/kibidango/engine/container"
	"github.com/tomocy/kibidango/engine/launcher"
)

func main() {
	conf := config.Parse()
	lcher := launcher.ForOS(conf.OS, os.Stdin, os.Stdout, os.Stderr)
	bter := booter.ForOS(conf.OS, "/root/container")
	cner := container.New(lcher, bter)
	if err := cner.Start(conf); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run: %s\n", err)
		os.Exit(1)
	}
}
