package main

import "github.com/BlackestDawn/gator/internal/config"

type state struct {
	conf *config.Config
}

func NewState() (*state, error) {
	cfg, err := config.Read()
	if err != nil {
		return nil, err
	}

	s := new(state)
	s.conf = cfg

	return s, nil
}
