package loader

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/tomocy/kibidango"
)

func ForLinux() *Linux {
	return new(Linux)
}

type Linux struct{}

func (l *Linux) Load(kibi *kibidango.Kibidango) error {
	srcName := filepath.Join("/run/kibidango", kibi.ID, "state.json")
	src, err := os.Open(srcName)
	if err != nil {
		return err
	}
	defer src.Close()

	return json.NewDecoder(src).Decode(kibi)
}
