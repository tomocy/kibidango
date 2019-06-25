package saver

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

func (l *Linux) Save(kibi *kibidango.Kibidango) error {
	destDir := filepath.Join("/run/kibidango", kibi.ID)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	destName := filepath.Join(destDir, "state.json")
	dest, err := os.OpenFile(destName, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}

	return json.NewEncoder(dest).Encode(kibi)
}
