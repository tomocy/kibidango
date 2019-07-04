package factory

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tomocy/kibidango"
)

func list() ([]*kibidango.Spec, error) {
	srces, err := read(workSpacesDir)
	if err != nil {
		return nil, err
	}

	specs := make([]*kibidango.Spec, len(srces))
	for i, src := range srces {
		if !src.IsDir() {
			continue
		}

		spec, err := load(src.Name())
		if err != nil {
			return nil, err
		}

		specs[i] = spec
	}

	return specs, nil
}

func read(dir string) ([]os.FileInfo, error) {
	var infos []os.FileInfo
	if _, err := os.Stat(dir); err != nil {
		return infos, nil
	}

	return ioutil.ReadDir(dir)
}

func load(id string) (*kibidango.Spec, error) {
	name := filepath.Join(workSpacesDir, id, "spec.json")
	src, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer src.Close()

	var spec *kibidango.Spec
	if err := json.NewDecoder(src).Decode(&spec); err != nil {
		return nil, err
	}

	return spec, nil
}

func createWorkspace(id string) error {
	dir := filepath.Join(workSpacesDir, id)
	return os.MkdirAll(dir, 0777)
}

func save(spec *kibidango.Spec) error {
	name := filepath.Join(workSpacesDir, spec.ID, "state.json")
	dest, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer dest.Close()

	return json.NewEncoder(dest).Encode(spec)
}

func delete(id string) error {
	name := filepath.Join(workSpacesDir, id)
	return os.RemoveAll(name)
}

const (
	workSpacesDir = "/run/kibidango"
)
