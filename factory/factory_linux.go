package factory

import (
	"github.com/tomocy/kibidango"
)

func ForLinux() *Linux {
	return new(Linux)
}

type Linux struct{}

func (l *Linux) List() ([]*kibidango.Linux, error) {
	states, err := list()
	if err != nil {
		return nil, err
	}

	kibis := make([]*kibidango.Linux, len(states))
	for i, state := range states {
		kibi, err := l.adapt(state)
		if err != nil {
			return nil, err
		}

		kibis[i] = kibi
	}

	return kibis, nil
}

func (l *Linux) Manufacture(id string) (*kibidango.Linux, error) {
	if err := createWorkspace(id); err != nil {
		return nil, err
	}

	return kibidango.ForLinux(id)
}

func (l *Linux) Load(id string) (*kibidango.Linux, error) {
	state, err := load(id)
	if err != nil {
		return nil, err
	}

	return l.adapt(state)
}

func (l *Linux) adapt(state *state) (*kibidango.Linux, error) {
	return kibidango.ForLinux(state.ID)
}

func (l *Linux) Save(kibi *kibidango.Linux) error {
	state := l.convert(kibi)
	return save(state)
}

func (l *Linux) convert(kibi *kibidango.Linux) *state {
	return &state{
		ID:      kibi.ID(),
		Process: kibi.Process(),
	}
}

func (l *Linux) Delete(id string) error {
	return delete(id)
}
