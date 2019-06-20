package container

import (
	"github.com/tomocy/kibidango/engine/creater"
	"github.com/tomocy/kibidango/engine/initializer"
)

func New(creater creater.Creater, initer initializer.Initializer) *Container {
	return &Container{
		creater: creater,
		initer:  initer,
	}
}

type Container struct {
	creater creater.Creater
	initer  initializer.Initializer
}

func (c *Container) Create() error {
	return create(c.creater)
}

func (c *Container) Init() error {
	return initialize(c.initer)
}

func create(creater creater.Creater) error {
	return creater.Create()
}

func initialize(initer initializer.Initializer) error {
	return initer.Init()
}
