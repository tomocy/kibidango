package factory

import "github.com/tomocy/kibidango"

func ForLinux() *Linux {
	return new(Linux)
}

type Linux struct{}

func (l *Linux) Manufacture(id string) *kibidango.Linux {
	return kibidango.ForLinux(id)
}
