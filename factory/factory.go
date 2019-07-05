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
	name := specFilename(id)
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
	dir := workspace(id)
	if err := os.MkdirAll(dir, 0777); err != nil {
		return err
	}

	return createSpecFile(id)
}

func createSpecFile(id string) error {
	name := specFilename(id)
	file, err := os.Create(name)
	if err != nil {
		return err
	}

	return file.Close()
}

func save(spec *kibidango.Spec) error {
	name := specFilename(spec.ID)
	dest, err := os.OpenFile(name, os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer dest.Close()

	return json.NewEncoder(dest).Encode(spec)
}

func delete(id string) error {
	name := workspace(id)
	return os.RemoveAll(name)
}

func specFilename(id string) string {
	return filepath.Join(workspace(id), "spec.json")
}

func pipeFilename(id string) string {
	return filepath.Join(workspace(id), "pipe.fifo")
}

func workspace(id string) string {
	return filepath.Join(workSpacesDir, id)
}

const (
	workSpacesDir = "/run/kibidango"
)
