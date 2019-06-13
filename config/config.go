package config

import (
	"flag"
	"runtime"
)

func Parse() *Config {
	flags := parseFlags()
	phase := PhaseLaunch
	if flags.boot {
		phase = PhaseBoot
	}

	return &Config{
		OS:      runtime.GOOS,
		Phase:   phase,
		Command: flags.command,
	}
}

func parseFlags() *flags {
	flags := new(flags)
	flag.BoolVar(&flags.boot, "boot", false, "boot container")
	flag.StringVar(&flags.command, "command", "/bin/sh", "a command to be executed after boot")
	flag.Parse()

	return flags
}

type flags struct {
	boot    bool
	command string
}

type Config struct {
	OS      string
	Phase   Phase
	Command string
}

const (
	OSLinux = "linux"

	_ Phase = iota
	PhaseLaunch
	PhaseBoot
)

type Phase int
