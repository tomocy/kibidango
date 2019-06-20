package launcher

import "io"

func ForOS(os string, input io.Reader, output, errput io.Writer) Launcher {
	switch os {
	case "linux":
		return ForLinux(input, output, errput)
	default:
		return nil
	}
}

type Launcher interface {
	Launch() error
}
