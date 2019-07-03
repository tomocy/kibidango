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

func load(id string) (*state, error) {
	name := filepath.Join(workSpacesDir, id, "state.json")
	src, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer src.Close()

	var state *state
	if err := json.NewDecoder(src).Decode(&state); err != nil {
		return nil, err
	}

	return state, nil
}

type state struct {
	ID string
}

const (
	workSpacesDir = "/run/kibidango"
)
