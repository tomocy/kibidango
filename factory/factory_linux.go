package factory

import (
	"github.com/tomocy/kibidango"
	errorPkg "github.com/tomocy/kibidango/error"
	"golang.org/x/sys/unix"
)

func ForLinux() *Linux {
	return new(Linux)
}

type Linux struct{}

func (l *Linux) List() ([]*kibidango.Linux, error) {
	states, err := list()
	if err != nil {
		return nil, errorPkg.Report("list", err)
	}

	kibis := make([]*kibidango.Linux, len(states))
	for i, state := range states {
		kibi, err := l.adapt(state)
		if err != nil {
			return nil, errorPkg.Report("list", err)
		}

		kibis[i] = kibi
	}

	return kibis, nil
}

func (l *Linux) Manufacture(spec *kibidango.Spec) (*kibidango.Linux, error) {
	if err := l.createWorkspace(spec.ID); err != nil {
		return nil, errorPkg.Report("manufacture", err)
	}

	kibi, err := kibidango.ForLinux(spec)
	if err != nil {
		return nil, errorPkg.Report("manufacture", err)
	}

	return kibi, nil
}

func (l *Linux) createWorkspace(id string) error {
	if err := createWorkspace(id); err != nil {
		return err
	}

	return l.createPipeFile(id)
}

func (l *Linux) createPipeFile(id string) error {
	name := pipeFilename(id)
	return unix.Mkfifo(name, 0777)
}

func (l *Linux) Load(id string) (*kibidango.Linux, error) {
	state, err := load(id)
	if err != nil {
		return nil, errorPkg.Report("load", err)
	}

	kibi, err := l.adapt(state)
	if err != nil {
		return nil, errorPkg.Report("load", err)
	}

	return kibi, nil
}

func (l *Linux) adapt(spec *kibidango.Spec) (*kibidango.Linux, error) {
	return kibidango.ForLinux(spec)
}

func (l *Linux) Save(kibi *kibidango.Linux) error {
	spec := kibi.Spec()
	if err := save(spec); err != nil {
		return errorPkg.Report("save", err)
	}

	return nil
}

func (l *Linux) Delete(id string) error {
	if err := delete(id); err != nil {
		return errorPkg.Report("delete", err)
	}

	return nil
}
