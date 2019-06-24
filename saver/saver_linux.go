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

func (l *Linux) Save(ctner *kibidango.Kibidango) error {
	destDir := filepath.Join("/run/kibidango", ctner.ID)
	if err := os.MkdirAll(destDir, 0777); err != nil {
		return err
	}
	destName := filepath.Join(destDir, "state.json")
	dest, err := os.OpenFile(destName, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}

	return json.NewEncoder(dest).Encode(ctner)
}
