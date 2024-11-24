package config

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	return write(*c)
}

func (c *Config) String() (string, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (c *Config) PrintMe() {
	data, err := c.String()
	if err != nil {
		fmt.Println("encountered error printing Config:", err)
	} else {
		fmt.Println("Current config content:", data)
	}
}
