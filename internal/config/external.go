package config

import (
	"encoding/json"
	"os"
)

func Read() (*Config, error) {
	fileName, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	configFile, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	data := new(Config)
	err = json.Unmarshal(configFile, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
