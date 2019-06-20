package container

import (
	"github.com/tomocy/kibidango/engine/booter"
	"github.com/tomocy/kibidango/engine/creater"
)

func New(creater creater.Creater, bter booter.Booter) *Container {
	return &Container{
		creater: creater,
		booter:  bter,
	}
}

type Container struct {
	creater creater.Creater
	booter  booter.Booter
}

func (c *Container) Create() error {
	return create(c.creater)
}

func (c *Container) Boot() error {
	return boot(c.booter)
}

func create(creater creater.Creater) error {
	return creater.Create()
}

func boot(bter booter.Booter) error {
	return bter.Boot()
}
