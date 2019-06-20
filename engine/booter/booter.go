package booter

import "github.com/tomocy/kibidango/config"

func ForOS(os string, root string) Booter {
	switch os {
	case config.OSLinux:
		return ForLinux(root)
	default:
		return nil
	}
}

type Booter interface {
	Boot() error
}
