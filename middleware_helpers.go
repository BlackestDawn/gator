package main

import (
	"context"
	"fmt"

	"github.com/BlackestDawn/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		userData, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error fetching user data: %w", err)
		}

		return handler(s, cmd, userData)
	}
}
