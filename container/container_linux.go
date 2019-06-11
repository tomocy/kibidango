package container

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
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
	if err := c.prepare(); err != nil {
		return err
	}
	if err := c.init(); err != nil {
		return err
	}

	return syscall.Exec(name, []string{name}, os.Environ())
}

func (c *LinuxContainer) prepare() error {
	if err := os.MkdirAll("/root/container/proc", 0777); err != nil {
		return err
	}
	if err := os.MkdirAll("/root/container/bin", 0777); err != nil {
		return err
	}
	if err := os.MkdirAll("/root/container/lib", 0777); err != nil {
		return err
	}
	if err := enable("/bin/sh", "/lib/ld-musl-x86_64.so.1"); err != nil {
		return err
	}

	return nil
}

func enable(name string, deps ...string) error {
	for _, dep := range deps {
		if err := copyFile(dep, filepath.Join("/root/container", dep)); err != nil {
			return err
		}
	}

	return copyFile(name, filepath.Join("/root/container", name))
}

func copyFile(src, dest string) error {
	input, err := os.Open(src)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer output.Close()

	if _, err := io.Copy(output, input); err != nil {
		return err
	}

	return nil
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
	if err := syscall.Chroot("/root/container"); err != nil {
		return err
	}
	if err := syscall.Chdir("/"); err != nil {
		return err
	}

	return nil
}
