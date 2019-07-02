package kibidango

import (
	"os"
	"os/exec"
	"syscall"
)

func ForLinux(id string) (*Linux, error) {
	kibi := new(kibidango)
	if err := kibi.updateID(id); err != nil {
		return nil, err
	}

	return &Linux{
		kibidango: kibi,
	}, nil
}

type Linux struct {
	*kibidango
}

func (l *Linux) Run(args ...string) error {
	return l.clone(args...)
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
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

	return cmd
}
