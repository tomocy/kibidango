package container

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/tomocy/kibidango/config"
)

type LinuxContainer struct{}

func (c *LinuxContainer) Run(conf *config.Config) error {
	switch conf.Command {
	case config.CommandLaunch:
		return c.launch()
	case config.CommandLoad:
		return c.load("/bin/sh")
	default:
		return nil
	}
}

func (c *LinuxContainer) launch() error {
	cmd := buildCloneCommand("-load")
	return cmd.Run()
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
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

	return cmd
}

func (c *LinuxContainer) load(name string) error {
	if err := c.init(); err != nil {
		return err
	}

	return syscall.Exec(name, []string{name}, os.Environ())
}

func (c *LinuxContainer) init() error {
	if err := syscall.Sethostname([]byte("container")); err != nil {
		return err
	}
	if err := syscall.Mount("/proc", "/proc", "proc", uintptr(
		syscall.MS_NOEXEC|syscall.MS_NOSUID|syscall.MS_NODEV,
	), ""); err != nil {
		return err
	}

	return nil
}
