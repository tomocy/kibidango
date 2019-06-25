package terminator

import (
	"os"
	"path/filepath"

	"github.com/tomocy/kibidango"
)

func ForLinux() *Linux {
	return new(Linux)
}

type Linux struct{}

func (l *Linux) Terminate(k *kibidango.Kibidango) error {
	dir := filepath.Join("/run/kibidango", k.ID)
	return os.RemoveAll(dir)
}
