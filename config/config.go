package config

import "flag"

func Parse() *Config {
	willLoad := flag.Bool("load", false, "load process as container")
	flag.Parse()
	cmd := CommandLaunch
	if *willLoad {
		cmd = CommandLoad
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
	CommandLoad
)

type Command int
