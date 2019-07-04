package kibidango

import (
	"errors"
	"fmt"
	"path/filepath"
)

type kibidango struct {
	id      string
	root    string
	process *Process
}

type Spec struct {
	ID      string
	Root    string
	Process *Process
}

type Process struct {
	Args []string
}

func (k *kibidango) Spec() *Spec {
	return &Spec{
		ID:      k.id,
		Root:    k.root,
		Process: k.process,
	}
}

func (k *kibidango) Meet(spec *Spec) error {
	if err := k.updateID(spec.ID); err != nil {
		return err
	}
	if err := k.updateRoot(spec.Root); err != nil {
		return err
	}
	if err := k.updateProcess(spec.Process); err != nil {
		return err
	}

	return nil
}

func (k *kibidango) updateID(id string) error {
	if id == "" {
		return errors.New("id should not be empty")
	}

	k.id = id

	return nil
}

func (k *kibidango) updateRoot(root string) error {
	k.root = root
	return nil
}

func (k *kibidango) updateProcess(proc *Process) error {
	k.process = proc
	return nil
}

func (k *kibidango) joinRoot(path string) string {
	return filepath.Join(k.root, path)
}

func report(did string, err error) error {
	return fmt.Errorf("failed to %s; %s", did, err)
}
