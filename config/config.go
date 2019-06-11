package config

import "flag"

func Parse() *Config {
	flags := parseFlags()
	phase := PhaseLaunch
	if flags.boot {
		phase = PhaseBoot
	}

	return &Config{
		Phase: phase,
	}
}

func parseFlags() *flags {
	flags := new(flags)
	flag.BoolVar(&flags.boot, "boot", false, "boot container")
	flag.Parse()

	return flags
}

type flags struct {
	boot bool
}

type Config struct {
	Phase Phase
}

const (
	_ Phase = iota
	PhaseLaunch
	PhaseBoot
)

type Phase int
