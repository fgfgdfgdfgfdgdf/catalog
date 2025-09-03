package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

func ReadYaml[T any](s *T, path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, s)
	if err != nil {
		return err
	}

	return nil
}
