package config

import (
	"os"
	"path/filepath"
)

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, configFileName), nil
}

func write(cfg Config) error {
	fileName, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := cfg.String()
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, []byte(data), configFileMode)
	if err != nil {
		return err
	}

	return nil
}
