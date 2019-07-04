package kibidango

import "errors"

type kibidango struct {
	id      string
	root    string
	process *Process
}

type Spec struct {
	ID      string
	Process *Process
}

type Process struct {
	Args []string
}

func (k *kibidango) ID() string {
	return k.id
}

func (k *kibidango) Process() *Process {
	return k.process
}

func (k *kibidango) Meet(spec *Spec) error {
	if err := k.UpdateID(spec.ID); err != nil {
		return err
	}
	if err := k.UpdateProcess(spec.Process); err != nil {
		return err
	}

	return nil
}

func (k *kibidango) UpdateID(id string) error {
	if id == "" {
		return errors.New("id should not be empty")
	}

	k.id = id

	return nil
}

func (k *kibidango) UpdateRoot(root string) error {
	k.root = root
	return nil
}

func (k *kibidango) UpdateProcess(proc *Process) error {
	k.process = proc
	return nil
}
