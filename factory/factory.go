package factory

import (
	"os"
	"path/filepath"
)

func createWorkspace(id string) error {
	dir := filepath.Join(workSpacesDir, id)
	return os.MkdirAll(dir, 0777)
}

const (
	workSpacesDir = "/run/kibidango"
)
