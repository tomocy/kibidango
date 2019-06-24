package kibidango

import "errors"

type Kibidango struct {
	ID string
}

func (k *Kibidango) UpdateID(id string) error {
	if id == "" {
		return errors.New("id should not be empty")
	}
	k.ID = id

	return nil
}

func (k *Kibidango) Create(creater Creater, args ...string) error {
	return create(creater, args...)
}

func create(creater Creater, args ...string) error {
	return creater.Create(args...)
}

type Creater interface {
	Create(args ...string) error
}

func (k *Kibidango) Init(initer Initializer) error {
	return initialize(initer)
}

func initialize(initer Initializer) error {
	return initer.Init()
}

type Initializer interface {
	Init() error
}

func (k *Kibidango) Save(saver Saver) error {
	return save(saver, k)
}

func save(saver Saver, kibi *Kibidango) error {
	return saver.Save(kibi)
}

type Saver interface {
	Save(kibi *Kibidango) error
}
