package config

import "flag"

func Parse() *Config {
	willBoot := flag.Bool("boot", false, "boot container")
	flag.Parse()
	phase := PhaseLaunch
	if *willBoot {
		phase = PhaseBoot
	}

	return &Config{
		Phase: phase,
	}
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
