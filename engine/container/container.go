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

func (c *Container) Create(args ...string) error {
	return create(c.creater, args...)
}

func (c *Container) Init() error {
	return initialize(c.initer)
}

func create(creater creater.Creater, args ...string) error {
	return creater.Create(args...)
}

func initialize(initer initializer.Initializer) error {
	return initer.Init()
}
