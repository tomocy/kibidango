package kibidango

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

func ForLinux(spec *Spec) (*Linux, error) {
	kibi := new(kibidango)
	if err := kibi.Meet(spec); err != nil {
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

func (l *Linux) Init() error {
	if err := syscall.Sethostname([]byte(l.id)); err != nil {
		return err
	}
	if err := l.limit(); err != nil {
		return reportErr("limit", err)
	}
	if err := l.enable(bins, libs...); err != nil {
		return reportErr("enable", err)
	}
	if err := l.mountProcs(); err != nil {
		return reportErr("mount procs", err)
	}
	if err := l.pivotRoot(); err != nil {
		return reportErr("pivot root", err)
	}

	return nil
}

var (
	bins = []string{"/bin/sh", "/bin/ls", "/bin/ps", "/bin/cat", "/bin/date"}
	libs = []string{"/lib/ld-musl-x86_64.so.1"}
)

func (l *Linux) limit() error {
	if err := l.limitCPUUsage(); err != nil {
		return reportErr("limit cpu usage", err)
	}

	return nil
}

func (l *Linux) limitCPUUsage() error {
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

const (
	cfsQuotaUS = 5000
)

func (l *Linux) enable(bins []string, libs ...string) error {
	if 1 <= len(libs) {
		if err := l.ensure(libs); err != nil {
			return reportErr("ensure", err)
		}
	}

	if err := os.MkdirAll(l.joinRoot("/bin"), 0755); err != nil {
		return err
	}

	for _, bin := range bins {
		if err := copyFile(bin, l.joinRoot(bin)); err != nil {
			return reportErr("copy file", err)
		}
	}

	return nil
}

func (l *Linux) ensure(libs []string) error {
	if err := os.MkdirAll(l.joinRoot("/lib"), 0755); err != nil {
		return err
	}

	for _, lib := range libs {
		if err := copyFile(lib, l.joinRoot(lib)); err != nil {
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

func (l *Linux) mountProcs() error {
	if err := os.MkdirAll(l.joinRoot("/proc"), 0755); err != nil {
		return err
	}
	if err := syscall.Mount("/proc", l.joinRoot("/proc"), "proc", syscall.MS_NOEXEC|syscall.MS_NOSUID|syscall.MS_NODEV, ""); err != nil {
		return err
	}

	return nil
}

func (l *Linux) pivotRoot() error {
	if err := os.MkdirAll(l.joinRoot("/oldfs"), 0755); err != nil {
		return err
	}
	if err := syscall.Mount(l.root, l.root, "", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return err
	}
	if err := syscall.PivotRoot(l.root, l.joinRoot("/oldfs")); err != nil {
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

func reportErr(did string, err error) error {
	return fmt.Errorf("failed to %s; %s", did, err)
}
