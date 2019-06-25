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

func (k *Kibidango) Clone(cloner Cloner, args ...string) error {
	return cloner.Clone(args...)
}

func cloner(cloner Cloner, args ...string) error {
	return cloner.Clone(args...)
}

type Cloner interface {
	Clone(args ...string) error
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

func List(lister Lister, loader Loader) ([]*Kibidango, error) {
	return lister.List(loader)
}

type Lister interface {
	List(loader Loader) ([]*Kibidango, error)
}

func (k *Kibidango) Load(loader Loader) error {
	return load(loader, k)
}

func load(loader Loader, kibi *Kibidango) error {
	return loader.Load(kibi)
}

type Loader interface {
	Load(kibi *Kibidango) error
}

func (k *Kibidango) Terminate(terminator Terminator) error {
	return terminate(terminator, k)
}

func terminate(terminator Terminator, kibi *Kibidango) error {
	return terminator.Terminate(kibi)
}

type Terminator interface {
	Terminate(kibi *Kibidango) error
}
