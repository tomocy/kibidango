package cloner

import (
	"io"
	"os"
	"os/exec"
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

func (l *Linux) Clone(args ...string) error {
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