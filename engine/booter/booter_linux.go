package booter

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)

const (
	cfsQuotaUS = 5000
)

var (
	bins = []string{"/bin/sh", "/bin/ls", "/bin/ps", "/bin/cat", "/bin/date"}
	libs = []string{"/lib/ld-musl-x86_64.so.1"}
)

func ForLinux(root string) *Linux {
	return &Linux{
		root: root,
	}
}

type Linux struct {
	root string
}

func (l *Linux) Boot() error {
	if err := syscall.Sethostname([]byte("container")); err != nil {
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

func (l *Linux) joinRoot(path string) string {
	return filepath.Join(l.root, path)
}

func reportErr(did string, err error) error {
	return fmt.Errorf("failed to %s; %s", did, err)
}
