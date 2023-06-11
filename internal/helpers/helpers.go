package helpers

import (
	"encoding/json"
	"os"
)

func ReadConfig[T any](path string) (*T, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var res T
	err = json.NewDecoder(configFile).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
