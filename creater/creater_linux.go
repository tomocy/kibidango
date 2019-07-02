package creater

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func ForLinux(
	input io.Reader,
	output, errput io.Writer,
) *Linux {
	return &Linux{
		input:  input,
		output: output,
		errput: errput,
	}
}

type Linux struct {
	input          io.Reader
	output, errput io.Writer
}

func (l *Linux) Create(id string, args ...string) error {
	if err := l.prepare(id); err != nil {
		return err
	}

	return l.clone(args...)
}

func (l *Linux) prepare(id string) error {
	if err := l.createWorkspace(id); err != nil {
		return err
	}

	return nil
}

func (l *Linux) createWorkspace(id string) error {
	dir := filepath.Join("/run/kibidango", id)
	if err := os.MkdirAll(dir, 0777); err != nil {
		return err
	}

	return nil
}

func (l *Linux) clone(args ...string) error {
	cmd := l.buildCloneCommand(args...)
	return cmd.Run()
}

func (l *Linux) buildCloneCommand(args ...string) *exec.Cmd {
	cmd := exec.Command("/proc/self/exe", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWIPC | syscall.CLONE_NEWNET | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWPID | syscall.CLONE_NEWUSER | syscall.CLONE_NEWUTS,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
	}
	cmd.Stdin, cmd.Stdout, cmd.Stderr = l.input, l.output, l.errput

	return cmd
}
