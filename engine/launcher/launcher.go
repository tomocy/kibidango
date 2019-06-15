package launcher

import (
	"io"

	"github.com/tomocy/kibidango/config"
)

func ForOS(os string, input io.Reader, output, errput io.Writer) Launcher {
	switch os {
	case config.OSLinux:
		return ForLinux(input, output, errput)
	default:
		return nil
	}
}

type Launcher interface {
	Launch(cmd string) error
}
