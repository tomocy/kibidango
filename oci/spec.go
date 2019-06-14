package oci

import (
	"encoding/json"
	"os"

	"github.com/opencontainers/runtime-spec/specs-go"
)

func LoadSpec(path string) (*specs.Spec, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var spec *specs.Spec
	if err := json.NewDecoder(file).Decode(&spec); err != nil {
		return nil, err
	}

	return spec, nil
}
