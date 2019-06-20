package container

import (
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

func (c *Container) Launch() error {
	return launch(c.launcher)
}

func (c *Container) Boot() error {
	return boot(c.booter)
}

func launch(lcher launcher.Launcher) error {
	return lcher.Launch()
}

func boot(bter booter.Booter) error {
	return bter.Boot()
}
