package factory

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func list() ([]*state, error) {
	srces, err := read(workSpacesDir)
	if err != nil {
		return nil, err
	}

	states := make([]*state, len(srces))
	for i, src := range srces {
		if !src.IsDir() {
			continue
		}

		state, err := load(src.Name())
		if err != nil {
			return nil, err
		}

		states[i] = state
	}

	return states, nil
}

func read(dir string) ([]os.FileInfo, error) {
	var infos []os.FileInfo
	if _, err := os.Stat(dir); err != nil {
		return infos, nil
	}

	return ioutil.ReadDir(dir)
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

func delete(id string) error {
	name := filepath.Join(workSpacesDir, id)
	return os.RemoveAll(name)
}

const (
	workSpacesDir = "/run/kibidango"
)
