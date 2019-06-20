package container

import (
	"github.com/tomocy/kibidango/engine/creater"
	"github.com/tomocy/kibidango/engine/initializer"
)

type Container struct{}

func (c *Container) Create(creater creater.Creater, args ...string) error {
	return create(creater, args...)
}

func (c *Container) Init(initer initializer.Initializer) error {
	return initialize(initer)
}

func create(creater creater.Creater, args ...string) error {
	return creater.Create(args...)
}

func initialize(initer initializer.Initializer) error {
	return initer.Init()
}
