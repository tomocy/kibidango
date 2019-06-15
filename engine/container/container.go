package container

import (
	"github.com/tomocy/kibidango/config"
	"github.com/tomocy/kibidango/engine/booter"
	"github.com/tomocy/kibidango/engine/launcher"
)

func New(lcher launcher.Launcher, bter booter.Booter) *Container {
	return &Container{
		launcher: lcher,
		booter:   bter,
	}
}

type Container struct {
	launcher launcher.Launcher
	booter   booter.Booter
}

func (c *Container) Start(conf *config.Config) error {
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

func launch(lcher launcher.Launcher, cmd string) error {
	return lcher.Launch(cmd)
}

func boot(bter booter.Booter, cmd string) error {
	return bter.Boot(cmd)
}
