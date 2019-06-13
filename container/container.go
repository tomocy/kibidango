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
	Launch() error
}

type booter interface {
	Boot() error
}

func (c *Container) Run(conf *config.Config) error {
	switch conf.Phase {
	case config.PhaseLaunch:
		return c.launch()
	case config.PhaseBoot:
		return c.boot()
	default:
		return nil
	}
}

func (c *Container) launch() error {
	return launch(c.launcher)
}

func (c *Container) boot() error {
	return boot(c.booter)
}

func launch(lcher launcher) error {
	return lcher.Launch()
}

func boot(bter booter) error {
	return bter.Boot()
}
