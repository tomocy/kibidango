package container

import "github.com/tomocy/kibidango/config"

func New(lcher launcher, bter booter) *Container {
	return &Container{
		launcher: lcher,
		booter:   bter,
	}
}

type Container struct {
	launcher launcher
	booter   booter
}

type launcher interface {
	Launch(cmd string) error
}

type booter interface {
	Boot(cmd string) error
}

func (c *Container) Run(conf *config.Config) error {
	switch conf.Phase {
	case config.PhaseLaunch:
		return c.launch(conf.Command)
	case config.PhaseBoot:
		return c.boot(conf.Command)
	default:
		return nil
	}
}

func (c *Container) launch(cmd string) error {
	return launch(c.launcher, cmd)
}

func (c *Container) boot(cmd string) error {
	return boot(c.booter, cmd)
}

func launch(lcher launcher, cmd string) error {
	return lcher.Launch(cmd)
}

func boot(bter booter, cmd string) error {
	return bter.Boot(cmd)
}
