package container

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/tomocy/kibidango/config"
)

const (
	cfsQuotaUS = 5000
)

var (
	bins = []string{"/bin/sh", "/bin/ls", "/bin/ps", "/bin/cat", "/bin/date"}
	libs = []string{"/lib/ld-musl-x86_64.so.1"}
)

type Linux struct {
	Root           string
	Input          io.Reader
	Output, Errput io.Writer
}

func (c *Linux) Run(conf *config.Config) error {
	switch conf.Phase {
	case config.PhaseLaunch:
		return c.launch(conf.Command)
	case config.PhaseBoot:
		return c.boot(conf.Command)
	default:
		return nil
	}
}

func (c *Linux) launch(cmd string) error {
	cloneCmd := c.buildCloneCommand("-boot", "-command="+cmd)
	return cloneCmd.Run()
}

func (c *Linux) buildCloneCommand(args ...string) *exec.Cmd {
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
	cmd.Stdin, cmd.Stdout, cmd.Stderr = c.Input, c.Output, c.Errput

	return cmd
}

func (c *Linux) boot(name string) error {
	if err := c.load(); err != nil {
		return reportErr("load", err)
	}

	return syscall.Exec(name, []string{name}, os.Environ())
}

func (c *Linux) load() error {
	if err := syscall.Sethostname([]byte("container")); err != nil {
		return err
	}
	if err := c.limit(); err != nil {
		return reportErr("limit", err)
	}
	if err := c.enable(bins, libs...); err != nil {
		return reportErr("enable", err)
	}
	if err := c.mountProcs(); err != nil {
		return reportErr("mount procs", err)
	}
	if err := c.pivotRoot(); err != nil {
		return reportErr("pivot root", err)
	}

	return nil
}

func (c *Linux) limit() error {
	if err := c.limitCPUUsage(); err != nil {
		return reportErr("limit cpu usage", err)
	}

	return nil
}

func (c *Linux) limitCPUUsage() error {
	if err := os.MkdirAll("/sys/fs/cgroup/cpu/container", 0755); err != nil {
		return err
	}
	if err := ioutil.WriteFile(
		"/sys/fs/cgroup/cpu/container/cpu.cfs_quota_us",
		[]byte(fmt.Sprintf("%d", cfsQuotaUS)),
		0755,
	); err != nil {
		return err
	}
	if err := ioutil.WriteFile(
		"/sys/fs/cgroup/cpu/container/tasks",
		[]byte(fmt.Sprintf("%d", os.Getpid())),
		0755,
	); err != nil {
		return err
	}

	return nil
}

func (c *Linux) enable(bins []string, libs ...string) error {
	if 1 <= len(libs) {
		if err := c.ensure(libs); err != nil {
			return reportErr("ensure", err)
		}
	}

	if err := os.MkdirAll(c.joinRoot("/bin"), 0755); err != nil {
		return err
	}

	for _, bin := range bins {
		if err := copyFile(bin, c.joinRoot(bin)); err != nil {
			return reportErr("copy file", err)
		}
	}

	return nil
}

func (c *Linux) ensure(libs []string) error {
	if err := os.MkdirAll(c.joinRoot("/lib"), 0755); err != nil {
		return err
	}

	for _, lib := range libs {
		if err := copyFile(lib, c.joinRoot(lib)); err != nil {
			return reportErr("copy file", err)
		}
	}

	return nil
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

func (c *Linux) mountProcs() error {
	if err := os.MkdirAll(c.joinRoot("/proc"), 0755); err != nil {
		return err
	}
	if err := syscall.Mount("/proc", c.joinRoot("/proc"), "proc", syscall.MS_NOEXEC|syscall.MS_NOSUID|syscall.MS_NODEV, ""); err != nil {
		return err
	}

	return nil
}

func (c *Linux) pivotRoot() error {
	if err := os.MkdirAll(c.joinRoot("/oldfs"), 0755); err != nil {
		return err
	}
	if err := syscall.Mount(c.Root, c.Root, "", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return err
	}
	if err := syscall.PivotRoot(c.Root, c.joinRoot("/oldfs")); err != nil {
		return err
	}
	if err := syscall.Chdir("/"); err != nil {
		return err
	}
	if err := syscall.Unmount("/oldfs", syscall.MNT_DETACH); err != nil {
		return err
	}
	if err := os.RemoveAll("/oldfs"); err != nil {
		return err
	}

	return nil
}

func (c *Linux) joinRoot(path string) string {
	return filepath.Join(c.Root, path)
}

func reportErr(did string, err error) error {
	return fmt.Errorf("failed to %s; %s", did, err)
}
