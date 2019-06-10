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
	switch config.command {
	case cmdClone:
		if err := clone("-init"); err != nil {
			fmt.Printf("failed to clone: %s\n", err)
			os.Exit(1)
		}
	case cmdInit:
		if err := initAndExec("/bin/sh"); err != nil {
			fmt.Printf("failed to init and exec: %s\n", err)
			os.Exit(1)
		}
	default:
		showUsage()
	}
}

func parseConfig() *config {
	isSetUp := flag.Bool("init", false, "init self as container")
	flag.Parse()
	cmd := cmdClone
	if *isSetUp {
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
		Cloneflags: syscall.CLONE_NEWIPC | syscall.CLONE_NEWNET | syscall.CLONE_NEWUSER | syscall.CLONE_NEWUTS,
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

	return nil
}
