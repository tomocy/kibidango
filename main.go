package main

import (
	"fmt"
	"os"

	"github.com/tomocy/kibidango/booter"
	"github.com/tomocy/kibidango/config"
	"github.com/tomocy/kibidango/container"
	"github.com/tomocy/kibidango/launcher"
)

func main() {
	conf := config.Parse()
	lcher := launcher.NewLinux(os.Stdin, os.Stdout, os.Stderr)
	bter := booter.NewLinux("/root/container")
	cner := container.New(lcher, bter)
	if err := cner.Start(conf); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run: %s\n", err)
		os.Exit(1)
	}
}
