package launcher

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

func NewLinux(
	input io.Reader,
	output, errput io.Writer,
	cmd string,
) *Linux {
	return &Linux{
		input:   input,
		output:  output,
		errput:  errput,
		command: cmd,
	}
}

type Linux struct {
	input          io.Reader
	output, errput io.Writer
	command        string
}

func (l *Linux) Launch() error {
	cmd := l.buildCloneCommand("-boot", fmt.Sprintf("-command=%s", l.command))
	return cmd.Run()
}

func (l *Linux) buildCloneCommand(args ...string) *exec.Cmd {
	cmd := buildCloneCommand(args...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = l.input, l.output, l.errput

	return cmd
}

func buildCloneCommand(args ...string) *exec.Cmd {
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

	return cmd
}
