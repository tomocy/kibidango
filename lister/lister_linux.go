package lister

import (
	"io/ioutil"
	"os"

	"github.com/tomocy/kibidango"
)

func ForLinux() *Linux {
	return new(Linux)
}

type Linux struct{}

func (l *Linux) List(loader kibidango.Loader) ([]*kibidango.Kibidango, error) {
	dir := "/run/kibidango"
	files, err := read(dir)
	if err != nil {
		return nil, err
	}

	var kibis []*kibidango.Kibidango
	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		kibi, err := load(loader, file.Name())
		if err != nil {
			return nil, err
		}

		kibis = append(kibis, kibi)
	}

	return kibis, nil
}

func read(dir string) ([]os.FileInfo, error) {
	var infos []os.FileInfo
	if _, err := os.Stat(dir); err != nil {
		return infos, nil
	}

	return ioutil.ReadDir(dir)
}

func load(loader kibidango.Loader, id string) (*kibidango.Kibidango, error) {
	kibi := new(kibidango.Kibidango)
	if err := kibi.UpdateID(id); err != nil {
		return nil, err
	}

	if err := kibi.Load(loader); err != nil {
		return nil, err
	}

	return kibi, nil
}
