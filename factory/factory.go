package factory

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func createWorkspace(id string) error {
	dir := filepath.Join(workSpacesDir, id)
	return os.MkdirAll(dir, 0777)
}

func save(state *state) error {
	name := filepath.Join(workSpacesDir, state.ID, "state.json")
	dest, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer dest.Close()

	return json.NewEncoder(dest).Encode(state)
}

type state struct {
	ID string
}

const (
	workSpacesDir = "/run/kibidango"
)
