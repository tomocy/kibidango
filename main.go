package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	config := parseConfig()
	if err := run(config); err != nil {
		fmt.Printf("failed to run: %s", err)
		os.Exit(1)
	}
}

func run(config *config) error {
	switch config.command {
	case cmdClone:
		return clone("-init")
	case cmdInit:
		return initAndExec("/bin/sh")
	default:
		showUsage()
		return nil
	}
}

func parseConfig() *config {
	isInit := flag.Bool("init", false, "init self as container")
	flag.Parse()
	cmd := cmdClone
	if *isInit {
		cmd = cmdInit
	}

	return &config{
		command: cmd,
	}
}

type config struct {
	command command
}

const (
	_ command = iota
	cmdClone
	cmdInit
)

type command int

func showUsage() {
	flag.Usage()
}

func clone(args ...string) error {
	cmd := exec.Command("/proc/self/exe", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWIPC | syscall.CLONE_NEWNET |
			syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUSER | syscall.CLONE_NEWUTS,
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

	return cmd.Run()
}

func initAndExec(name string) error {
	if err := initAsContainer(); err != nil {
		return err
	}

	return syscall.Exec(name, []string{name}, os.Environ())
}

func initAsContainer() error {
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
