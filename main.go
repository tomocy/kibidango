package main

import (
	"fmt"
	"os"

	"github.com/tomocy/kibidango/config"
	"github.com/tomocy/kibidango/container"
)

func main() {
	conf := config.Parse()
	cont := &container.Linux{
		Root:   "/root/container",
		Input:  os.Stdin,
		Output: os.Stdout,
		Errput: os.Stderr,
	}
	if err := cont.Run(conf); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run: %s\n", err)
		os.Exit(1)
	}
}
