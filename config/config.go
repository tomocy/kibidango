package config

import "flag"

func Parse() *Config {
	willBoot := flag.Bool("boot", false, "boot container")
	flag.Parse()
	cmd := CommandLaunch
	if *willBoot {
		cmd = CommandBoot
	}

	return &Config{
		Command: cmd,
	}
}

type Config struct {
	Command Command
}

const (
	_ Command = iota
	CommandLaunch
	CommandBoot
)

type Command int
