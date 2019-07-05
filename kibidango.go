package kibidango

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	errorPkg "github.com/tomocy/kibidango/error"
)

type kibidango struct {
	id      string
	root    string
	process *Process
	pipeFD  int
}

type Spec struct {
	ID      string
	Root    string
	Process *Process
	PipeFD  int
}

type Process struct {
	Args []string
}

func (k *kibidango) Spec() *Spec {
	return &Spec{
		ID:      k.id,
		Root:    k.root,
		Process: k.process,
		PipeFD:  k.pipeFD,
	}
}

func (k *kibidango) Meet(spec *Spec) error {
	if err := k.updateID(spec.ID); err != nil {
		return errorPkg.Report("meet", err)
	}
	if err := k.updateRoot(spec.Root); err != nil {
		return errorPkg.Report("meet", err)
	}
	if err := k.updateProcess(spec.Process); err != nil {
		return errorPkg.Report("meet", err)
	}
	if err := k.UpdatePipeFD(spec.PipeFD); err != nil {
		return errorPkg.Report("meet", err)
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

func (k *kibidango) UpdatePipeFD(fd int) error {
	k.pipeFD = fd
	return nil
}

func (k *kibidango) joinRoot(path string) string {
	return filepath.Join(k.root, path)
}

func (k *kibidango) writePipe() error {
	name := fmt.Sprintf("/proc/self/fd/%d", k.pipeFD)
	pipe, err := os.OpenFile(name, os.O_WRONLY, 0777)
	if err != nil {
		return err
	}

	return pipe.Close()
}

func (k *kibidango) readPipe() error {
	name := fmt.Sprintf("/proc/self/fd/%d", k.pipeFD)
	pipe, err := os.Open(name)
	if err != nil {
		return err
	}

	return pipe.Close()
}
