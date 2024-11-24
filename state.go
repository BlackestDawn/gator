package main

import (
	"fmt"

	"github.com/BlackestDawn/gator/internal/config"
	"github.com/BlackestDawn/gator/internal/database"
)

type state struct {
	conf *config.Config
	db   *database.Queries
}

func NewState() (*state, error) {
	cfg, err := config.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	s := new(state)
	s.conf = cfg

	return s, nil
}
