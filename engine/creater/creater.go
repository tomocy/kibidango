package creater

import "io"

func ForOS(os string, input io.Reader, output, errput io.Writer) Creater {
	switch os {
	case "linux":
		return ForLinux(input, output, errput)
	default:
		return nil
	}
}

type Creater interface {
	Create(args ...string) error
}
