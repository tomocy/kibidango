package factory

import "github.com/tomocy/kibidango"

func ForLinux() *Linux {
	return new(Linux)
}

type Linux struct{}

func (l *Linux) Manufacture(id string) (*kibidango.Linux, error) {
	if err := createWorkspace(id); err != nil {
		return nil, err
	}

	return kibidango.ForLinux(id)
}

func (l *Linux) Save(kibi *kibidango.Linux) error {
	state := l.convert(kibi)
	return save(state)
}

func (l *Linux) convert(kibi *kibidango.Linux) *state {
	return &state{
		ID: kibi.ID(),
	}
}

func (l *Linux) Load(id string) (*kibidango.Linux, error) {
	state, err := load(id)
	if err != nil {
		return nil, err
	}

	return l.adapt(state)
}

func (l *Linux) adapt(state *state) (*kibidango.Linux, error) {
	kibi := new(kibidango.Linux)
	if err := kibi.UpdateID(state.ID); err != nil {
		return nil, err
	}

	return kibi, nil
}
