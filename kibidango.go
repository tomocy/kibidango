package kibidango

import "errors"

type Kibidango struct {
	ID string
}

func (c *Kibidango) UpdateID(id string) error {
	if id == "" {
		return errors.New("id should not be empty")
	}
	c.ID = id

	return nil
}

func (c *Kibidango) Create(creater Creater, args ...string) error {
	return create(creater, args...)
}

func create(creater Creater, args ...string) error {
	return creater.Create(args...)
}

type Creater interface {
	Create(args ...string) error
}

func (c *Kibidango) Init(initer Initializer) error {
	return initialize(initer)
}

func initialize(initer Initializer) error {
	return initer.Init()
}

type Initializer interface {
	Init() error
}
