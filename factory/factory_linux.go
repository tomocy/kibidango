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
	state := l.convertToState(kibi)
	return save(state)
}

func (l *Linux) convertToState(kibi *kibidango.Linux) *state {
	return &state{
		ID: kibi.ID(),
	}
}
